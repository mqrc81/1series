package trakt

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
)

const (
	baseURL         = "https://api.trakt.tv"
	showsWatchedURL = baseURL + "/shows/watched/weekly"
)

func (c *Client) GetShowsWatchedWeekly(page int, limit int) ([]ShowsWatchedDto, error) {

	req, err := http.NewRequest(http.MethodGet, showsWatchedURL, nil)
	if err != nil {
		return nil, err
	}
	req.URL.Query().Set("page", strconv.Itoa(page))
	req.URL.Query().Set("limit", strconv.Itoa(limit))

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

	var showsWatched []ShowsWatchedDto
	err = json.Unmarshal(resBody, &showsWatched)
	if err != nil {
		return nil, err
	}

	return showsWatched, nil
}
