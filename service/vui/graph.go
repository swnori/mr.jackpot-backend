package vui

import (
	"errors"
	"fmt"

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
		preNodes: make(map[int]NodeInterface),
		proNodes: make(map[int]NodeInterface),
	}
	graph.db = db.NewVUIDB()

	return &graph
}



func (g *VUIGraph) GetTargetListBySeqId(seqTop int) (targetList []string, nxtIdList []int, err error) {
	nxtIdList, err = g.GetNxtProList(seqTop)
	if err != nil {
		return
	}

	for _, nxtId := range nxtIdList {
		node := g.proNodes[nxtId].GetNodeInfo()
		targetList = append(targetList, node.Target)
	}
	return
}

func (g *VUIGraph) GetNxtProList(seqId int) ([]int, error) {
	node, exist := g.preNodes[seqId]
	if !exist {
		return nil, errors.New(fmt.Sprintf("seqId %d Not Exist on %s", seqId, "GetHxtProList"))
	}
	return node.GetNxtSeqList(), nil
}

func (g *VUIGraph) GetNxtPreList(seqId int) ([]int, error) {
	node, exist := g.preNodes[seqId]
	if !exist {
		return nil, errors.New(fmt.Sprintf("seqId %d Not Exist on %s", seqId, "GetHxtPreList"))
	}
	return node.GetNxtSeqList(), nil
}

func (g *VUIGraph) GetAnswerInfo(seqId int) (model.OrderChoiceTable, error) {
	node, exist := g.preNodes[seqId]
	if !exist {
		return model.OrderChoiceTable{}, errors.New(fmt.Sprintf("seqId %d Not Exist on %s", seqId, "GetAnswerInfo"))
	}
	return node.GetNodeInfo(),nil
}

func (g *VUIGraph) GetQuestInfo(seqId int) (model.OrderChoiceTable, error) {
	node, exist := g.preNodes[seqId]
	if !exist {
		return model.OrderChoiceTable{}, errors.New(fmt.Sprintf("seqId %d Not Exist on %s", seqId, "GetQuestInfo"))
	}
	return node.GetNodeInfo(), nil
}



func (g *VUIGraph) Initialize() error {

	preOrderList, err := g.db.ReadAllPreOrderList()
	if err != nil {
		return err
	}

	for _, entity := range preOrderList {
		g.preNodes[entity.Id] = NewPreNode(entity)
	}

	entityInfoMap := make(map[int]*model.EntityInfo)


	entityInfoList, err := g.db.GetEntityInfoList()
	if err != nil {
		return err
	}

	for _, entityInfo := range entityInfoList {
		entityInfoMap[entityInfo.TargetId] = &entityInfo
	}

	proOrderList, err := g.db.ReadAllProOrderList()
	if err != nil {
		return err
	}


	for _, entity := range proOrderList {
		newentity := entityInfoMap[entity.Id]
		entity.EntityId = newentity.SpecId
		entity.EntityType = newentity.EntityType
		g.proNodes[entity.Id] = NewProNode(entity)
	}

	fmt.Println("prenode: ", len(g.preNodes))
	fmt.Println("pronode: ", len(g.proNodes))

	return nil
}