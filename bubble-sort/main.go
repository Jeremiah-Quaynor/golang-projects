package main

import "fmt"

func main() {
	//Swap, check order, one-round bubbling, repeated n times

	numbers := []int{1, 46, 23, 45, 34, 32}

	for j := 0; j< len(numbers)-1; j++ {
		for i := 0; i< len(numbers)-1; i++ {
			if numbers[i] > numbers[i+1]{
				swap(numbers,i,i+1)
			}
		}
	}


	fmt.Println(numbers)
}

func swap(numbers []int, i, j int) {

	temp := numbers[i]
	numbers[i] = numbers[j]
	numbers[j] = temp

}