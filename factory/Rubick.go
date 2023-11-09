package factory

type Rubick struct {
	Hero
}

func newRubick() IHero {
	return &Rubick{
		Hero: Hero{
			Name: "Rubick",
			Type: "Intelligence",
		},
	}
}
