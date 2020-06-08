package backtrack

import "testing"

func TestPattern_Match(t *testing.T) {
	patern := []rune("assh*666")
	text := []rune("asshsdfgdf的风格的发挥肺结核沙发上地方assh当时发生了来了bbb好666sjdkhfkjsdf的撒防守对方666")
	p := NewPattern(patern, len(patern))
	t.Log(p.Match(text, len(text)))
}
