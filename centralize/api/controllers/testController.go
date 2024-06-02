package controllers

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/nats-io/nats.go"
	"natsMicros/contracts/masterData/queries/provinceQueries/getProvinces"
	"natsMicros/contracts/masterData/responses"
	"net/http"
	"time"
)

type TestController struct {
	NatsConnection *nats.Conn
}

func NewTestController(echo *echo.Echo, conn *nats.Conn) *TestController {
	controller := &TestController{NatsConnection: conn}
	echo.POST("/testNats", controller.TestPostNats)
	return controller
}

func (testController *TestController) TestPostNats(c echo.Context) error {
	var getProvinceQuery getProvinces.GetProvincesQuery
	err := c.Bind(&getProvinceQuery)
	if err != nil {
		return err
	}
	query, err := json.Marshal(getProvinceQuery)
	response, err := testController.NatsConnection.Request("masterData.GetProvincesQuery", query, 10*time.Second)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	provinceResponse := responses.ProvinceResponse{}
	err = json.Unmarshal(response.Data, &provinceResponse)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, provinceResponse)
}
