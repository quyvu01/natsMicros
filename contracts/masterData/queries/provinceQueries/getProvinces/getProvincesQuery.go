package getProvinces

import "natsMicros/buildingBlocks/application/queries"

type GetProvincesQuery struct {
	SearchKey string
	queries.GetManyQuery
}
