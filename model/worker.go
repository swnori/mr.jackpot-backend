package model


type Task struct {
	TaskID int
	Status string
}

var (
	TaskStatusWaiting   string = "Waiting"
	TaskStatusWorking   string = "Working"
	TaaskStatusFinished string = "Finished"
)

var TaskStatus = []string{
	TaskStatusWaiting,
	TaskStatusWorking,
	TaaskStatusFinished,
}

