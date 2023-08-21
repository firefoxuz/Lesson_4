package main

type Human struct {
	name     string //имя человека
	mood     int    //настроение человека
	fullness int
}

func (h *Human) AddFullness(amount int) {
	h.fullness += amount
}

func (h *Human) MinusFullness(amount int) {
	h.fullness -= amount
}

func (h *Human) AddMood(amount int) {
	h.mood += amount
}

func (h *Human) MinusMood(amount int) {
	h.mood -= amount
}

func (h *Human) Eat(home *Home) {
	if home.food > 60 {
		h.AddFullness(30)
		home.MinusFood(30)
	} else {
		h.AddFullness(home.food / 2)
		home.MinusFood(home.food / 2)
	}
}

type Husband struct {
	Human
}

type Wife struct {
	Human
}

func (h *Husband) PlayComputer() {
	h.AddMood(20)
	h.MinusFullness(10)
}

func (h *Husband) GoWork(home *Home) {
	home.AddMoney(150)
	h.MinusFullness(10)
}

func (w *Wife) BuyProducts(home *Home) {
	if home.money > 60 {
		home.MinusMoney(60)
		home.AddFood(60)
	} else {
		home.AddFood(home.money)
		home.MinusMoney(home.money)
	}
	w.MinusFullness(10)
}

func (w *Wife) BuyClothes(home *Home) {
	if home.money >= 350 {
		home.MinusMoney(350)
		w.AddMood(60)
		w.MinusFullness(10)
	}
}

func (w *Wife) CleanHome(home *Home) {
	if home.dirt >= 100 {
		home.MinusDirt(100)
	}
	w.MinusFullness(10)
}
