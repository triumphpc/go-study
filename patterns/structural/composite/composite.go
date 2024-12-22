// Создание древовидных однотипных объектов

package main

import "fmt"

// Component — общий интерфейс для котиков и их групп

type Cat interface {
	Meow()
}

// Leaf — одиночный котик

type SimpleCat struct {
	name string
}

func (sc *SimpleCat) Meow() {
	fmt.Printf("%s: Мяу!\n", sc.name)
}

// NewSimpleCat Конструктор для котиков
func NewSimpleCat(name string) *SimpleCat {
	return &SimpleCat{name: name}
}

// Composite — группа котиков

type CatGroup struct {
	name string
	cats []Cat
}

func (cg *CatGroup) Meow() {
	fmt.Printf("%s: Начинаем общий концерт:\n", cg.name)
	for _, cat := range cg.cats {
		cat.Meow()
	}
}

func (cg *CatGroup) Add(cat Cat) {
	cg.cats = append(cg.cats, cat)
}

func NewCatGroup(name string) *CatGroup {
	return &CatGroup{name: name, cats: []Cat{}}
}

func main() {
	barsik := NewSimpleCat("Барсик")
	murzik := NewSimpleCat("Мурзик")

	dvorniki := NewCatGroup("Дворовая братва")
	dvorniki.Add(barsik)
	dvorniki.Add(murzik)

	aristokraty := NewCatGroup("Аристократы")
	aristokraty.Add(NewSimpleCat("Людовик"))
	aristokraty.Add(NewSimpleCat("Шарль"))

	zoo := NewCatGroup("Зоопарк")
	zoo.Add(dvorniki)
	zoo.Add(aristokraty)

	zoo.Meow()
}
