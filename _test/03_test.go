package problems_test

import (
	"testing"

	"github.com/myleshen/project-euler-go/problems"
	"github.com/stretchr/testify/assert"
)

func TestProblem3(t *testing.T) {
	assert.Equal(t, []int{11}, problems.Problem3(11))
	assert.Equal(t, []int{5, 7}, problems.Problem3(35))
	assert.Equal(t, []int{7, 11}, problems.Problem3(77))
	assert.Equal(t, []int{5, 7, 13, 29}, problems.Problem3(13195))
	assert.Equal(t, []int{71, 839, 1471, 6857}, problems.Problem3(600851475143))
}
