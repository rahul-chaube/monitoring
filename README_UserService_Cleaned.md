
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

📄 Monitoring - UserService API Documentation

📍 Base URL
http://<EC2-IP>:8080
Replace <EC2-IP> with your deployed EC2 instance IP.

________________________________________
🚀 API Endpoints

1️⃣ User Registration
• Method: POST
• Endpoint: /user/register
• Description: Registers a new user and triggers a welcome email.
• Request Body:
{
   "name": "John Doe",
   "email": "john@example.com",
   "password": "password123",
   "fcm_token": "xyz123token"
}
• Success Response:
{
   "status": "success",
   "message": "User registered successfully",
   "data": {
      "name": "John Doe",
      "email": "john@example.com"
   }
}

2️⃣ User Login
• Method: POST
• Endpoint: /user/login
• Request Body:
{
   "email": "john@example.com",
   "password": "password123"
}

3️⃣ Store Device Token
• Method: POST
• Endpoint: /user/device-token
• Request Body:
{
   "email": "john@example.com",
   "fcm_token": "new_token_456"
}

4️⃣ Get User by ID
• Method: GET
• Endpoint: /user/{id}
• Description: Fetch user details by MongoDB ObjectID.

5️⃣ Send Test Email
• Method: GET
• Endpoint: /user/send-test-email
• Description: Sends a test email to verify email service.

6️⃣ Send Forwarding Email
• Method: POST
• Endpoint: /user/send-forwarding-email
• Description: Sends a custom HTML-formatted email using the forwarding template.
• Request Body:
{
   "to": "recipient@example.com",
   "subject": "Forwarded Message",
   "header": "New User Contact",
   "body": "Hello team, this is a forwarded message from the user."
}
• Success Response:
{
   "status": "success",
   "message": "Templated email sent successfully",
   "data": null
}

________________________________________
📧 Auto-Triggered Welcome Email
• When a new user registers via /user/register, a welcome email is automatically sent using an HTML template (welcome_email.html).
• This improves onboarding and confirms user creation successfully.

________________________________________
📂 Postman Collection
• Import Monitoring_UserService_APIs_with_Welcome_Email.postman_collection.json into Postman.
• Set the {{base_url}} variable to your deployed server URL.

________________________________________
🚀 Deployment Guide on AWS EC2

⚙️ Prerequisites
• AWS EC2 instance (Amazon Linux)
• Port 8080 open in Security Group
• monitor-ec2.pem key file
• Go 1.18+ installed
• MongoDB connection URI & SMTP credentials ready

📦 Deployment Steps

1. SSH into EC2:
ssh -i "monitor-ec2.pem" ec2-user@<EC2-IP>

2. Install Go (if not already installed):
sudo yum update -y
sudo yum install golang -y

3. Clone the Repository:
git clone https://github.com/rahul-chaube/monitoring.git
cd monitoring

4. Configure Environment Variables:
• Create a .env file and add MongoDB & SMTP settings.

5. Build & Run the Service:
go mod tidy
go run main.go

6. Run in Background (Optional):
nohup go run main.go &

7. Verify Deployment:
• Test using Postman or browser:
http://<EC2-IP>:8080/user/send-test-email

8. Share Deployment Info:
• "Deployed UserService API on: http://<EC2-IP>:8080"
• Attach Postman collection & API documentation.

________________________________________
✅ Notes
• Ensure .env is correctly set before running.
• Use Postman for quick endpoint validation.
• Coordinate with frontend team for Dashboard API consumption.
• For further improvements or issues, contact the backend team.