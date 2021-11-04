package main

import (
	"fmt"
)

type mystrutc struct {
	id    int
	value int
	name  string
}

func (mystrutc mystrutc) String() string {
	return fmt.Sprintf(" %d %d %s ", mystrutc.id, mystrutc.value, mystrutc.name)
}

func say(mystrutc mystrutc, c chan mystrutc) {
	//defer wg.Done()
	c <- mystrutc

}
func main() {
	c := make(chan mystrutc, 5)

	say(mystrutc{1, 1, "Erix1"}, c)
	fmt.Println(len(c), cap(c))
	say(mystrutc{2, 1, "Erix2"}, c)
	fmt.Println(len(c), cap(c))
	say(mystrutc{3, 1, "Erix3"}, c)
	fmt.Println(len(c), cap(c))
	say(mystrutc{4, 1, "Erix4"}, c)
	fmt.Println(len(c), cap(c))
	say(mystrutc{5, 1, "Erix5"}, c)
	//fmt.Println(<-c)

	fmt.Println(len(c), cap(c))
	close(c)
	for mystrutc := range c {
		fmt.Println(mystrutc)
	}

	mail1 := make(chan mystrutc, 1)

	mail2 := make(chan mystrutc, 1)
	go say(mystrutc{1, 2, "E"}, mail1)
	go say(mystrutc{2, 3, "E2"}, mail2)
	for i := 0; i < 2; i++ {
		select {
		case m1 := <-mail1:
			fmt.Println("Desde email 1", m1)
		case m2 := <-mail2:
			fmt.Println("Desde email 1", m2)
		}
	}

	/*	var wg sync.WaitGroup
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
		}("Adios")*/
	/*
		wg.Add(3)

		go say("Hello", &wg)
		go say("Hello", &wg)
		wg.Add(2)
		go say("Hello", &wg)
		go say("Hello", &wg)
	*/
	//wg.Wait()
	//time.Sleep(time.Second * 1)
	//
}
