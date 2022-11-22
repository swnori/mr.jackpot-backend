package vui

import (
	"mr.jackpot-backend/database/db"
	"mr.jackpot-backend/model"
)



type VUIGraph struct {
	db db.VUILayer

	preNodes map[int]NodeInterface
	proNodes map[int]NodeInterface
}

func NewVUIGraph() *VUIGraph {
	graph := VUIGraph{
		db: db.NewVUIDB(),
		preNodes: make(map[int]NodeInterface),
		proNodes: make(map[int]NodeInterface),
	}
	return &graph
}



func (g *VUIGraph) GetTargetListBySeqId(seq_top int) (targetList []string, nxtIdList []int) {
	nxtIdList = g.GetNxtProList(seq_top)

	for _, nxtId := range nxtIdList {
		node := g.proNodes[nxtId].GetNodeInfo()
		targetList = append(targetList, node.Target)
	}
	return
}

func (g *VUIGraph) GetNxtProList(seqId int) []int {
	return g.preNodes[seqId].GetNxtSeqList()
}

func (g *VUIGraph) GetNxtPreList(seqId int) []int {
	return g.proNodes[seqId].GetNxtSeqList()
}

func (g *VUIGraph) GetAnswerInfo(seqId int) model.OrderChoiceTable {
	return g.proNodes[seqId].GetNodeInfo()
}

func (g *VUIGraph) GetQuestInfo(seqId int) model.OrderChoiceTable {
	return g.preNodes[seqId].GetNodeInfo()
}



func (g *VUIGraph) Initialize() error {

	preOrderList, err := g.db.ReadAllPreOrderList()
	if err != nil {
		return err
	}

	for _, entity := range preOrderList {
		g.preNodes[entity.Id] = NewPreNode(entity)
	}

	proOrderList, err := g.db.ReadAllProOrderList()
	if err != nil {
		return err
	}

	for _, entity := range proOrderList {
		g.proNodes[entity.Id] = NewProNode(entity)
	}
	return nil
}