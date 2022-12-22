package shows

import (
	"github.com/labstack/echo/v4"
	"github.com/mqrc81/zeries/domain"
	"net/http"
)

func (uc *useCase) GetGenres() ([]domain.Genre, error) {

	genres, err := uc.genreRepository.FindAll()
	if err != nil {
		return []domain.Genre{}, echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return genres, nil
}
