package string

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBruteForceMatch(t *testing.T) {
	str := "asdsdasdgwwcfghedcf"
	subStr := "ghe"
	assert.Equal(t, 13, BruteForceMatch(str, subStr))
}
