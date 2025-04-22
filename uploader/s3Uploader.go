package uploader

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"log"
	"mime/multipart"
	"time"
)

type S3Uploader struct {
	client     *s3.Client
	bucketName string
}

func NewS3Uploader(bucketName string) *S3Uploader {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	client := s3.NewFromConfig(cfg)

	//_, err = client.HeadObject(context.TODO(), &s3.HeadObjectInput{
	//	Bucket: aws.String(bucketName),
	//})
	//if err != nil {
	//	log.Fatal(err)
	//}

	return &S3Uploader{
		client:     client,
		bucketName: bucketName,
	}
}

func (uploader *S3Uploader) UploadFile(file *multipart.FileHeader) (string, string, error) {

	src, err := file.Open()
	if err != nil {
		return "", "", err
	}

	defer src.Close()

	key := fmt.Sprintf("%d_%s_%s", time.Now().UnixNano(), uploader.bucketName, file.Filename)
	fmt.Printf("Uploading file %s to %s\n", file.Filename, key)
	_, err = uploader.client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket:      aws.String(uploader.bucketName),
		Key:         aws.String(key),
		Body:        src,
		ContentType: aws.String(file.Header.Get("Content-Type")),
		//ACL:         types.ObjectCannedACLPublicRead
	})

	if err != nil {
		return "", "", err
	}

	presignClient := s3.NewPresignClient(uploader.client)
	resp, err := presignClient.PresignGetObject(context.TODO(), &s3.GetObjectInput{
		Bucket: aws.String(uploader.bucketName),
		Key:    aws.String(key),
	})
	fmt.Println("Presigned URL:", resp.URL)
	return key, fmt.Sprintf("%s/%s", uploader.bucketName, key), nil
}

func (uploader *S3Uploader) Presigned(files []string) []string {

	presignClient := s3.NewPresignClient(uploader.client)
	signedURLs := make([]string, len(files))
	for _, file := range files {
		resp, err := presignClient.PresignGetObject(context.TODO(), &s3.GetObjectInput{
			Bucket: aws.String(uploader.bucketName),
			Key:    aws.String(file),
		})
		if err != nil {
			fmt.Println("Failed to get presigned URL", file)
		}
		signedURLs = append(signedURLs, resp.URL)

	}

	fmt.Println("Presigned URL:", signedURLs)

	return signedURLs

}
