package main

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func createGsi(ctx context.Context, client *dynamodb.Client) {
	table, err := client.UpdateTable(ctx, &dynamodb.UpdateTableInput{
		TableName: aws.String(musicCollectionTable),
		AttributeDefinitions: []types.AttributeDefinition{
			{AttributeName: aws.String("Album"), AttributeType: types.ScalarAttributeTypeS},
			{AttributeName: aws.String("Song"), AttributeType: types.ScalarAttributeTypeS},
		},
		GlobalSecondaryIndexUpdates: []types.GlobalSecondaryIndexUpdate{
			{
				Create: &types.CreateGlobalSecondaryIndexAction{
					IndexName: aws.String("AlbumIndex"),
					KeySchema: []types.KeySchemaElement{
						{AttributeName: aws.String("Album"), KeyType: types.KeyTypeHash},
						{AttributeName: aws.String("Song"), KeyType: types.KeyTypeRange},
					},
					Projection: &types.Projection{
						ProjectionType: types.ProjectionTypeKeysOnly,
					},
					ProvisionedThroughput: &types.ProvisionedThroughput{
						ReadCapacityUnits:  aws.Int64(5),
						WriteCapacityUnits: aws.Int64(6),
					},
				},
			},
		},
	})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(table)
}
