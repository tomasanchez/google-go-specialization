package main

import "sync"

//There should be 5 philosophers sharing chopsticks, with one chopstick between each adjacent pair of philosophers.

var turns sync.WaitGroup

type Philosopher struct {
	name                          int
	leftChopstick, rightChopstick *Chopstick
}

type Chopstick struct {
	sync.Mutex
}

func (p Philosopher) Eat() {
	// Each philosopher picks up the chopsticks in any order, not lowest-numbered first (which we did in lecture).
	for i := 0; i < 3; i++ {
		p.leftChopstick.Lock()
		p.rightChopstick.Lock()
		println("starting to eat ", p.name)
		println("finishing eating ", p.name)
		p.rightChopstick.Unlock()
		p.leftChopstick.Unlock()
	}

	turns.Done()
}

func Host(wg *sync.WaitGroup, philos []Philosopher) {

	p1, p2 := philos[0], philos[1]

	turns.Add(2)

	go p1.Eat()
	go p2.Eat()

	turns.Wait()

	turns.Add(2)
	p3, p4 := philos[2], philos[3]
	go p3.Eat()
	go p4.Eat()
	turns.Wait()

	turns.Add(1)
	p5 := philos[4]
	go p5.Eat()
	turns.Wait()

	wg.Done()
}

func main() {

	// Philosophers in any order
	count := 5
	philos := make([]Philosopher, count)
	chops := make([]*Chopstick, count)

	for i := 0; i < count; i++ {
		chops[i] = new(Chopstick)
	}

	for i := 0; i < count; i++ {
		philos[i] = Philosopher{
			i + 1,
			chops[i],
			chops[(i+1)%count]}
	}

	var wg sync.WaitGroup

	wg.Add(1)
	go Host(&wg, philos[:])
	wg.Wait()

}
