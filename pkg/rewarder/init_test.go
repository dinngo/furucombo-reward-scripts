package rewarder

import (
	"path"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestLoadDotEnv(t *testing.T) {
	err := godotenv.Load(path.Join("..", "..", ".env.test"))
	assert.Nil(t, err)
}
