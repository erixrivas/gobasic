package main

import (
	"fmt"
	"strconv"
	"sync"
)

func say(text string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(text)

}
func main() {
	var wg sync.WaitGroup
	numero := 5

	fmt.Println("Hello, World!")

	wg.Add(numero)

	for i := 0; i <= numero; i++ {
		var stringA string = "Hello "
		go say(stringA+strconv.Itoa(i), &wg)
	}
	wg.Add(1)
	go func(text string) {
		fmt.Println(text)
	}("Adios")
	/*
		wg.Add(3)

		go say("Hello", &wg)
		go say("Hello", &wg)
		wg.Add(2)
		go say("Hello", &wg)
		go say("Hello", &wg)
	*/
	wg.Wait()
	//time.Sleep(time.Second * 1)
}
