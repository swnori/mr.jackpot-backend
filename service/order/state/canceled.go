package state


type CanceledState struct {
	State
	ID       int
	StateID  int
	Order    OrderInterface
	NextStep *OrderState
}

func (o *CanceledState) ProcessStep() error {
	return nil
}
func (o *CanceledState) CeaseStep() error {
	return nil
}

func (o *CanceledState) GetStateId() int {
	return o.StateID
}
