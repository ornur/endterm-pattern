package decorator

type Skadi struct {
	Hero IHeroInventory
}

func (i *Skadi) GetNetWorth() int {
	return i.Hero.GetNetWorth() + 5300
}
