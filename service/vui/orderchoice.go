package vui

import (
	"fmt"
	"sort"

	"mr.jackpot-backend/model"
)



func (v *VUIAccessor) HandleOrderChoice(request model.OrderChoiceRequest) (response model.OrderChoiceResponse, err error) {
	message, seqStack := request.Message, request.SeqStack

	if len(seqStack) == 0 {
		questNode, e := v.Graph.GetQuestInfo(v.startNode)
		if e != nil {
			err = e
			return
		}
		seqStack = append(seqStack, questNode.SeqList...)

		response = model.OrderChoiceResponse{
			Message: []string{"미스터 대박 음성 인식 서비스입니다.", questNode.Message},
			Decoded: "",
			EntityId: 0,
			EntityType: "message",
			SeqStack: seqStack,
		}
		return
	}

	seqStackTop, seqStack := seqStack[len(seqStack)-1], seqStack[:len(seqStack)-1]

	targetList, seqList, err := v.Graph.GetTargetListBySeqId(seqStackTop)
	if err != nil {
		return 
	}

	for _, t := range targetList {
		fmt.Println(t)
	}
	targetId := v.GetTargetId(message, targetList);

	if targetId == -1 {
		seqStack = append(seqStack, seqStackTop)

		questNode, e := v.Graph.GetQuestInfo(seqStackTop)
		if err != nil {
			err = e
			return
		}

		response = model.OrderChoiceResponse{
			Message: []string{"다시 한 번 선택해주세요", questNode.Message},
			Decoded: message,
			EntityId: 0,
			EntityType: "message",
			SeqStack: seqStack,
		}

	} else {
		proOrderChoiceId := seqList[targetId]

		ansNode, e := v.Graph.GetAnswerInfo(proOrderChoiceId)
		if e != nil {
			err = e
			return
		}

		prelist, e := v.Graph.GetNxtPreList(proOrderChoiceId)
		if e != nil {
			err = e
			return
		}

		sort.Sort(sort.Reverse(sort.IntSlice(prelist)))

		seqStack = append(seqStack, prelist...)

		if len(seqStack) == 0 {
			response = model.OrderChoiceResponse{
				Message: []string{ansNode.Message},
				Decoded: targetList[targetId],
				EntityId: ansNode.EntityId,
				EntityType: ansNode.EntityType,
				SeqStack:  seqStack,
			}

			return
		}

		seqStackTop = seqStack[len(seqStack)-1]
		questNode, e := v.Graph.GetQuestInfo(seqStackTop)
		if e != nil {
			err = e
			return
		}

		response = model.OrderChoiceResponse{
			Message: []string{ansNode.Message, questNode.Message},
			Decoded: targetList[targetId],
			EntityId: ansNode.EntityId,
			EntityType: ansNode.EntityType,
			SeqStack:  seqStack,
		}
	}

	return
}