package vui

import "mr.jackpot-backend/model"

type VUIController interface {
	HandleOrderChoice(request model.OrderChoiceRequest) (response model.OrderChoiceResponse, err error)
}

type VUIAccessor struct {
	Graph *VUIGraph
	threshold float64
	startNode int
}

var VUI = &VUIAccessor{
	startNode: 1,
	threshold: 0.8,
}

func NewVUIAccessor() *VUIAccessor {
	return VUI
}

func Initialize() error {
	VUI.Graph = NewVUIGraph()

	return VUI.Graph.Initialize()
}