package decorator

type Rapier struct {
	Hero IHeroInventory
}

func (i *Rapier) GetNetWorth() int {
	return i.Hero.GetNetWorth() + 5600
}
