package routes

import (
	"service/go-sanctum/controllers"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func SetupRoutes(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
	})
	r.GET("/get/:id", controllers.GetOrderById)
	r.POST("/create", controllers.CreateOrder)
	r.GET("/get", controllers.GetOrder)
	r.PUT("/update/:id", controllers.UpdateOrder)
	r.DELETE("delete/:id", controllers.DeleteOrder)
	return r
}

// router.HandleFunc("/create", controllers.CreateOrder).Methods("POST")
// router.HandleFunc("/get", controllers.GetAllOrder).Methods("GET")
// router.HandleFunc("/get/{id}", controllers.GetOrderByID).Methods("GET")
// router.HandleFunc("/update/{id}", controllers.UpdateOrderByID).Methods("PUT")
// router.HandleFunc("/delete/{id}", controllers.DeletOrderByID).Methods("DELETE")
