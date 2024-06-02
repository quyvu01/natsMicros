package responses

import (
	"natsMicros/buildingBlocks/application/responses"
)

type ProvinceResponse struct {
	responses.ModelResponse
	Name string `json:"name"`
}
