package main

func main() {
	var home = Home{
		money: 100,
		food:  50,
		dirt:  0,
	}

	var husband = Husband{Human{
		name:     "John",
		fullness: 30,
		mood:     100,
	}}
	var wife = Wife{Human{
		name:     "Ella",
		fullness: 30,
		mood:     100,
	}}

	for i := 0; i < 365; i++ {
		home.AddDirt(5)
		if home.dirt > 90 {
			husband.MinusMood(10)
			wife.MinusMood(10)
		}
	}
}

//func (h *Human) Eat(amountOfFood int) {
//	if amountOfFood > 30 {
//		panic("только человек может съесть меньше или 30")
//	}
//
//	h.AddFullness(amountOfFood)
//}

func () {

}
