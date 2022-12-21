package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/rafael-dev-21/myrpg/actor"
	"github.com/rafael-dev-21/myrpg/battle"
)

func main() {
	hero := actor.CreateActor("Hero", 100, 10, 8)
	var monsterList []actor.Actor

	monsterFile, err := os.ReadFile("data/monsters.json")
	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal(monsterFile, &monsterList)

	monster := monsterList[0]

	fmt.Println(monster.Describe())
	fmt.Println(hero.Describe())
	battle.Battle(*hero, monster)
}
