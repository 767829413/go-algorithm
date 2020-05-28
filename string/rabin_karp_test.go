package string

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRabinKarpMatch_1(t *testing.T) {
	str := "asdsdasdgwwcfghedcf"
	subStr := "ghe"
	assert.Equal(t, 13, RabinKarpMatch_1(str, subStr))
}
