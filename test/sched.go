
package main

import (
	"runtime"
)

func main() {
	// runtime.GOMAXPROCS(runtime.NumCPU())
	runtime.GOMAXPROCS(1)
	exit := make(chan struct{})

	go func() {
		defer close(exit)

		go func (){
			println("b")
		}()

		for i:=0; i<4; i++ {
			println("a:", i)

			if(i == 1) {
				runtime.Gosched()
			}
		}
	}()
	
	<-exit
	
}