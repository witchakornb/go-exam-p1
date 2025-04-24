package main

import (
	"fmt"
)

// 3. Write a Go program that defines a User struct with struct tags and a function that accepts a pointer to User and modifies one of its fields.

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

func updateUserEmail(user *User, newEmail string) {
	user.Email = newEmail
}

// 4. Write a validation function for User with fields: Email, Age
// The function should check if the email is valid and if the age is a positive integer.
func validateUser(user User) error {
	if user.Email == "" {
		return fmt.Errorf("email cannot be empty")
	}
	if user.Age <= 0 {
		return fmt.Errorf("age must be a positive integer")
	}
	return nil
}

func main() {
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
}
