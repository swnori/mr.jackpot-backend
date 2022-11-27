package worker

import (
	"errors"
)



type Worker struct {
	Role     string
	TaskList map[int][]int
}

func NewWorker(role string) *Worker {
	return &Worker{
		Role: role,
		TaskList: make(map[int][]int),
	}
}

func (w *Worker) AssignTask(taskid int) error {
	_, exist := w.TaskList[taskid]
	if exist {
		return errors.New("")
	}

	//w.TaskList[taskid].SetTaskStatus(model.TaskStatusWaiting)
	return nil
}

func (w *Worker) StartTask(taskid int) error {
	_, exist := w.TaskList[taskid]
	if !exist {
		return errors.New("")
	}

	//w.TaskList[taskid].SetTaskStatus(model.TaskStatusWorking)
	return nil
}

func (w *Worker) RemoveTask(taskid int) error {
	_, exist := w.TaskList[taskid]
	if !exist {
		return errors.New("")
	}

	delete(w.TaskList, taskid);
	return nil
}
/*
func (w *Worker) GetTaskList() []int {
	tasklist := make([]int, 0)
	
	for taskid, task := range w.TaskList {
		var status string
		if task.GetTaskStatus() == model.TaskStatusWaiting {
			status = model.TaskStatusWaiting
		} else {
			status = model.TaskStatusWorking
		}
		tasklist = append(tasklist,)
	}

	return tasklist
}*/