package main

import (
	"service/go-sanctum/database"
	"service/go-sanctum/model"
	"service/go-sanctum/routes"

	_ "github.com/jinzhu/gorm/dialects/mysql" //Required for MySQL dialect
)

func main() {
	// initDB()
	// log.Println("Starting the HTTP server on port 8090")
	// router := mux.NewRouter().StrictSlash(true)
	// initaliseHandlers(router)
	// log.Fatal(http.ListenAndServe(":8090", router))

	db := database.SetupDB()
	db.AutoMigrate(&model.Order{})

	r := routes.SetupRoutes(db)
	r.Run()

}

// func initaliseHandlers(router *mux.Router) {
// 	router.HandleFunc("/create", controllers.CreateOrder).Methods("POST")
// 	router.HandleFunc("/get", controllers.GetAllOrder).Methods("GET")
// 	router.HandleFunc("/get/{id}", controllers.GetOrderByID).Methods("GET")
// 	router.HandleFunc("/update/{id}", controllers.UpdateOrderByID).Methods("PUT")
// 	router.HandleFunc("/delete/{id}", controllers.DeletOrderByID).Methods("DELETE")
// }

// func initDB() {
// 	config :=
// 		database.Config{
// 			ServerName: "localhost:3306",
// 			User:       "root",
// 			Password:   "",
// 			DB:         "api-sanctum-order",
// 		}

// 	connectionString := database.GetConnectionString(config)
// 	err := database.Connect(connectionString)
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	database.Migrate(&model.Order{})
// }
