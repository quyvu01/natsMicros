package controllers

import (
	"github.com/labstack/echo/v4"
	"natsMicros/buildingBlocks/application/abstractions"
	commonResponse "natsMicros/buildingBlocks/application/responses"
	"natsMicros/contracts/masterData/queries/provinceQueries/getProvinces"
	"natsMicros/contracts/masterData/responses"
	"net/http"
)

type MasterDataController struct {
	getProvinceRequest abstractions.IMessageRequest[getProvinces.GetProvincesQuery, commonResponse.PaginationResponse[responses.ProvinceResponse]]
}

// NewMasterDataController Todo: using mediator
func NewMasterDataController(echo *echo.Echo, getProvinceRequest abstractions.IMessageRequest[getProvinces.GetProvincesQuery, commonResponse.PaginationResponse[responses.ProvinceResponse]]) *MasterDataController {
	controller := &MasterDataController{getProvinceRequest: getProvinceRequest}
	echo.POST("api/masterData/getProvinces", controller.getProvinces)
	return controller
}

// List Provinces
// @Description Get list of provinces
// @Summary List Provinces
// @Tags masterData
// @Router /api/masterData/getProvinces [get]
// @Accept json
// @Produce json
// @Param SearchKey query string false "Searching for province name"
// @Param PageSize query int32 false "Page Sizing"
// @Param PageIndex query int32 false "Page Index"
func (controller *MasterDataController) getProvinces(c echo.Context) error {
	getProvinceQuery := &getProvinces.GetProvincesQuery{}
	if err := c.Bind(getProvinceQuery); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	response, err := controller.getProvinceRequest.Request(*getProvinceQuery)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, response)
}
