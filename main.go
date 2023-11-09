package main

import (
	Command "EndtermProjet/command"
	Decorator "EndtermProjet/decorator"
	Factory "EndtermProjet/factory"
	Singleton "EndtermProjet/singleton"
	"fmt"
)

// Console application with 4 software design patterns
// 1. Singleton
// 2. Command
// 3. Decorator
// 4. Factory

func main() {
	fmt.Println("Welcome to Dota 2")
actionLoop:
	for {
		fmt.Println("Choose pattern: ")
		fmt.Println("1. Singleton")
		fmt.Println("2. Command")
		fmt.Println("3. Decorator")
		fmt.Println("4. Factory")
		action := 0
		fmt.Scanf("%d\n", &action)
		switch action {
		case 1:
			singleton()
		case 2:
			command()
		case 3:
			decorator()
		case 4:
			factory()
		default:
			fmt.Println("Wrong action")
		}
		fmt.Print("Press any key to continue or q to exit:")
		var answer string
		fmt.Scanf("%s\n", &answer)
		if answer == "q" {
			break actionLoop
		}
	}
}
func singleton() {
	// Singleton
	Singleton.GetRoshan()
	Singleton.GetRoshan()
	Singleton.GetRoshan()
}
func command() {

	// Command
	armlet := &Command.Armlet{}
	radiance := &Command.Radiance{}

	oncommand := &Command.OnCommand{
		CommandItem: armlet,
	}
	offcommand := &Command.OffCommand{
		CommandItem: armlet,
	}
	onKeyPress := &Command.KeyPress{
		KeyPressCommand: oncommand,
	}
	onKeyPress.Press()

	offKeyPress := &Command.KeyPress{
		KeyPressCommand: offcommand,
	}
	offKeyPress.Press()

	oncommand = &Command.OnCommand{
		CommandItem: radiance,
	}
	offcommand = &Command.OffCommand{
		CommandItem: radiance,
	}
	onKeyPress = &Command.KeyPress{
		KeyPressCommand: oncommand,
	}

	offKeyPress = &Command.KeyPress{
		KeyPressCommand: offcommand,
	}
	onKeyPress.Press()
	offKeyPress.Press()
}
func decorator() {
	// Decorator
	HeroInventory := &Decorator.HeroInventory{}

	InventoryWithSkadi := &Decorator.Skadi{Hero: HeroInventory}
	fmt.Println("Net worth with Skadi: ", InventoryWithSkadi.GetNetWorth())

	InventoryWithSkadiAndButterfly := &Decorator.Butterfly{Hero: InventoryWithSkadi}
	fmt.Println("Net worth with Skadi and Butterfly: ", InventoryWithSkadiAndButterfly.GetNetWorth())

	AddedTarrasque := &Decorator.Tarassque{Hero: InventoryWithSkadiAndButterfly}
	fmt.Println("Net worth with Skadi, Butterfly and Tarrasque: ", AddedTarrasque.GetNetWorth())

	AddedRapier := &Decorator.Rapier{Hero: AddedTarrasque}
	fmt.Println("Net worth with Skadi, Butterfly, Tarrasque and Rapier: ", AddedRapier.GetNetWorth())

	FinalInventory := &Decorator.Hex{Hero: AddedRapier}
	fmt.Println("Net worth with Skadi, Butterfly, Tarrasque, Rapier and Hex: ", FinalInventory.GetNetWorth())
	fmt.Println()
	fmt.Println("Total net worth: ", FinalInventory.GetNetWorth())
}
func factory() {
	// Factory
	rubick, _ := Factory.GetHero("Rubick")
	dazzle, _ := Factory.GetHero("Dazzle")
	drow, _ := Factory.GetHero("Drow Ranger")
	tiny, _ := Factory.GetHero("Tiny")

	Factory.PrintHero(rubick)
	Factory.PrintHero(dazzle)
	Factory.PrintHero(drow)
	Factory.PrintHero(tiny)
}
