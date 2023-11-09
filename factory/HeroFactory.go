package factory

import "fmt"

func GetHero(hero string) (IHero, error) {
	if hero == "Rubick" {
		return newRubick(), nil
	}
	if hero == "Dazzle" {
		return newDazzle(), nil
	}
	if hero == "Drow Ranger" {
		return newDrowRanger(), nil
	}
	if hero == "Tiny" {
		return newTiny(), nil
	}
	return nil, fmt.Errorf("hero not found")
}
