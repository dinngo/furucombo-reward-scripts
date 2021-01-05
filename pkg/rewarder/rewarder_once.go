package rewarder

import "path"

// OnceRewarderRequiredFieldNames once rewarder required field names
var OnceRewarderRequiredFieldNames = []string{"name", "rewards"}

// OnceRewarder struct
type OnceRewarder struct {
	config *Config
}

// NewOnceRewarder new once rewarder
func NewOnceRewarder(config *Config) *OnceRewarder {
	return &OnceRewarder{config: config}
}

// GenerateRewardMerkleTree generate rewards merkle tree
func (r *OnceRewarder) GenerateRewardMerkleTree() error {
	task := GenerateRewardMerkleTreeTask{
		rewardMap:                      r.config.RewardMap,
		rewardMerkleTreeLeavesFilepath: path.Join(r.config.RoundDir(), "merkle_tree_leaves.json"),
		rewardMerkleProofsFilepath:     path.Join(r.config.RoundDir(), "merkle_proofs.json"),
	}

	if err := task.CheckMerkleTreeFiles(); err != nil {
		if err := task.GenerateRewardMerkleTree(); err != nil {
			return err
		}

		if err := task.SaveRewardMerkleTreeLeavesToFile(); err != nil {
			return err
		}

		if err := task.SaveRewardMerkleProofsToFile(); err != nil {
			return err
		}
	}

	return nil
}
