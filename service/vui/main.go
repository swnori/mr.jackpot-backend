package vui

import "mr.jackpot-backend/model"

type VUIController interface {
	HandleOrderChoice(request model.OrderChoiceRequest) (response model.OrderChoiceResponse)
}

type VUIAccessor struct {
	Graph VUIGraph
	threshold float64
	startNode int
}

var VUI *VUIAccessor

func NewVUIAccessor() *VUIAccessor {
	return VUI
}

func Initialize() {
	VUI.Graph = *NewVUIGraph()

	if err := VUI.Graph.Initialize(); err != nil {
		panic(err)
	}
}