package pizza

//type PizzaNameRequest struct {name string}
//type PizzaNameResponse struct {pizza Pizza}

type Pizza struct {
	Name string
	Id   int64
}

func getPizzaByName(name string) Pizza {
	return Pizza{name, 66}
}

func NewPizza(name string, id int64) Pizza {
	return Pizza{name, id}
}

// type PizzaBaker struct {
// 	Name   string
// 	pizzas map[string]Pizza
// 	Id     int64
// }

// func (p PizzaBaker) GetPizzas() []Pizza {
// 	result := []Pizza{}
// 	for _, pizza := range p.pizzas {
// 		result = append(result, pizza)
// 	}
// 	return result
// }
