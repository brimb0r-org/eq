package eq_aws

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/brimb0r-org/eq/application/internal/metadata"
	"github.com/rs/zerolog/log"
	"os"
)

type AwsUtil struct {
	CachedSession *session.Session
	Sessions      map[string]*session.Session
}

var cachedSession *session.Session

func Session() *AwsUtil {
	var c aws.Config
	if localStack() {
		c = aws.Config{
			CredentialsChainVerboseErrors: aws.Bool(true),
			Credentials:                   credentials.NewStaticCredentials("not", "empty", ""),
			DisableSSL:                    aws.Bool(true),
			Endpoint:                      aws.String(os.Getenv("LOCALSTACK_ENDPOINT")),
			S3ForcePathStyle:              aws.Bool(true),
			Region:                        aws.String(metadata.MetaData().Active_Region()),
		}
	} else {
		c = aws.Config{
			CredentialsChainVerboseErrors: aws.Bool(true),
			Region:                        aws.String(metadata.MetaData().Active_Region()),
		}
	}

	if cachedSession == nil {
		var err error
		cachedSession, err = session.NewSession(&c)
		if err != nil {
			log.Fatal().Msgf("no aws session %s", err.Error())
		}
	}
	cachedSession.Copy(&c)
	awsUtils := &AwsUtil{
		CachedSession: cachedSession.Copy(&c),
		Sessions:      make(map[string]*session.Session),
	}
	return awsUtils
}

func (a *AwsUtil) GetSession() *session.Session {
	if sessionCheck, ok := a.Sessions[metadata.MetaData().Active_Region()]; ok {
		return sessionCheck
	}
	regionSession := a.CachedSession.Copy(&aws.Config{Region: aws.String(metadata.MetaData().Active_Region())})
	return regionSession
}

func localStack() bool {
	localstack := os.Getenv("LOCALSTACK")
	return localstack == "true"
}
