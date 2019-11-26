package helper

import "github.com/gin-gonic/gin"

var middleWareCollection []gin.HandlerFunc = make([]gin.HandlerFunc, 0)

// Deprecated
func AddMiddleWare(middle gin.HandlerFunc) {
	middleWareCollection = append(middleWareCollection, middle)
}

func AddMiddleWares(middle ...gin.HandlerFunc) {
	middleWareCollection = append(middleWareCollection, middle...)
}

func GetAllMiddleWares() []gin.HandlerFunc {
	return middleWareCollection
}
