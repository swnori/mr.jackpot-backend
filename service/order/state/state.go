package state

 

type OrderState interface {
	ProcessStep() error
	CeaseStep() error
	GetNextStep() *OrderState
	GetStateName() string
}



type State struct {
	ID int
	NextStep *OrderState
	CeasedStep *OrderState
}

func (s *State) GetNextStep() *OrderState {
	return s.NextStep
}
























