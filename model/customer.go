package model



type User struct {
	UserID   string `json:"userID"`
	Password string `json:"password"`
}



type PersonalInfo struct {
	ID      int    `json:"-"`
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
}



type UserInfo struct {
	ID     int
	UserID string
}

type CustomerRegister struct {
	ID       int
	UserID   string `json:"userID"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Address  string `json:"address"`
}

type CustomerInfo struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	UserID  string `json:"-"`
	Phone   string `json:"phone"`
	Address string `json:"address"`

	CreatedAt  string `json:"registeredAt"`
	Rating     int    `json:"rating"`
	OrderCount int    `json:"orders"`
	Paid       int    `json:"paid"`
}