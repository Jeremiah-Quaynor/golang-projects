package main 

import (
	"fmt"
	"math/rand"
	"time"
)


var source = rand.NewSource(time.Now().Unix())
var randN = rand.New(source)

func main() {

	x := generateValue()
	y := generateValue()

	sum := x +y 

	fmt.Print(sum)
}

func generateValue() int{
	sleepTime := rand.Intn(3)
	time.Sleep(time.Duration(sleepTime) * time.Second)

	return randN.Intn(10)
}