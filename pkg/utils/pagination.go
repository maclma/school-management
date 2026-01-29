package utils

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

type Pagination struct {
	Page   int
	Limit  int
	Offset int
	SortBy string
	Order  string
}

func GetPagination(c *gin.Context) Pagination {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}

	sortBy := c.DefaultQuery("sort_by", "id")
	order := c.DefaultQuery("order", "asc")

	if order != "asc" && order != "desc" {
		order = "asc"
	}

	offset := (page - 1) * limit

	return Pagination{
		Page:   page,
		Limit:  limit,
		Offset: offset,
		SortBy: sortBy,
		Order:  order,
	}
}
