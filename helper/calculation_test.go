package helper

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// yang perlu di test adalah hal2 yang menggangu flow kita sebagai developer
func TestSumRequire(t *testing.T) {
	// Use require to strictly enforce expectations
	res := Sum(1, 2)
	require.Equal(t, 3, res, "Result should be 3")
	require.NotNil(t, res, "Result should not be nil")
}

func TestSumAssert(t *testing.T) {
	// Use assert for soft assertions (execution continues on failure)
	res := Sum(1, 2)
	assert.Equal(t, 3, res, "Result should be 3")
	assert.NotNil(t, res, "Result should not be nil")
}

func TestSum(t * testing.T){
	tests := []struct {
		Expected int
		Actual int
		errMsg string
	}{
		{Actual: Sum(2, 2), Expected: 4, errMsg: "Invalid Sum Result"},
		{Actual: Sum(1, 3), Expected: 4, errMsg: "Invalid Sum Result"},
		{Actual: Sum(5, 12), Expected: 17, errMsg: "Invalid Sum Result"},
		{Actual: Sum(12, 20), Expected: 32, errMsg: "Invalid Sum Result"},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("Test Sum %d : ", test.Actual), func(t *testing.T) {
			assert.Equal(t, test.Expected, test.Actual, test.errMsg)
		})
	}
}