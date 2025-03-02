package local

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMemoryGetter_Get(t *testing.T) {
	g := NewMemoryGetter()
	m, err := g.Get(context.TODO())
	assert.NoError(t, err)
	fmt.Println(m.String())
}
