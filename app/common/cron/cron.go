package cron

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"github.com/spf13/viper"
	"sync"
	"temp/app/task"
	util "temp/app/utils"
)

var c *cron.Cron
var cronOnce sync.Once

func Setup() {
	if viper.GetString("cron.enabled") == "false" {
		return
	}
	cronOnce.Do(func() {
		go InitTask()
	})
}

func InitTask() {
	c = cron.New()
	TaskDemo()
	TaskDemo2()
	c.Start()
}

func TaskDemo() {
	e1, err := c.AddFunc("@every 5s", func() {
		util.Logger().Info("TaskDemo started")
		task.TaskDemo()
	})
	if err != nil {
		c.Stop()
		fmt.Println(e1, err)
	}
}

func TaskDemo2() {
	e1, err := c.AddFunc("@every 3s", func() {
		util.Logger().Info("TaskDemo2 started")
		fmt.Println("task demo 2")
	})
	if err != nil {
		c.Stop()
		fmt.Println(e1, err)
	}
}
