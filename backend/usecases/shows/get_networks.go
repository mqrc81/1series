package shows

import (
	"github.com/labstack/echo/v4"
	"github.com/mqrc81/zeries/domain"
	"net/http"
)

func (uc *useCase) GetNetworks() ([]domain.Network, error) {

	networks, err := uc.networkRepository.FindAll()
	if err != nil {
		return []domain.Network{}, echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return networks, nil
}
