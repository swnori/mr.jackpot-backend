package worker

import (
	"errors"

	"mr.jackpot-backend/model"
)



type Worker struct {
	Role     string
	TaskList map[int]bool
}

func NewWorker(role string) *Worker {
	return &Worker{
		Role: role,
		TaskList: make(map[int]bool),
	}
}

func (w *Worker) AssignTask(taskid int) error {
	_, exist := w.TaskList[taskid]
	if exist {
		return errors.New("")
	}

	w.TaskList[taskid] = false
	return nil
}

func (w *Worker) StartTask(taskid int) error {
	_, exist := w.TaskList[taskid]
	if !exist {
		return errors.New("")
	}

	w.TaskList[taskid] = true
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

func (w *Worker) GetTaskList() []model.Task {
	tasklist := make([]model.Task, 0)
	
	for taskid, statusbool := range w.TaskList {
		var status string
		if statusbool == false {
			status = model.TaskStatusWaiting
		} else {
			status = model.TaskStatusWorking
		}
		tasklist = append(tasklist, model.Task{
			TaskID: taskid,
			Status: status,
		})
	}

	return tasklist
}