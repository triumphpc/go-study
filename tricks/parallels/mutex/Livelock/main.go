package main

import (
	"fmt"
	"time"
)

//  Передают друг другу ложку и программа не продвигается дальше, некая рекурсия в процессе
// Livelock— это подмножество более широкого набора проблем, называемых Starvation.

func main() {
	husbend := &diner{
		name:     "Муж",
		isHungry: true,
	}

	wife := &diner{
		name:     "Жена",
		isHungry: true,
	}

	sp := &spoon{
		husbend,
	}

	go husbend.eatWith(sp, wife)
	wife.eatWith(sp, husbend)
}

type spoon struct {
	owner *diner
}

func (s spoon) use() {
	fmt.Printf("%s has eaten!\n", s.owner.name)
}

type diner struct {
	name     string
	isHungry bool
}

func (d *diner) eatWith(sp *spoon, spouse *diner) {
	for d.isHungry {
		// 1. Если ложке не у нас  - подождем
		if sp.owner != d {
			time.Sleep(1 * time.Second)
			continue
		}

		// 2. Если супруг(а) голодна, то уступим и передадим ложку ему/ей
		if spouse.isHungry {
			fmt.Printf("%s: You eat first my darling %s!\n", d.name, spouse.name)
			sp.owner = spouse
			continue
		}

		// 3. Используем ложку и наконец-то обедаем
		sp.use()
		d.isHungry = false
		fmt.Printf("%s: I'm stuffed? my darling %s!\n", d.name, spouse.name)
		sp.owner = spouse
	}
}
