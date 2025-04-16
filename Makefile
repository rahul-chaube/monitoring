# Makefile for Go app deployment

# Variables
GO_APP_NAME=monitor        # Change this to your app's binary name
GO_BUILD_DIR=./bin
EC2_USER=ec2-user            # Default user for Amazon Linux
EC2_HOST=13.234.202.23        # Replace with your EC2 instance IP address
REMOTE_PATH=/home/ec2-user/app
GOOS=linux
GOARCH=amd64

# Build the Go binary
build:
	@echo "Building Go binary..."
	GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o $(GO_BUILD_DIR)/$(GO_APP_NAME)

# Upload the binary to the EC2 instance
upload:
	@echo "Uploading binary to EC2 instance..."
	ssh -i /Users/r.chaube/Downloads/monitor-ec2.pem ec2-user@$(EC2_HOST) "mkdir -p $(REMOTE_PATH)"  # Create the directory if it doesn't exist
	scp -i /Users/r.chaube/Downloads/monitor-ec2.pem $(GO_BUILD_DIR)/$(GO_APP_NAME) $(EC2_USER)@$(EC2_HOST):$(REMOTE_PATH)

# Run the Go application on the EC2 instance
run:
	@echo "Running Go application on EC2 instance..."
	ssh -i /Users/r.chaube/Downloads/monitor-ec2.pem $(EC2_USER)@$(EC2_HOST) "cd $(REMOTE_PATH) && ./$(GO_APP_NAME)"

# Full process: Build, Upload, and Run
deploy: build upload run
	@echo "Deployment complete!"

# Clean up
clean:
	@echo "Cleaning up..."
	rm -rf $(GO_BUILD_DIR)/$(GO_APP_NAME)

.PHONY: build upload run deploy clean
