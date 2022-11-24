package order

import (
	"mr.jackpot-backend/model"
	"mr.jackpot-backend/service/order/state"
)

type Order struct {
	ID           int
	DeliveryInfo model.DeliveryInfo
	OrderInfo    model.Order

	TaskList map[int]map[int]int

	currentState state.OrderState

	created    state.OrderState
	accepted   state.OrderState
	started    state.OrderState
	styling    state.OrderState
	prepared   state.OrderState
	delivering state.OrderState
	delivered  state.OrderState
	requested  state.OrderState
	collected  state.OrderState

	rejected   state.OrderState
	canceled  state.OrderState
	finished   state.OrderState	
}

func NewOrder(id int) *Order {
	order := &Order{
		ID: id,
		TaskList: make(map[int]map[int]int),
	}

	order.created = &state.AcceptedState{
		ID: id,
		Order: order,
		NextStep: &order.accepted,
		CeasedStep: &order.rejected,
	}

	order.accepted = &state.AcceptedState{
		ID: id,
		Order: order,
		NextStep: &order.started,
		CeasedStep: &order.canceled,
	}

	order.started = &state.StartedState{
		ID: id,
		Order: order,
		NextStep: &order.prepared,
	}

	order.prepared = &state.PreparedState{
		ID: id,
		Order: order,
		NextStep: &order.delivered,
	}

	order.delivering = &state.DeliveringState{
		ID: id,
		Order: order,
		NextStep: &order.accepted,
	}

	order.delivered = &state.DeliveredState{
		ID: id,
		Order: order,
		NextStep: &order.requested,
	}

	order.requested = &state.RequestedState{
		ID: id,
		Order: order,
		NextStep: &order.collected,
	}

	order.collected = &state.CollectedState{
		ID: id,
		Order: order,
		NextStep: &order.finished,
	}

	order.currentState = order.created

	return order
}



func (o *Order) CreaetOrder(order model.Order, delivery model.DeliveryInfo) {
	o.DeliveryInfo = delivery
	o.OrderInfo = order
}

func (o *Order) CeaseOrder() error {
	return o.currentState.CeaseStep()
}

func (o *Order) ProcessStep() {
	o.currentState.ProcessStep()
}

func (o *Order) SetState(step state.OrderState) {
	o.currentState = step
}

func (o *Order) GetNextStep() *state.OrderState {
	return o.currentState.GetNextStep()
}

func (o *Order) GetOrderInfo() model.Order {
	return o.OrderInfo
}

func (o *Order) GetDeliveryInfo() model.DeliveryInfo {
	return o.DeliveryInfo
}

func (o *Order) GetOrderState() string {
	return o.currentState.GetStateName()
}

func (o *Order) GetAllTaskList() map[int]map[int]int {
	return o.TaskList
}