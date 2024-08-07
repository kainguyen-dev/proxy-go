package utils

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"svc/proxy-service/internal/common"
)

func ParseId(c *gin.Context, param string) int {
	id, err := strconv.Atoi(c.Param(param))
	if err != nil {
		panic(common.ClientError{Code: 400, Message: "Invalid format param " + param})
	}
	return id
}
