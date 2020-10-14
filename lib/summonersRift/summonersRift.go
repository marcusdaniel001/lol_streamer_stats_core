package sr

import (
	"igor.com.br/lol/lib"
	"igor.com.br/lol/lib/infra"
)

type SummonersRift struct {
	Account     *lib.LeagueAccount
	MatchStatus *lib.MatchStatus
	Match       *lib.Match
	Mongo       *infra.MongoDB
	Playing     bool
}

const NOW_MATCH = 0

func (s *SummonersRift) GetPlayers() []lib.Player {
	player1 := lib.NewPlayer("Toin", s.Match.Participants[1].Stats.TotalDamageDealt)
	player2 := lib.NewPlayer("Jubeleu", s.Match.Participants[2].Stats.TotalDamageDealt)

	players := []lib.Player{
		*player1,
		*player2,
	}

	return players
}
func (s *SummonersRift) FindMatch() *SummonersRift {
	if s.Playing {
		s.Match.GetMatch(s.MatchStatus.Matches[NOW_MATCH].GameID)
	} else {
		s.Queue()
	}

	return s

}
func (s *SummonersRift) Queue() *SummonersRift {
	s.MatchStatus.GetMatchStatusByID(s.Account.AccountID)
	if s.MatchStatus.Matches[NOW_MATCH].GameID != 0 {
		s.Playing = true
	}

	return s
}

func (s *SummonersRift) Login(userName string) *SummonersRift {
	account := s.Mongo.GetAccount(s.Account, userName)
	if account.ID != "" {
		s.Account = account
	} else {
		s.Account.GetAccount(userName)
		s.Mongo.AddNewAccount(s.Account.GetAccount(userName))
	}
	return s
}

func NewSummonerRift() *SummonersRift {
	return &SummonersRift{
		Account:     lib.NewLeagueAccount(),
		MatchStatus: lib.NewMatchStatus(),
		Match:       lib.NewMatch(),
		Mongo:       infra.NewMongoDB("root", "123"),
		Playing:     false,
	}
}
