package etcd

import (
	"reflect"
	"testing"
	"time"
)

//测试数据
func TestEtcdGet(t *testing.T) {
	ds, err := NewEtcdDataSource("http://192.168.139.134:2379", time.Second*3)
	if err != nil {
		panic(err)
	}

	ds.set("name", "kapok")
	v, err := ds.get("name")
	if err != nil {
		println("could not get key demo")
		t.Failed()
	}
	equal(t, v, "kapok")
}

func equal(t *testing.T, act, exp interface{}) {
	if !reflect.DeepEqual(exp, act) {
		t.FailNow()
	}
}
