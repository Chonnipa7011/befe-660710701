package main

import (
    "fmt"
    "net/http"
    "github.com/gin-gonic/gin"
)

// Food struct
type Food struct {
    ID       string  `json:"id"`
    FoodName string  `json:"name"`
    Category string  `json:"category"`
    Price    float64 `json:"price"`
    Rating   float64 `json:"rating"`
}

// In-memory database (ในโปรเจคจริงใช้ database)
var foods = []Food{
    {ID: "1", FoodName: "Pad Thai", Category: "Noodles", Price: 50.0, Rating: 4.5},
    {ID: "2", FoodName: "Green Curry", Category: "Curry", Price: 70.0, Rating: 4.7},
    {ID: "3", FoodName: "Som Tum", Category: "Salad", Price: 40.0, Rating: 4.3},
}

// GET /api/v1/foods
func getFoods(c *gin.Context) {
    // ตรวจสอบว่ามี query parameter "category" หรือไม่
    categoryQuery := c.Query("category")

    if categoryQuery != "" {
        // Filter foods by category
        var filtered []Food
        for _, food := range foods {
            if food.Category == categoryQuery {
                filtered = append(filtered, food)
            }
        }
        c.JSON(http.StatusOK, filtered)
        return
    }

    // Return all foods
    c.JSON(http.StatusOK, foods)
}

func main(){
	r:=gin.Default()

	r.GET("/health", func(c*gin.Context){
		c.JSON(200,  gin.H{"message" : "healthy"})
	})

	r.Run(":8080")
}