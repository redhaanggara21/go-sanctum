package controllers

import (
	"net/http"
	"service/go-sanctum/model"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type CreateOrderInput struct {
	Id_barang_order int    `json:"id_barang_order"`
	Barang_nama     string `json:"barang_nama"`
	User_order      int    `json:"user_order"`
	Quantity_order  int    `json:"quantity_order"`
	Total_order     int    `json:"total_order"`
}

type UpdateOrderInput struct {
	Id_barang_order int    `json:"id_barang_order"`
	Barang_nama     string `json:"barang_nama"`
	User_order      int    `json:"user_order"`
	Quantity_order  int    `json:"quantity_order"`
	Total_order     int    `json:"total_order"`
}

// GET /order
// Get all order
func GetOrder(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var order []model.Order
	db.Find(&order)

	c.JSON(http.StatusOK, gin.H{"data": order})
}

// POST /tasks
// Create new order
func CreateOrder(c *gin.Context) {
	// Validate input
	var input CreateOrderInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// date := "2006-01-02"
	// deadline, _ := time.Parse(date, input.Deadline)

	// Create order
	task := model.Order{
		Id_barang_order: input.Id_barang_order,
		Barang_nama:     input.Barang_nama,
		User_order:      input.User_order,
		Quantity_order:  input.Quantity_order,
		Total_order:     input.Total_order,
	}

	db := c.MustGet("db").(*gorm.DB)
	db.Create(&task)

	c.JSON(http.StatusOK, gin.H{"data": task})
}

// GET /order/:id
// Find a order
func GetOrderById(c *gin.Context) { // Get model if exist
	var order model.Order

	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("id = ?", c.Param("id")).First(&order).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": order})
}

// PATCH /order/:id
// Update a order
func UpdateOrder(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var order model.Order
	if err := db.Where("id = ?", c.Param("id")).First(&order).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data tidak ditemukan!!"})
		return
	}

	// Validate input
	var input UpdateOrderInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// date := "2006-01-02"
	// deadline, _ := time.Parse(date, input.Deadline)

	var updatedInput model.Order
	updatedInput.Id_barang_order = input.Id_barang_order
	updatedInput.Barang_nama = input.Barang_nama
	updatedInput.User_order = input.User_order
	updatedInput.Quantity_order = input.Quantity_order
	updatedInput.Total_order = input.Total_order

	db.Model(&order).Updates(updatedInput)

	c.JSON(http.StatusOK, gin.H{"data": order})
}

// DELETE /order/:id
// Delete a order
func DeleteOrder(c *gin.Context) {
	// Get model if exist
	db := c.MustGet("db").(*gorm.DB)
	var order model.Order
	if err := db.Where("id = ?", c.Param("id")).First(&order).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data tidak ditemukan!!"})
		return
	}

	db.Delete(&order)

	c.JSON(http.StatusOK, gin.H{"data": true})
}

// var response model.Response

// //GetAllPerson get all person data
// func GetAllOrder(w http.ResponseWriter, r *http.Request) {
// 	var order []model.Order
// 	database.Connector.Find(&order)
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusOK)
// 	json.NewEncoder(w).Encode(order)
// }

// //GetPersonByID returns person with specific ID
// func GetOrderByID(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	key := vars["id"]

// 	var order model.Order
// 	database.Connector.First(&order, key)
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(order)
// }

// //CreateOrder creates order
// func CreateOrder(w http.ResponseWriter, r *http.Request) {
// 	requestBody, _ := ioutil.ReadAll(r.Body)
// 	var order model.Order
// 	json.Unmarshal(requestBody, &order)

// 	database.Connector.Create(&order)
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusCreated)
// 	json.NewEncoder(w).Encode(order)
// }

// //UpdatePersonByID updates person with respective ID
// func UpdateOrderByID(w http.ResponseWriter, r *http.Request) {
// 	requestBody, _ := ioutil.ReadAll(r.Body)
// 	var order model.Order
// 	json.Unmarshal(requestBody, &order)
// 	database.Connector.Save(&order)

// 	response.Status = 1
// 	response.Message = "Success updated"

// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusOK)
// 	json.NewEncoder(w).Encode(order)
// }

// //DeletPersonByID delete's person with specific ID
// func DeletOrderByID(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	key := vars["id"]

// 	var order model.Order
// 	id, _ := strconv.ParseInt(key, 10, 64)
// 	database.Connector.Where("id = ?", id).Delete(&order)

// 	response.Status = 1
// 	response.Message = "Success Deleted"

// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusOK)
// 	json.NewEncoder(w).Encode(response)
// }
