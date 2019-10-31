package helper

import "github.com/gin-gonic/gin"

var middleWareCollection []gin.HandlerFunc = make([]gin.HandlerFunc, 0)

func AddMiddleWare(middle gin.HandlerFunc) {
	middleWareCollection = append(middleWareCollection, middle)
}

func GetAllMiddleWares() []gin.HandlerFunc {
	return middleWareCollection
}
