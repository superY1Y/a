package model

import (
	"gorm.io/gorm"
)

// 学生表
type Student struct {
	gorm.Model
	Sno        string `json:"sno"`
	Password   string `json:"password"`
	Sname      string `gorm:"type:nvarchar(8)"`
	Ssex       string `gorm:"type:nchar(1)"`
	Sage_month string `gorm:"type:text"`
	Sage_day   string `gorm:"type:text"`
	Sage_year  string `gorm:"type:text"`
	Sdept      string `gorm:"nvarchar(30)"`
}

// 课程表
type Post struct {
	gorm.Model
	Cno     string `gorm:"type:char(6)"`
	Cname   string `gorm:"type:nvarchar(50)"`
	Cpno    string `gorm:"type:char(6)"`
	Ccredit string `gorm:"type:smallint(2)"`
}

// 教师表
type Teacher struct {
	gorm.Model
	Tno        string `gorm:"type:char(8)"`
	Tname      string `gorm:"type:nvarchar(8)"`
	Tsex       string `gorm:"type:nchar(1)"`
	Teb        string `gorm:"type:nvarchar(10)"`
	Tpt        string `gorm:"type:nvarchar(10)"`
	Sage_month string `gorm:"type:text"`
	Sage_day   string `gorm:"type:text"`
	Sage_year  string `gorm:"type:text"`
}

// 院系信息表
type Department struct {
	gorm.Model
	Dno        string `gorm:"type:char(3)"`
	Dname      string `gorm:"type:nvarchar(30)"`
	Dmanagerno string `gorm:"type:char(8)"`
}

// 选课信息表
type Sct struct {
	gorm.Model
	Sno   string `gorm:"type:char(9)"`
	Cno   string `gorm:"type:char(6)"`
	Tno   string `gorm:"type:char(8)"`
	Grade string `gorm:"type:int(4)"`
}
