package main

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type Item struct {
	GameID   string `dynamodbav:"GameId"`
	PlayerID string `dynamodbav:"PlayerId"`
	Score    int    `dynamodbav:"Score"`
}

func insertItems(ctx context.Context, client *dynamodb.Client) {
	items := []Item{
		{GameID: "1", PlayerID: "a", Score: 150},
		{GameID: "1", PlayerID: "b", Score: 35},
		{GameID: "1", PlayerID: "c", Score: 50},
		{GameID: "1", PlayerID: "d", Score: 45},
		{GameID: "2", PlayerID: "d", Score: 10},
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

	_, err := client.BatchWriteItem(ctx, &dynamodb.BatchWriteItemInput{
		RequestItems: map[string][]types.WriteRequest{fightClubGameTalbe: writeReqs},
	})

	if err != nil {
		log.Fatal(err)
	}
}
