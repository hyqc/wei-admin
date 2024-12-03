package cronjob

import (
	"fmt"
	"github.com/antlabs/timer"
	"hope_city/lib/util"
	"sync"
	"time"
)

var (
	//DefaultTimer 全局时间轮
	DefaultTimer   = timer.NewTimer()
	DefaultCronJob = NewCropJob("DefaultCronJob")
)

// 单次延时任务，周期性任务定时管理器

type CronJob struct {
	name   string
	tw     timer.Timer
	jobs   *sync.Map
	mux    *sync.RWMutex
	isStop bool
}

func init() {
	go DefaultTimer.Run()
}

func NewCropJob(name string) *CronJob {
	return &CronJob{
		name: name,
		tw:   DefaultTimer,
		jobs: &sync.Map{},
		mux:  &sync.RWMutex{},
	}
}

func (c *CronJob) Stop() {
	c.Clear(func() {
		c.isStop = true
	})
}

func (c *CronJob) Clear(f ...func()) {
	c.mux.Lock()
	defer c.mux.Unlock()
	if c.isStop {
		return
	}
	c.jobs.Range(func(key, value any) bool {
		node := value.(timer.TimeNoder)
		if node != nil {
			c.Debug("【定时任务】 %v 执行停止任务-移除任务 Key: %v", c.name, key)
			node.Stop()
			c.jobs.Delete(key)
		}
		return true
	})
	c.jobs = &sync.Map{}
	if len(f) > 0 {
		f[0]()
	}
}

func (c *CronJob) ReStart() {
	c.mux.Lock()
	defer c.mux.Unlock()
	if !c.isStop {
		return
	}
	c.isStop = false
}

func (c *CronJob) AddJob(key string, node timer.TimeNoder) {
	c.mux.RLock()
	defer c.mux.RUnlock()
	if c.isStop {
		node.Stop()
		return
	}
	_, ok := c.jobs.LoadOrStore(key, node)
	if ok {
		node.Stop()
	}
}

func (c *CronJob) LoadJob(key string) bool {
	if c.isStop {
		return false
	}
	_, ok := c.jobs.Load(key)
	return ok
}

func (c *CronJob) RemoveJob(key string) {
	if c.isStop {
		return
	}
	timeNode, ok := c.jobs.LoadAndDelete(key)
	if ok {
		c.Debug("【定时任务】 %v 移除任务 Key: %v", c.name, key)
		timeNode.(timer.TimeNoder).Stop()
	}
}

// AfterFunc 单次任务
func (c *CronJob) AfterFunc(key string, expire time.Duration, callback func()) {
	if callback == nil {
		return
	}
	c.Debug("【单次任务】 %v 添加AfterFunc Key: %v", c.name, key)
	c.AddJob(key, c.tw.AfterFunc(expire, func() {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println("延迟任务panic, err: ", util.PanicTrace(err))
			}
			c.RemoveJob(key)
		}()
		c.Debug("【单次任务】 %v 执行任务 Key: %v", c.name, key)
		callback()
	}))
}

// ScheduleFunc  周期任务
func (c *CronJob) ScheduleFunc(key string, expire time.Duration, callback func()) {
	if callback == nil {
		return
	}
	c.Debug("【周期任务】 %v 添加ScheduleFunc Key: %v", c.name, key)
	c.AddJob(key, c.tw.ScheduleFunc(expire, func() {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println("周期任务panic, err: ", util.PanicTrace(err))
			}
		}()
		if c.isStop {
			return
		}
		//c.Debug("【周期任务】 %v 执行任务 Key: %v", c.name, key)
		callback()
	}))
}

func (c *CronJob) Debug(format string, a ...any) {
	params := []any{time.Now().Format("2006-01-02 15:04:05.000")}
	params = append(params, a...)
	fmt.Printf("[%v] "+format+"\n", params...)
}

func (c *CronJob) Error(format string, a ...any) {
	params := []any{time.Now().Format("2006-01-02 15:04:05.000")}
	params = append(params, a...)
	fmt.Printf("【ERROR】[%v] "+format+"\n", params...)
}
