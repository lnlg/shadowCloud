package task

import "shadowCloud/internal/crontab"

func Tasks() []crontab.TaskInterface {
	return []crontab.TaskInterface{
		// &TestTask{},
	}
}
