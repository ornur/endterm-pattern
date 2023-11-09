package command

import "fmt"

type Armlet struct {
	isOn bool
}

func (a *Armlet) on() {
	a.isOn = true
	fmt.Println("Armlet is on")
}
func (a *Armlet) off() {
	a.isOn = false
	fmt.Println("Armlet is off")
}
