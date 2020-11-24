package task

import (
	"fmt"
	"github.com/ebar-go/ego/app"
	"github.com/ebar-go/egu"
	"time"
)

type DemoJob struct {
}

func (job DemoJob) Run() {
	fmt.Println("task is running..", time.Now())
}

func Load() {
	err := app.Task().AddJob("*/2 * * * * ?", new(DemoJob))
	egu.SecurePanic(err)
}
