package state


type CreatedState struct {
	State
	ID         int
	StateID    int
	Order      OrderInterface
	NextStep   *OrderState
	CeasedStep *OrderState
}

func (o *CreatedState) ProcessStep() error {
	return nil
}
func (o *CreatedState) CeaseStep() error {
	return nil
}

func (o *CreatedState) GetStateId() int {
	return o.StateID
}
