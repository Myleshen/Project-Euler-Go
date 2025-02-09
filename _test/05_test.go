package problems_test

import (
	"testing"

	"github.com/myleshen/project-euler-go/problems"
	"github.com/stretchr/testify/assert"
)

func TestProblem5(t *testing.T) {
	assert.Equal(t, 232792560, problems.Problem5())
}
