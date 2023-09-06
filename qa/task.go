package qa

import "time"

type Task1 struct {
	ID              int64
	Name            string
	Type            string
	TriggerTime     time.Time
	EndTime         time.Time
	ConnectConfig   string
	IsStart         bool
	IsStop          bool
	Description     string
	ExecutionParams map[string]interface{}
	Priority        string
	Schedule        *TaskSchedule
	ReportFormat    string
	Status          string
	Logs            []*TaskLog
	Notifications   []*Notification
}
type TaskSchedule struct {
	Frequency   string
	Interval    int
	StartDate   time.Time
	EndDate     time.Time
}

type TaskLog struct {
	Timestamp   time.Time
	Message     string
}

type Notification struct {
	Type        string
	Recipients  []string
}
func (t *Task1) Start()  {

}
func (t *Task1) Run(){

}
func (t *Task1) Stop() {

}
