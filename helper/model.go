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

func ToCategoryResponses(categories []domain.Category) []web.CategoryResponse {
	var categoryResponses []web.CategoryResponse

	for _, category := range categories {
		categoryResponses = append(categoryResponses, CategoryResponse(category))
	}

	return categoryResponses
}
