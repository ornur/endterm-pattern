package Singleton

import (
	"fmt"
	"math/rand"
	"sync"
)

var lock = &sync.Mutex{}

type Roshan struct {
	currentHP    int
	maxHP        int
	hasCheese    bool
	hasAegis     bool
	hasAghanim   bool
	hasRefresher bool
}

var rosh *Roshan

func GetRoshan() *Roshan {
	if rosh == nil {
		lock.Lock()
		defer lock.Unlock()
		if rosh == nil {
			rosh = &Roshan{
				maxHP:    rand.Int()%8000 + 2000,
				hasAegis: true,
			}
			rosh.currentHP = rosh.maxHP
			fmt.Printf("Roshan spawned with %d HP and Aegis.\n", rosh.maxHP)
		} else {
			fmt.Printf("Roshan already spawned and has %d/%d HP.\n", rosh.currentHP, rosh.maxHP)
		}
	} else {
		fmt.Printf("Roshan already spawned and has %d/%d HP.\n", rosh.currentHP, rosh.maxHP)
	}

	return rosh
}
