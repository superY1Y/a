package dao

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"main.go/model"
)

var Mgr Manager

type manager struct {
	bd *gorm.DB
}

type Manager interface {
	//添加学生
	AddStudent(student *model.Student)
	UpdateStudent(student *model.Student)
	DeleteStudent(student *model.Student)
	GetStudent(sno string) *model.Student
	Login(username string) *model.Student
	GetAllStudent(sno string) []model.Student
	GetAllOfStudent() []model.Student
	//添加课程操作
	AddPost(post *model.Post)
	GetAllPost() []model.Post
	getPost(pid int) model.Post
	GetPost(cno string) *model.Post
	UpdatePost(post *model.Post)
	DeletePost(post *model.Post)

	//添加教师
	AddTeacher(teacher *model.Teacher)
	UpdateTeacher(teacher *model.Teacher)
	DeleteTeacher(teacher *model.Teacher)
	GetTeacher(tno string) *model.Teacher
	GetAllTeacher(tno string) []model.Teacher
	GetAllOfTeacher() []model.Teacher

	//添加院系
	AddDepartment(department *model.Department)
	UpdateDepartment(department *model.Department)
	DeleteDepartment(department *model.Department)
	GetDepartment(dno string) *model.Department
	GetAllDepartment(dno string) []model.Department
	GetAllOfDepartment() []model.Department

	//添加选课
	AddSct(sct *model.Sct)
	UpdateSct(sct *model.Sct)
	DeleteSct(sct *model.Sct)
	GetSct(cno string) *model.Sct
}

func init() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/subjectoptions?charset=utf8mb4&parseTime=True&Local"
	bd, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect database :", err)
	}
	Mgr = &manager{bd: bd}
	bd.AutoMigrate(&model.Student{})
	bd.AutoMigrate(&model.Post{})
	bd.AutoMigrate(&model.Teacher{})
	bd.AutoMigrate(&model.Department{})
	bd.AutoMigrate(&model.Sct{}) //自动创建表
}

// 学生操作
func (mgr *manager) AddStudent(student *model.Student) {
	mgr.bd.Create(student)

}

func (mgr *manager) UpdateStudent(student *model.Student) {
	mgr.bd.Model(&model.Student{}).Where("sno = ?", student.Sno).Updates(student)
}

func (mgr *manager) DeleteStudent(student *model.Student) {
	mgr.bd.Model(&model.Student{}).Where("sno = ?", student.Sno).Delete(student)
}

func (mgr *manager) GetStudent(sno string) *model.Student {
	var student model.Student
	mgr.bd.Where("sno=?", sno).First(&student)
	fmt.Println(sno)
	return &student
}

func (mgr *manager) GetAllStudent(sno string) []model.Student {
	var students []model.Student
	mgr.bd.Where("sno = ?", sno).Find(&students)
	return students
}

func (mgr *manager) GetAllOfStudent() []model.Student {
	var students = make([]model.Student, 10)
	mgr.bd.Find(&students)
	return students
}

// 添加课程
func (mgr *manager) AddPost(post *model.Post) {
	mgr.bd.Create(post)
}

func (mgr *manager) GetAllPost() []model.Post {
	var posts = make([]model.Post, 10)
	mgr.bd.Find(&posts)
	return posts
}

func (mgr *manager) getPost(pid int) model.Post {
	var post model.Post
	mgr.bd.First(&post, pid)
	return post
}

func (mgr *manager) UpdatePost(post *model.Post) {
	mgr.bd.Model(&model.Post{}).Where("cno = ?", post.Cno).Updates(post)
}

func (mgr *manager) GetPost(cno string) *model.Post {
	var post model.Post
	mgr.bd.Where("cno=?", cno).First(&post)
	return &post
}

func (mgr *manager) DeletePost(post *model.Post) {
	mgr.bd.Model(&model.Post{}).Where("cno = ?", post.Cno).Delete(post)
}

// 教师
func (mgr *manager) AddTeacher(teacher *model.Teacher) {
	mgr.bd.Create(teacher)
}

func (mgr *manager) UpdateTeacher(teacher *model.Teacher) {
	mgr.bd.Model(&model.Teacher{}).Where("tno = ?", teacher.Tno).Updates(teacher)
}

func (mgr *manager) DeleteTeacher(teacher *model.Teacher) {
	mgr.bd.Model(&model.Teacher{}).Where("tno = ?", teacher.Tno).Delete(teacher)
}

func (mgr *manager) GetTeacher(tno string) *model.Teacher {
	var teacher model.Teacher
	mgr.bd.Where("tno=?", tno).First(&teacher)
	fmt.Println(tno)
	return &teacher
}

func (mgr *manager) GetAllTeacher(tno string) []model.Teacher {
	var teachers []model.Teacher
	mgr.bd.Where("tno = ?", tno).Find(&teachers)
	return teachers
}

func (mgr *manager) GetAllOfTeacher() []model.Teacher {
	var teachers = make([]model.Teacher, 10)
	mgr.bd.Find(&teachers)
	return teachers
}

// 院系
func (mgr *manager) AddDepartment(department *model.Department) {
	mgr.bd.Create(department)

}

func (mgr *manager) UpdateDepartment(department *model.Department) {
	mgr.bd.Model(&model.Department{}).Where("dno = ?", department.Dno).Updates(department)
}

func (mgr *manager) DeleteDepartment(department *model.Department) {
	mgr.bd.Model(&model.Department{}).Where("tno = ?", department.Dno).Delete(department)
}

func (mgr *manager) GetDepartment(dno string) *model.Department {
	var department model.Department
	mgr.bd.Where("dno=?", dno).First(&department)
	fmt.Println(dno)
	return &department
}
func (mgr *manager) GetAllDepartment(dno string) []model.Department {
	var departments []model.Department
	mgr.bd.Where("dno = ?", dno).Find(&departments)
	return departments
}

func (mgr *manager) GetAllOfDepartment() []model.Department {
	var departments = make([]model.Department, 10)
	mgr.bd.Find(&departments)
	return departments
}

// 选课
func (mgr *manager) AddSct(sct *model.Sct) {
	mgr.bd.Create(sct)

}

func (mgr *manager) UpdateSct(sct *model.Sct) {
	mgr.bd.Model(&model.Sct{}).Where("cno = ?", sct.Cno).Updates(sct)
}

func (mgr *manager) DeleteSct(sct *model.Sct) {
	mgr.bd.Model(&model.Sct{}).Where("cno = ?", sct.Cno).Delete(sct)
}

func (mgr *manager) GetSct(cno string) *model.Sct {
	var sct model.Sct
	mgr.bd.Where("cno=?", cno).First(&sct)
	fmt.Println(cno)
	return &sct
}

// 登录
func (mgr *manager) Login(sno string) *model.Student {
	var student model.Student
	mgr.bd.Where("sno=?", sno).First(&student)
	return &student
}
