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
	ActionType string
	EntityID   int
}

func (n *VUIGraphProNode) GetNodeInfo() (node model.OrderChoiceTable) {
	node = model.OrderChoiceTable{
		Message: n.GetMessage(),
		Target: n.target,
	}
	/* ActionType 을 지웠는데 어떤 영향이 있는지 확인해볼 것 */
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
	node.ActionType = entity.ActionType
	node.nxtSeqList = entity.SeqList

	return &node
}
