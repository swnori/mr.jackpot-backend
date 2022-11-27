package state

 

type OrderState interface {
	ProcessStep() error
	CeaseStep() error
	GetNextStep() *OrderState
	GetStateName() string
}

type OrderInterface interface {
	GetAllTaskList() map[int][]int
}

type State struct {
	ID int
	StateID int
	Order OrderInterface
	NextStep *OrderState
	CeasedStep *OrderState
}

func (s *State) GetNextStep() *OrderState {
	return s.NextStep
}
























