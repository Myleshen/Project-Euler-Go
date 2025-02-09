package problems_test

import (
	"testing"

	"github.com/myleshen/project-euler-go/problems"
	"github.com/stretchr/testify/assert"
)

func TestProblem7(t *testing.T) {
	assert.Equal(t, true, problems.IsPrime(2))
	assert.Equal(t, true, problems.IsPrime(3))
	assert.Equal(t, false, problems.IsPrime(4))
	assert.Equal(t, true, problems.IsPrime(11))
	assert.Equal(t, 13, problems.Problem7(6))
	assert.Equal(t, 104743, problems.Problem7(10001))
}
