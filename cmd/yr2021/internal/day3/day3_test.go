package day3

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var _TEST_LINES = []string{
	"00100",
	"11110",
	"10110",
	"10111",
	"10101",
	"01111",
	"00111",
	"11100",
	"10000",
	"11001",
	"00010",
	"01010",
}

func TestBinToInt(t *testing.T) {
	i, err := btoi([]int{1, 0, 1, 1, 0})
	assert.NoError(t, err)
	assert.Equal(t, int64(22), i)
}

func TestSums(t *testing.T) {
	assert.Equal(t, []int{0}, sums([]string{"0"}))
	assert.Equal(t, []int{1}, sums([]string{"1"}))
	assert.Equal(t, []int{2}, sums([]string{"1", "1"}))
	assert.Equal(t, []int{1, 0}, sums([]string{"10"}))
	assert.Equal(t, []int{2, 0}, sums([]string{"10", "10"}))
	assert.Equal(t, []int{0, 2}, sums([]string{"01", "01"}))
	assert.Equal(t, []int{1, 1, 2}, sums([]string{"101", "011"}))
	assert.Equal(t, []int{7, 5, 8, 7, 5}, sums(_TEST_LINES))
}

func TestCommonBits(t *testing.T) {
	assert.Equal(t, []int{1, 0, 1, 1, 0}, commonBits(_TEST_LINES, 1))
	assert.Equal(t, []int{1, 0, 1, 0, 0}, commonBits([]string{
		"11110",
		"10110",
		"10111",
		"10101",
		"11100",
		"10000",
		"11001",
	}, 1))
}

func TestPower(t *testing.T) {
	p, err := power(_TEST_LINES)
	assert.NoError(t, err)
	assert.Equal(t, int64(198), p)
}

func TestOxygen(t *testing.T) {
	o, err := oxygen(_TEST_LINES)
	assert.NoError(t, err)
	assert.Equal(t, int64(23), o)
}

func TestCo2(t *testing.T) {
	o, err := co2(_TEST_LINES)
	assert.NoError(t, err)
	assert.Equal(t, int64(10), o)
}

func TestLifeSupport(t *testing.T) {
	l, err := lifeSupport(_TEST_LINES)
	assert.NoError(t, err)
	assert.Equal(t, int64(230), l)
}
