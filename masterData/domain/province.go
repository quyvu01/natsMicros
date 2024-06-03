package domain

import "natsMicros/contracts/masterData/modelIds"

type Province struct {
	Id   modelIds.ProvinceId
	Name string `json:"name"`
}
