package decorator

type IHeroInventory interface {
	GetNetWorth() int
}
type HeroInventory struct {
	Hero     IHeroInventory
	NetWorth int
}

func (i *HeroInventory) GetNetWorth() int {
	return i.NetWorth
}
