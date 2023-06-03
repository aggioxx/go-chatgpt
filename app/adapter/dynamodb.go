package adapter

import (
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type DynamoImpl struct {
	DClient *dynamodb.Client
}
