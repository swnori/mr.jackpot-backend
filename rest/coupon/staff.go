package coupon

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"mr.jackpot-backend/model"
)

type StaffCouponService interface {
	GetIssuedCouponList(c *gin.Context)
	IssueCoupon(c *gin.Context)
	DeleteCoupon(c *gin.Context)
}



func (h *CouponHandler) GetIssuedCouponList(c *gin.Context) {
	
}

func (h *CouponHandler) IssueCoupon(c *gin.Context) {
	couponInfo := model.CouponCreateRequest{}
	if err := c.ShouldBindJSON(&couponInfo); err != nil {
		c.JSON(http.StatusUnprocessableEntity, couponInfo)
		return
	}

	coupon, err := h.c.IssueCoupon(couponInfo);
	if err != nil {
		c.JSON(http.StatusInternalServerError, couponInfo)
		return
	}

	c.JSON(http.StatusOK, coupon)
}

func (h *CouponHandler) DeleteCoupon(c *gin.Context) {

}