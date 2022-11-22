package worker

import (
	"errors"

	"mr.jackpot-backend/database/db"
	"mr.jackpot-backend/model"
)



type WorkerController interface {
	AddWorker(workerid int) error
	RemoveWorker(workerid int) error
	GetSubtaskList(workerid int) ([]model.Task, error)

	StartTaskProcess(taskid int, subtask []int) error
	FinishTaskProcess(taskid int) error

	StartSubTask(workerid int, taskid int) error
	FinishSubTask(workerid int, taskid int) error
}

type WorkerManager struct {
	db db.StaffLayer
	Workers map[int]*Worker
	Tasklist map[int][]int
}



func (w *WorkerManager) AddWorker(workerid int) error {
	_, exist :=  w.Workers[workerid]
	if !exist {
		return errors.New("")
	}
	
	role, err := w.db.GetStaffRole(workerid)
	if err != nil {
		return err
	}

	w.Workers[workerid] = NewWorker(role)

	if len(w.Workers) == 1 {
		for _, subtastlist := range w.Tasklist {
			for _, subtask := range subtastlist {
				if err := w.AssignSubTask(workerid, subtask); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (w *WorkerManager) RemoveWorker(workerid int) error {
	worker, exist :=  w.Workers[workerid]
	if exist {
		return errors.New("")
	}

	tasklist := make([]int, 0)
	for key := range worker.TaskList {
		tasklist = append(tasklist, key)
	}

	delete(w.Workers, workerid);

	for taskid := range tasklist {
		workerid, err := w.GetLeastWorker()
		if err != nil {
			return nil
		}

		w.AssignSubTask(workerid, taskid)
	}

	return nil
}

func (w *WorkerManager) GetLeastWorker() (workerid int, err error) {
	count := 10

	if len(w.Workers) == 0 {
		return 0, errors.New("")
	}

	for id, worker := range w.Workers {
		if count < len(worker.TaskList) {
			workerid = id
		}
	}
	return
}

func (w *WorkerManager) AssignSubTask(workerid int, taskid int) error {
	worker, exist :=  w.Workers[workerid]
	if exist {
		return errors.New("")
	}

	return worker.AssignTask(taskid)
}

func (w *WorkerManager) RemoveSubTask(workerid int, taskid int) error {
	worker, exist :=  w.Workers[workerid]
	if exist {
		return errors.New("")
	}

	return worker.AssignTask(taskid)
}

func (w *WorkerManager) StartSubTask(workerid int, taskid int) error {
	worker, exist :=  w.Workers[workerid]
	if exist {
		return errors.New("")
	}

	return worker.StartTask(taskid)
}

func (w *WorkerManager) FinishSubTask(workerid int, taskid int) error {
	if err := w.RemoveSubTask(workerid, taskid); err != nil {
		return err
	}
	
	for _, taskid := range w.CheckAllTaskFinished() {
		return w.FinishTaskProcess(taskid)
	}
	return nil
}

func (w *WorkerManager) CheckAllTaskFinished() []int {
	finished := make([]int, 0)
	for key, value := range w.Tasklist {
		if len(value) == 0 {
			finished = append(finished, key)
		}
	}
	return finished
}

func (w *WorkerManager) StartTaskProcess(taskid int, subtaskid []int) error {
	_, exist := w.Tasklist[taskid]
	if exist {
		return errors.New("")
	}

	w.Tasklist[taskid] = subtaskid
	return nil
}

func (w *WorkerManager) FinishTaskProcess(taskid int) error {
	// 다음에 task를 어디로 넘겨줄건지 판단하는거라 개별적으로 구현이 필요함
	return errors.New("")
}

func (w *WorkerManager) GetSubtaskList(workerid int) ([]model.Task, error) {
	worker, exist := w.Workers[workerid]
	if !exist {
		return nil, errors.New("")
	}
	return worker.GetTaskList(), nil
}