package src

import (
	"fmt"
	"log"
	"time"
)

func Trace(funcName string) func(){
	start := time.Now()
	fmt.Printf("func s% enter \n",funcName)
	return func() {
		log.Printf("func s% exit s% ",funcName,time.Since(start))
	}
}


func foo(){
	defer Trace("foo()")()
	time.Sleep(2*time.Second)
}

func main(){
	foo()
	foo()
}
