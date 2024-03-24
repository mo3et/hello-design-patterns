package main

import "fmt"

type EventBusE interface {
	Subs(obs ObserverE)
	Unsubs(obs ObserverE)
	Publ()
}

type ObserverE interface {
	Update(hour int)
}

type Clock struct {
	obss []ObserverE
	hour int
}

func (c *Clock) Subs(obs ObserverE) {
	c.obss = append(c.obss, obs)
}

func (c *Clock) Unsubs(obs ObserverE) {
	for i, obser := range c.obss {
		if obser == obs {
			c.obss = append(c.obss[:i], c.obss[i+1:]...)
			break
		}
	}
}

func (c *Clock) Publ() {
	for _, obser := range c.obss {
		obser.Update(c.hour)
	}
}

func (c *Clock) Tick() {
	c.hour = (c.hour + 1) % 24 // 模拟世界的推移
	c.Publ()
}

// 具体 Observer 实现
type StuObserver struct {
	name string
}

func NewStudent(name string) *StuObserver {
	return &StuObserver{name: name}
}

func (s *StuObserver) Update(hour int) {
	fmt.Println(s.name, hour)
}

func main() {
	// Get student numbers
	var N int
	fmt.Scan(&N)

	// create clock
	clock := &Clock{}

	// subscribe stuObserver in stdin.Scan
	for i := 0; i < N; i++ {
		var stuName string
		fmt.Scan(&stuName)
		clock.Subs(NewStudent(stuName))
	}

	// read clock update times
	var updates int
	fmt.Scan(&updates)

	// mock clock run (when even hour update)
	for i := 0; i < updates; i++ {
		clock.Tick()
	}
}
