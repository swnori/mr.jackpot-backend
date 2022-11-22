package order

import "mr.jackpot-backend/model"

type Order struct {
	ID           int
	DeliveryInfo model.DeliveryInfo
	OrderInfo    model.Order

	currentState OrderState

	created    OrderState
	accepted   OrderState
	started    OrderState
	cooking    OrderState
	styling    OrderState
	prepared   OrderState
	delivering OrderState
	delivered  OrderState
	requested  OrderState
	collected  OrderState

	rejected   OrderState
	cancelled  OrderState
	finished   OrderState	
}

func NewOrder(id int) *Order {
	order := &Order{
		ID: id,
	}

	order.created    = &CreatedState   {ID: id, NextStep: &order.accepted  }
	order.accepted   = &AcceptedState  {ID: id, NextStep: &order.started   }
	order.started    = &StartedState   {ID: id, NextStep: &order.cooking   }
	order.cooking    = &CookingState   {ID: id, NextStep: &order.styling   }
	order.prepared   = &PreparedState  {ID: id, NextStep: &order.delivering}
	order.delivering = &DeliveringState{ID: id, NextStep: &order.accepted  }
	order.delivered  = &DeliveredState {ID: id, NextStep: &order.requested }
	order.requested  = &RequestedState {ID: id, NextStep: &order.collected }
	order.collected  = &CollectedState {ID: id, NextStep: &order.finished  }

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

func (o *Order) SetState(step OrderState) {
	o.currentState = step
}

func (o *Order) GetNextStep() *OrderState {
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
