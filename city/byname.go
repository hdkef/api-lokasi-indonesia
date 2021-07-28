package city

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type byName struct {
}

func (b *byName) FromProvince(value *string, ginctx *gin.Context) {
	fmt.Println(*value)
}
