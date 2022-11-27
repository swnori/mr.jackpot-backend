package state

type OrderState interface {
	ProcessStep() error
	CeaseStep() error
	GetNextStep() *OrderState
	GetStateId() int
}

type OrderInterface interface {
	GetAllTaskList() map[int][]int
}

type State struct {
	ID         int
	OrderID    int
	Order      OrderInterface
	NextStep   *OrderState
	CeasedStep *OrderState
}

func (s *State) GetNextStep() *OrderState {
	return s.NextStep
}
