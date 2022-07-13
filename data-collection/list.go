package main

import "fmt"

func main () {
	var productNames [4]string; 
	prices := [4]float32{45.665,43.54,343.45,56.34}
	productNames[3] = "Books"
	featuredPrice := prices[1:]
	highlightedPrices := featuredPrice[:1]
	fmt.Println(prices[3])
	fmt.Println(productNames[3])
	fmt.Println(featuredPrice)
	fmt.Print(len(highlightedPrices), cap(featuredPrice))
}
