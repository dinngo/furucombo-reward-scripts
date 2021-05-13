package rewarder

// UsageRewarderRequiredFieldNames usage rewarder required field names
var UsageRewarderRequiredFieldNames = []string{"name", "startBlock", "endBlock", "blocks", "reward_amount", "max_gas_used", "nfts", "nft_boost", "nft_max_boost"}

// UsageRewarder struct
type UsageRewarder struct {
	config *Config

	tradingGasComboMap TradingGasComboMap
	nftCountMap        NftCountMap
	rewardsMap         RewardMap
}

// NewUsageRewarder new usage rewarder
func NewUsageRewarder(config *Config) *UsageRewarder {
	return &UsageRewarder{config: config}
}

// LoadTradingGasCombo load trading gas
func (r *UsageRewarder) LoadTradingGasCombo() error {
	task := LoadTradingGasComboTask{
		rootpath:       r.config.RoundDir(),
		startBlock:     r.config.StartBlock,
		endBlock:       r.config.EndBlock,
		startTimestamp: r.config.startTimestamp,
		endTimestamp:   r.config.endTimestamp,
		maxGasUsed:     r.config.MaxGasUsed,
	}

	if err := task.Execute(); err != nil {
		return err
	}

	r.tradingGasComboMap = task.tradingGasComboMap

	return nil
}

func (r *UsageRewarder) LoadNftCounts() error {
	task := LoadNftCountsTask{
		rootpath: r.config.RoundDir(),
		endBlock: r.config.EndBlock,
		nfts:     r.config.Nfts,
	}

	if err := task.Execute(); err != nil {
		return err
	}

	r.nftCountMap = task.nftCountMap

	return nil
}

// LoadRewards load rewards
func (r *UsageRewarder) LoadGasRewards() error {
	r.rewardsMap = make(RewardMap)

	task := LoadGasRewardsTask{
		rootpath:           r.config.RoundDir(),
		rewardAmount:       r.config.RewardAmount,
		tradingGasComboMap: r.tradingGasComboMap,
		nftCountMap:        r.nftCountMap,
		nftBoost:           r.config.NftBoost,
		nftMaxBoost:        r.config.NftMaxBoost,
	}

	if err := task.Execute(); err != nil {
		return err
	}

	r.rewardsMap = task.rewardMap

	return nil
}

// GenerateRewardsMerkleTree generate rewards merkle tree
func (r *UsageRewarder) GenerateRewardsMerkleTree() error {
	task := GenerateRewardMerkleTreeTask{
		rootpath:  r.config.RoundDir(),
		rewardMap: r.rewardsMap,
	}

	if err := task.Execute(); err != nil {
		return err
	}

	return nil
}
