package sstore

import (
	"fmt"
	"gSheets/application/pkg/eq_aws"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
	"github.com/aws/aws-sdk-go/service/secretsmanager/secretsmanageriface"
)

type SecretStore struct {
	client      secretsmanageriface.SecretsManagerAPI
	secretCache map[string]cachedValue
	region      string
	environment string
}

type cachedValue struct {
	value string
}

func NewSStoreClient(util *eq_aws.AwsUtil) *SecretStore {
	session := util.GetSession()
	sstore := &SecretStore{
		client:      secretsmanager.New(session),
		secretCache: make(map[string]cachedValue),
	}
	return sstore
}

func (sstore *SecretStore) Get(key string) (interface{}, error) {
	v, ok := sstore.secretCache[key]
	if ok {
		return v.value, nil
	}
	return sstore.retrieveSecret(key)
}

func (sstore *SecretStore) retrieveSecret(key string) (string, error) {
	resp, err := sstore.client.GetSecretValue(&secretsmanager.GetSecretValueInput{
		SecretId: aws.String(key),
	})
	if err != nil {
		return "", fmt.Errorf("error on getting key %s [%w]", key, err)
	}
	if resp.SecretString == nil {
		return "", fmt.Errorf("secret value for key not present %s", key)
	}
	value := *resp.SecretString
	sstore.secretCache[key] = cachedValue{
		value: value,
	}
	return value, err
}
