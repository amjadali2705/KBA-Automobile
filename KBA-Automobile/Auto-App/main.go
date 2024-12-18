package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

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

type Order struct {
	OrderId string `json:"orderId"`
	Make    string `json:"make"`
	Model   string `json:"model"`
	Color   string `json:"color"`
	Dealer  string `json:"dealerName"`
}

type CarData struct {
	AssetType         string `json:"AssetType"`
	CarId             string `json:"CarId"`
	Color             string `json:"Color"`
	DateOfManufacture string `json:"DateOfManufacture"`
	OwnedBy           string `json:"OwnedBy"`
	Make              string `json:"Make"`
	Model             string `json:"Model"`
	Status            string `json:"Status"`
}

type OrderData struct {
	AssetType  string `json:"assetType"`
	Color      string `json:"color"`
	DealerName string `json:"dealerName"`
	Make       string `json:"make"`
	Model      string `json:"model"`
	OrderID    string `json:"orderID"`
}

type Match struct {
	OrderId string `json:"orderId"`
	CarId   string `json:"carId"`
}

type CarHistory struct {
	Record    *CarData `json:"record"`
	TxId      string   `json:"txId"`
	Timestamp string   `json:"timestamp"`
	IsDelete  bool     `json:"isDelete"`
}

type Register struct {
	CarId     string `json:"carId"`
	CarOwner  string `json:"carOwner"`
	RegNumber string `json:"regNumber"`
}

func main() {
	router := gin.Default()

	var wg sync.WaitGroup
	wg.Add(1)
	go ChaincodeEventListener("manufacturer", "autochannel", "KBA-Automobile", &wg)

	router.Static("/public", "./public")
	router.LoadHTMLGlob("templates/*")

	router.GET("/", func(ctx *gin.Context) {
		result := submitTxnFn("manufacturer", "autochannel", "KBA-Automobile", "CarContract", "query", make(map[string][]byte), "GetAllCars")

		var cars []CarData

		if len(result) > 0 {
			// Unmarshal the JSON array string into the cars slice
			if err := json.Unmarshal([]byte(result), &cars); err != nil {
				fmt.Println("Error:", err)
				return
			}
		}

		ctx.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Auto App", "carList": cars,
		})
	})

	router.GET("/manufacturer", func(ctx *gin.Context) {
		result := submitTxnFn("manufacturer", "autochannel", "KBA-Automobile", "CarContract", "query", make(map[string][]byte), "GetAllCars")

		var cars []CarData

		if len(result) > 0 {
			// Unmarshal the JSON array string into the cars slice
			if err := json.Unmarshal([]byte(result), &cars); err != nil {
				fmt.Println("Error:", err)
				return
			}
		}

		ctx.HTML(http.StatusOK, "manufacturer.html", gin.H{
			"title": "Manufacturer Dashboard", "carList": cars,
		})
	})

	router.POST("/api/car", func(ctx *gin.Context) {
		var req Car
		if err := ctx.BindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
			return
		}

		fmt.Printf("car response %s", req)
		submitTxnFn("manufacturer", "autochannel", "KBA-Automobile", "CarContract", "invoke", make(map[string][]byte), "CreateCar", req.CarId, req.Make, req.Model, req.Color, req.Manufacturer, req.Date)

		ctx.JSON(http.StatusOK, req)
	})

	router.GET("/api/car/:id", func(ctx *gin.Context) {
		carId := ctx.Param("id")

		result := submitTxnFn("manufacturer", "autochannel", "KBA-Automobile", "CarContract", "query", make(map[string][]byte), "ReadCar", carId)

		ctx.JSON(http.StatusOK, gin.H{"data": result})
	})

	router.GET("/api/order/match-car", func(ctx *gin.Context) {
		carID := ctx.Query("carId")
		result := submitTxnFn("manufacturer", "autochannel", "KBA-Automobile", "CarContract", "query", make(map[string][]byte), "GetMatchingOrders", carID)

		// fmt.Printf("result %s", result)

		var orders []OrderData

		if len(result) > 0 {
			// Unmarshal the JSON array string into the orders slice
			if err := json.Unmarshal([]byte(result), &orders); err != nil {
				fmt.Println("Error:", err)
				return
			}
		}

		ctx.HTML(http.StatusOK, "matchOrder.html", gin.H{
			"title": "Matching Orders", "orderList": orders, "carId": carID,
		})
	})

	router.POST("/api/car/match-order", func(ctx *gin.Context) {
		var req Match
		if err := ctx.BindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
			return
		}

		fmt.Printf("match  %s", req)
		submitTxnFn("manufacturer", "autochannel", "KBA-Automobile", "CarContract", "invoke", make(map[string][]byte), "MatchOrder", req.CarId, req.OrderId)

		ctx.JSON(http.StatusOK, req)
	})

	router.GET("/api/event", func(ctx *gin.Context) {
		result := getEvents()
		fmt.Println("result:", result)

		ctx.JSON(http.StatusOK, gin.H{"carEvent": result})

	})

	router.GET("/dealer", func(ctx *gin.Context) {

		ctx.HTML(http.StatusOK, "dealer.html", gin.H{
			"title": "Dealer Dashboard",
		})
	})

	//Get all orders
	router.GET("/api/order/all", func(ctx *gin.Context) {

		result := submitTxnFn("dealer", "autochannel", "KBA-Automobile", "OrderContract", "query", make(map[string][]byte), "GetAllOrders")

		var orders []OrderData

		if len(result) > 0 {
			// Unmarshal the JSON array string into the orders slice
			if err := json.Unmarshal([]byte(result), &orders); err != nil {
				fmt.Println("Error:", err)
				return
			}
		}

		ctx.HTML(http.StatusOK, "orders.html", gin.H{
			"title": "All Orders", "orderList": orders,
		})
	})

	router.POST("/api/order", func(ctx *gin.Context) {
		var req Order
		if err := ctx.BindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
			return
		}

		fmt.Printf("order  %s", req)

		privateData := map[string][]byte{
			"make":       []byte(req.Make),
			"model":      []byte(req.Model),
			"color":      []byte(req.Color),
			"dealerName": []byte(req.Dealer),
		}

		submitTxnFn("dealer", "autochannel", "KBA-Automobile", "OrderContract", "private", privateData, "CreateOrder", req.OrderId)

		ctx.JSON(http.StatusOK, req)
	})

	router.GET("/api/order/:id", func(ctx *gin.Context) {
		orderId := ctx.Param("id")

		result := submitTxnFn("dealer", "autochannel", "KBA-Automobile", "OrderContract", "query", make(map[string][]byte), "ReadOrder", orderId)

		ctx.JSON(http.StatusOK, gin.H{"data": result})
	})

	router.GET("/mvd", func(ctx *gin.Context) {
		result := submitTxnFn("mvd", "autochannel", "KBA-Automobile", "CarContract", "query", make(map[string][]byte), "GetAllCars")

		var cars []CarData

		if len(result) > 0 {
			// Unmarshal the JSON array string into the cars slice
			if err := json.Unmarshal([]byte(result), &cars); err != nil {
				fmt.Println("Error:", err)
				return
			}
		}

		ctx.HTML(http.StatusOK, "mvd.html", gin.H{
			"title": "MVD Dashboard", "carList": cars,
		})
	})

	router.GET("/api/car/history", func(ctx *gin.Context) {
		carID := ctx.Query("carId")
		result := submitTxnFn("mvd", "autochannel", "KBA-Automobile", "CarContract", "query", make(map[string][]byte), "GetCarHistory", carID)

		// fmt.Printf("result %s", result)

		var cars []CarHistory

		if len(result) > 0 {
			// Unmarshal the JSON array string into the orders slice
			if err := json.Unmarshal([]byte(result), &cars); err != nil {
				fmt.Println("Error:", err)
				return
			}
		}

		ctx.HTML(http.StatusOK, "history.html", gin.H{
			"title": "Car History", "itemList": cars,
		})
	})

	router.POST("/api/car/register", func(ctx *gin.Context) {
		var req Register
		if err := ctx.BindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
			return
		}

		fmt.Printf("car response %s", req)
		submitTxnFn("mvd", "autochannel", "KBA-Automobile", "CarContract", "invoke", make(map[string][]byte), "RegisterCar", req.CarId, req.CarOwner, req.RegNumber)

		ctx.JSON(http.StatusOK, req)
	})

	router.Run("localhost:8080")
}
