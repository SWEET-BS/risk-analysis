package main

import (
	"fmt"
	"riskanly/pkg"
	"riskanly/qa"
	"time"
)

func main() {
	// 创建一个定时器，每隔 30 分钟触发一次
	ticker := time.NewTicker(30 * time.Minute)

	stop := make(chan bool)

	go func() {
		for {
			select {
			case <-ticker.C:
				makeRequest()
			case <-stop:
				// 收到停止信号时停止定时任务
				ticker.Stop()

				return
			}
		}
	}()

	// 等待程序退出信号
	<-make(chan struct{})
}

func makeRequest() {
	qa.Taskindex.Start()
	msg := qa.Taskindex.CheckCount()
	fmt.Println(time.Now().Format("2006-01-02 15:04:05")," 指标巡检结果",msg)
	defer qa.Taskindex.Stop()
	if msg != "" {
		msg = "警告！指标表rule_id存在数据量为空" + msg
		checkDate()
		pkg.RquestDingTalkBot(msg)
	}
}
func checkDate(){
	qa.TaskDate.Start()
	msg :=qa.TaskDate.CheckLatestDate()
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"),qa.TaskDate.Name," 及时性检查结果",msg)
	defer qa.Taskindex.Stop()
	if msg != "" {
		msg = "警告！"+qa.TaskDate.Name + "不通过" + msg
		pkg.RquestDingTalkBot(msg)
	}
}