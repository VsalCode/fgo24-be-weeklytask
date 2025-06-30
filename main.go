package main

import (
	"be-weeklytask/routers"
	"be-weeklytask/utils"
	"log"
	"github.com/gin-gonic/gin"
) 

// @title           Koda E-Wallet Backend
// @version         1.0
// @description     This is the backend service for Koda E-Wallet, providing user authentication, profile management, wallet transactions, and more.
func main() {
	db, err := utils.DBConnect()

	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close()

	r := gin.Default()
	routers.CombineRouters(r)

	log.Println("Server starting on port 8080...")
	r.Run(":8080")
}
