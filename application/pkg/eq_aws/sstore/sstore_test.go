package sstore

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
	"github.com/aws/aws-sdk-go/service/secretsmanager/secretsmanageriface"
	"testing"
)

type mockSecretsManager struct {
	secretsmanageriface.SecretsManagerAPI
	getSecretCounter int
	secretString     string
}

func (sm *mockSecretsManager) GetSecretValue(input *secretsmanager.GetSecretValueInput) (*secretsmanager.GetSecretValueOutput, error) {
	sm.getSecretCounter++
	return &secretsmanager.GetSecretValueOutput{SecretString: aws.String(sm.secretString)}, nil
}

func (sm *mockSecretsManager) setSecretString(s string) *mockSecretsManager {
	sm.secretString = s
	return sm
}

func TestSecretManager(t *testing.T)      {}
func TestSecretManagerCache(t *testing.T) {}
