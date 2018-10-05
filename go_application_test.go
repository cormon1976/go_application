package sample

import (
  "testing"
  "github.com/stretchr/testify/assert"
)
func Test_adder(t *testing.T) {
  assert := assert.New(t)
  assert.Equal(adder(3, 5), 8)
}

func Test_suber(t *testing.T) {
  assert := assert.New(t)
  assert.Equal(subtracter(2, 1), 1)
}