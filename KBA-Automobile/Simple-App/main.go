package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Car struct {
	CarId        string `json:"carId"`
	Make         string `json:"make"`
	Model        string `json:"model"`
	Color        string `json:"color"`
	Date         string `json:"dateOfManufacture"`
	Manufacturer string `json:"manufacturerName"`
}

func main() {
	router := gin.Default()

	router.Static("/public", "./public")
	router.LoadHTMLGlob("templates/*")
	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Manufacturer Dashboard",
		})

	})

	router.POST("/api/car", func(ctx *gin.Context) {

		var req Car
		if err := ctx.BindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
			return
		}
		fmt.Println("request", req)
		result := submitTxnFn(
			"manufacturer",
			"autochannel",
			"KBA-Automobile",
			"CarContract",
			"invoke",
			make(map[string][]byte),
			"CreateCar",
			req.CarId,
			req.Make,
			req.Model,
			req.Color,
			req.Manufacturer,
			req.Date,
		)
		ctx.JSON(http.StatusOK, gin.H{"message": "Created new car", "result": result})

	})

	router.GET("/api/car/:id", func(ctx *gin.Context) {
		carId := ctx.Param("id")

		result := submitTxnFn("manufacturer", "autochannel", "KBA-Automobile", "CarContract", "query", make(map[string][]byte), "ReadCar", carId)

		ctx.JSON(http.StatusOK, gin.H{"data": result})
	})

	router.Run(":3000")

}
