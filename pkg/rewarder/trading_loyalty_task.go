package rewarder

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"math"
	"math/big"
	"net/http"
	"os"
	"path"
	"time"

	"github.com/dinngodev/furucombo-reward-scripts/pkg/ethereum"
	"github.com/dinngodev/furucombo-reward-scripts/pkg/ethereum/erc721"
	"github.com/dinngodev/furucombo-reward-scripts/pkg/etherscan"
	"github.com/dinngodev/furucombo-reward-scripts/pkg/polygon/aavev2"
	"github.com/dinngodev/furucombo-reward-scripts/pkg/polygon/curve"
	"github.com/dinngodev/furucombo-reward-scripts/pkg/polygon/furucombo"
	"github.com/dinngodev/furucombo-reward-scripts/pkg/polygon/quickswap"
	"github.com/dinngodev/furucombo-reward-scripts/pkg/polygon/sushiswap"
	"github.com/dinngodev/furucombo-reward-scripts/pkg/polygon/wmatic"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/shopspring/decimal"
)

// LoadTradingLoyaltyTask load trading loyalty task
type LoadTradingLoyaltyTask struct {
	// Exterior
	rootpath        string
	startBlock      uint64
	endBlock        uint64
	maxLoyalTxCount decimal.Decimal
	nft             NftConfig
	bridgeTxMap     BridgeTxMap

	// Intermedium
	filepath           string
	polygonTasteNftMap map[common.Address]int

	// File
	tradingLoyaltyMap TradingLoyaltyMap
}

// LoadFromFile load from file
func (t *LoadTradingLoyaltyTask) LoadFromFile() error {
	t.filepath = path.Join(t.rootpath, "trading_loyalty.json")

	log.Printf("loading trading loyalty from ./%s", t.filepath)

	if _, err := os.Stat(t.filepath); err != nil {
		log.Println("trading loyalty file not found")
		return err
	}

	data, err := ioutil.ReadFile(t.filepath)
	if err != nil {
		return err
	}

	if err := jsonex.Unmarshal(data, &t.tradingLoyaltyMap); err != nil {
		return err
	}

	return nil
}

// GetNftPolygonTaste get polygon paste nft
func (t *LoadTradingLoyaltyTask) GetPolygonTasteNft() error {
	log.Printf("getting user list of polygon paste nft")

	// get user list of polygon paste nft
	t.polygonTasteNftMap = make(map[common.Address]int)
	nft, err := erc721.NewERC721Contract(t.nft.Polygon[0], ethereum.ClientPolygon())
	if err != nil {
		return err
	}

	tokenIdStart := 1
	tokenIdEnd := 30
	opts := &bind.CallOpts{
		BlockNumber: big.NewInt(int64(t.endBlock)),
	}

	for i := tokenIdStart; i <= tokenIdEnd; i++ {
		owner, err := nft.OwnerOf(opts, big.NewInt(int64(i)))
		if err != nil {
			return err
		}
		t.polygonTasteNftMap[owner]++
	}

	return nil
}

// CalcLoyalty calculate loyalty = nftBoost * loyalTxCount
func (t *LoadTradingLoyaltyTask) CalcLoyalty() error {
	log.Printf("getting trading loyalty from etherscan")

	apiKey := os.Getenv("POLYGONSCAN_API_KEY")

	if len(apiKey) == 0 {
		log.Fatalln("env POLYGONSCAN_API_KEY can't be blank")
	}

	client := etherscan.NewClient(&http.Client{Timeout: 10 * time.Second}, apiKey)

	t.tradingLoyaltyMap = make(TradingLoyaltyMap)
	for user := range t.bridgeTxMap {
		// Calculate polygon taste nftBoost
		var nftBoost float64 = 1
		if count, ok := t.polygonTasteNftMap[user]; ok {
			nftBoost = math.Pow(t.nft.Boost, float64(count))
			if nftBoost > t.nft.MaxBoost {
				nftBoost = t.nft.MaxBoost
			}
		}

		// Calculate loyalTxCount
		params := etherscan.Params{
			"address":    user.String(),
			"startBlock": t.startBlock,
			"endBlock":   t.endBlock,
			"sort":       "asc",
		}
		txs1, err := client.AccountTxsPolygon(params)
		if err != nil {
			return err
		}

		var txCount, officialTxCount, furucomboTxCount decimal.Decimal
		for _, tx := range txs1 {
			// Ignore failed tx
			if tx.IsError == 1 {
				continue
			}

			// Ignore user as polygon receiver
			if user != tx.From {
				continue
			}

			// Count total tx
			txCount = txCount.Add(decimal.NewFromInt(1))

			// Count official tx
			if aavev2.IsLendingPoolAddress(tx.To) ||
				curve.IsSwapAddress(tx.To) ||
				sushiswap.IsRouter02Address(tx.To) ||
				quickswap.IsRouter02Address(tx.To) ||
				(wmatic.IsWmaticAddress(tx.To) && wmatic.HasDepositOrWithdrawalEvent(tx.Hash)) { // only count tx including two events
				officialTxCount = officialTxCount.Add(decimal.NewFromInt(1))
			}

			// Count furucombo tx
			if furucombo.IsProxyAddress(tx.To) {
				furucomboTxCount = furucomboTxCount.Add(decimal.NewFromInt(1))
			}
		}

		if txCount.IsZero() {
			continue
		}

		// Calculate loyalTxCount = min(furu * (1 - official / total), maxLoyalTxCount)
		loyalTxCount := furucomboTxCount.Mul(decimal.NewFromInt(1).Sub(officialTxCount.Div(txCount))).Truncate(2)
		if loyalTxCount.GreaterThan(t.maxLoyalTxCount) {
			loyalTxCount = t.maxLoyalTxCount
		}

		// Calculate loyalty = nftBoost * loyalTxCount
		loyalty := decimal.NewFromFloat(nftBoost).Mul(loyalTxCount).Truncate(2)
		log.Printf(
			"found user %s loyalty %s (nftBoost: %.2f loyalTxCount %s furucomboTxCount: %s officialTxCount: %s txCount: %s)",
			user.String(), loyalty, nftBoost, loyalTxCount, furucomboTxCount, officialTxCount, txCount)

		if loyalty.IsPositive() {
			t.tradingLoyaltyMap[user] = &Loyalty{
				NftBoost:         decimal.NewFromFloat(nftBoost),
				FurucomboTxCount: furucomboTxCount,
				OfficialTxCount:  officialTxCount,
				TxCount:          txCount,
				LoyalTxCount:     loyalTxCount,
				Loyalty:          loyalty,
			}
		}
	}

	return nil
}

// SaveToFile save trading counts to file
func (t *LoadTradingLoyaltyTask) SaveToFile() error {
	log.Printf("saving trading counts to ./%s", t.filepath)

	data, err := json.MarshalIndent(t.tradingLoyaltyMap, "", "  ")
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(t.filepath, append(data, '\n'), 0644); err != nil {
		return err
	}

	return nil
}

// Execute execute
func (t *LoadTradingLoyaltyTask) Execute() error {
	if err := t.LoadFromFile(); err != nil {
		if err := t.GetPolygonTasteNft(); err != nil {
			return err
		}

		if err := t.CalcLoyalty(); err != nil {
			return err
		}

		if err := t.SaveToFile(); err != nil {
			return err
		}
	}

	return nil
}
