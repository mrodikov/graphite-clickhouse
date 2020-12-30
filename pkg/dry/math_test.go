package dry

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMax(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(int64(1), Max(1, -1))
	assert.Equal(int64(2), Max(1, 2))
	assert.Equal(int64(3), Max(3, 3))
}

func TestMin(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(int64(-1), Min(1, -1))
	assert.Equal(int64(1), Min(1, 2))
	assert.Equal(int64(3), Min(3, 3))
}

func TestCeil(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(int64(0), Ceil(0, -1))
	assert.Equal(int64(3), Ceil(5, 2))
	assert.Equal(int64(1), Ceil(5, 5))           // if quotient is integer we should get quotient without +1
	assert.Equal(int64(2), Ceil(100001, 100000)) // if quotient is any fraction bigger than integer then we get +1
}

func TestCeilToMultiplier(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(int64(0), CeilToMultiplier(0, -1))
	assert.Equal(int64(0), CeilToMultiplier(1, 0))
	assert.Equal(int64(0), CeilToMultiplier(1, -1))
	assert.Equal(int64(2), CeilToMultiplier(1, 2))
	assert.Equal(int64(6), CeilToMultiplier(4, 3))
	assert.Equal(int64(6), CeilToMultiplier(6, 3))
}

func TestFloorToMultiplier(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(int64(0), FloorToMultiplier(0, -1))
	assert.Equal(int64(0), FloorToMultiplier(1, 0))
	assert.Equal(int64(0), FloorToMultiplier(1, -1))
	assert.Equal(int64(0), FloorToMultiplier(1, 2))
	assert.Equal(int64(3), FloorToMultiplier(4, 3))
	assert.Equal(int64(6), FloorToMultiplier(6, 3))
}

func TestGCD(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(int64(1), GCD(1, -1))
	assert.Equal(int64(1), GCD(-1, 1))
	assert.Equal(int64(1), GCD(-1, -1))
	assert.Equal(int64(1), GCD(1, 2))
	assert.Equal(int64(1), GCD(4, 3))
	assert.Equal(int64(3), GCD(6, 3))
}

func TestLCM(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(int64(1), LCM(1, -1))
	assert.Equal(int64(1), LCM(-1, 1))
	assert.Equal(int64(1), LCM(-1, -1))
	assert.Equal(int64(2), LCM(1, 2))
	assert.Equal(int64(6), LCM(6, 3))
	assert.Equal(int64(12), LCM(4, 3))
}
