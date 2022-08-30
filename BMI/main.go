//package main
//import ( 
//	"fmt"
//	"strings" 
//	"time"
//)
//func main() {
//	const col = 40
//	bar := fmt.Sprintf("\x0c[%%-%vs]",col)
//	for i :=0; i<col; i++ {
//		fmt.Printf(bar, strings.Repeat("=",i)+">")
//		time.Sleep(100 * time.Millisecond)
//	}
//	fmt.Printf(bar+" Done!",strings.Repeat("=",col))
//}
package main

import "fmt"


type Car interface {
	Drive()
	Stop()
}

type Ferrari struct {
	FerrariModel string
}

type Chevy struct {
	ChevyModel string
}

func (F *Ferrari) Drive() {
	fmt.Println("Ferrari is no the move")
	fmt.Println(F.FerrariModel)
}

func (C *Chevy) Drive() {
	fmt.Println("Chevy is on the move")
	fmt.Println(C.ChevyModel)
}

func main() {
    L := Ferrari{"Stallion"}
    C := Chevy{"C98"}
	L.Drive()
	C.Drive()
}
