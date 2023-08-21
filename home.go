package main

type Home struct {
	money int // кол-во денег
	food  int //кол-во еды
	dirt  int //кол-во грязи
}

func (h *Home) AddMoney(amount int) {
	h.money += amount
}

func (h *Home) MinusMoney(amount int) {
	h.money -= amount
}

func (h *Home) AddDirt(amount int) {
	h.dirt += amount
}

func (h *Home) MinusDirt(amount int) {
	h.dirt -= amount
}

func (h *Home) AddFood(amount int) {
	h.food += amount
}

func (h *Home) MinusFood(amount int) {
	h.food -= amount
}
