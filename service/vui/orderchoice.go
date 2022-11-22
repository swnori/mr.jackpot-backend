package vui

import "mr.jackpot-backend/model"



func (v *VUIAccessor) HandleOrderChoice(request model.OrderChoiceRequest) (response model.OrderChoiceResponse) {
	message, seqStack := request.Message, request.SeqStack

	if len(seqStack) == 0 {
		questNode := v.Graph.GetQuestInfo(v.startNode)
		seqStack = append(seqStack, questNode.SeqList...)

		response = model.OrderChoiceResponse{
			Message: []string{questNode.Message},
			Decoded: "",
			ActionType: "message",
			SeqStack: seqStack,
		}
		return
	}

	seqStackTop, seqStack := seqStack[len(seqStack)-1], seqStack[:len(seqStack)-1]

	targetList, seqList := v.Graph.GetTargetListBySeqId(seqStackTop)
	targetId := v.GetTargetId(message, targetList);

	if targetId == -1 {
		seqStack = append(seqStack, seqStackTop)

		response = model.OrderChoiceResponse{
			Message: []string{"다시 한 번 선택해주세요"},
			Decoded: message,
			ActionType: "message",
			SeqStack: seqStack,
		}

	} else {
		proOrderChoiceId := seqList[targetId]
		ansNode := v.Graph.GetAnswerInfo(proOrderChoiceId)

		seqStack = append(seqStack, v.Graph.GetNxtPreList(proOrderChoiceId)...)

		if len(seqStack) == 0 {
			response = model.OrderChoiceResponse{
				Message: []string{ansNode.Message},
				Decoded: targetList[targetId],
				ActionType: ansNode.ActionType,
				SeqStack:  seqStack,
			}

			return
		}

		seqStackTop = seqStack[len(seqStack)-1]
		questNode := v.Graph.GetQuestInfo(seqStackTop)

		response = model.OrderChoiceResponse{
			Message: []string{ansNode.Message, questNode.Message},
			Decoded: targetList[targetId],
			ActionType: ansNode.ActionType,
			SeqStack:  seqStack,
		}
	}

	return
}