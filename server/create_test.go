package server

import (
	"testing"

	"github.com/baopham/reminder"
	"github.com/stretchr/testify/assert"
	context "golang.org/x/net/context"
)

func TestCreate(t *testing.T) {
	ctx := context.Background()
	req := &reminder.CreateRequest{}

	res, err := cli.Create(ctx, req)
	assert.Nil(t, err)
	assert.NotNil(t, res)
}
