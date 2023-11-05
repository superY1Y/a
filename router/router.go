package router

import (
	"github.com/gin-gonic/gin"
	"main.go/controller"
)

func Start() {
	r := gin.Default()
	r.LoadHTMLFiles("login.html")
	r.LoadHTMLGlob("templates/*")

	r.Static("/asset", "./asset") //访问该文件所有静态文件
	r.GET("/style.css", func(c *gin.Context) {
		c.File("./style.css")
	})
	r.GET("/style.scss", func(c *gin.Context) {
		c.File("./style.scss")
	})
	r.GET("/style.css.map", func(c *gin.Context) {
		c.File("./style.css.map")
	})

	r.GET("/index", controller.ListUser)
	r.GET("/", controller.Index)
	r.GET("/login", controller.GoLogin)
	r.POST("/login", controller.Login)
	r.GET("/post_index", controller.GetPostIndex)
	r.POST("/post", controller.AddPost)
	r.GET("/post", controller.GoAddPost)
	r.POST("/post_update", controller.UpdatePost)
	r.GET("/post_update", func(c *gin.Context) {
		c.HTML(200, "post_update.html", nil)
	})
	r.GET("/user", controller.User)

	r.POST("/student", controller.AddStudent)
	r.GET("/student", func(c *gin.Context) {
		c.HTML(200, "student.html", nil)
	})
	r.POST("/student_update", controller.UpdateStudent)
	r.GET("/student_update", func(c *gin.Context) {
		c.HTML(200, "student_update.html", nil)
	})
	r.POST("/delete", controller.DeleteStudent)
	r.GET("/delete", func(c *gin.Context) {
		c.HTML(200, "header.html", nil)
	})
	r.POST("/search_student", controller.GetStudent)

	r.POST("/teacher_add", controller.AddTeacher)
	r.GET("/teacher_add", func(c *gin.Context) {
		c.HTML(200, "teacher_add.html", nil)
	})
	r.POST("/teacher_update", controller.UpdateTeacher)
	r.GET("/teacher_update", func(c *gin.Context) {
		c.HTML(200, "teacher_update.html", nil)
	})
	r.POST("/teacher_delete", controller.DeleteTeacher)
	r.GET("/teacher_delete", func(c *gin.Context) {
		c.HTML(200, "header.html", nil)
	})
	r.POST("/teacher_check", controller.GetTeacher)

	r.POST("/department_add", controller.AddDepartment)
	r.GET("/department_add", func(c *gin.Context) {
		c.HTML(200, "department_add.html", nil)
	})
	r.POST("/department_update", controller.UpdateDepartment)
	r.GET("/department_update", func(c *gin.Context) {
		c.HTML(200, "department_update.html", nil)
	})
	r.POST("/department_delete", controller.DeleteDepartment)
	r.GET("/department_delete", func(c *gin.Context) {
		c.HTML(200, "header.html", nil)
	})
	r.POST("/department_check", controller.GetDepartment)

	r.GET("/ShowAll_students", controller.Getallstudents)
	r.GET("/ShowAll_teachers", controller.Getallteacher)
	r.GET("/ShowAll_departments", controller.Getalldepartment)

	r.Run(":8080")
}
