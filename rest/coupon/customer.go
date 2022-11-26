package coupon

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"mr.jackpot-backend/model"
)

type CustomerCouponService interface {
	GetCouponList(c *gin.Context)
	GainCoupon(c *gin.Context)
}



func (h *CouponHandler) GetCouponList(c *gin.Context) {
	userid := c.Keys["userid"].(int)

	couponlist, err := h.c.GetCouponList(userid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, couponlist)
}

func (h *CouponHandler) GainCoupon(c *gin.Context) {
	var (
		coupon model.CouponCode
		userid = c.Keys["userid"].(int)
	)

	if err := c.ShouldBindJSON(&coupon); err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	couponinfo, err := h.c.GainCoupon(userid, coupon.Code)
	if err != nil {
		c.JSON(http.StatusUnauthorized, err.Error())
		return
	}

	c.JSON(http.StatusOK, couponinfo)
}

