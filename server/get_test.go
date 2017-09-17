package server

import (
	"testing"

	"github.com/baopham/reminder"
	"github.com/stretchr/testify/assert"
	context "golang.org/x/net/context"
)

func TestGet(t *testing.T) {
	ctx := context.Background()
	req := &reminder.GetRequest{}

	res, err := cli.Get(ctx, req)
	assert.Nil(t, err)
	assert.NotNil(t, res)
}
