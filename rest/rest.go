package rest

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"mr.jackpot-backend/rest/authority"
	"mr.jackpot-backend/rest/coupon"
	"mr.jackpot-backend/rest/manager"
	"mr.jackpot-backend/rest/order"
	"mr.jackpot-backend/rest/orderinfo"
	"mr.jackpot-backend/rest/stock"
	"mr.jackpot-backend/rest/task"
)

func RunAPI(address string) error {

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{
			"http://localhost:3000",
			"https://mr-jackpot.run.goorm.io",
			"https://mr-jackpot.run.goorm.io:5173",
			"http://mr-jackpot.run.goorm.io",
			"http://mr-jackpot.run.goorm.io:5173",
		},
		AllowMethods:     []string{"PUT", "PATCH", "POST", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Accept", "Content-Type", "Content-Length", "Accept-Encoding","Authorization" , "Authorization,X-CSRF-Token"},
		AllowCredentials: true,
	}))

	var mh authority.AuthMiddlewareService = authority.NewAuthMiddlewareHandler()
	r.Use(mh.SetAuthority)

	Auth := r.Group("/auth")
	{
		Visitor := Auth.Group("/visitor")
		{
			var h authority.VisitorAuthService = authority.NewVisitorAuthHandler()

			Visitor.POST("/signin", h.Signin)
			Visitor.POST("/signout", h.Signout)
		}

		Customer := Auth.Group("/customer")
		{
			var h authority.CustomerAuthService = authority.NewCustomerAuthHandler()

			Customer.POST("/signin", h.Signin)
			Customer.POST("/signout", h.Signout)
			Customer.POST("/register", h.Register)
			Customer.POST("/unregister", h.Unregister)
		}

		Staff := Auth.Group("/staff")
		{
			var h authority.StaffAuthService = authority.NewStaffAuthHandler()

			Staff.POST("/signin", h.Signin)
			Staff.POST("/signout", h.Signout)
		}
	}


	Public := r.Group("/public")
	Public.Use(mh.CheckCient)
	{
		Orderinfo := Public.Group("/orderinfo")
		{
			var h orderinfo.OrderInfoService = orderinfo.NewOrderInfoHandler()

			Orderinfo.GET("/orderboard", h.GetOrderBoard)
			Orderinfo.GET("/statelist", h.GetStateList)
		}
	}

	Customer := r.Group("/customer")
	Customer.Use(mh.CheckCustomer)
	{
		Orderinfo := Customer.Group("/orderinfo")
		{
			var h orderinfo.OrderInfoService = orderinfo.NewOrderInfoHandler()
			Orderinfo.POST("/vuistep", h.HandleVUIStep)
		}
		{
			var h orderinfo.OrderHistoryHandler = *orderinfo.NewOrderHistoryHandler()
			Orderinfo.GET("/history", mh.CheckCustomerOnly, h.GetOrderHistory)
		}

		Order := Customer.Group("/order")
		{
			var h order.CustomerOrderService = order.NewOrderHandler()

			Order.GET("/info", h.GetOrderInfo)
			Order.POST("/create", h.CreateOrder)
			Order.POST("/cancle", h.CancleOrder)
			Order.POST("/requestcollecting", h.RequestCollecting)
		}

		Coupon := Customer.Group("/coupon")
		{
			var h coupon.CouponService = coupon.NewCouponHandler()

			Coupon.GET("/list", mh.CheckCustomerOnly, h.GetCouponList)
			Coupon.POST("/gain", mh.CheckCustomerOnly, h.GainCoupon)
		}

		Personal := Customer.Group("/personalinfo")
		{
			var h manager.ManagerService = manager.NewManagerHandler()

			Personal.GET("",  mh.CheckCustomerOnly, h.GetPersonalInfo)
			Personal.POST("", mh.CheckCustomerOnly, h.UpdatePersonalInfo)
		}
	}

	Staff := r.Group("/staff")
	Staff.Use(mh.CheckStaff)
	{
		Stock := Staff.Group("/stock")
		{
			var h stock.StockService = stock.NewStockHandler()

			Stock.GET("/itemlist", h.GetAllStockList)
			Stock.POST("/update", h.UpdateStockItem)
			Stock.POST("/add", h.AddStockItem)
			Stock.POST("/delete", h.DeleteStockItem)
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
	Ceo.Use(mh.CheckCEO)
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
