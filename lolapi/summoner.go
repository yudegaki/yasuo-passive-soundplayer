package lolapi


type Summoner struct {
	SummonerName string `json:"summmonerName"`
	Level int `json:"level"`
	Abilities map[string]Ability `json:"abilities"`
	ChampionStats ChampionStats `json:"championStats"`
}
