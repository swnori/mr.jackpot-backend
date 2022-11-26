package model



type DinnerFormed struct {
	Name string
	Desc string
	Price string
	MenuFormed []MenuFormed
}

type MenuFormed struct {
	Name string
	Price string	
	Optionlist []OptionFormed
}

type OptionFormed struct {
	Name string
}

