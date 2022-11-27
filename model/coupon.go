package model

type CouponCode struct {
	Code string `json:"code"`
}

type CouponInfo struct {
	ID        int    `json:"id"`
	Code      string `json:"code"`
	Title     string `json:"title"`
	Amount    int    `json:"amount"`
	Message   string `json:"message"`
	CreatedAt string `json:"createdAt"`
	ExpiresAt string `json:"expiresAt"`
}
