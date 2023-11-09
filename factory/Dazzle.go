package factory

type Dazzle struct {
	Hero
}

func newDazzle() IHero {
	return &Dazzle{
		Hero: Hero{
			Name: "Dazzle",
			Type: "Universal",
		},
	}
}
