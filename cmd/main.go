package main

import (
	"riskanly/pkg"
	"riskanly/store"
)

func main()  {
	store.Tasktmp.Start()
	store.Tasktmp.Run()
	defer store.Tasktmp.Stop()
	pkg.RquestDingTalkBot(store.Tasktmp.Jsontask())
}
