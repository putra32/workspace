package main

import (
	"log"
	"sync"
)

func sum(from, to int, wg *sync.WaitGroup, res *int) {
	*res = 0
	for i := from; i <= to; i++ {
		*res += i
	}

	wg.Done()
	return
}

func main() {
	s1 := 0
	s2 := 0
	wg := &sync.WaitGroup{}

	wg.Add(1)
	go sum(1, 100, wg, &s1)
	wg.Wait()
	wg.Add(1)
	go sum(1, 10, wg, &s2)
	wg.Wait()
	log.Println(s1)
	log.Println(s2)

}
