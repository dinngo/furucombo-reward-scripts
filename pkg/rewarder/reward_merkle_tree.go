package rewarder

import (
	"bytes"
	"math/big"
	"sort"

	"github.com/dinngodev/furucombo-reward-scripts/pkg/ethereum"
	"github.com/dinngodev/furucombo-reward-scripts/pkg/ethereum/furucombo"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ksin751119/merkletree"
	"golang.org/x/crypto/sha3"
)

var rewardMerkleTreeConfig = merkletree.TreeConfig{
	HashStrategy: sha3.NewLegacyKeccak256,
	HashSortFunc: func(left int, leftHash []byte, right int, rightHash []byte) (int, int) {
		if bytes.Compare(leftHash, rightHash) > 0 {
			return right, left
		}

		return left, right
	},
}

// -----------------------
// -     Merkle Tree     -
// -----------------------

// RewardMerkleTree struct
type RewardMerkleTree struct {
	MerkleTree       *merkletree.MerkleTree
	MerkleTreeLeaves RewardMerkleTreeLeaves
	MerkleProofs     *RewardMerkleProofs
}

// NewRewardMerkleTree new reward merkle tree
func NewRewardMerkleTree(rewardMap RewardMap) (*RewardMerkleTree, error) {
	merkleTreeLeaves := make(RewardMerkleTreeLeaves, len(rewardMap))
	i := 0
	for account, amount := range rewardMap {
		merkleTreeLeaves[i] = RewardMerkleTreeLeaf{
			Account: account,
			Amount:  ethereum.ToSmallUnit(amount, furucombo.COMBODecimals),
		}
		i++
	}

	sort.Sort(merkleTreeLeaves)

	merkleTree, err := merkletree.NewTreeWithConfig(merkleTreeLeaves.ToMerkleTreeContents(), &rewardMerkleTreeConfig)
	if err != nil {
		return nil, err
	}

	rewardMerkleTree := &RewardMerkleTree{
		MerkleTree:       merkleTree,
		MerkleTreeLeaves: merkleTreeLeaves,
	}

	return rewardMerkleTree, nil
}

// GenerateMerkleProofs generate merkle proofs
func (t *RewardMerkleTree) GenerateMerkleProofs() error {
	if t.MerkleProofs == nil {
		rewards := make(map[common.Address]RewardMerkleProofsReward)

		for _, leaf := range t.MerkleTreeLeaves {
			proofs, _, err := t.MerkleTree.GetMerklePath(leaf)
			if err != nil {
				return err
			}

			proofsItem := RewardMerkleProofsReward{
				Amount: leaf.Amount,
				Proofs: make([]common.Hash, len(proofs)),
			}

			for i, proof := range proofs {
				proofsItem.Proofs[i] = common.BytesToHash(proof)
			}

			rewards[leaf.Account] = proofsItem
		}

		t.MerkleProofs = &RewardMerkleProofs{
			MerkleRoot: common.BytesToHash(t.MerkleTree.MerkleRoot()),
			Rewards:    rewards,
		}
	}

	return nil
}

// ----------------------------
// -     Merkle Tree Leaf     -
// ----------------------------

// RewardMerkleTreeLeaf struct
type RewardMerkleTreeLeaf struct {
	Account common.Address `json:"account"`
	Amount  *big.Int       `json:"amount"`
}

// CalculateHash calculate content hash
func (t RewardMerkleTreeLeaf) CalculateHash() ([]byte, error) {
	hash := crypto.Keccak256(t.Account.Bytes(), common.BytesToHash(t.Amount.Bytes()).Bytes())

	return hash, nil
}

// Equals implement equals function
func (t RewardMerkleTreeLeaf) Equals(other merkletree.Content) (bool, error) {
	return t.Account == other.(RewardMerkleTreeLeaf).Account &&
		t.Amount.Cmp(other.(RewardMerkleTreeLeaf).Amount) == 0, nil
}

// RewardMerkleTreeLeaves type
type RewardMerkleTreeLeaves []RewardMerkleTreeLeaf

// ToMerkleTreeContents to merkle tree contents
func (s RewardMerkleTreeLeaves) ToMerkleTreeContents() []merkletree.Content {
	merkleTreeContents := make([]merkletree.Content, len(s))
	for i, merkleTreeLeaf := range s {
		merkleTreeContents[i] = merkleTreeLeaf
	}

	return merkleTreeContents
}

func (s RewardMerkleTreeLeaves) Len() int { return len(s) }

func (s RewardMerkleTreeLeaves) Less(i, j int) bool {
	return bytes.Compare(s[i].Account.Bytes(), s[j].Account.Bytes()) < 0
}

func (s RewardMerkleTreeLeaves) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

// ------------------------
// -     Merkle Proof     -
// ------------------------

// RewardMerkleProofs struct
type RewardMerkleProofs struct {
	MerkleRoot common.Hash                                 `json:"merkle_root"`
	Rewards    map[common.Address]RewardMerkleProofsReward `json:"rewards"`
}

// RewardMerkleProofsReward struct
type RewardMerkleProofsReward struct {
	Amount *big.Int      `json:"amount"`
	Proofs []common.Hash `json:"proofs"`
}
