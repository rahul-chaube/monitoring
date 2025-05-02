package notificationService

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sns"
	"log"
)

func SendSMS() string {
	ctx := context.Background()

	// Load AWS configuration
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	snsClient := sns.NewFromConfig(cfg)
	// Send SMS
	phoneNumber := "+918976898022" // must include country code, e.g., +91 for India
	message := "Hello from AWS SNS!"

	input := &sns.PublishInput{
		Message:     aws.String(message),
		PhoneNumber: aws.String(phoneNumber),
	}

	result, err := snsClient.Publish(ctx, input)
	if err != nil {
		log.Fatalf("failed to send SMS, %v", err)
	}

	fmt.Println("Message ID:", *result.MessageId)

	return *result.MessageId
}
