package main

import (
	"fmt"
	"sync"
)

type ChopS struct {
	sync.Mutex
}

type Philo struct {
	leftCS, rightCs *ChopS
	number          int
}

var on sync.Once
var hostOutput, hostInput chan bool

func (p Philo) eat(wg *sync.WaitGroup) {
	fmt.Printf("Running eat for Philosopher: %d \n", p.number)
	count := 1
	for {
		fmt.Println("Waiting Signal To Eat")
		<-hostOutput
		fmt.Println("Recieved Signal To Eat")
		p.leftCS.Lock()
		p.rightCs.Lock()

		fmt.Printf("Philosopher %d starting to eat, Course: %d\n", p.number, count)
		fmt.Printf("Philosopher %d is finishing eating, Course: %d\n", p.number, count)
		p.rightCs.Unlock()
		p.leftCS.Unlock()

		hostInput <- true
		count++
		if count == 4 {
			break
		}
	}
	fmt.Printf("Philosopher %d has finished his meal.\n", p.number)
	wg.Done()
}

func host() {
	eatCounter := 0
	
	go func() {
		for {
			if eatCounter < 1 {
				fmt.Println("Sending Signal To Eat")
				hostOutput <- true
				eatCounter++
			}
		}
	}()
	go func() {
		for {
			<-hostInput
			eatCounter--
		}
	}()

}

func main() {
	var wg sync.WaitGroup
	CSticks := make([]*ChopS, 5)
	for i := 0; i < 5; i++ {
		CSticks[i] = new(ChopS)
	}
	philos := make([]*Philo, 5)
	for i := 0; i < 5; i++ {
		philos[i] = &Philo{CSticks[i], CSticks[(i+1)%5], i}
	}

	wg.Add(5)
	for i := 0; i < 5; i++ {
		go philos[i].eat(&wg)
	}
	go host()
	wg.Wait()
	fmt.Printf("Everyone finished their meals!\n")
}
