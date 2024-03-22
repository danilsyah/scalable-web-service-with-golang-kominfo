package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Order adalah model untuk pesanan
type Order struct {
	ID           uint           `json:"id" gorm:"primaryKey"`
	OrderedAt    time.Time      `gorm:"autoCreateTime" json:"orderedAt"`
	CustomerName string         `json:"customerName"`
	Items        []OrderItem    `json:"items" gorm:"foreignKey:OrderID"`
	CreatedAt    time.Time      `json:"createdAt"`
	UpdatedAt    time.Time      `json:"updatedAt"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"-"`
}

// OrderItem adalah model untuk item dalam pesanan
type OrderItem struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	OrderID     uint   `json:"-"`
	ItemCode    string `json:"itemCode"`
	Description string `json:"description"`
	Quantity    uint   `json:"quantity"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

var (
	db *gorm.DB
)

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/orders", createOrder)
	r.GET("/orders", getOrders)
	r.GET("/orders/:id", getOrderByID)
	r.PUT("/orders/:id", updateOrder)
	r.DELETE("/orders/:id", deleteOrder)

	return r
}

func createOrder(c *gin.Context) {
	var order Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	db.Create(&order)

	c.JSON(201, order)
}

func getOrders(c *gin.Context) {
	var orders []Order
	db.Preload("Items").Find(&orders)

	c.JSON(200, orders)
}

func getOrderByID(c *gin.Context) {
	id := c.Param("id")

	var order Order
	if err := db.Preload("Items").First(&order, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(200, order)
}

func updateOrder(c *gin.Context) {
	id := c.Param("id")

	var order Order
	if err := db.First(&order, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Record not found!"})
		return
	}

	var updatedOrder Order
	if err := c.ShouldBindJSON(&updatedOrder); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// Update order fields
	order.OrderedAt = updatedOrder.OrderedAt
	order.CustomerName = updatedOrder.CustomerName

	db.Save(&order)

	c.JSON(200, order)
}

func deleteOrder(c *gin.Context) {
	id := c.Param("id")

	var order Order
	if err := db.First(&order, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Record not found!"})
		return
	}

	db.Delete(&order)

	c.JSON(204, nil)
}

func connectDB() {
	dsn := "host=localhost user=postgres dbname=db_penjualan port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database!")
	}

	// Auto Migrate the schema
	db.Debug().AutoMigrate(&Order{}, &OrderItem{})
	fmt.Println("Connected to database")
}

func main() {
	connectDB()

	r := setupRouter()
	r.Run(":8080")
}
