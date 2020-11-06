package main

import (
	"fmt"
	"sync"
)

var lock = &sync.Mutex{}

type single struct {

}

func (s *single)printMyName() {
	fmt.Println("MyName")
}

var singleInstance *single

func getInstance() *single {
	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleInstance == nil {
			fmt.Println("Creating single instance now")
			singleInstance = &single{}
		} else {
			fmt.Println("Singe instance already created")
		}
	} else {
		fmt.Println("Single instance already created")
	}
	return singleInstance
}


func main() {
	for i:=0; i < 3; i++ {
		getInstance()
	}
}