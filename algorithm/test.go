package main

import (
	"fmt"
	"runtime"
  "time"
)

func main() {
	names := []string{"chandler","oscars","names","perls"}
	for _,name := range(names){
		go func(foo string){
			fmt.Println(foo)
		}(name)

}
time.Sleep(1*time.Second)
runtime.Gosched()
}
