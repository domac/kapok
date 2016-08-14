package hb

import (
	"encoding/json"
	"fmt"
	"net"
	"sort"
	"strings"
	"sync"
	"time"
)

var (
	DEFAULT_RECYCLE     = 2 * time.Second
	LEVEL_FULL      int = 9
	LEVEL_ALIVE     int = 6
	LEVEL_WARNING   int = 3
	LEVEL_DEAD      int = 0
)

type Host struct {
	Ip    string
	Time  time.Time
	Alive bool
	HP    int
}

type Doctor struct {
	sync.Mutex
	Addr            string
	RecycleDuration time.Duration
	Hosts           map[string]*Host
}

func NewDoctor(addr string) *Doctor {
	return &Doctor{
		Addr:            addr,
		RecycleDuration: DEFAULT_RECYCLE,
		Hosts:           make(map[string]*Host, 100),
	}
}

// recover HP
func (this *Doctor) fix(ip string) {
	this.Lock()
	defer this.Unlock()
	h, ok := this.Hosts[ip]
	if !ok {
		this.Hosts[ip] = &Host{Ip: ip, Time: time.Now(), HP: LEVEL_FULL, Alive: true}
		return
	}
	h.HP += 1
	if h.HP > LEVEL_FULL {
		h.HP = LEVEL_FULL
	}
	if h.HP < LEVEL_WARNING {
		h.HP = LEVEL_WARNING
	}
	h.Time = time.Now()
	this.updateStatus(ip)
}

//更新状态
func (this *Doctor) updateStatus(ip string) {
	host, ok := this.Hosts[ip]
	if !ok {
		return
	}
	if host.HP >= LEVEL_ALIVE {
		host.Alive = true
	}
	if host.HP == 0 {
		host.Alive = false
	}
}

// cut down HP
func (this *Doctor) hurt(ip string) {
	h, ok := this.Hosts[ip]
	if !ok {
		return
	}
	if h.HP -= 1; h.HP < 0 {
		h.HP = 0
	}
	this.updateStatus(ip)
}

//工作处理
func (this *Doctor) oncall(notify chan Host) {
	for {
		this.Lock()

		if len(this.Hosts) == 0 {
			println("empty host to connect")
		}

		for _, h := range this.Hosts {

			var state = h.Alive

			this.hurt(h.Ip)

			if h.Alive != state {
				delete(this.Hosts, h.Ip)
				notify <- *h
			}
		}
		this.Unlock()
		time.Sleep(this.RecycleDuration)
	}
}

//心跳诊断
func (this *Doctor) Watch() (chan Host, error) {
	ch, err := this.listenUDP()
	if err != nil {
		return nil, err
	}
	go func() {
		for {
			ip := <-ch
			this.fix(ip)
		}
	}()
	notify := make(chan Host, 50)
	go this.oncall(notify)
	return notify, nil
}

//监听UDP包
func (this *Doctor) listenUDP() (chan string, error) {
	packetConn, err := net.ListenPacket("udp", this.Addr)
	if err != nil {
		return nil, err
	}
	fmt.Printf("start listen packet from : %s \n", this.Addr)
	ch := make(chan string)
	go func() {
		defer packetConn.Close()
		buf := make([]byte, 1000)
		for {
			n, addr, err := packetConn.ReadFrom(buf)
			msg := string(buf[:n])
			fmt.Printf("Receive from %s : %s \n", addr, msg)
			if err != nil {
				fmt.Errorf(err.Error())
				continue
			}
			ip := strings.Split(addr.String(), ":")[0]
			ch <- ip
		}
	}()
	return ch, nil
}

//输出当前监控数据
func (this *Doctor) JSONMessage() []byte {
	this.Lock()
	defer this.Unlock()
	aliveds := make([]string, 0, len(this.Hosts))
	deads := make([]string, 0, len(this.Hosts))

	for ip, host := range this.Hosts {
		if host.Alive {
			aliveds = append(aliveds, LookupName(ip))
		} else {
			deads = append(deads, LookupName(ip))
		}
	}

	sort.Strings(aliveds)
	sort.Strings(deads)

	//输出JSON格式
	data, err := json.Marshal(struct {
		Aliveds []string
		Deads   []string
	}{
		aliveds,
		deads,
	})

	if err != nil {
		panic(err)
	}

	return data
}
