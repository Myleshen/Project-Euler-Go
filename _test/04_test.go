package problems_test

import (
	"testing"

	"github.com/myleshen/project-euler-go/problems"
	"github.com/stretchr/testify/assert"
)

func TestProblem4(t *testing.T) {
	// assert.Equal(t, []int{125, 800}, problems.Problem4())
	assert.Equal(t, true, problems.IsPalindrome(121))
	assert.Equal(t, false, problems.IsPalindrome(1000))
	assert.Equal(t, 906609, problems.Problem4())
}
