package rewarder

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// Cube represents valid volume of a token
type Cube struct {
	Name         string
	TokenAddress common.Address
	TokenAmount  *big.Int
}

// CubeFinder type
type CubeFinder func(txLog *types.Log) (*Cube, error)

// cubeFinderMap find cube map
var cubeFinderMap = map[string]CubeFinder{}

// GetCubeFinders get cube finders
func GetCubeFinders(cubeNames []string) ([]CubeFinder, error) {
	cubeFinders := make([]CubeFinder, len(cubeNames))

	for i, name := range cubeNames {
		cubeFinder, ok := cubeFinderMap[name]
		if !ok {
			return nil, fmt.Errorf("unsupported cube: %s", name)
		}

		cubeFinders[i] = cubeFinder
	}

	return cubeFinders, nil
}

// CubeFinders type
type CubeFinders []CubeFinder

// Find find cube
func (finders CubeFinders) Find(txLog *types.Log) (*Cube, error) {
	for _, finder := range finders {
		cube, err := finder(txLog)
		if err != nil {
			return nil, err
		}

		if cube != nil {
			return cube, nil
		}
	}

	return nil, nil
}
