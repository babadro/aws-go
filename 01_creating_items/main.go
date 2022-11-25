package main

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type BlogItem struct {
	Author     string `dynamodbav:"Author"`
	TopicTitle string `dynamodbav:"Topic_Title"`
	Website    string `dynamodbav:"Website"`
}

const BlogTable = "Blog"

func main() {
	// Creating a single item
	ctx := context.TODO()
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		log.Fatal(err)
	}

	client := dynamodb.NewFromConfig(cfg)

	item := BlogItem{
		Author:     "Seth Godin",
		TopicTitle: "Marketing, This is Marketing",
	}

	mItem, err := attributevalue.MarshalMap(item)
	if err != nil {
		log.Fatal(err)
	}

	_, err = client.PutItem(ctx, &dynamodb.PutItemInput{
		Item:      mItem,
		TableName: aws.String(BlogTable),
	})

	if err != nil {
		log.Fatal(err)
	}

	// Calling the BatchWriteItem API
	items := []BlogItem{
		{
			Author:     "Anshul",
			TopicTitle: "DynamoDB, Introduction To AWS CLI"},
		{
			Author:     "Anshul",
			TopicTitle: "Marketing, 33 Things to Learn from 33 Top Marketing Blogs",
			Website:    "www.neilpatel.com",
		},
		{
			Author:     "Neil Patel",
			TopicTitle: "Entrepreneurship, Bluehost Vs Hostgator",
			Website:    "www.neilpatel.com",
		},
	}

	var writeReqs []types.WriteRequest
	for _, item := range items {
		mItem, err := attributevalue.MarshalMap(item)
		if err != nil {
			log.Fatal(err)
		}

		writeReqs = append(writeReqs,
			types.WriteRequest{
				PutRequest: &types.PutRequest{Item: mItem},
			})
	}

	_, err = client.BatchWriteItem(ctx, &dynamodb.BatchWriteItemInput{
		RequestItems: map[string][]types.WriteRequest{BlogTable: writeReqs},
	})

	if err != nil {
		log.Fatal(err)
	}
}
