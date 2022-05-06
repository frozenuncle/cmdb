package connectivity_test

import (
	"testing"

	"github.com/infraboard/cmdb/provider/aliyun/connectivity"
	"github.com/stretchr/testify/assert"
)

func TestClient(t *testing.T) {
	should := assert.New(t)

	err := connectivity.LoadClientFromEnv()
	if should.NoError(err) {
		connectivity.C().EcsClient()
	}
}
