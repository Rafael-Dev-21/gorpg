package actor

import (
	"fmt"
	"strings"
)

type Actor struct {
	Name    string `json:"name"`
	Life    int    `json:"life"`
	Attack  int    `json:"attack"`
	Defense int    `json:"defense"`
}

func (a *Actor) TakeHit(damageRaw int) string {
	damage := damageRaw - a.Defense/4

	if damage <= 0 {
		return fmt.Sprintf("%v defends against the attack and blocks the damage!", a.Name)
	}

	a.Life -= damage
	return fmt.Sprintf("%v tries to defend but even so receives %d damage!", a.Name, damage)
}

func (a Actor) Hit() int {
	return a.Attack / 2
}

func (a Actor) Describe() string {
	bars := strings.Repeat("-", 32)

	format := `%v
%v
%v
Life:    %d
Attack:  %d
Defense: %d
%v`

	return fmt.Sprintf(format, bars, a.Name, bars, a.Life, a.Attack, a.Defense, bars)
}

func (a Actor) IsAlive() bool {
	return a.Life > 0
}

func CreateActor(name string, life int, attack int, defense int) *Actor {
	return &Actor{name, life, attack, defense}
}
