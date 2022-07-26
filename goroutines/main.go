package main 

import (
	"fmt"
	"math/rand"
	"time"
)


var source = rand.NewSource(time.Now().Unix())
var randN = rand.New(source)

func main() {
	x := make(chan int)
	y := make(chan int)
	limiter := make(chan int, 3)


	go generateValue(x, limiter)
	go generateValue(y, limiter)


	var a int 
	var b int

	select {
	case a = <- x:
		fmt.Printf("X finished first. value is: %v",a)
	case b = <-y:
		fmt.Printf("Y finished first. value is: %v",b)

	}


	// go generateValue(c, limiter)
	// go generateValue(c, limiter)
	// sum := 0
	// i := 0
	// for num := range c {
	// 	sum += num
	// 	i ++
	// 	if i ==4 {
	// 	close(c)
	// 	}
	// }
	// fmt.Print(sum)
}

func generateValue(c chan int, limit chan int) int{
	limit <- 1
	fmt.Println("Generating value...")
	sleepTime := rand.Intn(3)
	time.Sleep(time.Duration(sleepTime) * time.Second)
	result := randN.Intn(10)

	c <- result
	<- limit
	return result
}