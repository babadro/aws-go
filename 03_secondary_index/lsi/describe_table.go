package main

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func describeTable(ctx context.Context, client *dynamodb.Client) {
	res, err := client.DescribeTable(ctx,
		&dynamodb.DescribeTableInput{TableName: aws.String(fightClubGameTalbe)})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(res)
}
