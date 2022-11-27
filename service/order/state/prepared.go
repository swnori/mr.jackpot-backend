package state

import (
	"errors"

	//"mr.jackpot-backend/service/order"
)

type PreparedState struct {
	State
	ID       int
	StateID  int
	Order    OrderInterface
	NextStep *OrderState
}

func (o *PreparedState) ProcessStep() error {
	//order := order.OrderManagers.Orders[o.ID]
	//for i := range order.Order.DinnerList {
	//	order.Order.DinnerList[i].StateId = 3
	//}

	

	return nil
}
func (o *PreparedState) CeaseStep() error {
	return errors.New("")
}

func (o *PreparedState) GetStateId() int {
	return o.StateID
}
