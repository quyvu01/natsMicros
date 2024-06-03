package modelIds

import "github.com/google/uuid"

type ProvinceId struct {
	Id uuid.UUID
}

func NewProvinceId(id uuid.UUID) ProvinceId { return ProvinceId{Id: id} }
func (x ProvinceId) String() string         { return x.Id.String() }
