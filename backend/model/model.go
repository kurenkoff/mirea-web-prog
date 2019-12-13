package model

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type Player struct {
	ID            string `json:"id"`
	FirstName     string `json:"first_name"`
	LastName      string `json:"last_name"`
	Position      string `json:"position"`
	DateOfBirth   string `json:"date_of_birth"`
	Nationality   string `json:"nationality"`
	CountryCode   string `json:"country_code"`
	Height        int    `json:"height"`
	Weight        int    `json:"weight"`
	PreferredFoot string `json:"preferred_foot"`
	Gender        string `json:"gender"`
}

var players []Player

func init() {
	players = make([]Player, 0)

	raw, err := ioutil.ReadFile("players.json")
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(raw, &players)
	if err != nil {
		log.Fatal(err)
	}
}

func GetCollection() ([]byte, error) {
	return json.Marshal(players)

}

func Remove(ids []string) {
	fmt.Println(players, ids)
	for _, id := range ids{
		for i, player := range players{
			if player.ID == id {
				players = append(players[:i], players[(i+1):]...)
				break
			}
		}
	}
}
