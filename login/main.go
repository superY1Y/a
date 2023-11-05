package main

import (
	"fmt"
	"time"
)

func calculateAge(birthDate time.Time) int {
	now := time.Now()
	age := now.Year() - birthDate.Year()
	if now.Month() < birthDate.Month() || (now.Month() == birthDate.Month() && now.Day() < birthDate.Day()) {
		age--
	}
	return age
}

func main() {
	var year, month, day int

	fmt.Println("请输入您的出生年份：")
	fmt.Scanln(&year)

	fmt.Println("请输入您的出生月份：")
	fmt.Scanln(&month)

	fmt.Println("请输入您的出生日期：")
	fmt.Scanln(&day)

	birthDate := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
	age := calculateAge(birthDate)
	fmt.Printf("年龄：%d岁\n", age)
}
