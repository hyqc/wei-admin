package cronjob

import (
	"context"
	"hope_city/lib/log"
	"sync"
	"time"
)

//周期定时管理  等上次调用运行完才会执行下次调用

var (
	GlobalLtMgr = &LoopTimerManager{Lts: sync.Map{}}
)

type LoopTimer struct {
	Name   string
	timer  *time.Timer
	t      time.Duration
	ctx    context.Context
	cancel context.CancelFunc
}

type LoopTimerManager struct {
	Lts sync.Map
}

// NewLoopTimer 确保name唯一
func (m *LoopTimerManager) NewLoopTimer(name string, t time.Duration, fn func()) {
	timer := m.Get(name)
	if timer != nil {
		log.Infof("循环任务：%s,已经存在,未做任何操作", timer.Name)
		return
	}
	ctx, cancel := context.WithCancel(context.Background())
	timer = &LoopTimer{
		Name:   name,
		timer:  time.NewTimer(t),
		t:      t,
		ctx:    ctx,
		cancel: cancel,
	}
	m.Set(timer)
	go func(lt *LoopTimer) {
		for {
			select {
			case <-lt.timer.C:
				fn()
				if lt.timer != nil {
					lt.timer.Reset(lt.t)
				} else {
					return
				}
			case <-lt.ctx.Done():
				return
			}
		}
	}(timer)
}

func (m *LoopTimerManager) Get(name string) *LoopTimer {
	val, ok := m.Lts.Load(name)
	if ok {
		return val.(*LoopTimer)
	}
	return nil
}

func (m *LoopTimerManager) Set(timer *LoopTimer) {
	log.Infof("循环任务：%s,添加成功", timer.Name)
	m.Lts.Store(timer.Name, timer)
}

func (m *LoopTimerManager) Delete(name string) {
	loop := m.Get(name)
	if loop == nil {
		return
	}
	loop.cancel()
	loop.timer.Stop()
	m.Lts.Delete(loop.Name)
}

func (m *LoopTimerManager) Stop() {
	m.Lts.Range(func(key, value any) bool {
		timer := value.(*LoopTimer)
		defer m.Lts.Delete(key)
		if timer == nil || timer.timer == nil {
			return true
		}
		log.Infof("循环任务：%v,移除成功", key)
		timer.cancel()
		timer.timer.Stop()
		return true
	})
}
