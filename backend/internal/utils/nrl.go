package utils

import "strconv"

func ParseMatchID(id string) (season, competition, round, game int) {
	season, _ = strconv.Atoi(id[0:4])
	competition, _ = strconv.Atoi(id[4:7])
	round, _ = strconv.Atoi(id[7:9])
	game, _ = strconv.Atoi(string(id[9]))

	return
}
