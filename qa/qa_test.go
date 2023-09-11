package qa

import (
	"riskanly/conf"
	"testing"
)

func TestA(t *testing.T) {
	Taskindex.Start()
	Taskindex.CheckCount()
}
func TestB(t *testing.T) {
	TaskDate.Start()
	TaskDate.CheckLatestDate()
}
func TestC(t *testing.T) {
	TaskDate.SetDSN(conf.DsnLocal)
	TaskDate.Start()
}
