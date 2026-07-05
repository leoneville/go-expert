package main

import (
	"fmt"
	"io"
	"os"
	"sync"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

var (
	s3Client *s3.S3
	s3Bucket string
)

func init() {
	sess, err := session.NewSession(
		&aws.Config{
			Region: aws.String("us-east-1"),
			Credentials: credentials.NewStaticCredentials(
				"id_client",
				"id_secret",
				"",
			),
		},
	)
	if err != nil {
		panic(err)
	}

	s3Client = s3.New(sess)
	s3Bucket = "goexpert-bucket-exemplo-nb"
}

func main() {
	dir, err := os.Open("./tmp")
	if err != nil {
		panic(err)
	}
	defer dir.Close()

	uploadControl := make(chan struct{}, 100)
	errorFileUpload := make(chan string, 10)
	wg := sync.WaitGroup{}

	go func() {
		for filename := range errorFileUpload {
			uploadControl <- struct{}{}
			wg.Go(func() { uploadFile(filename, uploadControl, errorFileUpload) })
		}
	}()

	for {
		files, err := dir.ReadDir(1)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Printf("Error reading directory: %s\n", err)
			continue
		}

		uploadControl <- struct{}{}
		wg.Go(func() { uploadFile(files[0].Name(), uploadControl, errorFileUpload) })
	}
	wg.Wait()
	close(errorFileUpload)
}

func uploadFile(filename string, uploadControl <-chan struct{}, errorFileUpload chan<- string) {
	defer func() {
		<-uploadControl
	}()

	completeFilename := fmt.Sprintf("./tmp/%s", filename)
	fmt.Printf("Uploading file %s to bucket %s\n", completeFilename, s3Bucket)

	f, err := os.Open(completeFilename)
	if err != nil {
		fmt.Printf("Error opening file %s: %v\n", completeFilename, err)
		errorFileUpload <- completeFilename
		return
	}

	defer f.Close()

	_, err = s3Client.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(s3Bucket),
		Key:    aws.String(filename),
		Body:   f,
	})
	if err != nil {
		fmt.Printf("Error uploading file %s: %v\n", completeFilename, err)
		errorFileUpload <- completeFilename
		return
	}

	fmt.Printf("File %s uploaded successfully\n", completeFilename)
}
