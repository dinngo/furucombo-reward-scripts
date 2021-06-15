package rewarder

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"math/big"
	"os"
	"path"

	"github.com/dinngodev/furucombo-reward-scripts/pkg/ethereum"
	"github.com/dinngodev/furucombo-reward-scripts/pkg/ethereum/erc721"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

// LoadNftCountsTask load nft counts task
type LoadNftCountsTask struct {
	rootpath        string
	endBlock        uint64
	nft             NftConfig
	polygonEndBlock uint64

	filepath    string
	nftCountMap NftCountMap
}

// LoadFromFile load from file
func (t *LoadNftCountsTask) LoadFromFile() error {
	t.filepath = path.Join(t.rootpath, "nft_counts.json")

	log.Printf("loading nft counts from ./%s", t.filepath)

	if _, err := os.Stat(t.filepath); err != nil {
		log.Println("nft counts file not found")
		return err
	}

	data, err := ioutil.ReadFile(t.filepath)
	if err != nil {
		return err
	}

	if err := jsonex.Unmarshal(data, &t.nftCountMap); err != nil {
		return err
	}

	return nil
}

// GetNftCounts get nft counts
func (t *LoadNftCountsTask) GetNftCounts() error {
	t.nftCountMap = make(NftCountMap)

	// Count ethereum nft
	for _, nft := range t.nft.Ethereum {
		log.Printf("getting ethereum nft counts")
		nft, err := erc721.NewERC721Contract(nft, ethereum.Client())
		if err != nil {
			return err
		}
		t.accumulateNft(nft, t.endBlock)
	}

	// Count polygon nft
	for _, nft := range t.nft.Polygon {
		log.Printf("getting polygon nft counts")
		nft, err := erc721.NewERC721Contract(nft, ethereum.ClientPolygon())
		if err != nil {
			return err
		}
		t.accumulateNft(nft, t.polygonEndBlock)
	}

	return nil
}

func (t *LoadNftCountsTask) accumulateNft(nft *erc721.ERC721Contract, block uint64) error {
	opts := &bind.CallOpts{
		BlockNumber: big.NewInt(int64(block)),
	}

	totalSupply, err := nft.TotalSupply(opts)
	if err != nil {
		return err
	}

	for i := big.NewInt(1); i.Cmp(totalSupply) <= 0; i = new(big.Int).Add(i, big.NewInt(1)) {
		owner, err := nft.OwnerOf(opts, i)
		if err != nil {
			return err
		}
		t.nftCountMap[owner]++
	}

	return nil
}

// SaveToFile save nft counts to file
func (t *LoadNftCountsTask) SaveToFile() error {
	log.Printf("saving nft counts to ./%s", t.filepath)

	data, err := json.MarshalIndent(t.nftCountMap, "", "  ")
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(t.filepath, append(data, '\n'), 0644); err != nil {
		return err
	}

	return nil
}

// Execute execute
func (t *LoadNftCountsTask) Execute() error {
	if err := t.LoadFromFile(); err != nil {
		if err := t.GetNftCounts(); err != nil {
			return err
		}

		if err := t.SaveToFile(); err != nil {
			return err
		}
	}

	return nil
}
