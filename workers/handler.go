package workers

import (
	"github.com/8tomat8/GoRepost/counter"
	"github.com/8tomat8/GoRepost/task"
)

// Handler of all incoming tasks
func Handler(t *task.Task) {
	c := counter.GetCounter()
	c.JobStarted()
	defer c.JobFinished()
	for _, dest := range t.Destinations {
		switch social := dest.Social; social {
		case "vk":
			vk(t)
		case "fb":
		case "gp":
		case "tw":
		}
	}
}
