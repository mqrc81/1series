package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/cyruzin/golang-tmdb"
	"github.com/mqrc81/zeries/domain"
	"github.com/mqrc81/zeries/trakt"
)

const (
	tmdbImageUrl = "https://image.tmdb.org/t/p/original"
)

type ShowHandler struct {
	store domain.Store
	tmdb  *tmdb.Client
	trakt *trakt.Client
}

func (h *ShowHandler) PopularShows() http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		var shows []domain.Show

		page := 1
		if req.URL.Query().Has("page") {
			page, _ = strconv.Atoi(req.URL.Query().Get("page"))
		}

		traktShows, err := h.trakt.ShowsWatchedWeekly(page, 20)
		if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}

		for _, traktShow := range traktShows {
			tmdbShow, err := h.tmdb.GetTVDetails(traktShow.TmdbId(), nil)
			if err != nil {
				http.Error(res, err.Error(), http.StatusInternalServerError)
				return
			}

			shows = append(shows, showFromDto(tmdbShow))
		}

		if err = respond(res, shows); err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func (h *ShowHandler) Show() http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		// TODO
	}
}

func (h *ShowHandler) SearchShows() http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		// TODO
	}
}

func respond(res http.ResponseWriter, body interface{}) error {
	bodyJSON, err := json.Marshal(body)
	if err != nil {
		return err
	}

	res.Header().Add("Content-Type", "application/json")
	if _, err = res.Write(bodyJSON); err != nil {
		return err
	}

	return nil
}

func showFromDto(show *tmdb.TVDetails) domain.Show {

	var genres []domain.Genre
	for _, genre := range show.Genres {
		fmt.Println(genre.Name)
		genres = append(genres, domain.Genre{
			Id:   int(genre.ID),
			Name: genre.Name,
		})
	}

	var networks []domain.Network
	for _, network := range show.Networks {
		networks = append(networks, domain.Network{
			Id:   int(network.ID),
			Name: network.Name,
			Logo: tmdbImageUrl + network.LogoPath,
		})
	}

	airDate, err := time.Parse("2006-01-02", show.FirstAirDate)
	if err != nil {
		airDate, _ = time.Parse("2006-01-02", show.Seasons[0].AirDate)
	}

	return domain.Show{
		Id:            int(show.ID),
		Name:          show.Name,
		Description:   show.Overview,
		Year:          airDate.Year(),
		Poster:        tmdbImageUrl + show.PosterPath,
		Rating:        show.VoteAverage,
		Genres:        genres,
		Networks:      networks,
		Homepage:      show.Homepage,
		SeasonsCount:  show.NumberOfSeasons,
		EpisodesCount: show.NumberOfEpisodes,
	}
}
