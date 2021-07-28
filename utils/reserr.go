package utils

import "github.com/gin-gonic/gin"

func ResErr(ginctx *gin.Context, errcode int, err error) {
	ginctx.Writer.WriteHeader(errcode)
	ginctx.Writer.Write([]byte(err.Error()))
}
