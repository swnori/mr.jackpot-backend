package model

var (
	StateCreated    int = 4
	StateAccepted   int = 5
	StateStarted    int = 6
	StateCooking    int = 7
	StatePrepared   int = 8
	StateDelivering int = 9
	StateDelivered  int = 10
	StateRequested  int = 11
	StateCollected  int = 12
	
	StateCeased   int = 13
	StateCanceled int = 14
	StateFinished int = 15
)


var StateList = []int{
	StateCreated,
	StateAccepted,
	StateStarted,
	StateCooking,
	StatePrepared,
	StateDelivering,
	StateDelivered,
	StateRequested,
	StateCollected,

	StateCeased,
	StateCanceled,
	StateFinished,
}

type OrderState struct {
	Id int       `json:"id"`
	State string `json:"state"`
}
