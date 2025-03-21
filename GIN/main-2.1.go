package main

import "github.com/gin-gonic/gin"
import (
	//"time"
	//"net/http"
)
//2.1 
func main(){
	router :=gin.Default() 
	router.GET("/getData", getData)
	router.GET("/getQueryString", getQueryString)
	router.GET("/getUrlData/:name/:age", getUrlData)

	//1)custom http config properties (L12-19)
	server := &http.Server {
		Addr: ":9091",
		Handler: router,
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	server.ListenAndServe()
}

// 2) for route grouping in gin, the lines (L26- 45) of code are meant to be uncommented.----------------------------------------------------
/*
func main(){
	router := gin.Default()
	
	router.GET("/getUrlData/:name/:age", getUrlData)


	//admin group
	admin := router.Group("/admin")
	{
		admin.GET("/getData", getData)
	}

	//client group
	client := router.Group("/client")
	{
		client.GET("/getQueryString", getQueryString)

	}
}
*/
// When running the above, the url to access getQueryString: http://localhost:9091/client/getQueryString and similarly for getData. 

//-----------------------------------------------------------------------------------------------------------------------------------------

//From this part till the end of the code, it is taken from main-1.go 
//GET func
func getData(c *gin.Context){
	c.JSON(200, gin.H{
		"data": "Hi, I am GIN Framework",
	})
}
//from main-1.go
func getQueryString(c *gin.Context){
	name := c.Query("name")
	age := c.Query("age")
	c.JSON(200, gin.H{
		"data": "Hi, I am get Query String Framework",
		"name": name,
		"age": age,
	})
}
// soo for the above function, we are direcctly pasting line 45 URL in google.

//Incase in the above section, they do not give URL and tell u to fetch data from URL param, like given below. 
// http://localhost:8080/getUrlData/Budbag/73

func getUrlData(c *gin.Context){
	name := c.Param("name")
	age := c.Param("age")
	c.JSON(200, gin.H{
		"data": "Hi, I am Get URL method",
		"name": name,
		"age": age,
	})
}