
*Student API with Gin*

โครงสร้างโปรเจกต์

go-api-gin/

    * main.go
    * students.db
    * models/
    * repositories/
    * services/
    * handlers/
    * config/

main.go
ใช้สำหรับตั้งค่า Router และเริ่มต้นการทำงานของเซิร์ฟเวอร์

students.db
ไฟล์ฐานข้อมูล SQLite จะถูกสร้างอัตโนมัติเมื่อรันโปรเจกต์ครั้งแรก

models/
เก็บโครงสร้างข้อมูล เช่น struct Student และฟังก์ชัน validation

repositories/
ติดต่อกับฐานข้อมูล SQLite โดยตรง (SQL logic)

services/
จัดการ business logic และควบคุมการทำงานระหว่าง handler กับ repository

handlers/
จัดการ HTTP request และ response

config/
ตั้งค่าการเชื่อมต่อฐานข้อมูล

---

วิธีรันโปรเจกต์

1. ติดตั้ง dependency (ถ้ายังไม่ได้ติดตั้ง)

go mod tidy

2. รันโปรแกรม

go run main.go

เมื่อรันสำเร็จ เซิร์ฟเวอร์จะทำงานที่

[http://localhost:8080](http://localhost:8080)

---

API ที่มีในระบบ

1. ดึงข้อมูลนักศึกษาทั้งหมด
   Method: GET
   Endpoint: /students

2. ดึงข้อมูลนักศึกษาตามรหัส
   Method: GET
   Endpoint: /students/{id}

3. เพิ่มข้อมูลนักศึกษา
   Method: POST
   Endpoint: /students

ตัวอย่าง JSON ที่ใช้ส่งข้อมูล

{
"id": "66090001",
"name": "John",
"major": "CS",
"gpa": 3.0
}

4. แก้ไขข้อมูลนักศึกษา
   Method: PUT
   Endpoint: /students/{id}

ตัวอย่าง JSON

{
"name": "John02",
"major": "PH",
"gpa": 3.5
}

5. ลบนักศึกษา
   Method: DELETE
   Endpoint: /students/{id}

หากลบสำเร็จจะได้สถานะ 204 No Content

---

วิธีทดสอบระบบเบื้องต้นด้วย Postman

1. ทดสอบ GET ทั้งหมด

* เลือก Method: GET
* URL: [http://localhost:8080/students](http://localhost:8080/students)
* กด Send
  คาดหวัง: ได้สถานะ 200 OK

2. ทดสอบ POST เพิ่มข้อมูล

* เลือก Method: POST
* URL: [http://localhost:8080/students](http://localhost:8080/students)
* ไปที่ Body เลือก raw และเลือก JSON
* ใส่ข้อมูลตัวอย่างตามที่กำหนด
* กด Send
  คาดหวัง: ได้สถานะ 201 Created และคืนข้อมูลนักศึกษาที่เพิ่ม

3. ทดสอบ GET ตาม ID

* เลือก Method: GET
* URL: [http://localhost:8080/students/66090001](http://localhost:8080/students/66090001)
  คาดหวัง: ได้สถานะ 200 OK และข้อมูลนักศึกษาตรงกับที่เพิ่มไว้

4. ทดสอบ PUT แก้ไขข้อมูล

* เลือก Method: PUT
* URL: [http://localhost:8080/students/66090001](http://localhost:8080/students/66090001)
* ใส่ JSON สำหรับแก้ไข
  คาดหวัง: ได้สถานะ 200 OK และข้อมูลถูกอัปเดต

5. ทดสอบ DELETE

* เลือก Method: DELETE
* URL: [http://localhost:8080/students/66090001](http://localhost:8080/students/66090001)
  คาดหวัง: ได้สถานะ 204 No Content

6. ทดสอบกรณีไม่พบข้อมูล (Not Found)
   ลองเรียก
   GET [http://localhost:8080/students/999999](http://localhost:8080/students/999999)
   คาดหวัง:
   สถานะ 404 Not Found
   และได้ข้อความ
   {
   "error": "Student not found"
   }

7. ทดสอบ Validation
   กรณี GPA มากกว่า 4.00 หรือ id/name ว่าง
   ระบบจะตอบกลับ 400 Bad Request พร้อมข้อความ error ในรูปแบบ JSON

---


