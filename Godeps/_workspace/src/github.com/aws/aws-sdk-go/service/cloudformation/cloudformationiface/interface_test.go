// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package cloudformationiface_test

import (
	"testing"

	"github.com/convox/cli/Godeps/_workspace/src/github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/convox/cli/Godeps/_workspace/src/github.com/aws/aws-sdk-go/service/cloudformation/cloudformationiface"
	"github.com/stretchr/testify/assert"
)

func TestInterface(t *testing.T) {
	assert.Implements(t, (*cloudformationiface.CloudFormationAPI)(nil), cloudformation.New(nil))
}
