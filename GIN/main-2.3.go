//2nd video 3rd part
package main

import "github.com/gin-gonic/gin"
func main(){
	router := gin.Default()

	//authentication grp
	auth := gin.BasicAuth(gin.Accounts{
		"user": "pass",
		"user2": "pass2",
		"user3": "pass3",
	})
	// From here it is same as the code of 2) Route grouping in gin.
	router.GET("/getUrlData/:name/:age", getUrlData)


	//admin group
	admin := router.Group("/admin", auth) // Yaha pe auth lag raha hai user ko.
	{
		admin.GET("/getData", getData) 
	}

	//client group
	client := router.Group("/client")
	{
		client.GET("/getQueryString", getQueryString)

	}
	router.Run(":9091")
}

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