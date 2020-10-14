package lib

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type MatchStatus struct {
	Matches []struct {
		PlatformID string `json:"platformId"`
		GameID     int    `json:"gameId"`
		Champion   int    `json:"champion"`
		Queue      int    `json:"queue"`
		Season     int    `json:"season"`
		Timestamp  int64  `json:"timestamp"`
		Role       string `json:"role"`
		Lane       string `json:"lane"`
	} `json:"matches"`
	StartIndex int `json:"startIndex"`
	EndIndex   int `json:"endIndex"`
	TotalGames int `json:"totalGames"`
}

func (m *MatchStatus) GetMatchStatusByID(accountID string) *MatchStatus {
	fmt.Println("accountID", accountID)
	resp, err := http.Get("https://br1.api.riotgames.com/lol/match/v4/matchlists/by-account/" + accountID + "?season=13&api_key=RGAPI-f8342a15-eff7-41ba-98e0-29f09e221688")
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	json.Unmarshal([]byte(body), &m)

	return m
}
func NewMatchStatus() *MatchStatus {
	return &MatchStatus{}
}
