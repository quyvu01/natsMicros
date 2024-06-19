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
	echo.POST("/getProvinces", controller.getProvinces)
	return controller
}

func (controller *MasterDataController) getProvinces(c echo.Context) error {
	var getProvinceQuery getProvinces.GetProvincesQuery
	err := c.Bind(&getProvinceQuery)
	if err != nil {
		return err
	}
	response, err := controller.getProvinceRequest.Request(getProvinceQuery)
	return c.JSON(http.StatusOK, response)
}
