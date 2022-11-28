package main

import (
	"context"
	"log"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

const musicCollectionTable = "MusicCollection"

type Song struct {
	Song   string `dynamodbav:"Song"`
	Artist string `dynamodbav:"Artist"`
	Genre  string `dynamodbav:"Genre"`
}

func createTableWithIndex(ctx context.Context, client *dynamodb.Client) {
	table, err := client.CreateTable(ctx, &dynamodb.CreateTableInput{
		AttributeDefinitions: []types.AttributeDefinition{
			{AttributeName: aws.String("Song"), AttributeType: types.ScalarAttributeTypeS},
			{AttributeName: aws.String("Artist"), AttributeType: types.ScalarAttributeTypeS},
			{AttributeName: aws.String("Genre"), AttributeType: types.ScalarAttributeTypeS},
		},
		TableName: aws.String(musicCollectionTable),
		KeySchema: []types.KeySchemaElement{
			{AttributeName: aws.String("Artist"), KeyType: types.KeyTypeHash},
			{AttributeName: aws.String("Song"), KeyType: types.KeyTypeRange},
		},
		ProvisionedThroughput: &types.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(5),
			WriteCapacityUnits: aws.Int64(5),
		},
		GlobalSecondaryIndexes: []types.GlobalSecondaryIndex{
			{
				IndexName: aws.String("GenreIndex"),
				KeySchema: []types.KeySchemaElement{
					{AttributeName: aws.String("Genre"), KeyType: types.KeyTypeHash},
					{AttributeName: aws.String("Artist"), KeyType: types.KeyTypeRange},
				},
				Projection: &types.Projection{
					NonKeyAttributes: []string{"Song"},
					ProjectionType:   types.ProjectionTypeInclude,
				},
				ProvisionedThroughput: &types.ProvisionedThroughput{
					ReadCapacityUnits:  aws.Int64(10),
					WriteCapacityUnits: aws.Int64(6),
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
			TableName: aws.String(musicCollectionTable),
		},
		5*time.Minute,
	)

	if err != nil {
		log.Fatal(err)
	}

	_ = table
}
