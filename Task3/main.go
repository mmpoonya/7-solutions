package main

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type Request struct {
	Data string `json:"data"`
}
type Response struct {
	Beef map[string]int `json:"beef"`
}

func main() {
	router := gin.Default()
	router.GET("/beef/summary", getBeefSummary)
	router.Run("localhost:8080")
}

func getBeefSummary(c *gin.Context) {
	req := Request{}
	if err := c.BindJSON(&req); err != nil {
		return
	}
	rsp := Response{
		Beef: getBeef(req.Data),
	}
	c.JSON(http.StatusOK, rsp)
}

func getBeef(data string) map[string]int {
	result := make(map[string]int)
	for i := 0; i < len(data); i++ {
		if isIgnoreChar(data[i]) {
			continue
		}
		start := i
		for j := i; j <= len(data); j++ {
			if j == len(data) || isIgnoreChar(data[j]) {
				end := j
				beef := strings.ToLower(data[start:end])
				if _, has := result[beef]; !has {
					result[beef] = 0
				}
				result[beef] = result[beef] + 1
				i = j
				break
			}
		}
	}
	return result
}

func isIgnoreChar(c byte) bool {
	return c == ',' || c == '.' || c == ' '
}
