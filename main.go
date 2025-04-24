package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// กำหนดโครงสร้างข้อมูล User
// struct นี้ใช้สำหรับเก็บข้อมูลผู้ใช้ เช่น ชื่อ อีเมล และอายุ
// struct tag ใช้สำหรับกำหนดรูปแบบ JSON
// เช่น `json:"name"` หมายถึงฟิลด์นี้จะถูกแปลงเป็น key "name" ใน JSON
type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

// ฟังก์ชัน updateUserEmail ใช้สำหรับอัปเดตค่า email ของ struct User
// input: user (pointer ไปยัง struct User), newEmail (string ที่เป็น email ใหม่)
// output: ไม่มี output แต่จะเปลี่ยนค่า email ใน struct User โดยตรง
func updateUserEmail(user *User, newEmail string) {
	user.Email = newEmail
}

// ฟังก์ชัน validateUser ใช้สำหรับตรวจสอบความถูกต้องของ struct User
// input: user (struct User ที่ต้องการตรวจสอบ)
// output: error (ถ้าข้อมูลไม่ถูกต้องจะ return error, ถ้าถูกต้องจะ return nil)
// - ตรวจสอบว่า email ไม่เป็นค่าว่าง
// - ตรวจสอบว่า age เป็นจำนวนเต็มบวก
func validateUser(user User) error {
	if user.Email == "" {
		return fmt.Errorf("email cannot be empty")
	}
	if user.Age <= 0 {
		return fmt.Errorf("age must be a positive integer")
	}
	return nil
}

// โครงสร้างข้อมูลสำหรับรับ JSON จากผู้ใช้
// ใช้ใน endpoint POST /user
// Name และ Email เป็นฟิลด์ที่จำเป็น
// Email ต้องเป็นอีเมลที่ถูกต้อง
type CreateUserRequest struct {
	Name  string `json:"name" binding:"required"`        // Name is required
	Email string `json:"email" binding:"required,email"` // Email is required and must be valid
}

func main() {
	// สร้างตัวแปร user และกำหนดค่าเริ่มต้น
	user := User{Name: "John Doe", Email: "john@example.com", Age: 30}

	// ฟังก์ชัน updateUserEmail ใช้สำหรับอัปเดตค่า email ของ struct User
	// input: user (pointer ไปยัง struct User), newEmail (string ที่เป็น email ใหม่)
	// output: ไม่มี output แต่จะเปลี่ยนค่า email ใน struct User โดยตรง
	updateUserEmail(&user, "john.doe@example.com")

	// แสดงค่า email ที่ถูกอัปเดตเพื่อยืนยันว่าการเปลี่ยนแปลงสำเร็จ
	fmt.Println(user.Email)

	// ฟังก์ชัน validateUser ใช้สำหรับตรวจสอบความถูกต้องของ struct User
	// input: user (struct User ที่ต้องการตรวจสอบ)
	// output: error (ถ้าข้อมูลไม่ถูกต้องจะ return error, ถ้าถูกต้องจะ return nil)
	// - ตรวจสอบว่า email ไม่เป็นค่าว่าง
	// - ตรวจสอบว่า age เป็นจำนวนเต็มบวก
	if err := validateUser(user); err != nil {
		// ถ้าการตรวจสอบล้มเหลว แสดงข้อความ error
		fmt.Println("Validation error:", err)
	} else {
		// ถ้าการตรวจสอบผ่าน แสดงข้อความยืนยัน
		fmt.Println("User is valid")
	}

	// สร้าง router ของ Gin framework
	r := gin.Default()

	// กำหนด endpoint POST /user
	// ใช้สำหรับรับข้อมูล JSON {name, email} จากผู้ใช้
	r.POST("/user", func(c *gin.Context) {
		var req CreateUserRequest

		// Bind และ validate ข้อมูล JSON ที่รับเข้ามา
		if err := c.ShouldBindJSON(&req); err != nil {
			// ถ้าข้อมูลไม่ถูกต้อง ส่งสถานะ 400 พร้อมข้อความ error
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// ถ้าข้อมูลถูกต้อง ส่งสถานะ 201 พร้อมข้อมูลผู้ใช้ที่สร้างขึ้น
		c.JSON(http.StatusCreated, gin.H{
			"name":  req.Name,
			"email": req.Email,
		})
	})

	// เริ่มต้นเซิร์ฟเวอร์ที่พอร์ต 8080
	r.Run(":8080")
}
