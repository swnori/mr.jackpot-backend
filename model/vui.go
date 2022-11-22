package model



type OrderChoiceRequest struct {
	SeqStack []int  `json:"seq_id"`
	Message string 	`json:"message"`
}

type OrderChoiceResponse struct {
	SeqStack   []int
	Message    []string
	Decoded    string
	ActionType string
}



type PreOrderTable struct {
	Id int
	SeqList []int
	Message string
}

type ProOrderTable struct {
	Id int
	SeqList    []int
	Message    string
	Target     string
	Action
}

type OrderChoiceTable struct {
	Id int
	SeqList    []int
	Message    string
	Target     string
	Action
}
