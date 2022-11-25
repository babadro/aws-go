package main

import (
	"context"
	"log"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

const medicinesTableName = "Medicines"

type MedicinesItem struct {
	ChemicalFormula string   `dynamodbav:"ChemicalFormula"`
	Dosage          int      `dynamodbav:"Dosage"`
	AvailableBrands []string `dynamodbav:"AvailableBrands"`
}

func creatingTable(ctx context.Context, client *dynamodb.Client) {
	table, err := client.CreateTable(ctx, &dynamodb.CreateTableInput{
		AttributeDefinitions: []types.AttributeDefinition{
			{AttributeName: aws.String("ChemicalFormula"), AttributeType: types.ScalarAttributeTypeS},
			{AttributeName: aws.String("Dosage"), AttributeType: types.ScalarAttributeTypeN},
		},
		TableName: aws.String(medicinesTableName),

		KeySchema: []types.KeySchemaElement{
			{AttributeName: aws.String("ChemicalFormula"), KeyType: types.KeyTypeHash},
			{AttributeName: aws.String("Dosage"), KeyType: types.KeyTypeRange},
		},
		ProvisionedThroughput: &types.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(4),
			WriteCapacityUnits: aws.Int64(4),
		},
	})

	if err != nil {
		log.Fatal(err)
	}

	waiter := dynamodb.NewTableExistsWaiter(client)
	err = waiter.Wait(
		ctx,
		&dynamodb.DescribeTableInput{
			TableName: aws.String("Medicines"),
		},
		5*time.Minute)

	if err != nil {
		log.Fatal(err)
	}

	_ = table
}

func insertItems(ctx context.Context, client *dynamodb.Client) {
	items := []MedicinesItem{
		{
			ChemicalFormula: "Esomeprazole",
			Dosage:          25,
			AvailableBrands: []string{"NexIUM"},
		},
		{
			ChemicalFormula: "Rosuvastatin",
			Dosage:          25,
			AvailableBrands: []string{"Jupiros", "Lipirose"},
		},
		{
			ChemicalFormula: "Levothyroxine",
			Dosage:          25,
			AvailableBrands: []string{"Levoxyl", "Synthroid"},
		},
		{
			ChemicalFormula: "Levothyroxine",
			Dosage:          75,
			AvailableBrands: []string{"Levoxyl", "Unithroid"},
		},
		{
			ChemicalFormula: "Rosuvastatin",
			Dosage:          10,
			AvailableBrands: []string{"Arvast", "Crestor"},
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

	_, err := client.BatchWriteItem(ctx, &dynamodb.BatchWriteItemInput{
		RequestItems: map[string][]types.WriteRequest{medicinesTableName: writeReqs},
	})

	if err != nil {
		log.Fatal(err)
	}
}
