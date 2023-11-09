package decorator

type Hex struct {
	Hero IHeroInventory
}

func (i *Hex) GetNetWorth() int {
	return i.Hero.GetNetWorth() + 5550
}
