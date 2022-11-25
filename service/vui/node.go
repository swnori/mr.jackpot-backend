package vui

import "mr.jackpot-backend/model"



type NodeInterface interface {
	GetNxtSeqList() []int
	GetMessage() string
	GetNodeInfo() model.OrderChoiceTable
}



type VUIGraphNode struct {
	nxtSeqList []int
	message    string
}

func (n *VUIGraphNode) GetNxtSeqList() []int {
	return n.nxtSeqList
}

func (n *VUIGraphNode) GetMessage() string {
	return n.message
}





type VUIGraphPreNode struct {
	VUIGraphNode
}

func (n *VUIGraphPreNode) GetNodeInfo() (node model.OrderChoiceTable) {
	node = model.OrderChoiceTable{
		Message: n.GetMessage(),
	}
	return
}



type VUIGraphProNode struct {
	VUIGraphNode
	target     string
	EntityType string
	EntityID   int
}

func (n *VUIGraphProNode) GetNodeInfo() (node model.OrderChoiceTable) {
	node = model.OrderChoiceTable{
		Message: n.GetMessage(),
		Target: n.target,
		EntityId: n.EntityID,
		EntityType: n.EntityType,
	}
	return
}



func NewPreNode(entity model.PreOrderTable) *VUIGraphPreNode {
	node := VUIGraphPreNode{}
	node.message = entity.Message
	node.nxtSeqList = entity.SeqList

	return &node
}

func NewProNode(entity model.ProOrderTable) *VUIGraphProNode {
	node := VUIGraphProNode{}
	node.message = entity.Message
	node.target = entity.Target
	node.nxtSeqList = entity.SeqList
	node.EntityID = entity.EntityId
	node.EntityType = entity.EntityType

	return &node
}
