package main

import (
	"log"
	"time"

	"github.com/robfig/cron/v3"
)

func main() {
	shanghaiLC, _ := time.LoadLocation("Asia/Shanghai")
	cronRunner := cron.New(cron.WithLocation(shanghaiLC))
	if _, err := cronRunner.AddFunc("@daily", Test1); err != nil {
		log.Fatal("failed register cron job")
	}
	cronRunner.Start()
	defer cronRunner.Stop()
}

func Test1() {

}

func Test2() {

}
