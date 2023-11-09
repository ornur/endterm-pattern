package decorator

type Tarassque struct {
	Hero IHeroInventory
}

func (i *Tarassque) GetNetWorth() int {
	return i.Hero.GetNetWorth() + 5000
}
