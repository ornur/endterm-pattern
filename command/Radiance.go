package command

import "fmt"

type Radiance struct {
	isOn bool
}

func (r *Radiance) on() {
	fmt.Println("Radiance is on")
	r.isOn = true
}
func (r *Radiance) off() {
	fmt.Println("Radiance is off")
	r.isOn = false
}
