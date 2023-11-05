package controller

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"main.go/dao"
	"main.go/model"
)

func GoLogin(c *gin.Context) {
	c.HTML(200, "login.html", nil)
}

func Login(c *gin.Context) {
	sno := c.PostForm("sno")
	password := c.PostForm("password")
	fmt.Println(sno)
	u := dao.Mgr.Login(sno)

	if u.Sno == "" {
		fmt.Println("用户名不存在")
		c.HTML(200, "login.html", "用户不存在")
	} else {
		if u.Password != password {
			fmt.Println("密码错误")
			c.HTML(200, "login.html", "密码错误")
		} else {
			fmt.Println("登陆成功！")
			time.Sleep(2 * time.Second)
			c.Redirect(301, "/post_index")
		}
	}

	fmt.Println("登录成功！")
	c.Redirect(http.StatusMovedPermanently, "/post_index")
}

//学生信息操作

// 获取系统时间计算年龄
//
//	func calculateAge(birthDate time.Time) int {
//		now := time.Now()
//		age := now.Year() - birthDate.Year()
//		if now.Month() < birthDate.Month() || (now.Month() == birthDate.Month() && now.Day() < birthDate.Day()) {
//			age--
//		}
//		return age
//	}
//
// -----------------------------------------------------------------------------------------------------------------
func AddStudent(c *gin.Context) {
	sno := c.PostForm("sno")
	password := c.PostForm("password")
	sname := c.PostForm("sname")
	ssex := c.PostForm("ssex")
	sage_year := c.PostForm("sage_year")
	sage_month := c.PostForm("sage_month")
	sage_day := c.PostForm("sage_day")
	sdept := c.PostForm("sdept")

	// year, _ := strconv.Atoi(sage_year) //强转为整形
	// month, _ := strconv.Atoi(sage_month)
	// day, _ := strconv.Atoi(sage_day)
	// birthDate := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
	// sage := strconv.Itoa(calculateAge(birthDate))
	student := model.Student{
		Sno:        sno,
		Password:   password,
		Sname:      sname,
		Ssex:       ssex,
		Sage_year:  sage_year,
		Sage_month: sage_month,
		Sage_day:   sage_day,
		Sdept:      sdept,
		// Sage:       sage,
	}
	dao.Mgr.AddStudent(&student)

	c.Redirect(302, "/student")
}

func UpdateStudent(c *gin.Context) {
	sno := c.PostForm("sno")
	student := dao.Mgr.GetStudent(sno) // 根据学号查询学生信息

	if student != nil {
		// 根据表单中的其他字段更新学生信息
		student.Password = c.PostForm("password")
		student.Sname = c.PostForm("sname")
		student.Ssex = c.PostForm("ssex")
		student.Sage_year = c.PostForm("sage_year")
		student.Sage_month = c.PostForm("sage_month")
		student.Sage_day = c.PostForm("sage_day")
		student.Sdept = c.PostForm("sdept")

		dao.Mgr.UpdateStudent(student) // 更新学生信息

		fmt.Println("更新成功")
		c.Redirect(302, "/student_update")
	} else {
		panic("查找失败")
	}
}

func DeleteStudent(c *gin.Context) {
	sno := c.PostForm("sno")
	student := dao.Mgr.GetStudent(sno)
	dao.Mgr.DeleteStudent(student)
	if student.Sno == "" {
		fmt.Println("用户名不存在")
		c.HTML(http.StatusOK, "index.html", gin.H{
			"errorI": "用户名不存在",
		})
		return
	} else {
		fmt.Println("删除成功！")
		c.HTML(200, "/", "删除成功！")
		time.Sleep(1 * time.Second)
	}
}

func GetStudent(c *gin.Context) {
	sno := c.PostForm("sno")

	student := dao.Mgr.GetAllStudent(sno)

	c.HTML(http.StatusOK, "student_check.html", student)
}

func Getallstudents(c *gin.Context) {
	student := dao.Mgr.GetAllOfStudent()
	c.HTML(http.StatusOK, "ShowAll_students.html", student)
}

// --------------------------------------------------------------------------------------------------------------------
// 课程操作
func GetPostIndex(c *gin.Context) {
	posts := dao.Mgr.GetAllPost()
	c.HTML(200, "post_index.html", posts)
}

func AddPost(c *gin.Context) {
	cno := c.PostForm("cno")
	cname := c.PostForm("cname")
	cpno := c.PostForm("cpno")
	ccredit := c.PostForm("ccredit")

	post := model.Post{
		Cno:     cno,
		Cname:   cname,
		Cpno:    cpno,
		Ccredit: ccredit,
	}

	dao.Mgr.AddPost(&post)
	c.Redirect(302, "/post")
}

func UpdatePost(c *gin.Context) {
	cno := c.PostForm("cno")
	post := dao.Mgr.GetPost(cno) // 根据学号查询学生信息

	if post != nil {
		// 根据表单中的其他字段更新学生信息
		post.Cno = c.PostForm("cno")
		post.Cname = c.PostForm("cname")
		post.Cpno = c.PostForm("cpno")
		post.Ccredit = c.PostForm("ccredit")

		dao.Mgr.UpdatePost(post) // 更新学生信息

		fmt.Println("更新成功")
		c.Redirect(302, "/post_update")
	} else {
		panic("查找失败")
	}
}

func DeletePost(c *gin.Context) {
	cno := c.PostForm("cno")
	post := dao.Mgr.GetPost(cno)
	dao.Mgr.DeletePost(post)
	if post.Cno == "" {
		fmt.Println("课程号不存在")
		c.HTML(http.StatusOK, "index.html", gin.H{
			"errorI": "课程号不存在",
		})
		return
	} else {
		fmt.Println("删除成功！")
		c.HTML(200, "/", "删除成功！")
		time.Sleep(1 * time.Second)
	}
}

// --------------------------------------------------------------------------------------------------------------------
// 教师信息操作
func AddTeacher(c *gin.Context) {
	// 根据表单中的其他字段更新学生信息
	tno := c.PostForm("tno")
	tname := c.PostForm("tname")
	tsex := c.PostForm("tsex")
	sage_year := c.PostForm("sage_year")
	sage_month := c.PostForm("sage_month")
	sage_day := c.PostForm("sage_day")
	teb := c.PostForm("teb")
	tpt := c.PostForm("tpt")
	teacher := model.Teacher{
		Tno:        tno,
		Tname:      tname,
		Tsex:       tsex,
		Sage_year:  sage_year,
		Sage_month: sage_month,
		Sage_day:   sage_day,
		Teb:        teb,
		Tpt:        tpt,
	}

	dao.Mgr.AddTeacher(&teacher)

	fmt.Println("添加成功")
	c.Redirect(302, "/teacher_add")
}

func UpdateTeacher(c *gin.Context) {
	tno := c.PostForm("tno")
	teacher := dao.Mgr.GetTeacher(tno) // 根据学号查询学生信息

	if teacher != nil {
		// 根据表单中的其他字段更新学生信息
		teacher.Tno = c.PostForm("tno")
		teacher.Tname = c.PostForm("tname")
		teacher.Tsex = c.PostForm("tsex")
		teacher.Sage_year = c.PostForm("sage_year")
		teacher.Sage_month = c.PostForm("sage_month")
		teacher.Sage_day = c.PostForm("sage_day")
		teacher.Teb = c.PostForm("teb")
		teacher.Tpt = c.PostForm("tpt")

		dao.Mgr.UpdateTeacher(teacher) // 更新学生信息

		fmt.Println("更新成功")
		c.Redirect(302, "/teacher_update")
	} else {
		fmt.Println("查找失败")
	}
}

func DeleteTeacher(c *gin.Context) {
	tno := c.PostForm("tno")
	teacher := dao.Mgr.GetTeacher(tno)
	dao.Mgr.DeleteTeacher(teacher)
	if teacher.Tno == "" {
		fmt.Println("教工号不存在")
		c.HTML(http.StatusOK, "index.html", gin.H{
			"errorI": "教工号不存在",
		})
		return
	} else {
		fmt.Println("删除成功！")
		c.HTML(200, "/", "删除成功！")
		time.Sleep(1 * time.Second)
	}
}

func GetTeacher(c *gin.Context) {

	tno := c.PostForm("tno")

	teacher := dao.Mgr.GetAllTeacher(tno)

	c.HTML(http.StatusOK, "teacher_check.html", teacher)
}

func Getallteacher(c *gin.Context) {
	teachers := dao.Mgr.GetAllOfTeacher()
	c.HTML(200, "ShowAll_teachers.html", teachers)
}

// ---------------------------------------------------------------------------------------------------------------------
// 院系操作
func AddDepartment(c *gin.Context) {
	// 根据表单中的其他字段更新学生信息
	dno := c.PostForm("dno")
	dname := c.PostForm("dname")
	dmanagerno := c.PostForm("dmanagerno")

	department := model.Department{
		Dno:        dno,
		Dname:      dname,
		Dmanagerno: dmanagerno,
	}
	dao.Mgr.AddDepartment(&department)

	fmt.Println("添加成功")
	c.Redirect(302, "/department_add")
}

func UpdateDepartment(c *gin.Context) {
	dno := c.PostForm("dno")
	department := dao.Mgr.GetDepartment(dno)

	if department != nil {
		// 根据表单中的其他字段更新学生信息
		department.Dno = c.PostForm("dno")
		department.Dname = c.PostForm("dname")
		department.Dmanagerno = c.PostForm("dmanagerno")

		dao.Mgr.UpdateDepartment(department)

		fmt.Println("更新成功")
		c.Redirect(302, "/department_update")
	} else {
		panic("查找失败")
	}
}

func DeleteDepartment(c *gin.Context) {
	dno := c.PostForm("dno")
	department := dao.Mgr.GetDepartment(dno)

	if department != nil {
		// 根据表单中的其他字段更新学生信息
		department.Dno = c.PostForm("dno")
		department.Dname = c.PostForm("dname")
		department.Dmanagerno = c.PostForm("dmanagerno")

		dao.Mgr.DeleteDepartment(department)

		fmt.Println("删除成功")
		c.Redirect(302, "/index")
	} else {
		fmt.Println("查找失败")
	}
}

func GetDepartment(c *gin.Context) {
	dno := c.PostForm("dno")

	department := dao.Mgr.GetAllDepartment(dno)

	c.HTML(http.StatusOK, "department_check.html", department)
}

func Getalldepartment(c *gin.Context) {
	departments := dao.Mgr.GetAllOfDepartment()
	c.HTML(200, "ShowAll_departments.html", departments)
}

// ---------------------------------------------------------------------------------------------------------------------
// 选课操作
func AddSct(c *gin.Context) {
	cno := c.PostForm("cno")
	sct := dao.Mgr.GetSct(cno)

	if sct != nil {
		// 根据表单中的其他字段更新学生信息
		sct.Tno = c.PostForm("tno")
		sct.Cno = c.PostForm("cno")
		sct.Sno = c.PostForm("sno")

		dao.Mgr.AddSct(sct)

		fmt.Println("添加成功")
		c.Redirect(302, "/student_update")
	} else {
		panic("查找失败")
	}
}

func UpdateSct(c *gin.Context) {
	cno := c.PostForm("cno")
	sct := dao.Mgr.GetSct(cno)

	if sct != nil {
		// 根据表单中的其他字段更新学生信息
		sct.Tno = c.PostForm("tno")
		sct.Cno = c.PostForm("cno")
		sct.Sno = c.PostForm("sno")

		dao.Mgr.AddSct(sct)

		fmt.Println("添加成功")
		c.Redirect(302, "/student_update")
	} else {
		panic("查找失败")
	}
}

func DeleteSct(c *gin.Context) {
	cno := c.PostForm("cno")
	sct := dao.Mgr.GetSct(cno)

	if sct != nil {
		// 根据表单中的其他字段更新学生信息
		sct.Tno = c.PostForm("tno")
		sct.Cno = c.PostForm("cno")
		sct.Sno = c.PostForm("sno")

		dao.Mgr.AddSct(sct)

		fmt.Println("添加成功")
		c.Redirect(302, "/student_update")
	} else {
		panic("查找失败")
	}
}

func GetSct(c *gin.Context) {
	cno := c.PostForm("cno")

	sct := dao.Mgr.GetSct(cno)

	c.HTML(http.StatusOK, "teacher_check.html", sct)
}

// 跳转
func GoAddPost(c *gin.Context) {
	c.HTML(200, "post.html", nil)
}

func User(c *gin.Context) {
	c.HTML(200, "user.html", nil)
}

func ListUser(c *gin.Context) {
	c.HTML(200, "index.html", nil)
}

func GoRegister(c *gin.Context) {
	c.HTML(200, "register.html", nil)
}

func Index(c *gin.Context) {
	c.HTML(200, "index.html", nil)
}
