package coupon

import "github.com/gin-gonic/gin"

type CustomerCouponService interface {
	GetCouponList(c *gin.Context)
	GainCoupon(c *gin.Context)
}



func (h *CouponHandler) GetCouponList(c *gin.Context) {
	couponList, err := h.c
}

func (h *CouponHandler) GainCoupon(c *gin.Context) {

}