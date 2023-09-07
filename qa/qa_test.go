package qa

import "testing"

func TestA(t *testing.T) {
	Taskindex.Start()
	Taskindex.CheckCount()
}
func TestB(t *testing.T) {
	TaskDate.Start()
	TaskDate.CheckLatestDate()
}
