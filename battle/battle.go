package battle

import (
	"fmt"

	"github.com/rafael-dev-21/myrpg/actor"
)

func Battle(a actor.Actor, b actor.Actor) {
	current, enemy := &a, &b
	for doTurn(current, enemy) {
		fmt.Println(enemy.Describe())
		current, enemy = enemy, current
	}

	fmt.Printf("%v is the winner! \n", current.Name)
}

func doTurn(a *actor.Actor, b *actor.Actor) bool {
	fmt.Println(b.TakeHit(a.Hit()))
	return b.IsAlive()
}
