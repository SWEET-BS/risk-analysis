package main

import (
	"riskanly/report"
	"riskanly/store"
)

func main()  {
	store.Tasktmp.Start()
	store.Tasktmp.Run()
	defer store.Tasktmp.Stop()
	report.Run(store.Tasktmp.Jsontask())
}
