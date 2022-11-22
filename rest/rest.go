package rest

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"mr.jackpot-backend/rest/coupon"
	"mr.jackpot-backend/rest/inventory"
	"mr.jackpot-backend/rest/manager"
	"mr.jackpot-backend/rest/order"
	"mr.jackpot-backend/rest/orderinfo"
	"mr.jackpot-backend/rest/task"
)

func RunAPI(address string) error {

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "https://mr-jackpot.run.goorm.io/"},
		AllowMethods:     []string{"PUT", "PATCH", "POST", "DELETE", "OPTIONS"},
		AllowCredentials: true,
	}))

	r.Use()

	Auth := r.Group("/auth")
	{
		Customer := Auth.Group("/customer")
		{
			Customer.POST("/signin")
			Customer.POST("/signout")
			Customer.POST("/Register")
			Customer.POST("/Unegister")
		}
		Staff := Auth.Group("/staff")
		{
			Staff.POST("/login")
			Staff.POST("/logout")
		}
	}

	Customer := r.Group("/customer")
	{
		Order := Customer.Group("/orderinfo")
		{
			var h orderinfo.OrderInfoService = orderinfo.NewOrderInfoHandler()

			Order.GET("/dinnerboard", h.GetDinnerBoard)
			Order.GET("/menuboard", h.GetMenuBoard)
			Order.POST("/vuistep", h.HandleVUIStep)
			Order.GET("/history", h.GetOrderHistory)
		}

		Orderstep := Customer.Group("/order")
		{
			var h order.CustomerOrderService = order.NewOrderHandler()

			Orderstep.POST("/create", h.CreateOrder)
			Orderstep.POST("/cancle", h.CancleOrder)
			Orderstep.POST("/requestcollecting", h.RequestCollecting)
		}

		Coupon := Customer.Group("/coupon")
		{
			var h coupon.CouponService = coupon.NewCouponHandler()

			Coupon.GET("/list", h.GetCouponList)
			Coupon.POST("/asdf", h.GainCoupon)
		}
	}

	Staff := r.Group("/staff")
	{
		Stock := Staff.Group("/inventory")
		{
			var h inventory.InventoryService = inventory.NewInventoryHandler()

			Stock.GET("/itemlist", h.GetAllInventoryList)
			Stock.POST("/update", h.UpdateInventoryItem)
			Stock.POST("/add", h.AddInventoryItem)
			Stock.POST("/delete", h.DeleteInventoryItem)
		}

		Order := Staff.Group("/order")
		{
			var h order.StaffOrderService = order.NewOrderHandler()

			Order.GET("/list", h.GetAllOrderList)
		}

		Task := Staff.Group("/task")
		{
			var h task.TaskService = task.NewTaskHandler()

			Task.GET("/list", h.GetAllTaskList)
			Task.POST("/nextstatus")
			Task.POST("/previousstatus")
		}
	}

	Ceo := r.Group("/ceo")
	{
		Order := Ceo.Group("/order")
		{
			var h order.StaffOrderService = order.NewOrderHandler()

			Order.POST("/accept", h.AcceptOrder)
			Order.POST("/reject", h.RejectOrder)
		}

		Staff := Ceo.Group("/staff")
		{
			var h manager.StaffManagerService = manager.NewManagerHandler()

			Staff.GET("/list", h.GetStaffList)
			Staff.POST("/register", h.RegisterStaff)
			Staff.POST("/update", h.UpdateStaffInfo)
		}

		Customer := Ceo.Group("/customer")
		{
			var h manager.CustomerManagerService = manager.NewManagerHandler()

			Customer.GET("/list", h.GetCustomerList)
		}

		Coupon := Ceo.Group("/coupon")
		{
			var h coupon.StaffCouponService = coupon.NewCouponHandler()

			Coupon.GET("/list", h.GetIssuedCouponList)
			Coupon.POST("/issue", h.IssueCoupon)
			Coupon.POST("/delete", h.DeleteCoupon)
		}
	}

	return r.Run(address)
}
