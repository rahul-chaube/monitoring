
# Monitoring Project - UserService Module Contribution
## Summary
This contribution delivers a complete **UserService** module as part of the Monitoring system, providing essential backend APIs to support frontend dashboard functionalities.

## Features
- - RESTful APIs built with **Gin** framework
- **MongoDB** integration via environment variables
- Modular architecture (routes, controllers, models, utils)
- Email service integration using SMTP
- FCM token management
- Ready for expansion with AWS services (S3, SNS)
- MongoDB connection using environment variable
- Modular code (routes, controllers, models)

## ⚙️ Setup Instructions

### Prerequisites
- Go 1.18+
- MongoDB (Local or Atlas)
- (Optional) AWS CLI for future enhancements

### Installation Steps
1. Clone the repository:
   ```bash
   git clone https://github.com/rahul-chaube/monitoring.git
   cd monitoring
   ```
2. Copy `.env` file and configure:
   - Set MongoDB URI
   - Set SMTP credentials
3. Install dependencies:
   ```bash
   go mod tidy
   ```
4. Start the server:
   ```bash
   go run main.go
   ```

---

## 🚀 API Endpoints

### 1️⃣ **User Registration**  
`POST /user/register`  
Registers a new user.

### 2️⃣ **User Login**  
`POST /user/login`  
Authenticate user with email and password.

### 3️⃣ **Store Device Token**  
`POST /user/device-token`  
Update user's FCM token.

### 4️⃣ **Get User by ID**  
`GET /user/{id}`  
Fetch user details.

### 5️⃣ **Send Test Email**  
`GET /user/send-test-email`  
Trigger a test email.

_For detailed request/response samples, refer to the Postman collection._

---

## 📂 Postman Collection
- Use: `Monitoring_UserService_APIs.postman_collection.json`
- Set `{{base_url}}` to your deployed server (e.g., `http://<EC2-IP>:8080`)

---

## 🚀 Deployment Guide (AWS EC2)

1. **SSH into EC2:**
   ```bash
   ssh -i "monitor-ec2.pem" ec2-user@<EC2-IP>
   ```
2. **Install Go:**
   ```bash
   sudo yum update -y
   sudo yum install golang -y
   ```
3. **Clone & Run:**
   ```bash
   git clone https://github.com/rahul-chaube/monitoring.git
   cd monitoring
   go mod tidy
   go run main.go
   ```
4. Ensure port **8080** is open in AWS Security Group.
5. (Optional) Run in background:
   ```bash
   nohup go run main.go &
   ```

---

## ✅ Notes
- Configure `.env` properly (MongoDB, SMTP).
- Use Postman for API testing.
- Coordinate with frontend team for Dashboard API consumption.
- For enhancements, contact backend team.

