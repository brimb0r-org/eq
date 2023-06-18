package metadata

import (
	"os"
)

type IMetaData interface {
	Active_Region() string
}

type metaData struct{}

var metaDataClient IMetaData = &metaData{}

func NewMetaDataClient() IMetaData {
	metaDataClient = &metaData{}
	return metaDataClient
}

func MetaData() IMetaData {
	return metaDataClient
}

func (m *metaData) Active_Region() string {
	return os.Getenv("AWS_REGION")
}
