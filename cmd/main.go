package main

import (
	"fmt"
	"github.com/duguying/gomtr"
	"github.com/gogather/com/log"
	"time"
)

func main() {
	mtr := gomtr.NewMtrService("./mtr-packet")
	go mtr.Start()

	time.Sleep(time.Second * 5)

	fmt.Println(mtr.GetServiceStartupTime())

	iplist := []string{"4.4.4.4", "183.131.7.130", "127.0.0.1", "114.215.151.25", "111.13.101.208"}

	for i := 0; i < len(iplist); i++ {
		id := i
		go mtr.Request(iplist[i], 10, func(response interface{}) {
			task := response.(*gomtr.MtrTask)
			log.Bluef("[ID] %d cost: %d ms\n", id, task.CostTime/1000000)
			fmt.Println(task.GetSummaryDecorateString())
		})
	}

	time.Sleep(time.Minute)
}
