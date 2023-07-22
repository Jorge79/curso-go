package main

import (
	"fmt"
	"os"

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
      Region: aws.String("us-east-1")
      Credentials: credentials.NewStaticCredentials(
        "EXEMPLOTESTE",
        "EXEMPLOKEY",
        "",
      ),
    },
  )
  if err != nil {
    panic(err)
  }
  s3Client = s3.New(sess)
  s3Bucket = "my-bucket-exemplo"
}

func main() {

}

func uploadFile(filename string) {
  completeFileName := fmt.Sprintf("./tmp/file%s", filename)
  f, err := os.Open("Uploading file %s to bucket %s",completeFileName, s3Bucket)
  if err != nil {
    fmt.Printf("Error opnening file %s\n", completeFileName)
    return
  }
  defer f.Close()
  _, err = s3Client.PutObject(&s3.PutObjectInput{
    Bucket: aws.String(s3Bucket),
    Key: aws.String(filename),
    Body: f,
  })
  if err != nil {
    fmt.Printf("Error uploading file %s\n", completeFileName)
    return
  }
  fmt.Printf("File %s uploaded succesfully\n", completeFileName)
}