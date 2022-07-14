package main

import "fmt"



func main() {
	prices := []float64{45.665,43.54,343.45,56.34}
	fmt.Println((prices[0:1]))

	prices[1] = 99.9


	prices = append(prices, 5.66,234.34,344.34)

	prices = prices[1:]

	fmt.Println(prices)


}


// func main () {
// 	var productNames [4]string; 
// 	prices := [4]float32{45.665,43.54,343.45,56.34}
// 	productNames[3] = "Books"
// 	featuredPrice := prices[1:]
// 	highlightedPrices := featuredPrice[:1]
// 	fmt.Println(prices[3])
// 	fmt.Println(productNames[3])
// 	fmt.Println(featuredPrice)
// 	fmt.Print(len(highlightedPrices), cap(featuredPrice))
// }
