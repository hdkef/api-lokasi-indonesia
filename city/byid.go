package city

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type byID struct {
}

func (b *byID) FromProvince(value *string, ginctx *gin.Context) {
	fmt.Println(*value)
}
