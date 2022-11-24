package db

import (
	"context"

	"mr.jackpot-backend/model"
)

type VUILayer interface {
	ReadAllPreOrderList() ([]model.PreOrderTable, error)
	ReadAllProOrderList() ([]model.ProOrderTable, error)
}

type VUIDB struct {
	DBAccessor
}

func NewVUIDB() *VUIDB {
	db := &VUIDB{}
	db.q = NewAccessor()

	return db
}


func (db *VUIDB) ReadAllPreOrderList() ([]model.PreOrderTable, error) {

	ctx := context.Background()

	PreOrderList, err := db.q.ReadPreOrderChoice(ctx)
	if err != nil {
		return nil, err
	}


	preOrderNxtList, err := db.q.ReadPreOrderChoiceNxtSeq(ctx)
	if err != nil {
		return nil, err
	}

	adj := make(map[int][]int)
	for _, node := range preOrderNxtList {
		adj[int(node.SeqID)] = append(adj[int(node.SeqID)], int(node.NxtID))
	}

	preOrderList := make([]model.PreOrderTable, 0, len(PreOrderList))

	for _, PreOrder := range PreOrderList {
		Id := int(PreOrder.SeqID)

		preOrder := model.PreOrderTable{
			Id: Id,
			Message: PreOrder.Message,
			SeqList: adj[Id],
		}

		preOrderList = append(preOrderList, preOrder)
	}

	return preOrderList, nil

}



func (db *VUIDB) ReadAllProOrderList() ([]model.ProOrderTable, error) {
	ctx := context.Background()

	ProOrderList, err := db.q.ReadProOrderChoice(ctx)
	if err != nil {
		return nil, err
	}

	proOrderNxtList, err := db.q.ReadProOrderChoiceNxtSeq(ctx)
	if err != nil {
		return nil, err
	}

	adj := make(map[int][]int)
	for _, node := range proOrderNxtList {
		adj[int(node.SeqID)] = append(adj[int(node.SeqID)], int(node.NxtID))
	}

	proOrderList := make([]model.ProOrderTable, 0, len(ProOrderList))

	for _, ProOrder := range ProOrderList {
		Id := int(ProOrder.SeqID)

		proOrder := model.ProOrderTable{
			Id: Id,
			Message: ProOrder.Message,
			SeqList: adj[Id],
			Target: ProOrder.Target,
		}

		proOrderList = append(proOrderList, proOrder)
	}

	return proOrderList, nil
}