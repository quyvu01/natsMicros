package controllers

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/nats-io/nats.go"
	commonResponse "natsMicros/buildingBlocks/application/responses"
	"natsMicros/contracts/masterData/commands/provinceCommands/createProvince"
	"natsMicros/contracts/masterData/queries/provinceQueries/getProvinces"
	"natsMicros/contracts/masterData/responses"
	"net/http"
	"time"
)

type MasterDataController struct {
	NatsConnection *nats.Conn
}

func NewMasterDataController(echo *echo.Echo, conn *nats.Conn) *MasterDataController {
	controller := &MasterDataController{NatsConnection: conn}
	echo.GET("/getProvinces", controller.GetProvinces)
	echo.POST("/createProvince", controller.CreateProvince)
	return controller
}

func (controller *MasterDataController) GetProvinces(c echo.Context) error {
	var getProvinceQuery getProvinces.GetProvincesQuery
	err := c.Bind(&getProvinceQuery)
	if err != nil {
		return err
	}
	query, err := json.Marshal(getProvinceQuery)
	response, err := controller.NatsConnection.Request("masterData.GetProvincesQuery", query, 10*time.Second)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	provincesResponse := commonResponse.PaginationResponse[responses.ProvinceResponse]{}
	err = json.Unmarshal(response.Data, &provincesResponse)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, provincesResponse)
}

func (controller *MasterDataController) CreateProvince(c echo.Context) error {
	var createProvinceCommand createProvince.CreateProvinceCommand
	err := c.Bind(&createProvinceCommand)
	if err != nil {
		return err
	}
	command, err := json.Marshal(createProvinceCommand)
	_, err = controller.NatsConnection.Request("masterData.CreateProvinceCommand", command, 10*time.Second)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	return c.NoContent(http.StatusOK)
}
