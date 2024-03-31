package main

import (
	"fmt"
	"sync"
)

var mutex = &sync.Mutex{}

type config struct {
}

var singleInstance *config
var counter int = 1

func getInstance() *config {

	if singleInstance == nil {
		mutex.Lock()
		// defer mutex.Unlock()
		if singleInstance == nil {
			fmt.Println("new instance created", counter)
			singleInstance = &config{}
			counter = counter + 1

		} else {
			fmt.Println("Instance-1 already exist")
		}
	} else {
		fmt.Println("Instance-2 already exist")
	}
	return singleInstance
}

func main() {
	for i := 0; i < 100; i++ {
		go getInstance()
	}
	fmt.Scanln()

}
