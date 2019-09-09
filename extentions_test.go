package extentions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func testIsEmpty(t *testing.T) {
	str1 := ""
	str2 := "I have a value"

	isEmpty := IsEmpty(str1)
	assert.Equal(t, true, isEmpty, "This variable should result emtpy is true")

	isEmpty = IsEmpty(str2)
	assert.Equal(t, false, isEmpty, "This variable should result emtpy is false")

}

func testIsEqual(t *testing.T) {
	str1 := "I am"
	str2 := "You are"
	str3 := "I am"

	isEqual := IsEqual(str1, str2)
	assert.Equal(t, false, isEqual, "str1 and str2 should equal false")

	isEqual = IsEqual(str2, str3)
	assert.Equal(t, true, isEqual, "str2 and str3 should equal true")

}
