package data

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/WeatherGod3218/mlc-project-template/internal/sentences"
)

// func replaceNullWithString(obj []map[string]interface{}) {
// 	for key, value := range obj {
// 		if value == nil {
// 			obj[key] = "null"
// 		} else if nested, ok := value.(map[string]interface{}); ok {
// 			replaceNullWithString(nested)
// 		}
// 	}
// }

func GetAllData(c *gin.Context) {
	var sentences = sentences.GetAllSentences()

	c.JSON(http.StatusOK, gin.H{
		"data": sentences,
	})
}

func Routes(r *gin.RouterGroup) {
	data := r.Group("/data")

	data.GET("/", GetAllData)
}
