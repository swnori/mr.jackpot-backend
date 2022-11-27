package model




 
type Staff struct {
	Code    string
	Name    string
	Role    string
	Part    string
}

type StaffResponse struct {
	ID        int    `json:"id"`
	Code      string `json:"code"`
	Name      string `json:"name"`
	Role      string `json:"role"`
	Part      string `json:"part"`
	Score     int    `json:"score"`
	CreatedAt string `json:"createat"`
}

type StaffInfo struct {
	ID        int   
	Status    bool
	Code      string
	Name      string
	Role      string
	Part      string
	Score     int   
	CreatedAt string
}

type StaffRegister struct {
	Code string
	Name string
	RoleID int
}