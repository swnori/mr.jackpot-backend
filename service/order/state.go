package order

import (
	"errors"

	"mr.jackpot-backend/model"
)


 

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


type CreatedState struct {
	ID int
	NextStep *OrderState
}

func (o *CreatedState) ProcessStep() error {
	return nil
}
func (o *CreatedState) CeaseStep() error {
	return nil
}
func (o *CreatedState) GetNextStep() *OrderState {
	return o.NextStep
}
func (o *CreatedState) GetStateName() string {
	return model.StateCreated
}



type AcceptedState struct {
	ID int
	NextStep *OrderState
}

func (o *AcceptedState) ProcessStep() error {
	return nil
}
func (o *AcceptedState) CeaseStep() error {
	return nil
}
func (o *AcceptedState) GetNextStep() *OrderState {
	return o.NextStep
}
func (o *AcceptedState) GetStateName() string {
	return model.StateAccepted
}



type StartedState struct {
	ID int
	NextStep *OrderState
}

func (o *StartedState) ProcessStep() error {
	return nil
}
func (o *StartedState) CeaseStep() error {
	return errors.New("")
}
func (o *StartedState) GetNextStep() *OrderState {
	return o.NextStep
}
func (o *StartedState) GetStateName() string {
	return model.StateStarted
}


type CookingState struct {
	ID int
	NextStep *OrderState
}

func (o *CookingState) ProcessStep() error {
	return nil
}
func (o *CookingState) CeaseStep() error {
	return errors.New("")
}

func (o *CookingState) GetNextStep() *OrderState {
	return o.NextStep
}
func (o *CookingState) GetStateName() string {
	return model.StateCooking
}



type PreparedState struct {
	ID int
	NextStep *OrderState
}

func (o *PreparedState) ProcessStep() error {
	return nil
}
func (o *PreparedState) CeaseStep() error {
	return errors.New("")
}

func (o *PreparedState) GetNextStep() *OrderState {
	return o.NextStep
}
func (o *PreparedState) GetStateName() string {
	return model.StatePrepared
}



type DeliveringState struct {
	ID int
	NextStep *OrderState
}

func (o *DeliveringState) ProcessStep() error {
	return nil
}
func (o *DeliveringState) CeaseStep() error {
	return errors.New("")
}
func (o *DeliveringState) GetNextStep() *OrderState {
	return o.NextStep
}
func (o *DeliveringState) GetStateName() string {
	return model.StateDelivering
}



type DeliveredState struct {
	ID int
	NextStep *OrderState
}

func (o *DeliveredState) ProcessStep() error {
	return nil
}
func (o *DeliveredState) CeaseStep() error {
	return errors.New("")
}

func (o *DeliveredState) GetNextStep() *OrderState {
	return o.NextStep
}
func (o *DeliveredState) GetStateName() string {
	return model.StateDelivered
}



type RequestedState struct {
	ID int
	NextStep *OrderState
}

func (o *RequestedState) ProcessStep() error {
	return nil
}
func (o *RequestedState) CeaseStep() error {
	return errors.New("")
}

func (o *RequestedState) GetNextStep() *OrderState {
	return o.NextStep
}
func (o *RequestedState) GetStateName() string {
	return model.StateRequested
}



type CollectedState struct {
	ID int
	NextStep *OrderState
}

func (o *CollectedState) ProcessStep() error {
	return nil
}
func (o *CollectedState) CeaseStep() error {
	return errors.New("")
}

func (o *CollectedState) GetNextStep() *OrderState {
	return o.NextStep
}
func (o *CollectedState) GetStateName() string {
	return model.StateCollected
}



type FinishedState struct {
	ID int
	NextStep *OrderState
}

func (o *FinishedState) ProcessStep() error {
	return nil
}
func (o *FinishedState) CeaseStep() error {
	return errors.New("")
}

func (o *FinishedState) GetNextStep() *OrderState {
	return o.NextStep
}
func (o *FinishedState) GetStateName() string {
	return model.StateFinished
}