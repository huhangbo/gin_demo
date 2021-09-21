package main

import (
	"github.com/gin-gonic/gin"
)


func main() {
	r := gin.Default()
	r = CollectRouter(r)
	panic(r.Run())
}
