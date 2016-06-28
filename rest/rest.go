package rest

import (
	"github.com/gin-gonic/gin"
)

// All of the methods are the same type as HandlerFunc
// if you don't want to support any methods of CRUD, then don't implement it
type CreateSupported interface {
	CreateHandler(*gin.Context)
}
type ListSupported interface {
	ListHandler(*gin.Context)
}
type TakeSupported interface {
	TakeHandler(*gin.Context)
}
type UpdateSupported interface {
	UpdateHandler(*gin.Context)
}
type DeleteSupported interface {
	DeleteHandler(*gin.Context)
}

// It defines
//   POST: /path
//   GET:  /path
//   PUT:  /path/:id
//   POST: /path/:id
func CRUD(group *gin.RouterGroup, path string, resource interface{}, handlers ...gin.HandlerFunc) {
	if resource, ok := resource.(CreateSupported); ok {
		res := append([]gin.HandlerFunc{resource.CreateHandler}, handlers...)
		group.POST(path, res...)
	}
	if resource, ok := resource.(ListSupported); ok {
		res := append([]gin.HandlerFunc{resource.ListHandler}, handlers...)
		group.GET(path, res...)
	}
	if resource, ok := resource.(TakeSupported); ok {
		res := append([]gin.HandlerFunc{resource.TakeHandler}, handlers...)
		group.GET(path+"/:id", res...)
	}
	if resource, ok := resource.(UpdateSupported); ok {
		res := append([]gin.HandlerFunc{resource.UpdateHandler}, handlers...)
		group.PUT(path+"/:id", res...)
	}
	if resource, ok := resource.(DeleteSupported); ok {
		res := append([]gin.HandlerFunc{resource.DeleteHandler}, handlers...)
		group.DELETE(path+"/:id", res...)
	}
}
