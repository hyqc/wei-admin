package cronjob

import (
	"fmt"
	"testing"
	"time"
)

func TestNewCropJob(t *testing.T) {
	job := NewCropJob("a")
	job.AfterFunc("1", time.Second*0, func() {
		fmt.Println("=====================")
	})

	time.Sleep(time.Second * 5)
}
