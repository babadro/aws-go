package main

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)
import "github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"

func query(ctx context.Context, client *dynamodb.Client) {
	keyEx := expression.Key("ChemicalFormula").Equal(expression.Value("Levothyroxine")).
		And(expression.Key("Dosage").Equal(expression.Value(25)))

	expr, err := expression.NewBuilder().WithKeyCondition(keyEx).Build()
	if err != nil {
		log.Fatal(err)
	}

	response, err := client.Query(ctx, &dynamodb.QueryInput{
		TableName:                 aws.String(medicinesTableName),
		ProjectionExpression:      aws.String("AvailableBrands"),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		KeyConditionExpression:    expr.KeyCondition(),
	})

	if err != nil {
		log.Fatal(err)
	}

	var medicinesItems []MedicinesItem
	err = attributevalue.UnmarshalListOfMaps(response.Items, &medicinesItems)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(medicinesItems)
}

func scan(ctx context.Context, client *dynamodb.Client) {
	filtEx := expression.Name("Dosage").Equal(expression.Value(25))
	expr, err := expression.NewBuilder().WithFilter(filtEx).Build()
	if err != nil {
		log.Fatal(err)
	}

	response, err := client.Scan(ctx, &dynamodb.ScanInput{
		TableName:                 aws.String(medicinesTableName),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		FilterExpression:          expr.Filter(),
	})

	if err != nil {
		log.Fatal(err)
	}

	var items []MedicinesItem
	err = attributevalue.UnmarshalListOfMaps(response.Items, &items)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(items)
}
