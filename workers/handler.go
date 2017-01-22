package workers

import (
	"github.com/8tomat8/GoRepost/counter"
	"github.com/8tomat8/GoRepost/task"
	"strings"
	"github.com/golang/glog"
	"github.com/8tomat8/GoRepost/logging"
)

// Handler of all incoming tasks
func Handler(t *task.Task) {
	c := counter.GetCounter()
	c.JobStarted()
	defer c.JobFinished()
	for social := range t.Destinations {
		switch social = strings.ToLower(social); social {
		case "vk":
			vk(t)
		case "fb":
		case "gp":
		case "tw":
		}
	}

	err := logging.WriteLog(t)
	if err != nil {
		glog.Error(err, t)
	}
}
