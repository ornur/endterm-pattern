package factory

type Tiny struct {
	Hero
}

func newTiny() IHero {
	return &Tiny{
		Hero: Hero{
			Name: "Tiny",
			Type: "Strength",
		},
	}
}
