package save

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/tsundata/assistant/internal/app/bot/plugin/end"
	"testing"
)

func TestSave(t *testing.T) {
	p := Save{
		Next: end.End{},
	}
	input := "test"
	out, err := p.Run(context.Background(), input)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, input, out)
}
