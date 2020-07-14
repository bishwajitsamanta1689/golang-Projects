package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"os"
)

func main(){
	sess, err:= session.NewSession(&aws.Config{
		Region: aws.String("us-west-2"),

		// Below is the Way to declare the Credentials Directly
		Credentials: credentials.NewStaticCredentials("AKIAYJY7MRJKJMAOUTH2","ggnDpz6LAcqh+N/2KH0Kv+PMqyqXt8ChluyWJG5y", ""),
	})

	svc:= s3.New(sess)

	result, err:= svc.ListBuckets(nil)
	if err != nil {
		exitErrof("Unable to List Buckets, %v", err)
	}
	fmt.Println("Buckets:")

	for _, b:= range result.Buckets{
		fmt.Printf("* %s created on => %s \n", aws.StringValue(b.Name), aws.TimeValue(b.CreationDate))
	}
}

func exitErrof(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stderr,msg+"\n",args...)
	os.Exit(1)
}