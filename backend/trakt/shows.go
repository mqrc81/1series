package trakt

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	baseURL = "https://api.trakt.tv"
)

func (c *Client) GetShowsWatchedWeekly(page int, limit int) (showsWatched []ShowsWatchedDto, err error) {

	url := fmt.Sprintf(baseURL+"/shows/watched/weekly?page=%d&limit=%d", page, limit)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("trakt-api-version", "2")
	req.Header.Add("trakt-api-key", c.apiKey)

	res, err := c.http.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(resBody, &showsWatched)
	if err != nil {
		return nil, err
	}

	return showsWatched, nil
}

func (c *Client) GetSeasonPremieres(startDate time.Time, days int) (seasonPremieres []SeasonPremieresDto, err error) {

	url := fmt.Sprintf(baseURL+"/calendars/all/shows/premieres/%v/%d?extended=full", startDate.Format("2006-01-02"),
		days)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("trakt-api-version", "2")
	req.Header.Add("trakt-api-key", c.apiKey)

	res, err := c.http.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(resBody, &seasonPremieres)
	if err != nil {
		return nil, err
	}

	return seasonPremieres, nil
}
