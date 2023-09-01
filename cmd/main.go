package main

import (
	"riskanly/pkg"
	"riskanly/store"
)

func main() {
	store.Taskindex.Start()
	msg := store.Taskindex.CheckCount()
	defer store.Taskindex.Stop()
	if msg != "" {
		msg = "警告！规则数据存在数据量为空" + msg
		pkg.RquestDingTalkBot(msg)
	} else {
		msg = "一切正常"
		pkg.RquestDingTalkBot(msg)
	}
}
