package renderer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_escapeFirstAsterixInEachPair(t *testing.T) {
	assert.Equal(t, "[a-z]", escapeFirstAsterixInEachPair("[a-z]"))
	assert.Equal(t, "[a-z]*", escapeFirstAsterixInEachPair("[a-z]*"))
	assert.Equal(t, `0\*[a-z]*`, escapeFirstAsterixInEachPair(`0*[a-z]*`))
	assert.Equal(t, `0\*[a-z]*[a-z]*[0-9]`, escapeFirstAsterixInEachPair(`0*[a-z]*[a-z]*[0-9]`))
	assert.Equal(t, `0\*[a-z]*[a-z]\*[0-9]*`, escapeFirstAsterixInEachPair(`0*[a-z]*[a-z]*[0-9]*`))
}
