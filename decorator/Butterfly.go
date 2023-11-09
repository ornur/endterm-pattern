package decorator

type Butterfly struct {
	Hero IHeroInventory
}

func (i *Butterfly) GetNetWorth() int {
	return i.Hero.GetNetWorth() + 4975
}
