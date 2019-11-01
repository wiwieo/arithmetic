// 无限父子级联，不支持环形
package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"reflect"
	"time"
	"unsafe"
)

func main() {
	data, err := json.Marshal(tree(generateData2()))
	if err != nil {
		panic(err)
	}
	println(fmt.Sprintf("%s", data))
}

// 该数据有问题，一个环形数据。父-子-孙-父
func generateData2() []*info {
	temp := make([]*info, 0)
	data := []byte(`[{"Id":0,"Pid":8,"Title":"","Children":null},{"Id":1,"Pid":6,"Title":"","Children":null},{"Id":2,"Pid":6,"Title":"","Children":null},{"Id":3,"Pid":9,"Title":"","Children":null},{"Id":4,"Pid":6,"Title":"","Children":null},{"Id":5,"Pid":4,"Title":"","Children":null},{"Id":6,"Pid":9,"Title":"","Children":null},{"Id":7,"Pid":2,"Title":"","Children":null},{"Id":8,"Pid":2,"Title":"","Children":null},{"Id":9,"Pid":8,"Title":"","Children":null}]`)
	err := json.Unmarshal(data, &temp)
	if err != nil{
		panic(err)
	}
	return temp
}

func generateData() []*info {
	length := 10
	temp := make([]*info, 0, length)
	for i := 0; i < length; i++ {
		pid := int(rand.NewSource(time.Now().UnixNano()).Int63())%length + 1
		for i == pid {
			pid = int(rand.NewSource(time.Now().UnixNano()).Int63())%length + 1
		}
		temp = append(temp, &info{
			Id:  i,
			Pid: pid,
		})
		time.Sleep(time.Nanosecond)
	}
	data, err := json.Marshal(temp)
	if err != nil {
		panic(err)
	}
	println(fmt.Sprintf("%s", data))
	return temp
}

type info struct {
	Id       int
	Pid      int
	Title    string
	Children []*info
}

func tree(origins []*info) []*info {
	nodes := make([]*info, 0)
	done := (*bool)(unsafe.Pointer(reflect.ValueOf(new(bool)).Pointer()))
	for _, c := range origins {
		node(nodes, c, done)
		if !*done {
			nodes = append(nodes, c)
		}
		*done = false
	}
	for i := len(nodes) - 1; i >= 0; i-- {
		if nodes[i] == nil {
			nodes = append(nodes[:i], nodes[i+1:]...)
		}
	}
	return nodes
}

func node(infos []*info, c *info, done *bool) {
	for idx, n := range infos {
		if n == nil {
			continue
		}
		// 当前菜单是子菜单
		if c.Pid == n.Id {
			n.Children = append(n.Children, c)
			*done = true
			continue
		}
		// 当前菜单是父菜单
		if c.Id == n.Pid {
			c.Children = append(c.Children, n)
			infos[idx] = nil
		}
		if len(n.Children) > 0 {
			node(n.Children, c, done)
		}
	}
}
