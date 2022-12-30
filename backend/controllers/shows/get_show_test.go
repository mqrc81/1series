package shows

import (
	"encoding/json"
	"github.com/cyruzin/golang-tmdb"
	"github.com/h2non/gock"
	"github.com/labstack/echo/v4"
	"github.com/mqrc81/zeries/domain"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetShow_valid(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
	ctx.SetPath("/api/shows/:showId")
	ctx.SetParamNames("showId")
	ctx.SetParamValues("13")

	defer gock.Off()

	gock.New("https://api.themoviedb.org/3").
		Get("/tv/13").
		Reply(http.StatusOK).
		JSON(tmdb.TVDetails{
			ID:           13,
			Name:         "Game of Thrones",
			FirstAirDate: "2022-12-30",
			VoteAverage:  7.344902,
		})

	c := &showsController{
		tmdbClient: &tmdb.Client{},
	}

	// Assertions
	if assert.NoError(t, c.GetShow(ctx)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		expectedShow := domain.Show{
			Id:            13,
			Name:          "Game of Thrones",
			Year:          2022,
			Rating:        7.3,
			SeasonsCount:  0,
			EpisodesCount: 0,
			Genres:        []domain.Genre{},
			Networks:      []domain.Network{},
		}
		assert.Equal(t, jsonStringify(expectedShow), rec.Body.String())
	}
}

func TestGetShow_invalidShowId(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
	ctx.SetPath("/api/shows/:showId")
	ctx.SetParamNames("showId")

	c := &showsController{}

	// Assertions
	result := c.GetShow(ctx)
	if assert.Error(t, result) {
		assert.Equal(t, http.StatusBadRequest, result.(*echo.HTTPError).Code)
		assert.Equal(t, "invalid show-id", result.(*echo.HTTPError).Message)
	}
}

func TestGetShow_tmdbError(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
	ctx.SetPath("/api/shows/:showId")
	ctx.SetParamNames("showId")
	ctx.SetParamValues("13")

	defer gock.Off()

	gock.New("https://api.themoviedb.org/3").
		Get("/tv/13").
		Reply(http.StatusBadRequest).
		JSON(tmdb.Error{
			StatusMessage: "show not found",
		})

	c := &showsController{
		tmdbClient: &tmdb.Client{},
	}

	// Assertions
	result := c.GetShow(ctx)
	if assert.Error(t, result) {
		assert.Equal(t, http.StatusConflict, result.(*echo.HTTPError).Code)
		assert.ErrorContains(t, result, "show not found")
	}
}

func jsonStringify(i interface{}) string {
	bytes, err := json.Marshal(i)
	if err != nil {
		panic(err)
	}
	return string(bytes) + "\n"
}
