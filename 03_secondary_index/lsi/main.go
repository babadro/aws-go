package main

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func main() {
	ctx := context.TODO()
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		log.Fatal(err)
	}

	client := dynamodb.NewFromConfig(cfg)

	//createTable(ctx, client)
	//describeTable(ctx, client)
	//insertItems(ctx, client)

	//getPlayer(ctx, client)
	getTopScores(ctx, client)
}
