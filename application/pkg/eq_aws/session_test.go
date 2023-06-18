package eq_aws

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func Test_localStack(t *testing.T) {
	t.Run("unset", func(t *testing.T) {
		value := localStack()
		assert.Equal(t, false, value)
	})
	t.Run("set", func(t *testing.T) {
		os.Setenv("LOCALSTACK", "true")
		value := localStack()
		assert.Equal(t, true, value)
	})
}
