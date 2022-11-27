package state


type AcceptedState struct {
	State
	ID         int
	StateID    int
	Order      OrderInterface
	NextStep   *OrderState
	CeasedStep *OrderState
}

func (o *AcceptedState) ProcessStep() error {

	return nil
}
func (o *AcceptedState) CeaseStep() error {
	return nil
}
func (o *AcceptedState) GetStateId() int {
	return o.StateID
}
