package problems_test

import (
	"testing"

	"github.com/myleshen/project-euler-go/problems"
	"github.com/stretchr/testify/assert"
)

func TestProblem6(t *testing.T) {
	assert.Equal(t, 2640, problems.Problem6(10))
	assert.Equal(t, 25164150, problems.Problem6(100))
}
