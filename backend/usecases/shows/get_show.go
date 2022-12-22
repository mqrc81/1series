package shows

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mqrc81/zeries/domain"
)

func (uc *useCase) GetShow(showId int) (domain.Show, error) {

	tmdbShow, err := uc.tmdbClient.GetTVDetails(showId, nil)
	if err != nil {
		return domain.Show{}, echo.NewHTTPError(http.StatusConflict, err.Error())
	}

	return showFromTmdbShow(tmdbShow), nil
}
