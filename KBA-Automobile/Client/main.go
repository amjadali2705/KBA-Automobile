package main

import "fmt"

func main() {

	result := submitTxnFn(
		"manufacturer",
		"autochannel",
		"KBA-Automobile",
		"CarContract",
		"invoke",
		make(map[string][]byte),
		"CreateCar",
		"Car-20",
		"Audi",
		"Q6",
		"Red",
		"Audi Manufacturers",
		"11/11/2024",
	)

	// result := submitTxnFn("manufacturer", "autochannel", "KBA-Automobile", "CarContract", "query", make(map[string][]byte), "ReadCar", "Car-06")

	// privateData := map[string][]byte{
	// 	"make":       []byte("Maruti"),
	// 	"model":      []byte("Alto"),
	// 	"color":      []byte("Red"),
	// 	"dealerName": []byte("Popular"),
	// }

	// result := submitTxnFn("dealer", "autochannel", "KBA-Automobile", "OrderContract", "private", privateData, "CreateOrder", "ORD-03")

	// result := submitTxnFn("dealer", "autochannel", "KBA-Automobile", "OrderContract", "query", make(map[string][]byte), "ReadOrder", "ORD-03")

	// result := submitTxnFn("manufacturer", "autochannel", "KBA-Automobile", "CarContract", "query", make(map[string][]byte), "GetAllCars")

	// result := submitTxnFn("manufacturer", "autochannel", "KBA-Automobile", "OrderContract", "query", make(map[string][]byte), "GetAllOrders")

	// result := submitTxnFn("manufacturer", "autochannel", "KBA-Automobile", "CarContract", "query", make(map[string][]byte), "GetMatchingOrders", "Car-06")

	// result := submitTxnFn("manufacturer", "autochannel", "KBA-Automobile", "CarContract", "invoke", make(map[string][]byte), "MatchOrder", "Car-06", "ORD-03")

	// result := submitTxnFn("mvd", "autochannel", "KBA-Automobile", "CarContract", "invoke", make(map[string][]byte), "RegisterCar", "Car-06", "Dani", "KL-01-CD-01")




	// result := submitTxnFn("manufacturer", "autochannel", "KBA-Automobile", "CarContract", "query", make(map[string][]byte), "ReadCar", "Car-06")


	fmt.Println(result)

}
