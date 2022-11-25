package vui

import (
	"fmt"

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
			Message: []string{questNode.Message},
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

		response = model.OrderChoiceResponse{
			Message: []string{"다시 한 번 선택해주세요"},
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