package main 

import (
	"fmt"
	"os"
)

func main () {
	greet()
	storeData("This is some dummy data!", "dummy-data.txt")

	channel:= make(chan int)
	// channel2 := make(chan int)

	go storeMoreData(50000, "500_1.txt", channel)
	// go storeMoreData(50000, "500_2.txt", channel2)

	<-channel
	<-channel
	// <-channel2 

}

func greet() {
	fmt.Println("Hi there!")
}

func storeData(storableText string, fileName string) {
	file, err := os.OpenFile(fileName, 
	os.O_CREATE|os.O_APPEND|os.O_WRONLY,
0666,
)	

	if err != nil {
	fmt.Println("Creating the file failed. Exiting")
	}

	defer file.Close()
	_, err = file.WriteString(storableText)

	if err != nil {
	fmt.Println("Writing to the file failed. Exiting")
	}

}

func storeMoreData(lines int, fileName string, c chan int) {
	for i := 0; i <lines; i++ {
		text := fmt.Sprint("Line %v - Dummy Data\n", i)
		storeData(text, fileName)
	}

	fmt.Printf("-- Done storing %v lines of text --\n", lines)
	c <- 1
}