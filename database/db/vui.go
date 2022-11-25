package db

import (
	"context"

	"mr.jackpot-backend/model"
	"mr.jackpot-backend/utility/util"
)

type VUILayer interface {
	ReadAllPreOrderList() ([]model.PreOrderTable, error)
	ReadAllProOrderList() ([]model.ProOrderTable, error)
	GetEntityInfoList() ([]model.EntityInfo, error)
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


func (db *VUIDB) GetEntityInfoList() ([]model.EntityInfo, error) {
	ctx := context.Background()

	entityInfoList := make([]model.EntityInfo, 0)

	all, err := db.q.GetAllEntityIdList(ctx)
	targetlist := util.IntAll(all)
	if err != nil {
		return nil, err
	}

	dinnerlist, err := db.q.GetDinnerEntity(ctx)
	if err != nil {
		return nil, err
	}

	menulist, err := db.q.GetMenuEntity(ctx)
	if err != nil {
		return nil, err
	}

	optionlist, err := db.q.GetOptionEntity(ctx)
	if err != nil {
		return nil, err
	}

	stylelist, err := db.q.GetStyleEntity(ctx)
	if err != nil {
		return nil, err
	}

	for _, entity := range targetlist {
		var flag bool = true

		for _, dinner := range dinnerlist {
			if dinner.TargetID == int32(entity) {
				entityInfoList = append(entityInfoList, model.EntityInfo{
					TargetId: int(dinner.TargetID),
					SpecId: int(dinner.DinnerID),
					EntityType: dinner.Typename,
				})

				flag = false
				break
			}
		}

		if flag == false {
			continue
		}

		for _, menu := range menulist {
			if menu.TargetID == int32(entity) {
				entityInfoList = append(entityInfoList, model.EntityInfo{
					TargetId: int(menu.TargetID),
					SpecId: int(menu.MenuID),
					EntityType: menu.Typename,
				})

				flag = false
				break
			}
		}

		if flag == false {
			continue
		}

		for _, option := range optionlist {
			if option.TargetID == int32(entity) {
				entityInfoList = append(entityInfoList, model.EntityInfo{
					TargetId: int(option.TargetID),
					SpecId: int(option.OptionID),
					EntityType: option.Typename,
				})

				flag = false
				break
			}
		}

		if flag == false {
			continue
		}

		for _, style := range stylelist {
			if style.TargetID == int32(entity) {
				entityInfoList = append(entityInfoList, model.EntityInfo{
					TargetId: int(style.TargetID),
					SpecId: int(style.StyleID),
					EntityType: style.Typename,
				})

				flag = false
				break
			}
		}

		if flag == false {
			continue
		}

		entityInfoList = append(entityInfoList, model.EntityInfo{
			TargetId: entity,
			EntityType: "message",
		})
	}

	return entityInfoList, nil
}