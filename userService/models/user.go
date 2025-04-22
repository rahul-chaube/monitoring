package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Main User Model for DB
type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name      string             `bson:"name" json:"name"`
	Email     string             `bson:"email" json:"email"`
	Password  string             `bson:"password,omitempty" json:"password,omitempty"`
	FCMToken  string             `bson:"fcm_token,omitempty" json:"fcm_token,omitempty"`
	CreatedAt time.Time          `bson:"created_at,omitempty" json:"created_at,omitempty"`
}

// Request struct for creating user
type UserCreateRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
	FCMToken string `json:"fcm_token"`
}

// Response struct after user creation
type UserCreateResponse struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

// Login Request struct
type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

// Login Response struct (token placeholder)
type LoginResponse struct {
	Message string `json:"message"`
}

type DeviceTokenRequest struct {
	Email    string `json:"email" binding:"required,email"`
	FCMToken string `json:"fcm_token" binding:"required"`
}
