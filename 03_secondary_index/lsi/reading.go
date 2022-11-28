package main

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func getPlayer(ctx context.Context, client *dynamodb.Client) {
	keyEx := expression.Key(gameID).Equal(expression.Value("2")).
		And(expression.Key(playerID).Equal(expression.Value("d")))

	expr, err := expression.NewBuilder().WithKeyCondition(keyEx).Build()
	if err != nil {
		log.Fatal(err)
	}

	response, err := client.Query(ctx, &dynamodb.QueryInput{
		TableName:                 aws.String(fightClubGameTalbe),
		ProjectionExpression:      aws.String(score),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		KeyConditionExpression:    expr.KeyCondition(),
	})

	if err != nil {
		log.Fatal(err)
	}

	var items []Item
	err = attributevalue.UnmarshalListOfMaps(response.Items, &items)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(items)
}

func getTopScores(ctx context.Context, client *dynamodb.Client) {
	keyEx := expression.Key(gameID).Equal(expression.Value("1"))

	expr, err := expression.NewBuilder().WithKeyCondition(keyEx).Build()
	if err != nil {
		log.Fatal(err)
	}

	response, err := client.Query(ctx, &dynamodb.QueryInput{
		TableName:                 aws.String(fightClubGameTalbe),
		IndexName:                 aws.String("ScoreIndex"),
		Limit:                     aws.Int32(2),
		ScanIndexForward:          aws.Bool(false),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		KeyConditionExpression:    expr.KeyCondition(),
	})

	if err != nil {
		log.Fatal(err)
	}

	var items []Item
	err = attributevalue.UnmarshalListOfMaps(response.Items, &items)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(items)
}
