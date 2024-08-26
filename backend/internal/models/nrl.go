package models

type NRLFixture struct {
	ID             string  `json:"matchId"`
	RoundTitle     string  `json:"roundTitle"`
	MatchState     string  `json:"matchState"`
	Venue          string  `json:"venue"`
	VenueCity      string  `json:"venueCity"`
	MatchCentreURL string  `json:"matchCentreURL"`
	HomeTeam       NRLTeam `json:"homeTeam"`
	AwayTeam       NRLTeam `json:"awayTeam"`
	KickOffTime    string  `json:"startTime"`
}

type NRLTeam struct {
	ID    int       `json:"teamId"`
	Name  string    `json:"nickName"`
	Odds  *string   `json:"odds,omitempty"`
	Score *int      `json:"score,omitempty"`
	Form  []NRLForm `json:"recentForm,omitempty"`
}

type NRLForm struct {
	Result string `json:"result"`
	Score  string `json:"score"`
}
