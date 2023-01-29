package helper

import (
	"Restfull-api/model/domain"
	"Restfull-api/model/web"
)

func CategoryResponse(category domain.Category) web.CategoryResponse {
	return web.CategoryResponse{
		Id:   category.Id,
		Name: category.Name,
	}
}
