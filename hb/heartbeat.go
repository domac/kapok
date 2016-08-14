package hb

import (
	"fmt"
	"net"
	"time"
)

//开启心跳
func OpenHeartBeat() {
	time.Sleep(3 * time.Second)
	doctor := NewDoctor(":8888")
	notify, err := doctor.Watch()
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		host := <-notify
		fmt.Println(string(doctor.JSONMessage()))

		if !host.Alive {
			fmt.Printf("%s is dead \n", host.Ip)
		}
	}
}

//发送心跳
func SendHeartBeat(addr string) {
	go func() {
		for {
			Send(addr, "kapok action")
			time.Sleep(time.Second)
		}
	}()
}

//发送消息
func Send(addr string, msg string) {
	conn, err := net.DialTimeout("udp", addr, time.Second*2)
	if err != nil {
		fmt.Errorf(err.Error())
		return
	}
	defer conn.Close()
	fmt.Printf(" %s send msg: %s \n", addr, msg)
	_, err = conn.Write([]byte(msg))
	return
}
