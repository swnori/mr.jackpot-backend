package model



type OrderChoiceRequest struct {
	SeqStack []int  `json:"seqStack"`
	Message string 	`json:"message"`
}

type OrderChoiceResponse struct {
	SeqStack   []int
	Message    []string
	Decoded    string
	EntityId   int
	EntityType string
}



type PreOrderTable struct {
	Id      int
	SeqList []int
	Message string
}

type ProOrderTable struct {
	Id      int
	SeqList []int
	Message string

	Target  string
	EntityId   int
	EntityType string
}

type OrderChoiceTable struct {
	Id      int
	SeqList []int
	Message string

	Target  string
	EntityId   int
	EntityType string
}

type EntityByType struct {
	Dinner  []int
	Menu    []int
	Style   []int
	Option  []int
	Message []int
}

type EntityInfo struct {
	TargetId int
	SpecId   int
	EntityType string
}