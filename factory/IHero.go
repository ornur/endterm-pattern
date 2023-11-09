package factory

type IHero interface {
	GetName() string
	SetName(name string)
	GetType() string
	SetType(heroType string)
}
