package config

// Match States
const (
	MatchStateUpcoming = "Upcoming" // Match has not started yet
	MatchStateFullTime = "FullTime" // Match has ended

	// TBD: Add more states as needed
)

// Competition IDs
const (
	CompetitionNRL                 = 111 // National Rugby League
	CompetitionNRLW                = 161 // National Rugby League Women's
	CompetitionStateOfOrigin       = 116 // State of Origin
	CompetitionStateOfOriginWomens = 156 // State of Origin Women's
)

// Scheduling Constants
const (
	CheckInterval     = 5 * 60  // Interval in seconds to recheck match status if not "FullTime"
	InitialCheckDelay = 80 * 60 // Delay in seconds for initial check after kickoff
)
