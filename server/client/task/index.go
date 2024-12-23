package task

import (
	"carambola/client/global"
	"fmt"
	"slices"

	cron "github.com/robfig/cron/v3"
)

var tc *TaskCron

func GetTaskCron() *TaskCron {
	if tc == nil {
		tc = NewTaskCron()
	}
	return tc
}

func InitTask() {
	GetTaskCron()
	tc.Start()
}

type TaskRun struct {
	ID            cron.EntryID `json:"cronId"`        // 定时任务ID
	TaskID        string       `json:"taskId"`        // 任务ID
	ScriptContent string       `json:"scriptContent"` // 脚本内容
	Job           func()       // 方法
}

func (tr TaskRun) Run() {
	// 每次执行任务前，先记录日志
	fmt.Printf("TaskRun: %d\n", tr.ID)
	tr.Job()
}

type TaskCron struct {
	*cron.Cron
	list []*TaskRun
}

func (tc *TaskCron) GetTask(id string) *TaskRun {
	for _, info := range tc.list {
		if info.TaskID == id {
			return info
		}
	}

	return nil
}

func (tc *TaskCron) GetEntry(id string) cron.Entry {
	for _, info := range tc.list {
		if info.TaskID == id {
			return tc.Entry(info.ID)
		}
	}

	return cron.Entry{}
}

func NewTaskRun(id string, job func()) TaskRun {
	return TaskRun{
		TaskID: id,
		Job:    job,
	}
}

func NewTaskCron() *TaskCron {
	c := cron.New(cron.WithSeconds())
	return &TaskCron{
		Cron: c,
		list: make([]*TaskRun, 0),
	}
}

func (tc *TaskCron) Start() {
	tc.Cron.Start()
}

func (tc *TaskCron) Run(name string, fn func()) {
	global.Logger.Sugar().Infof("执行任务: %s", name)
	fn()
}

func (tc *TaskCron) AddFunc(id string, cmd func()) error {
	tRun := NewTaskRun(id, cmd)
	entryID, err := tc.Cron.AddJob("", tRun)
	if err != nil {
		return err
	}

	tc.list = append(tc.list, &TaskRun{TaskID: id, ID: entryID})
	return nil
}

// 添加到指定时间点执行
func (tc *TaskCron) AddFuncAt(id string, spec string, cmd func()) error {
	tRun := NewTaskRun(id, cmd)
	entryID, err := tc.Cron.AddJob(spec, tRun)
	if err != nil {
		return err
	}
	tc.list = append(tc.list, &TaskRun{TaskID: id, ID: entryID})
	return nil
}

func (tc *TaskCron) Remove(id string) {
	job := tc.GetTask(id)
	if job == nil {
		return
	}

	tc.list = slices.DeleteFunc(tc.list, func(task *TaskRun) bool {
		return task.TaskID == id
	})

	tc.Cron.Remove(job.ID)
}

func (tc *TaskCron) CronList() []*TaskRun {
	return tc.list
}
