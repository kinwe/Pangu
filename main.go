package main

import (
	"fmt"
	"sync"
	_ "test/core"
	"test/global"
	"test/model"
	"test/service"
)

func main() {

	mserver := model.Mysql{
		Host: "127.0.0.1:3306",
		Password: "hx19870527",
		Username: "kinwe",
		Dbname: "test",
		Relo: model.Relo(0),
	}
	service.ShowMaster(&mserver)

	var ips []string

	ips = make([]string,2)
	ips[0] = "11.42.100.65"
	ips[1] =  "11.40.225.123"

	wg1 := sync.WaitGroup{}
	wg1.Add(len(ips))
	for _,ip := range ips{
		url := fmt.Sprintf(global.GVA_Serve.Dbhost.Dbsurl + "%s",ip)
		go test(url,&wg1)
	}
	wg1.Wait()
}

func test(url string,wg1 *sync.WaitGroup )  {
	size := 10
	wg := sync.WaitGroup{}
	wg.Add(size)
	for i := 0; i != size;i++ {
		go service.Changetask(url,&wg)
	}
	wg.Wait()
	wg1.Done()
}