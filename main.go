package main

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

func main() {
	router := gin.New()
	router.GET("/getData", getData)
	router.GET("/getQueryString", getQueryString)
	router.GET("/getUrlData/:name/:age/:address", getUrlData)
	router.POST("/getDataPost", getDataPost)
	router.Run()
}

func getData(c *gin.Context) {

	c.JSON(200, gin.H{

		"data": "Hi I am GIN Framework",
	})

}

// http://localhost:8080/getUrlData/Mark/30
func getUrlData(c *gin.Context) {

	name := c.Param("name")
	age := c.Param("age")
	address := c.Param("phnom penh")

	c.JSON(200, gin.H{

		"data":    "Hi I am a getUrlMethod method ",
		"name":    name,
		"age":     age,
		"address": address,
	})
}

// http://localhost:8080/getQueryString?name=Mark&age=30
func getQueryString(c *gin.Context) {

	name := c.Query("name")
	age := c.Query("age")

	c.JSON(200, gin.H{

		"data": "Hi I am a query method ",
		"name": name,
		"age":  age,
	})

}

func getDataPost(c *gin.Context) {

	body := c.Request.Body
	value, _ := ioutil.ReadAll(body)

	c.JSON(200, gin.H{

		"data":     "Hi I am post method from GIN Framework",
		"bodyData": string(value),
	})

}

/*

What is gin ?
How to use intall gin ?
Gin basic
Handle query and string with GIN
Hanle Url parameters with GIN

*/
