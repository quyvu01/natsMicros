package domain

import "natsMicros/contracts/masterData/modelIds"

type Province struct {
	Id   modelIds.ProvinceId `bson:"id"`
	Name string              `json:"name"`
}
