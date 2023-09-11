package main

import (
	"fmt"
	"riskanly/conf"
	"riskanly/pkg"
	"riskanly/qa"
	"time"
)

//	func main() {
//		// 创建一个定时器，每隔 30 分钟触发一次
//		ticker := time.NewTicker(30 * time.Minute)
//
//		stop := make(chan bool)
//
//		go func() {
//			for {
//				select {
//				case <-ticker.C:
//					makeRequest()
//					if time.Now().Format(time.DateTime) == time.Now().Format(conf.Y_M_D)+conf.CronTime {
//						msg := makeRequest()
//						if strings.Contains(msg, conf.CheckCount) {
//							pkg.RquestDingTalkBot(msg)
//						}
//					}
//				case <-stop:
//					// 收到停止信号时停止定时任务
//					ticker.Stop()
//
//					return
//				}
//			}
//		}()
//
//		// 等待程序退出信号
//		<-make(chan struct{})
//	}
func main() {
	makeRequest()
}
func makeRequest() error {
	qa.Taskindex.Start()
	msg, err := qa.Taskindex.CheckCount()
	fmt.Println(time.Now().Format(time.DateTime), " 指标巡检结果", msg)
	defer qa.Taskindex.Stop()
	if err != nil && err != fmt.Errorf(conf.ErromsgConnectionDb) {
		datemsg := checkDate()
		msg = "警告！" + qa.Taskindex.Name + msg + " " + datemsg
		pkg.RquestDingTalkBot(msg)
		return err
	}
	pkg.RquestDingTalkBot(" 指标巡检结果" + msg)
	return nil
}
func checkDate() string {
	qa.TaskDate.Start()
	msg, err := qa.TaskDate.CheckLatestDate()
	fmt.Println(time.Now().Format(time.DateTime), qa.TaskDate.Name, " 及时性检查结果 ", msg)
	defer qa.Taskindex.Stop()
	if err != nil && err != fmt.Errorf(conf.ErromsgConnectionDb) {
		msg = qa.TaskDate.Name + " " + msg
		return msg
	} else {
		return qa.TaskDate.Name + msg
	}
}
