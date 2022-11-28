package main

import (
	"context"
	"log"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

const (
	fightClubGameTalbe = "FightClubGame"

	gameID   = "GameId"
	playerID = "PlayerId"
	score    = "Score"
)

func createTable(ctx context.Context, client *dynamodb.Client) {
	table, err := client.CreateTable(ctx, &dynamodb.CreateTableInput{
		AttributeDefinitions: []types.AttributeDefinition{
			{AttributeName: aws.String(gameID), AttributeType: types.ScalarAttributeTypeS},
			{AttributeName: aws.String(playerID), AttributeType: types.ScalarAttributeTypeS},
			{AttributeName: aws.String(score), AttributeType: types.ScalarAttributeTypeN},
		},
		TableName: aws.String(fightClubGameTalbe),
		KeySchema: []types.KeySchemaElement{
			{AttributeName: aws.String(gameID), KeyType: types.KeyTypeHash},
			{AttributeName: aws.String(playerID), KeyType: types.KeyTypeRange},
		},
		ProvisionedThroughput: &types.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(5),
			WriteCapacityUnits: aws.Int64(5),
		},
		LocalSecondaryIndexes: []types.LocalSecondaryIndex{
			{
				IndexName: aws.String("ScoreIndex"),
				KeySchema: []types.KeySchemaElement{
					{AttributeName: aws.String(gameID), KeyType: types.KeyTypeHash},
					{AttributeName: aws.String(score), KeyType: types.KeyTypeRange},
				},
				Projection: &types.Projection{
					ProjectionType: types.ProjectionTypeKeysOnly,
				},
			},
		},
	})

	if err != nil {
		log.Fatal(err)
	}

	waiter := dynamodb.NewTableExistsWaiter(client)
	err = waiter.Wait(
		ctx,
		&dynamodb.DescribeTableInput{
			TableName: aws.String(fightClubGameTalbe),
		},
		5*time.Minute,
	)

	if err != nil {
		log.Fatal(err)
	}

	_ = table
}
