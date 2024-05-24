package main

import "fmt"

// refactoring Guru
//

// 定义一个行为接口
// 拥有一个具体组件，无任何字段(BaseDecorator)
// 具体修饰器结构体包含行为接口字段，可以让组件放入(concrete Decorator)

type IPizza interface {
	getPrice() int
}

type peppyPaneer struct{}

func (p *peppyPaneer) getPrice() int {
	return 20
}

type veggeMania struct{}

func (p *veggeMania) getPrice() int {
	return 15
}

type cheeseTopping struct {
	pizza IPizza
}

func (c *cheeseTopping) getPrice() int {
	pizzaPrice := c.pizza.getPrice()
	return pizzaPrice + 10
}

type tomatoTopping struct {
	pizza IPizza
}

func (t *tomatoTopping) getPrice() int {
	pizzPrice := t.pizza.getPrice()
	return pizzPrice + 7
}

func main() {
	veggiePizza := &veggeMania{}

	// Add chesse topping
	veggiePizzaWithCheese := &cheeseTopping{
		pizza: veggiePizza,
	}

	// Add tomato topping
	veggiePizzaWithCheeseAndTomato := &tomatoTopping{
		pizza: veggiePizzaWithCheese,
	}

	fmt.Printf("Price of veggie with tomato and cheese is %d\n", veggiePizzaWithCheeseAndTomato.getPrice())

	peppyPaneerPizza := &peppyPaneer{}

	// Add cheese topping
	peppyPaneerPizzaWithCheese := &cheeseTopping{
		pizza: peppyPaneerPizza,
	}

	fmt.Printf("Price of peppyPaneer with cheese topping is %d\n", peppyPaneerPizzaWithCheese.getPrice())
}
