package rest

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/sklinkert/go-ddd/internal/application/interfaces"
	"github.com/sklinkert/go-ddd/internal/interface/api/rest/request"
	"net/http"
)

type SellerController struct {
	service interfaces.SellerService
}

func NewSellerController(e *echo.Echo, service interfaces.SellerService) *SellerController {
	controller := &SellerController{
		service: service,
	}

	e.POST("/sellers", controller.CreateSellerController)
	e.GET("/sellers", controller.GetAllSellersController)
	e.GET("/sellers/:id", controller.GetSellerByIdController)

	return controller
}

func (sc *SellerController) CreateSellerController(c echo.Context) error {
	var createSellerRequest request.CreateSellerRequest

	if err := c.Bind(&createSellerRequest); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Failed to parse request body",
		})
	}

	sellerCommand, err := createSellerRequest.ToCreateSellerCommand()
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid seller ID format",
		})
	}

	commandResult, err := sc.service.CreateSeller(sellerCommand)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to create seller",
		})
	}

	return c.JSON(http.StatusCreated, commandResult.Result)
}

func (sc *SellerController) GetAllSellersController(c echo.Context) error {
	sellers, err := sc.service.FindAllSellers()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to fetch sellers",
		})
	}

	return c.JSON(http.StatusOK, sellers)
}

func (sc *SellerController) GetSellerByIdController(c echo.Context) error {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid seller ID format",
		})
	}

	seller, err := sc.service.FindSellerById(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to fetch seller",
		})
	}

	if seller == nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": "Seller not found",
		})
	}

	return c.JSON(http.StatusOK, seller)
}
