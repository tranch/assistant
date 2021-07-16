package classifier

import (
	"github.com/stretchr/testify/require"
	"github.com/tsundata/assistant/api/model"
	"testing"
)

func TestRule(t *testing.T) {
	r := NewRule()
	r.Format = "test1 > str"
	a, err := r.Do("test1")
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, model.StrengthAttr, a)
}

func TestRule2(t *testing.T) {
	r := NewRule()
	r.Format = "test1 > cul"
	a, err := r.Do("test1")
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, model.CultureAttr, a)
}

func TestRule3(t *testing.T) {
	r := NewRule()
	r.Format = "test1 > env"
	a, err := r.Do("test1")
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, model.EnvironmentAttr, a)
}

func TestRule4(t *testing.T) {
	r := NewRule()
	r.Format = "test1 > cha"
	a, err := r.Do("test1")
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, model.CharismaAttr, a)
}

func TestRule5(t *testing.T) {
	r := NewRule()
	r.Format = "test1 > tal"
	a, err := r.Do("test1")
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, model.TalentAttr, a)
}

func TestRule6(t *testing.T) {
	r := NewRule()
	r.Format = "test1 > int"
	a, err := r.Do("test1")
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, model.IntellectAttr, a)
}

func TestRule7(t *testing.T) {
	r := NewRule()
	r.Format = "test1|test2|test3 > int"
	a, err := r.Do("test3")
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, model.IntellectAttr, a)
}

func TestRule8(t *testing.T) {
	r := NewRule()
	r.Format = "test1|test2|test3 > int"
	_, err := r.Do("test0")
	require.ErrorIs(t, ErrEmpty, err)
}
