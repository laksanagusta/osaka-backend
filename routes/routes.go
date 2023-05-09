package routes

import (
	"net/http"
	"strings"
	"tokokecilkita-go/auth"
	"tokokecilkita-go/handler"
	"tokokecilkita-go/helper"
	"tokokecilkita-go/order"
	"tokokecilkita-go/organization"
	"tokokecilkita-go/product"
	"tokokecilkita-go/user"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(db *gorm.DB) gin.Engine {
	//Version 1.0

	//Repository
	userRepository := user.NewRepository(db)
	organizationRepository := organization.NewRepository(db)
	productRepository := product.NewRepository(db)
	orderRepository := order.NewRepository(db)

	//Service
	userService := user.NewService(userRepository)
	organizationService := organization.NewService(organizationRepository)
	productService := product.NewService(productRepository)
	orderService := order.NewService(orderRepository)
	authService := auth.NewService()

	userJob := user.NewJob(userService)

	//Handler
	userHandler := handler.NewUserHandler(userService, authService, userJob)
	organizationHandler := handler.NewOrganizationHandler(organizationService)
	productHandler := handler.NewProductHandler(productService)
	orderHandler := handler.NewOrderHandler(orderService)

	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders: []string{"Content-Type,access-control-allow-origin, access-control-allow-headers, Authorization"},
	}))

	router.Static("/images", "./images")
	api := router.Group("/api/v1")

	//api users
	api.POST("/users", userHandler.RegisterUser)
	api.POST("/sessions", userHandler.Login)
	api.GET("/users/:id", userHandler.UserDetails)
	api.GET("/users", userHandler.UserFindAll)
	api.PUT("/users/:id", authMiddleware(authService, userService), userHandler.UpdateUser)
	api.DELETE("/users/:id", authMiddleware(authService, userService), userHandler.DeleteUser)

	//api organizations
	api.POST("/organizations", authMiddleware(authService, userService), organizationHandler.Save)
	api.GET("/organizations/:id", authMiddleware(authService, userService), organizationHandler.FindById)

	//api products
	api.POST("/products", authMiddleware(authService, userService), productHandler.Save)
	api.POST("/products/image", authMiddleware(authService, userService), productHandler.SaveImage)
	api.PUT("/products/:id", authMiddleware(authService, userService), productHandler.UpdateProduct)
	api.DELETE("/products/:id", authMiddleware(authService, userService), productHandler.Delete)
	api.GET("/products/:id", authMiddleware(authService, userService), productHandler.FindById)
	api.GET("/products/code/:code", authMiddleware(authService, userService), productHandler.FindByCode)
	api.GET("/products", authMiddleware(authService, userService), productHandler.FindAll)

	//api order & basket
	api.POST("/orders", authMiddleware(authService, userService), orderHandler.Save)
	api.GET("/orders", orderHandler.FindAll)
	api.PUT("/orders/:id", authMiddleware(authService, userService), orderHandler.UpdateOrder)
	api.GET("/orders/:id", authMiddleware(authService, userService), orderHandler.BasketFindByOrderId)
	api.GET("/basket/:id", authMiddleware(authService, userService), orderHandler.BasketFindByOrderId)
	api.POST("/order-basket/place", authMiddleware(authService, userService), orderHandler.SaveBasket)

	return *router
}

func authMiddleware(authService auth.Service, userService user.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if !strings.Contains(authHeader, "Bearer") {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		tokenString := ""
		arrayToken := strings.Split(authHeader, " ")
		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}

		token, err := authService.ValidateToken(tokenString)
		if err != nil {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		claim, ok := token.Claims.(jwt.MapClaims)

		if !ok || !token.Valid {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		userID := int(claim["user_id"].(float64))

		user, err := userService.UserDetails(userID)
		if err != nil {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		c.Set("currentUser", user)

	}
}
