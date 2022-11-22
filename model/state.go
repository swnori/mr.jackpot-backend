package model

var (
	StateCreated    string = "Created"
	StateAccepted   string = "Accepted"
	StateStarted   string = "Started"
	StateCooking    string = "Cooking"
	StatePrepared   string = "Prepared"
	StateDelivering string = "Delivering"
	StateDelivered  string = "Delivered"
	StateRequested  string = "Requested"
	StateCollected  string = "Collected"
	
	StateCeased   string = "Ceased"
	StateCanceled string = "Canceled"
	StateFinished string = "Finished"
)


var StateList []string = []string{
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
