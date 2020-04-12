package main

import (
	"fmt"
	"sync"
 
)
type ChopS struct {
	sync.Mutex
}
type Philos struct {
	num int
	leftcs,rightcs *ChopS
}

func (p Philos) eat(c chan *Philos, wg *sync.WaitGroup) {
	for i:=0;i<3;i++ {
			c <- &p
			
			p.leftcs.Lock()
			p.rightcs.Lock()

			fmt.Println("starting to eat ",p.num)
			fmt.Println("finishing eating",p.num)
			p.rightcs.Unlock()
			p.leftcs.Unlock()
	        <- c		// consume itself
		

	}
		wg.Done()
}


func main() {
	var i int
	var wg sync.WaitGroup
	c := make(chan *Philos,2)
    fmt.Println(len(c))
	wg.Add(5)

	ChopSticks := make([] *ChopS,5)
	for i=0;i<5;i++ {
		ChopSticks[i] = new(ChopS)
	}

	Philosophers := make([] *Philos,5)
	for i=0;i<5;i++ {
		Philosophers[i] = &Philos{i+1,ChopSticks[i],ChopSticks[(i+1)%5]}
	}


	for i=4;i>=0;i-- {
		go Philosophers[i].eat(c,&wg)
	}
	wg.Wait()
}