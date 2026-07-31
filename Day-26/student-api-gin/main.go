package main

import "github.com/gin-gonic/gin"

func main() {

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.String(200, "Home Page")
	})

	router.GET("/about", func(c *gin.Context) {
		c.String(200, "About Page")
	})

	router.GET("/contact", func(c *gin.Context) {
		c.String(200, "Contact Page")
	})
	router.GET("/students/:id", func(c *gin.Context) {
		id := c.Param("id")
		c.String(200, "Student ID: "+id)
	})
	router.GET("/students/:id/:courses/:courseId", func(c *gin.Context) {
		studentID := c.Param("id")
		courseID := c.Param("courseId")
		course := c.Param("courses")
		c.String(200, "Student "+studentID+" "+course+" "+courseID)
	})
	router.GET("/students", func(c *gin.Context) {
		page := c.Query("page")
		limit := c.Query("limit")
		c.String(200, "Page: "+page+" Limit: "+limit)
	})
	router.Run(":8080")

}
