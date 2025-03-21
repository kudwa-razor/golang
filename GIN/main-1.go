//VIDEO 1 
/* Includes:
- What is GIN?
-how to install/use GIN
-GIN basics
-Routing with GIN i.e GET and POST
- Handle Query string with GIN
- Handle URL parameters with GIN
*/

package main

import "github.com/gin-gonic/gin"
import "io" // line 35 in getDataPost has io.ReadAll(body) -> it was ioutil earlier, but my version is newer.

func main(){
	router :=gin.Default() //gin.Default is a gin logger -> bcuz of which we are getting in terminal the 200| 133.33 ms| ::1| GET| /getData.

	// if we use gin.New() -> logger is not recorded i.e logs not recorded. 

/*	router.GET("/getData", func(c *gin.Context){
		c.JSON(200, gin.H{
			"data": "Hi, I am GIN Framework",
		})
	})
	router.Run()
}*/
	//better way to write the above commented code, but extra POST, 
	router.GET("/getData", getData)
	router.GET("/getQueryString", getQueryString)
	router.GET("/getUrlData/:name/:age", getUrlData) // in this URL data, we've to tell ki there are 2 dynamic param: name, age with value Budbag and 73.
	// the above :name and :age will internally create key-value pair in GIN, so we can get name and age in the func using Param.
	router.POST("/getDataPost", getDataPost)
	

	router.Run()

}
//GET func
func getData(c *gin.Context){
	c.JSON(200, gin.H{
		"data": "Hi, I am GIN Framework",
	})
}
//POST func
func getDataPost(c *gin.Context){
	//Reading the data into body as a string ig.
	body := c.Request.Body
	value, _ := io.ReadAll(body)
	c.JSON(200, gin.H{
		"data": 	"Hi, I am POST method from GIN Framework",
		"bodyData": string(value),
	})
}
//Handle query string with gin - usually a URL.
// In the URL, anything after '?' is a string and queries can be seperated by &
// URL: http://localhost:8080/getQueryString?name=Budbag&age=73

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