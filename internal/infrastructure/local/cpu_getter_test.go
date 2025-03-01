package local

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewCpuGetter(t *testing.T) {
	g := NewLocalCpuGetter()
	c, err := g.Get(context.TODO())
	assert.NoError(t, err)
	fmt.Println(c.String())
}
