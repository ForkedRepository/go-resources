// test_gin
package main

import (
	"fmt"
	"github.com/bitly/go-simplejson"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"path"
	"reflect"
)

func main() {
	router := gin.Default()

	router.GET("/data/:resource", func(c *gin.Context) {
		res := c.Params.ByName("resource")
		m, err := simplejson.NewJson(getMissions(res))
		if err != nil {
			log.Fatal(err)
		}
		c.JSON(http.StatusOK, m)
	})
	router.Run(":8080")
}

func getMissions(res string) []byte {
	abs, err := os.Getwd()
	dir := path.Join(abs, "xploration-data", res+".json")
	file, err := os.Open(dir) // For read access.
	if err != nil {
		log.Fatal(err)
	}
	data := make([]byte, 9999)
	count, err := file.Read(data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("read %d bytes: %q\n", count, data[:count])
	fmt.Println(reflect.TypeOf(data))
	fmt.Println(dir)

	return data[:count]

}
