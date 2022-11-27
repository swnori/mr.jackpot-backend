package order

import (
	"mr.jackpot-backend/service/order/state"
	"mr.jackpot-backend/model"
)

type Order struct {
	ID        int
	OrderInfo model.AllOrderInfo
	Order     model.Order

	TaskList map[int][]int

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
		TaskList: make(map[int][]int),
	}

	order.created = &state.AcceptedState{
		ID: id,
		StateID: 4,
		Order: order,
		NextStep: &order.accepted,
		CeasedStep: &order.rejected,
	}

	order.accepted = &state.AcceptedState{
		ID: id,
		StateID: 5,
		Order: order,
		NextStep: &order.started,
		CeasedStep: &order.canceled,
	}

	order.started = &state.StartedState{
		ID: id,
		StateID: 6,
		Order: order,
		NextStep: &order.prepared,
	}

	order.prepared = &state.PreparedState{
		ID: id,
		StateID: 7,
		Order: order,
		NextStep: &order.delivered,
	}

	order.delivering = &state.DeliveringState{
		ID: id,
		StateID: 8,
		Order: order,
		NextStep: &order.accepted,
	}

	order.delivered = &state.DeliveredState{
		ID: id,
		StateID: 9,
		Order: order,
		NextStep: &order.requested,
	}

	order.requested = &state.RequestedState{
		ID: id,
		StateID: 10,
		Order: order,
		NextStep: &order.collected,
	}

	order.collected = &state.CollectedState{
		ID: id,
		StateID: 11,
		Order: order,
		NextStep: &order.finished,
	}

	order.rejected = &state.RejectedState{
		ID: id,
		StateID: 12,
		Order: order,
	}

	order.canceled = &state.CanceledState{
		ID: id,
		StateID: 13,
		Order: order,
	}

	order.finished = &state.FinishedState{
		ID: id,
		StateID: 14,
		Order: order,
	}

	order.currentState = order.created

	return order
}



func (o *Order) CreateOrder(order model.Order, info model.AllOrderInfo) {
	o.OrderInfo = info
	o.Order = order
}

func (o *Order) GetOrderInfo() model.AllOrderInfo {
	return o.OrderInfo
}

func (o *Order) GetOrder() model.Order {
	return o.Order
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

func (o *Order) GetOrderState() int {
	return o.currentState.GetStateId()
}

func (o *Order) GetAllTaskList() map[int][]int {
	return o.TaskList
}