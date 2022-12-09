package domain

type Show struct {
	Id            int
	Name          string
	Overview      string
	Year          int
	Poster        string
	Rating        float32
	Homepage      string
	SeasonsCount  int
	EpisodesCount int
	Genres        []Genre
	Networks      []Network
}
