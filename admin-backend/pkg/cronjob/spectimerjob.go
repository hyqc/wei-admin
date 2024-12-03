package cronjob

import (
	"github.com/robfig/cron/v3"
	"hope_city/lib/log"
	"sync"
)

var (
	GlobalSpecMgr = NewSpecTimerJobManager()
)

//定点定时任务管理器 AddTask("0 0 5 * * *", ManagerObj.ClearManorCache, "定时清除庄园node缓存")

type Task struct {
	id   cron.EntryID
	name string
	spec string
}

type SpecTimerJobManager struct {
	tw     *cron.Cron
	tasks  sync.Map
	isStop bool
	mux    *sync.RWMutex
}

func NewSpecTimerJobManager(opts ...cron.Option) *SpecTimerJobManager {
	if len(opts) == 0 {
		opts = append(opts, cron.WithSeconds())
	}
	return &SpecTimerJobManager{
		tw:    cron.New(opts...),
		tasks: sync.Map{},
		mux:   &sync.RWMutex{},
	}
}

func (m *SpecTimerJobManager) AddTask(spec string, job cron.FuncJob, name ...string) (entryId cron.EntryID, err error) {
	m.mux.RLock()
	defer m.mux.RUnlock()
	if m.isStop {
		return
	}
	taskName := "未知任务"
	if len(name) > 0 {
		taskName = name[0]
	}
	entryId, err = m.tw.AddJob(spec, job)
	task := &Task{
		id:   entryId,
		name: taskName,
		spec: spec,
	}
	if err == nil {
		log.Infof("添加定时任务：taskId:%d,name:%s,spec:%s", task.id, task.name, task.spec)
		m.tasks.Store(entryId, task)
	} else {
		log.Errorf("添加定时任务失败：taskId:%d,name:%s,spec:%s，err:%s", task.id, task.name, task.spec, err.Error())
	}
	return
}

func (m *SpecTimerJobManager) RemoveTask(entryId cron.EntryID) {
	m.mux.RLock()
	defer m.mux.RUnlock()
	if m.isStop {
		return
	}
	val, ok := m.tasks.LoadAndDelete(entryId)
	if ok {
		task := val.(*Task)
		m.tw.Remove(entryId)
		log.Infof("删除定时任务：taskId:%d,name:%s,spec:%s", task.id, task.name, task.spec)
		task = nil
	}
}

func (m *SpecTimerJobManager) Start() {
	m.tw.Start()
	m.isStop = false
}

func (m *SpecTimerJobManager) Stop() {
	m.mux.Lock()
	defer m.mux.Unlock()
	if m.isStop {
		return
	}
	m.isStop = true
	ctx := m.tw.Stop()
	<-ctx.Done()
	log.Info("SpecTimerJobManager Stopped")
	m.tasks = sync.Map{}
}
