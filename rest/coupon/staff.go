package coupon

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"mr.jackpot-backend/model"
)

type StaffCouponService interface {
	GetIssuedCouponList(c *gin.Context)
	IssueCoupon(c *gin.Context)
	DeleteCoupon(c *gin.Context)
}

func (h *CouponHandler) GetIssuedCouponList(c *gin.Context) {
	couponlist, err := h.c.GetIssuedCouponList()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, couponlist)
}

func (h *CouponHandler) IssueCoupon(c *gin.Context) {
	couponInfo := model.CouponInfo{}
	if err := c.ShouldBindJSON(&couponInfo); err != nil {
		c.JSON(http.StatusUnprocessableEntity, couponInfo)
		return
	}

	coupon, err := h.c.IssueCoupon(couponInfo)
	coupon.CreatedAt = time.Now().Format(model.TimeDayFormat)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, coupon)
}

func (h *CouponHandler) DeleteCoupon(c *gin.Context) {
	c.JSON(http.StatusBadGateway, "not yet recovered")
/*
	var coupon model.CouponInfo

	if err := c.ShouldBindJSON(&coupon); err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	if err := h.c.DeleteCoupon(coupon.ID); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, "")
*/
}
