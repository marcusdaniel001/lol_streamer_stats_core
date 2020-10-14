package lib

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type LeagueAccount struct {
	ID            string `json:"id"`
	AccountID     string `json:"accountId"`
	Puuid         string `json:"puuid"`
	Name          string `json:"name"`
	ProfileIconID int    `json:"profileIconId"`
	RevisionDate  int64  `json:"revisionDate"`
	SummonerLevel int    `json:"summonerLevel"`
}

func (l *LeagueAccount) GetAccount(accountName string) *LeagueAccount {
	resp, err := http.Get("https://br1.api.riotgames.com/lol/summoner/v4/summoners/by-name/LuisaOdonto?api_key=RGAPI-f8342a15-eff7-41ba-98e0-29f09e221688")
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	json.Unmarshal([]byte(body), &l)

	return l
}

func NewLeagueAccount() *LeagueAccount {
	return &LeagueAccount{}
}
