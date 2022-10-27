package specification

import (
	"github.com/jinzhu/gorm"
)

//PaginationStruct stores return info
type PaginationStruct struct {
	TableName string
	PageSize  int32
	Index     int32
}

//PaginationStruct stores return info
type PaginationInfo struct {
	CurrentPage  int32
	NextPage     int32
	PreviousPage int32
	TotalPages   int32
	TotalRows    int32
}

// NewPaginationSpecification generate pagination query
func NewPaginationSpecification(pagination *PaginationStruct) func(db *gorm.DB) *gorm.DB {
	return func(tx *gorm.DB) *gorm.DB {
		return tx.Offset(pagination.Index).Limit(pagination.PageSize)
	}
}

//SetPaginationDto set pagination dto
func SetWalletPaginationDto(page int32, pageSize int32, count int32, index int32) *PaginationInfo {

	var nextPage, previousPage, totalPages int32
	//fmt.Println("pageSize", pageSize)
	if index+pageSize < count {
		nextPage = page + 1
	} else {
		nextPage = page
	}

	if page > 1 {
		previousPage = page - 1
	} else {
		previousPage = page
	}

	if count%pageSize != 0 {
		totalPages = count/pageSize + 1
	} else {
		totalPages = count / pageSize
	}

	return &PaginationInfo{
		CurrentPage:  page,
		NextPage:     nextPage,
		PreviousPage: previousPage,
		TotalPages:   totalPages,
		TotalRows:    count,
	}
}
