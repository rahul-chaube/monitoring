package controllers

import (
	"context"
	"github.com/rahul-chaube/monitoring/util"
	"net/http"
	"time"

	"github.com/rahul-chaube/monitoring/userService/config"
	"github.com/rahul-chaube/monitoring/userService/models"
	"github.com/rahul-chaube/monitoring/userService/utils"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(c *gin.Context) {
	var input models.UserCreateRequest

	if err := c.ShouldBindJSON(&input); err != nil {
		utils.ErrorResponse(c, "Invalid input")
		return
	}

	collection := config.GetCollection(util.UserRepository)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Check if user already exists
	var existingUser models.User
	err := collection.FindOne(ctx, bson.M{"email": input.Email}).Decode(&existingUser)
	if err == nil {
		utils.ErrorResponse(c, "Email already registered")
		return
	}

	// Hash password
	hashedPassword, err := utils.HashPassword(input.Password)
	if err != nil {
		utils.ErrorResponse(c, "Error hashing password")
		return
	}

	newUser := models.User{
		ID:        primitive.NewObjectID(),
		Name:      input.Name,
		Email:     input.Email,
		Password:  hashedPassword,
		FCMToken:  input.FCMToken,
		CreatedAt: time.Now(),
	}

	_, err = collection.InsertOne(ctx, newUser)
	if err != nil {
		utils.ErrorResponse(c, "Failed to register user")
		return
	}

	// After inserting the new user into the database
	// err = utils.SendEmail(newUser.Email, "Welcome to Our Service", "Hello "+newUser.Name+",\n\nThank you for registering with us!")
	// if err != nil {
	// 	log.Printf("Failed to send welcome email: %v", err)
	// 	// Proceed without halting the registration process
	// }
	err = utils.SendTemplatedEmail(
		newUser.Email,
		"Welcome to Monitoring!",
		utils.EmailTemplateData{
			Subject: "Welcome to Monitoring",
			Header:  "Thanks for joining, " + newUser.Name + "!",
			Body:    "We’re excited to have you on board. Let us know if you need any help getting started.",
		},
		"templates/welcome_email.html", // 👈 use welcome template
	)

	response := models.UserCreateResponse{
		Name:  newUser.Name,
		Email: newUser.Email,
	}

	utils.SuccessResponse(c, "User registered successfully", response)
}

func LoginUser(c *gin.Context) {
	var loginReq models.LoginRequest

	if err := c.ShouldBindJSON(&loginReq); err != nil {
		utils.ErrorResponse(c, "Invalid login request")
		return
	}

	collection := config.GetCollection(util.UserRepository)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var user models.User
	err := collection.FindOne(ctx, bson.M{"email": loginReq.Email}).Decode(&user)
	if err != nil {
		utils.ErrorResponse(c, "Email not registered")
		return
	}

	// Verify Password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginReq.Password))
	if err != nil {
		utils.ErrorResponse(c, "Incorrect password")
		return
	}

	utils.SuccessResponse(c, "Login successful", models.LoginResponse{
		Message: "Welcome " + user.Name,
	})
}

func StoreDeviceToken(c *gin.Context) {
	var req models.DeviceTokenRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, "Invalid input")
		return
	}

	collection := config.GetCollection(util.UserRepository)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"email": req.Email}
	update := bson.M{"$set": bson.M{"fcm_token": req.FCMToken}}

	result, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		utils.ErrorResponse(c, "Failed to update FCM token")
		return
	}

	if result.MatchedCount == 0 {
		utils.ErrorResponse(c, "User not found")
		return
	}

	utils.SuccessResponse(c, "FCM token updated successfully", nil)
}

func SendTestEmail(c *gin.Context) {
	err := utils.SendEmail("recipient@example.com", "Test Email", "This is a test email.")
	if err != nil {
		utils.ErrorResponse(c, "Failed to send email")
		return
	}
	utils.SuccessResponse(c, "Email sent successfully", nil)
}

func GetUserByID(c *gin.Context) {
	userID := c.Param("id")
	c.JSON(http.StatusOK, gin.H{"user_id": userID})
}

func SendForwardingEmail(c *gin.Context) {
	var payload struct {
		To      string `json:"to"`
		Subject string `json:"subject"`
		Header  string `json:"header"`
		Body    string `json:"body"`
	}

	if err := c.ShouldBindJSON(&payload); err != nil {
		utils.ErrorResponse(c, "Invalid request")
		return
	}

	err := utils.SendTemplatedEmail(
		payload.To,
		payload.Subject,
		utils.EmailTemplateData{
			Subject: payload.Subject,
			Header:  payload.Header,
			Body:    payload.Body,
		},
		"templates/forwarding_email.html",
	)

	if err != nil {
		utils.ErrorResponse(c, "Failed to send email: "+err.Error())
		return
	}

	utils.SuccessResponse(c, "Templated email sent successfully", nil)
}
