package main

import (
	"fmt"
	"github.com/robfig/cron"
	"time"
)

func main()  {
	fmt.Println("starting...")

	c:=cron.New()

	c.AddFunc("* * * * * *", func() {
		fmt.Println("每秒触发一次")
	})

	c.Start()

	t1:=time.NewTimer(time.Second*10)
	for {
		select {
		case <-t1.C:
			fmt.Println("time的定时任务")
			t1.Reset(time.Second*10)
		}
	}
}