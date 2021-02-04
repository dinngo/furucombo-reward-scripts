package rewarder

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path"
)

// MerkleTreeHashSort merkle tree hash sort
func MerkleTreeHashSort(left int, leftHash []byte, right int, rightHash []byte) (int, int) {
	if bytes.Compare(leftHash, rightHash) > 0 {
		return right, left
	}

	return left, right
}

// GenerateRewardMerkleTreeTask generate reward merkle tree task
type GenerateRewardMerkleTreeTask struct {
	rootpath  string
	rewardMap RewardMap

	rewardMerkleTreeLeavesFilepath string
	rewardMerkleProofsFilepath     string
	rewardMerkleTree               *RewardMerkleTree
}

// CheckMerkleTreeFiles check merkle tree files
func (t *GenerateRewardMerkleTreeTask) CheckMerkleTreeFiles() error {
	t.rewardMerkleTreeLeavesFilepath = path.Join(t.rootpath, "merkle_tree_leaves.json")
	t.rewardMerkleProofsFilepath = path.Join(t.rootpath, "merkle_proofs.json")

	log.Printf("check merkle tree leaves file ./%s", t.rewardMerkleTreeLeavesFilepath)

	if _, err := os.Stat(t.rewardMerkleTreeLeavesFilepath); err != nil {
		log.Println("merkle tree leaves file not found")
		return err
	}

	log.Printf("check merkle proofs file ./%s", t.rewardMerkleProofsFilepath)

	if _, err := os.Stat(t.rewardMerkleProofsFilepath); err != nil {
		log.Println("merkle proofs file not found")
		return err
	}

	return nil
}

// GenerateRewardMerkleTree generate reward merkle tree
func (t *GenerateRewardMerkleTreeTask) GenerateRewardMerkleTree() error {
	log.Println("generating reward merkle tree")

	rewardMerkleTree, err := NewRewardMerkleTree(t.rewardMap)
	if err != nil {
		return err
	}

	if err := rewardMerkleTree.GenerateMerkleProofs(); err != nil {
		return err
	}

	t.rewardMerkleTree = rewardMerkleTree

	return nil
}

// SaveRewardMerkleTreeLeavesToFile save rewards to file
func (t *GenerateRewardMerkleTreeTask) SaveRewardMerkleTreeLeavesToFile() error {
	log.Printf("saving merkle tree leaves to ./%s", t.rewardMerkleTreeLeavesFilepath)

	data, err := json.MarshalIndent(t.rewardMerkleTree.MerkleTreeLeaves, "", "  ")
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(t.rewardMerkleTreeLeavesFilepath, append(data, '\n'), 0644); err != nil {
		return err
	}

	return nil
}

// SaveRewardMerkleProofsToFile save reward merkle proofs to file
func (t *GenerateRewardMerkleTreeTask) SaveRewardMerkleProofsToFile() error {
	log.Printf("saving merkle proofs to ./%s", t.rewardMerkleProofsFilepath)

	data, err := json.MarshalIndent(t.rewardMerkleTree.MerkleProofs, "", "  ")
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(t.rewardMerkleProofsFilepath, append(data, '\n'), 0644); err != nil {
		return err
	}

	return nil
}

// Execute execute
func (t *GenerateRewardMerkleTreeTask) Execute() error {
	if err := t.CheckMerkleTreeFiles(); err != nil {
		if err := t.GenerateRewardMerkleTree(); err != nil {
			return err
		}

		if err := t.SaveRewardMerkleTreeLeavesToFile(); err != nil {
			return err
		}

		if err := t.SaveRewardMerkleProofsToFile(); err != nil {
			return err
		}
	}

	return nil
}
