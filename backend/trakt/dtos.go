package trakt

// ShowsWatchedDto represents Trakt's DTO, of which we only need the TMDb-ID, since we fetch all the details from TMDb.
// The rest can be ignored.
type ShowsWatchedDto struct {
	Show struct {
		Ids struct {
			Tmdb int `json:"tmdb"`
		} `json:"ids"`
	} `json:"show"`
}

func (show ShowsWatchedDto) TmdbId() int {
	return show.Show.Ids.Tmdb
}
