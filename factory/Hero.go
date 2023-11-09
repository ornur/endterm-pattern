package factory

import "fmt"

type Hero struct {
	Name string
	Type string
}

func (h *Hero) GetName() string {
	return h.Name
}
func (h *Hero) GetType() string {
	return h.Type
}

func (h *Hero) SetName(name string) {
	h.Name = name
}

func (h *Hero) SetType(heroType string) {
	h.Type = heroType
}
func PrintHero(hero IHero) {
	fmt.Println(hero.GetName() + " is a " + hero.GetType() + " hero.")
}
