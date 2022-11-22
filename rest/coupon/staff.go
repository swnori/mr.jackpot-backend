package coupon

import "github.com/gin-gonic/gin"

type StaffCouponService interface {
	GetIssuedCouponList(c *gin.Context)
	IssueCoupon(c *gin.Context)
	DeleteCoupon(c *gin.Context)
}



func (h *CouponHandler) GetIssuedCouponList(c *gin.Context) {
	
}

func (h *CouponHandler) IssueCoupon(c *gin.Context) {

}

func (h *CouponHandler) DeleteCoupon(c *gin.Context) {

}