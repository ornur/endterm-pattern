package factory

type DrowRanger struct {
	Hero
}

func newDrowRanger() IHero {
	return &DrowRanger{
		Hero: Hero{
			Name: "Drow Ranger",
			Type: "Agility",
		},
	}
}
