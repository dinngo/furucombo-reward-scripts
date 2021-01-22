package rewarder

import (
	"fmt"
	"sort"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

func TestRewardMerkleTreeLeavesSort(t *testing.T) {
	leaves := RewardMerkleTreeLeaves{
		{
			Account: common.HexToAddress("0x24570C86f029E562Bdb05236bC416e8465a49A02"),
			Amount:  decimal.NewFromInt(2),
		},
		{
			Account: common.HexToAddress("0x52Ad61Fee56e1Dc0BF2C2Da48eDB23d4066DBB3a"),
			Amount:  decimal.NewFromInt(4),
		},
		{
			Account: common.HexToAddress("0x2ca5bd4EcBa9d159Bb4Ad035832cb60a0024D722"),
			Amount:  decimal.NewFromInt(3),
		},
		{
			Account: common.HexToAddress("0x061517CbA46D7c9230C2caa560d1D33680AEE63e"),
			Amount:  decimal.NewFromInt(1),
		},
	}

	sort.Sort(leaves)

	assert.Equal(t, "0x061517CbA46D7c9230C2caa560d1D33680AEE63e", leaves[0].Account.String())
	assert.Equal(t, "0x24570C86f029E562Bdb05236bC416e8465a49A02", leaves[1].Account.String())
	assert.Equal(t, "0x2ca5bd4EcBa9d159Bb4Ad035832cb60a0024D722", leaves[2].Account.String())
	assert.Equal(t, "0x52Ad61Fee56e1Dc0BF2C2Da48eDB23d4066DBB3a", leaves[3].Account.String())
}

func TestRewardMerkleTree(t *testing.T) {
	testCases := []struct {
		rewardMap RewardMap
		expected  RewardMerkleProofs
	}{
		{
			rewardMap: RewardMap{
				common.HexToAddress("0x2f2fcc239a6c5c28f62a4ebe4ffdd180e00663b9"): decimal.RequireFromString("30026.172489673112739215"),
				common.HexToAddress("0x6bc8b3a2773a92c0941eda83611f3063a090c77b"): decimal.RequireFromString("18274.355650261064137294"),
				common.HexToAddress("0x74b746c8d2652cb198a9b0d87bd2315f639caed0"): decimal.RequireFromString("3515.642126116526628746"),
				common.HexToAddress("0x851f9de75c021c90c40e209ffba0488b0320c46f"): decimal.RequireFromString("26869.104389306808913956"),
				common.HexToAddress("0xaa293a146aaf9e05bedd1ff29b0da5bd8be70955"): decimal.RequireFromString("4128.323938714978113081"),
				common.HexToAddress("0xb066c21e2d4ace41bc82a1fdb3d53e385d6fc24b"): decimal.RequireFromString("26872.771430477964988377"),
				common.HexToAddress("0xe02de734d92293ee5725b5e82404aab6d19ad5e9"): decimal.RequireFromString("12311.414367069265458088"),
				common.HexToAddress("0xf849cfcf783f094e66a18da660938013f903817a"): decimal.RequireFromString("3002.215608380279021243"),
			},
			expected: RewardMerkleProofs{
				MerkleRoot: common.HexToHash("0x38cfc04a0694813f36299ff8981f4e1754e0cd21464d1e0c14d814a0b00f07bd"),
				Rewards: map[common.Address]RewardMerkleProofsReward{
					common.HexToAddress("0x2f2fcc239a6c5c28f62a4ebe4ffdd180e00663b9"): {
						Amount: decimal.RequireFromString("30026172489673112739215"),
						Proofs: []common.Hash{
							common.HexToHash("0x867f191b764c2c41539f3b9f58cdda61bcd13d1d9a3e5a8c2fc5ee6b260547f4"),
							common.HexToHash("0x314008acffb90f2819f65a98f9961a5fb30dac175caa7c1340c3499feec5f375"),
							common.HexToHash("0x0f1a570432833316307848d37fcd29b95aaefa4255d27a72a5e2f9a065fad2cb"),
						},
					},
					common.HexToAddress("0x6bc8b3a2773a92c0941eda83611f3063a090c77b"): {
						Amount: decimal.RequireFromString("18274355650261064137294"),
						Proofs: []common.Hash{
							common.HexToHash("0xc172b2320d0c1d677d9c4dda7fcc5688ebd84f69da7661ec4a7212d6bb3bf227"),
							common.HexToHash("0x314008acffb90f2819f65a98f9961a5fb30dac175caa7c1340c3499feec5f375"),
							common.HexToHash("0x0f1a570432833316307848d37fcd29b95aaefa4255d27a72a5e2f9a065fad2cb"),
						},
					},
					common.HexToAddress("0x74b746c8d2652cb198a9b0d87bd2315f639caed0"): {
						Amount: decimal.RequireFromString("3515642126116526628746"),
						Proofs: []common.Hash{
							common.HexToHash("0xb2df49cbb1937051e4c355ac20dbe335bcc65f601d33b9ce0c5c0ededc77673d"),
							common.HexToHash("0x9296a4ff7a896c910ce5a8e0167ef755137d7ffc53cd757fc95b24f9adfc40a9"),
							common.HexToHash("0x0f1a570432833316307848d37fcd29b95aaefa4255d27a72a5e2f9a065fad2cb"),
						},
					},
					common.HexToAddress("0x851f9de75c021c90c40e209ffba0488b0320c46f"): {
						Amount: decimal.RequireFromString("26869104389306808913956"),
						Proofs: []common.Hash{
							common.HexToHash("0xd1c099841be469e1c55459644ff057f8992491b013f2a4401297c6669b7cd20d"),
							common.HexToHash("0x9296a4ff7a896c910ce5a8e0167ef755137d7ffc53cd757fc95b24f9adfc40a9"),
							common.HexToHash("0x0f1a570432833316307848d37fcd29b95aaefa4255d27a72a5e2f9a065fad2cb"),
						},
					},
					common.HexToAddress("0xaa293a146aaf9e05bedd1ff29b0da5bd8be70955"): {
						Amount: decimal.RequireFromString("4128323938714978113081"),
						Proofs: []common.Hash{
							common.HexToHash("0xdcfd8c8bceeed55ec4ae5367d97687251cbb82532183fba3e2ac6257f020ebb9"),
							common.HexToHash("0xce41f96f431f6a3a5a6ee9b355d18986bd0a8aeeeec3a9be8b2577ffd967a37c"),
							common.HexToHash("0x28c968f6e05241a80dcac92c94c61f439a4ab2bde173288927cb3616c8a19b36"),
						},
					},
					common.HexToAddress("0xb066c21e2d4ace41bc82a1fdb3d53e385d6fc24b"): {
						Amount: decimal.RequireFromString("26872771430477964988377"),
						Proofs: []common.Hash{
							common.HexToHash("0xe05ca0996eaf3b3c96ada2be9be39fff0f9ed8bb7239f8c9f43ebd5919922a92"),
							common.HexToHash("0xce41f96f431f6a3a5a6ee9b355d18986bd0a8aeeeec3a9be8b2577ffd967a37c"),
							common.HexToHash("0x28c968f6e05241a80dcac92c94c61f439a4ab2bde173288927cb3616c8a19b36"),
						},
					},
					common.HexToAddress("0xe02de734d92293ee5725b5e82404aab6d19ad5e9"): {
						Amount: decimal.RequireFromString("12311414367069265458088"),
						Proofs: []common.Hash{
							common.HexToHash("0x7434c8daef0b1353d60baef8a5e0e48cdb798eb81d3e45ad01c51abfc0022f5a"),
							common.HexToHash("0x5ed47c8ad5f44354769602aaecc2d44cd6d76024155ae4bad4ca279bf0fb6da2"),
							common.HexToHash("0x28c968f6e05241a80dcac92c94c61f439a4ab2bde173288927cb3616c8a19b36"),
						},
					},
					common.HexToAddress("0xf849cfcf783f094e66a18da660938013f903817a"): {
						Amount: decimal.RequireFromString("3002215608380279021243"),
						Proofs: []common.Hash{
							common.HexToHash("0x970920cd59c71016482aab19e2336d544b8dbfe2bfed01af58f2c9e4ca04b200"),
							common.HexToHash("0x5ed47c8ad5f44354769602aaecc2d44cd6d76024155ae4bad4ca279bf0fb6da2"),
							common.HexToHash("0x28c968f6e05241a80dcac92c94c61f439a4ab2bde173288927cb3616c8a19b36"),
						},
					},
				},
			},
		},
		{
			rewardMap: RewardMap{
				common.HexToAddress("0x28c5afd0eee456bf62a8d1187f633ca639b4d9ed"): decimal.RequireFromString("16488.311649245234056167"),
				common.HexToAddress("0x2f2fcc239a6c5c28f62a4ebe4ffdd180e00663b9"): decimal.RequireFromString("16729.101643405008986624"),
				common.HexToAddress("0x6bc8b3a2773a92c0941eda83611f3063a090c77b"): decimal.RequireFromString("13570.679277767782370838"),
				common.HexToAddress("0x74b746c8d2652cb198a9b0d87bd2315f639caed0"): decimal.RequireFromString("14709.682979240536703059"),
				common.HexToAddress("0x851f9de75c021c90c40e209ffba0488b0320c46f"): decimal.RequireFromString("28171.252009881750777824"),
				common.HexToAddress("0xaa293a146aaf9e05bedd1ff29b0da5bd8be70955"): decimal.RequireFromString("18045.877974535505493763"),
				common.HexToAddress("0xb066c21e2d4ace41bc82a1fdb3d53e385d6fc24b"): decimal.RequireFromString("13955.468885464348203264"),
				common.HexToAddress("0xe02de734d92293ee5725b5e82404aab6d19ad5e9"): decimal.RequireFromString("3329.446834599978209057"),
				common.HexToAddress("0xf849cfcf783f094e66a18da660938013f903817a"): decimal.RequireFromString("0.178745859855199404"),
			},
			expected: RewardMerkleProofs{
				MerkleRoot: common.HexToHash("0x69c794981f1902c092e1e54b1efb365d7c166204007695a4405ce92a0888d646"),
				Rewards: map[common.Address]RewardMerkleProofsReward{
					common.HexToAddress("0x28c5afd0eee456bf62a8d1187f633ca639b4d9ed"): {
						Amount: decimal.RequireFromString("16488311649245234056167"),
						Proofs: []common.Hash{
							common.HexToHash("0xb276ab158ec8b3b13ff0fb58a21db74480b0d78354fa041d9a3dcbd130079a5f"),
							common.HexToHash("0x5546ac1e098ec9c839f3466135aec01f6323ae2065115de34a9ef9a6731f3188"),
							common.HexToHash("0x1288f19bfa35381cf7bcf6a6b5efa1057907d978dc70daacbbd6a08f3feb5fc8"),
							common.HexToHash("0x4acba57647b935d471333df39ad0b0e5468b36a172f425dc4bd261f2003e7640"),
						},
					},
					common.HexToAddress("0x2f2fcc239a6c5c28f62a4ebe4ffdd180e00663b9"): {
						Amount: decimal.RequireFromString("16729101643405008986624"),
						Proofs: []common.Hash{
							common.HexToHash("0xcd7366ed1e4de5393f72bb95c10d05f9813f3089ead3d9865e8bf17d4940e008"),
							common.HexToHash("0x5546ac1e098ec9c839f3466135aec01f6323ae2065115de34a9ef9a6731f3188"),
							common.HexToHash("0x1288f19bfa35381cf7bcf6a6b5efa1057907d978dc70daacbbd6a08f3feb5fc8"),
							common.HexToHash("0x4acba57647b935d471333df39ad0b0e5468b36a172f425dc4bd261f2003e7640"),
						},
					},
					common.HexToAddress("0x6bc8b3a2773a92c0941eda83611f3063a090c77b"): {
						Amount: decimal.RequireFromString("13570679277767782370838"),
						Proofs: []common.Hash{
							common.HexToHash("0x42310dfe3b2fcc22bb99871d4843294cf29418e7e87ebdc5ba90ecaf19c3ff6a"),
							common.HexToHash("0x4a19f00688937e1d780aca58a363946c40cf3cbe55b362d51db69fd0ac74b1b6"),
							common.HexToHash("0x1288f19bfa35381cf7bcf6a6b5efa1057907d978dc70daacbbd6a08f3feb5fc8"),
							common.HexToHash("0x4acba57647b935d471333df39ad0b0e5468b36a172f425dc4bd261f2003e7640"),
						},
					},
					common.HexToAddress("0x74b746c8d2652cb198a9b0d87bd2315f639caed0"): {
						Amount: decimal.RequireFromString("14709682979240536703059"),
						Proofs: []common.Hash{
							common.HexToHash("0x1dffe2ab070d8acc33e03a27771d1f3adf4e03b393278d3d985e54ce92ff021a"),
							common.HexToHash("0x4a19f00688937e1d780aca58a363946c40cf3cbe55b362d51db69fd0ac74b1b6"),
							common.HexToHash("0x1288f19bfa35381cf7bcf6a6b5efa1057907d978dc70daacbbd6a08f3feb5fc8"),
							common.HexToHash("0x4acba57647b935d471333df39ad0b0e5468b36a172f425dc4bd261f2003e7640"),
						},
					},
					common.HexToAddress("0x851f9de75c021c90c40e209ffba0488b0320c46f"): {
						Amount: decimal.RequireFromString("28171252009881750777824"),
						Proofs: []common.Hash{
							common.HexToHash("0xea04ec31d39379361f90acd0baf8cf78e2c20b6a6fa9dafb1813a99c16142ad3"),
							common.HexToHash("0x2dac9a3fa80d5c03189cb599e7ceb3740796209ecc12e68dd2dc1368151ef6f5"),
							common.HexToHash("0x8fb4b2d6e8cd4d00bfd10c625109f3a19be59b9c5c2cbe68e184e3e3e705c6bd"),
							common.HexToHash("0x4acba57647b935d471333df39ad0b0e5468b36a172f425dc4bd261f2003e7640"),
						},
					},
					common.HexToAddress("0xaa293a146aaf9e05bedd1ff29b0da5bd8be70955"): {
						Amount: decimal.RequireFromString("18045877974535505493763"),
						Proofs: []common.Hash{
							common.HexToHash("0x397f3857d1e0527d617b05b8f5dfe826c081b934d6690d81afc1ed0e299139e4"),
							common.HexToHash("0x2dac9a3fa80d5c03189cb599e7ceb3740796209ecc12e68dd2dc1368151ef6f5"),
							common.HexToHash("0x8fb4b2d6e8cd4d00bfd10c625109f3a19be59b9c5c2cbe68e184e3e3e705c6bd"),
							common.HexToHash("0x4acba57647b935d471333df39ad0b0e5468b36a172f425dc4bd261f2003e7640"),
						},
					},
					common.HexToAddress("0xb066c21e2d4ace41bc82a1fdb3d53e385d6fc24b"): {
						Amount: decimal.RequireFromString("13955468885464348203264"),
						Proofs: []common.Hash{
							common.HexToHash("0x45a9e8f43bbab53bf1ba90293a97ef60b087f57ffc3e72a89d839befff6e753b"),
							common.HexToHash("0xdcd74bb14d6a42dac4a41e32dc4d9a18553d8398ed42087ca0a832f890702db7"),
							common.HexToHash("0x8fb4b2d6e8cd4d00bfd10c625109f3a19be59b9c5c2cbe68e184e3e3e705c6bd"),
							common.HexToHash("0x4acba57647b935d471333df39ad0b0e5468b36a172f425dc4bd261f2003e7640"),
						},
					},
					common.HexToAddress("0xe02de734d92293ee5725b5e82404aab6d19ad5e9"): {
						Amount: decimal.RequireFromString("3329446834599978209057"),
						Proofs: []common.Hash{
							common.HexToHash("0xd35bfca6200a00232832dd89d1ba342c6ea1288524c706d85667ba09a758074e"),
							common.HexToHash("0xdcd74bb14d6a42dac4a41e32dc4d9a18553d8398ed42087ca0a832f890702db7"),
							common.HexToHash("0x8fb4b2d6e8cd4d00bfd10c625109f3a19be59b9c5c2cbe68e184e3e3e705c6bd"),
							common.HexToHash("0x4acba57647b935d471333df39ad0b0e5468b36a172f425dc4bd261f2003e7640"),
						},
					},
					common.HexToAddress("0xf849cfcf783f094e66a18da660938013f903817a"): {
						Amount: decimal.RequireFromString("178745859855199404"),
						Proofs: []common.Hash{
							common.HexToHash("0xe72177780d76c4605d3b2fcc29854f2a4bbd60a1fbf6d3621885a49d7d84db51"),
							common.HexToHash("0xcee9c32c7644e978874619d370db2fe2c938c36657fa28d7cab7d7ec73b5286d"),
							common.HexToHash("0xc4d0a884aeb763e97b5de092b29c4248dd4d70030bf600d0f26168beda630135"),
							common.HexToHash("0xf984426d201924f5bfea2c1cd11428328ce851a510bb455fe3db0845f411529a"),
						},
					},
				},
			},
		},
		{
			rewardMap: RewardMap{
				common.HexToAddress("0x061517CbA46D7c9230C2caa560d1D33680AEE63e"): decimal.RequireFromString("196.07843137254902"),
				common.HexToAddress("0x24570C86f029E562Bdb05236bC416e8465a49A02"): decimal.RequireFromString("196.07843137254902"),
				common.HexToAddress("0x2ca5bd4EcBa9d159Bb4Ad035832cb60a0024D722"): decimal.RequireFromString("196.07843137254902"),
				common.HexToAddress("0xf13B5d905F60C0DB27cCB2d0D7048a26a8DA8868"): decimal.RequireFromString("196.07843137254902"),
				common.HexToAddress("0x52Ad61Fee56e1Dc0BF2C2Da48eDB23d4066DBB3a"): decimal.RequireFromString("196.07843137254902"),
				common.HexToAddress("0x71fa77d9aff47DfDfDe0509af06f004f9cf08138"): decimal.RequireFromString("196.07843137254902"),
				common.HexToAddress("0x99b76164BeEadF1d5D98163C4919BF50d6f10480"): decimal.RequireFromString("196.07843137254902"),
				common.HexToAddress("0xA915763E53903937394F35a56eA891970BEbfbB4"): decimal.RequireFromString("196.07843137254902"),
				common.HexToAddress("0xC7B8d5785c3d049A7Aaa1f959a2167da6a03Ceb5"): decimal.RequireFromString("196.07843137254902"),
				common.HexToAddress("0xD366e67f4fe62F1Cdd0c339aDe5919a70a46906f"): decimal.RequireFromString("196.07843137254902"),
				common.HexToAddress("0xE168e55619D7e43223B2e9cA825e4846Ff79817e"): decimal.RequireFromString("196.07843137254902"),
				common.HexToAddress("0xF3e49D7e1Eb938b925ab19F7493948b76524aD81"): decimal.RequireFromString("196.07843137254902"),
			},
			expected: RewardMerkleProofs{
				MerkleRoot: common.HexToHash("0xb97fa19bc5ec25a33f12a83d672f4e4d80efc8d4aafc2fdb1e1c9cb0a447ea01"),
				Rewards: map[common.Address]RewardMerkleProofsReward{
					common.HexToAddress("0x061517CbA46D7c9230C2caa560d1D33680AEE63e"): {
						Amount: decimal.RequireFromString("196078431372549020000"),
						Proofs: []common.Hash{
							common.HexToHash("0x3336ec04f2a0953763449c86a0968b6717f1c7e539da6f80e1e7bca524cae8f0"),
							common.HexToHash("0x933b2ea9a2d3567f58c1826f0bf7c0a601da8b7b1923017c41eca65b4c37432a"),
							common.HexToHash("0x2332226b9f3af052c4d89adcbe10b11933e94748a839e5adad7be04a9b80bdaa"),
							common.HexToHash("0xcbc13f149c2c24a3029ad00fe02057b71e5fe0418133a394ae839bbf28e675eb"),
						},
					},
					common.HexToAddress("0x24570C86f029E562Bdb05236bC416e8465a49A02"): {
						Amount: decimal.RequireFromString("196078431372549020000"),
						Proofs: []common.Hash{
							common.HexToHash("0x3104d2dd9254b8a62ff5a2e5b2c71d47d9b68b2faf0de05869d959a68d3c6a8b"),
							common.HexToHash("0x933b2ea9a2d3567f58c1826f0bf7c0a601da8b7b1923017c41eca65b4c37432a"),
							common.HexToHash("0x2332226b9f3af052c4d89adcbe10b11933e94748a839e5adad7be04a9b80bdaa"),
							common.HexToHash("0xcbc13f149c2c24a3029ad00fe02057b71e5fe0418133a394ae839bbf28e675eb"),
						},
					},
					common.HexToAddress("0x2ca5bd4EcBa9d159Bb4Ad035832cb60a0024D722"): {
						Amount: decimal.RequireFromString("196078431372549020000"),
						Proofs: []common.Hash{
							common.HexToHash("0xb7f0cb28d946da0f8fdff8a9a4e0604811fe4430a5a36223f3707c5d364ba2c8"),
							common.HexToHash("0x55d402a1476e00a0f48031cbd6c3a56113b299c911a09ee1111fef9bd523fbd7"),
							common.HexToHash("0x2332226b9f3af052c4d89adcbe10b11933e94748a839e5adad7be04a9b80bdaa"),
							common.HexToHash("0xcbc13f149c2c24a3029ad00fe02057b71e5fe0418133a394ae839bbf28e675eb"),
						},
					},
					common.HexToAddress("0x52Ad61Fee56e1Dc0BF2C2Da48eDB23d4066DBB3a"): {
						Amount: decimal.RequireFromString("196078431372549020000"),
						Proofs: []common.Hash{
							common.HexToHash("0x1417bd34b98e0c1f61364fc555cba1386f4125cbd4f1f8fee4b36f3552bb6812"),
							common.HexToHash("0x55d402a1476e00a0f48031cbd6c3a56113b299c911a09ee1111fef9bd523fbd7"),
							common.HexToHash("0x2332226b9f3af052c4d89adcbe10b11933e94748a839e5adad7be04a9b80bdaa"),
							common.HexToHash("0xcbc13f149c2c24a3029ad00fe02057b71e5fe0418133a394ae839bbf28e675eb"),
						},
					},
					common.HexToAddress("0x71fa77d9aff47DfDfDe0509af06f004f9cf08138"): {
						Amount: decimal.RequireFromString("196078431372549020000"),
						Proofs: []common.Hash{
							common.HexToHash("0x38597ed8467f00528c55a02d3f1d4242a1ffb74b5b15b7b7c5233f06433172a5"),
							common.HexToHash("0x08a6ec014965a741e17da1ad25cf5744aa5132d08f15260aa88a5b512eeb6d0f"),
							common.HexToHash("0xb34dca08548133afa02d9ba2c1a980dbbfa326f940a9faefb65138225d0744ce"),
							common.HexToHash("0xcbc13f149c2c24a3029ad00fe02057b71e5fe0418133a394ae839bbf28e675eb"),
						},
					},
					common.HexToAddress("0x99b76164BeEadF1d5D98163C4919BF50d6f10480"): {
						Amount: decimal.RequireFromString("196078431372549020000"),
						Proofs: []common.Hash{
							common.HexToHash("0x435ff5480ae277a4e61d2fd57297df547ac588c31a4b350bc91c63ae38af0c85"),
							common.HexToHash("0x08a6ec014965a741e17da1ad25cf5744aa5132d08f15260aa88a5b512eeb6d0f"),
							common.HexToHash("0xb34dca08548133afa02d9ba2c1a980dbbfa326f940a9faefb65138225d0744ce"),
							common.HexToHash("0xcbc13f149c2c24a3029ad00fe02057b71e5fe0418133a394ae839bbf28e675eb"),
						},
					},
					common.HexToAddress("0xA915763E53903937394F35a56eA891970BEbfbB4"): {
						Amount: decimal.RequireFromString("196078431372549020000"),
						Proofs: []common.Hash{
							common.HexToHash("0x2f524dfffce3fa8b493fac4c3f764fb9d10e46f6daaa909ef0c5d2eb331abff5"),
							common.HexToHash("0x91e14796e8e377de74e6f1a08363b57c58c21dbf9e00eeae1219b2eeda15e067"),
							common.HexToHash("0xb34dca08548133afa02d9ba2c1a980dbbfa326f940a9faefb65138225d0744ce"),
							common.HexToHash("0xcbc13f149c2c24a3029ad00fe02057b71e5fe0418133a394ae839bbf28e675eb"),
						},
					},
					common.HexToAddress("0xC7B8d5785c3d049A7Aaa1f959a2167da6a03Ceb5"): {
						Amount: decimal.RequireFromString("196078431372549020000"),
						Proofs: []common.Hash{
							common.HexToHash("0xc32b56277227669d8f8502139bee024e68629f9297ecb93c2c3f2bd9a1db6d5f"),
							common.HexToHash("0x91e14796e8e377de74e6f1a08363b57c58c21dbf9e00eeae1219b2eeda15e067"),
							common.HexToHash("0xb34dca08548133afa02d9ba2c1a980dbbfa326f940a9faefb65138225d0744ce"),
							common.HexToHash("0xcbc13f149c2c24a3029ad00fe02057b71e5fe0418133a394ae839bbf28e675eb"),
						},
					},
					common.HexToAddress("0xD366e67f4fe62F1Cdd0c339aDe5919a70a46906f"): {
						Amount: decimal.RequireFromString("196078431372549020000"),
						Proofs: []common.Hash{
							common.HexToHash("0xbf099bc4aa018d0a64598a2053c0562bd26a2490cdbc1376b587e58b5924cc63"),
							common.HexToHash("0xfea554996f67c560f2c01ed3deacc89b35ea11c97853f827e8c1129a08064a61"),
							common.HexToHash("0xa0cd9cafd84d040084bcba5c7463340731d673237ff3c506f330736ef615fa41"),
							common.HexToHash("0x05d9acf22a5e10d34f15b41cd124a3bee90ce1f41b176f3ac23dec2aa23f99fd"),
						},
					},
					common.HexToAddress("0xE168e55619D7e43223B2e9cA825e4846Ff79817e"): {
						Amount: decimal.RequireFromString("196078431372549020000"),
						Proofs: []common.Hash{
							common.HexToHash("0x951799da3ffc6c08edd89ac10a6c72cc51d6163b3ea89d1d588ce70092eeeff0"),
							common.HexToHash("0xfea554996f67c560f2c01ed3deacc89b35ea11c97853f827e8c1129a08064a61"),
							common.HexToHash("0xa0cd9cafd84d040084bcba5c7463340731d673237ff3c506f330736ef615fa41"),
							common.HexToHash("0x05d9acf22a5e10d34f15b41cd124a3bee90ce1f41b176f3ac23dec2aa23f99fd"),
						},
					},
					common.HexToAddress("0xF3e49D7e1Eb938b925ab19F7493948b76524aD81"): {
						Amount: decimal.RequireFromString("196078431372549020000"),
						Proofs: []common.Hash{
							common.HexToHash("0x20021444dfc219577d0b496663ec70067f5d28f9f07ab3aff650fb99a9d02bcd"),
							common.HexToHash("0xaffb4b62dbab179a3147edf616d6f223148daa04a813cf33680336b0fdc18035"),
							common.HexToHash("0xa0cd9cafd84d040084bcba5c7463340731d673237ff3c506f330736ef615fa41"),
							common.HexToHash("0x05d9acf22a5e10d34f15b41cd124a3bee90ce1f41b176f3ac23dec2aa23f99fd"),
						},
					},
					common.HexToAddress("0xf13B5d905F60C0DB27cCB2d0D7048a26a8DA8868"): {
						Amount: decimal.RequireFromString("196078431372549020000"),
						Proofs: []common.Hash{
							common.HexToHash("0x2c9626aeac39f1bb23506a37e413f8169eae5e4103768acca68977e28ed9d05c"),
							common.HexToHash("0xaffb4b62dbab179a3147edf616d6f223148daa04a813cf33680336b0fdc18035"),
							common.HexToHash("0xa0cd9cafd84d040084bcba5c7463340731d673237ff3c506f330736ef615fa41"),
							common.HexToHash("0x05d9acf22a5e10d34f15b41cd124a3bee90ce1f41b176f3ac23dec2aa23f99fd"),
						},
					},
				},
			},
		},
		{
			rewardMap: RewardMap{
				common.HexToAddress("0x52Ad61Fee56e1Dc0BF2C2Da48eDB23d4066DBB3a"): decimal.RequireFromString("20833.333333333333333333"),
				common.HexToAddress("0x02b46fB0b491D517E5854fc329522e4F1Fb08B13"): decimal.RequireFromString("20833.333333333333333333"),
				common.HexToAddress("0xB066C21E2D4ace41BC82a1FdB3d53E385D6fc24B"): decimal.RequireFromString("20833.333333333333333333"),
				common.HexToAddress("0xAA293A146aAf9E05BeDD1Ff29B0da5bD8BE70955"): decimal.RequireFromString("20833.333333333333333333"),
				common.HexToAddress("0x2F2fcc239A6C5C28F62a4Ebe4fFdD180e00663b9"): decimal.RequireFromString("20833.333333333333333333"),
				common.HexToAddress("0xf849cFcf783F094E66A18da660938013F903817a"): decimal.RequireFromString("20833.333333333333333333"),
				common.HexToAddress("0x6Bc8B3a2773A92c0941eDa83611F3063a090c77b"): decimal.RequireFromString("20833.333333333333333333"),
				common.HexToAddress("0xe02de734D92293Ee5725b5E82404AaB6d19Ad5e9"): decimal.RequireFromString("20833.333333333333333333"),
				common.HexToAddress("0x74b746C8d2652CB198a9b0D87bd2315f639caeD0"): decimal.RequireFromString("20833.333333333333333333"),
				common.HexToAddress("0x851f9de75c021c90c40E209FfBa0488B0320c46f"): decimal.RequireFromString("20833.333333333333333333"),
				common.HexToAddress("0xFf3a09C21b54912Bb7A727e6060b051d7bcC602d"): decimal.RequireFromString("20833.333333333333333333"),
				common.HexToAddress("0x28C5aFD0EEE456BF62a8d1187F633ca639B4d9eD"): decimal.RequireFromString("20833.333333333333333333"),
			},
			expected: RewardMerkleProofs{
				MerkleRoot: common.HexToHash("0x9dd98168b6236ba34da16d9ca92ece5fe71d1b93bcc750fec8b2871234aa9a14"),
				Rewards: map[common.Address]RewardMerkleProofsReward{
					common.HexToAddress("0x02b46fb0b491d517e5854fc329522e4f1fb08b13"): {
						Amount: decimal.RequireFromString("20833333333333333333333"),
						Proofs: []common.Hash{
							common.HexToHash("0xd7b494fdc9b0de36e5a1a38a95b939f609fe947f10c73796897b46b4337fd12c"),
							common.HexToHash("0xe06d515b1ff9d974fa8e0daaf79eb02e290467cd1c192ec8a933d28e35a93e5d"),
							common.HexToHash("0x4a7cd74416cbd2c32bd2e32dd685e48523f6875da23bd679f0b6fda16b73b1f4"),
							common.HexToHash("0xf8425a8641b5395cba8b64962653a433990b9aa095ff2b8cd76173acf3912f44"),
						},
					},
					common.HexToAddress("0x28c5afd0eee456bf62a8d1187f633ca639b4d9ed"): {
						Amount: decimal.RequireFromString("20833333333333333333333"),
						Proofs: []common.Hash{
							common.HexToHash("0xa1d4b50cac252edfbe76d982caf4c78106a5873279a8b4f599d732044b3292a5"),
							common.HexToHash("0xe06d515b1ff9d974fa8e0daaf79eb02e290467cd1c192ec8a933d28e35a93e5d"),
							common.HexToHash("0x4a7cd74416cbd2c32bd2e32dd685e48523f6875da23bd679f0b6fda16b73b1f4"),
							common.HexToHash("0xf8425a8641b5395cba8b64962653a433990b9aa095ff2b8cd76173acf3912f44"),
						},
					},
					common.HexToAddress("0x2f2fcc239a6c5c28f62a4ebe4ffdd180e00663b9"): {
						Amount: decimal.RequireFromString("20833333333333333333333"),
						Proofs: []common.Hash{
							common.HexToHash("0x81798092f85eb7abbfa23628a8a795fde5d4f40861fb8cd3e0b9ec98cedb03f6"),
							common.HexToHash("0x8b2314d4cad700fccd1f436ae5792a94abfd5b2c4c49efb04dab6fbc9a07323a"),
							common.HexToHash("0x4a7cd74416cbd2c32bd2e32dd685e48523f6875da23bd679f0b6fda16b73b1f4"),
							common.HexToHash("0xf8425a8641b5395cba8b64962653a433990b9aa095ff2b8cd76173acf3912f44"),
						},
					},
					common.HexToAddress("0x52ad61fee56e1dc0bf2c2da48edb23d4066dbb3a"): {
						Amount: decimal.RequireFromString("20833333333333333333333"),
						Proofs: []common.Hash{
							common.HexToHash("0xb2266d25505b67230eec62e881e81042542742a39f8e2228fae264c0d3d973d9"),
							common.HexToHash("0x8b2314d4cad700fccd1f436ae5792a94abfd5b2c4c49efb04dab6fbc9a07323a"),
							common.HexToHash("0x4a7cd74416cbd2c32bd2e32dd685e48523f6875da23bd679f0b6fda16b73b1f4"),
							common.HexToHash("0xf8425a8641b5395cba8b64962653a433990b9aa095ff2b8cd76173acf3912f44"),
						},
					},
					common.HexToAddress("0x6bc8b3a2773a92c0941eda83611f3063a090c77b"): {
						Amount: decimal.RequireFromString("20833333333333333333333"),
						Proofs: []common.Hash{
							common.HexToHash("0x89ad16a181cdcfe8c655b6dfd9ca6f62d564286ff67385777204c17d86c5dd14"),
							common.HexToHash("0x24f0cf89dd85e0bd87ac524156151e506a4b4ba8a90adf30d6cf4aba417a3c67"),
							common.HexToHash("0x67ed959c4d59466052eced78f049b84afd2ea3ef5fcb1b3ebbe15808786367d1"),
							common.HexToHash("0xf8425a8641b5395cba8b64962653a433990b9aa095ff2b8cd76173acf3912f44"),
						},
					},
					common.HexToAddress("0x74b746c8d2652cb198a9b0d87bd2315f639caed0"): {
						Amount: decimal.RequireFromString("20833333333333333333333"),
						Proofs: []common.Hash{
							common.HexToHash("0x003d99a1fd73ca8e9f0b6b1e0355bfd93d079ca22b06ee911d477a82c2d96519"),
							common.HexToHash("0x24f0cf89dd85e0bd87ac524156151e506a4b4ba8a90adf30d6cf4aba417a3c67"),
							common.HexToHash("0x67ed959c4d59466052eced78f049b84afd2ea3ef5fcb1b3ebbe15808786367d1"),
							common.HexToHash("0xf8425a8641b5395cba8b64962653a433990b9aa095ff2b8cd76173acf3912f44"),
						},
					},
					common.HexToAddress("0x851f9de75c021c90c40e209ffba0488b0320c46f"): {
						Amount: decimal.RequireFromString("20833333333333333333333"),
						Proofs: []common.Hash{
							common.HexToHash("0xfbab5f9941bc6957c12636bd6f73608011e5bdfccbfd65a558512e96f7d7f9cd"),
							common.HexToHash("0xd40f28637ad771e8416def83cfd63f43ae1e9e0ec576b561529792bb8e0068fe"),
							common.HexToHash("0x67ed959c4d59466052eced78f049b84afd2ea3ef5fcb1b3ebbe15808786367d1"),
							common.HexToHash("0xf8425a8641b5395cba8b64962653a433990b9aa095ff2b8cd76173acf3912f44"),
						},
					},
					common.HexToAddress("0xaa293a146aaf9e05bedd1ff29b0da5bd8be70955"): {
						Amount: decimal.RequireFromString("20833333333333333333333"),
						Proofs: []common.Hash{
							common.HexToHash("0xa7635716fbc72841c55b9e0b69b742125f203ab9c493ffaacdb163768a7558b0"),
							common.HexToHash("0xd40f28637ad771e8416def83cfd63f43ae1e9e0ec576b561529792bb8e0068fe"),
							common.HexToHash("0x67ed959c4d59466052eced78f049b84afd2ea3ef5fcb1b3ebbe15808786367d1"),
							common.HexToHash("0xf8425a8641b5395cba8b64962653a433990b9aa095ff2b8cd76173acf3912f44"),
						},
					},
					common.HexToAddress("0xb066c21e2d4ace41bc82a1fdb3d53e385d6fc24b"): {
						Amount: decimal.RequireFromString("20833333333333333333333"),
						Proofs: []common.Hash{
							common.HexToHash("0x1ba3afd210d259b56ed8fb8045897f054c7826e0b59f4d7ea5bb549c8f15eccc"),
							common.HexToHash("0xdba17a02ff3ffae76585caac83466db185535f2e65774c2274f2367419c525fd"),
							common.HexToHash("0x9e4f703fef5b6f50b15aa63abb8af24ab3baeb531ccb09e8ed0ff6fc729cb604"),
							common.HexToHash("0xad8d94ed4896893456302797f39eb4a435946e21a27ba320c6febfb8d003a716"),
						},
					},
					common.HexToAddress("0xe02de734d92293ee5725b5e82404aab6d19ad5e9"): {
						Amount: decimal.RequireFromString("20833333333333333333333"),
						Proofs: []common.Hash{
							common.HexToHash("0x443db1effc6555eedd0e1d6fa35ec1c8b5c13b68326d0bd3b4e02c324593c57f"),
							common.HexToHash("0xdba17a02ff3ffae76585caac83466db185535f2e65774c2274f2367419c525fd"),
							common.HexToHash("0x9e4f703fef5b6f50b15aa63abb8af24ab3baeb531ccb09e8ed0ff6fc729cb604"),
							common.HexToHash("0xad8d94ed4896893456302797f39eb4a435946e21a27ba320c6febfb8d003a716"),
						},
					},
					common.HexToAddress("0xf849cfcf783f094e66a18da660938013f903817a"): {
						Amount: decimal.RequireFromString("20833333333333333333333"),
						Proofs: []common.Hash{
							common.HexToHash("0xc72496e7a2a6602f556464a3e0624253660b3a7c5171fe848727b83ef67383a5"),
							common.HexToHash("0xde927066c9b118f9e372964ecba8e831d813a195961cdeee781bde1f227285f1"),
							common.HexToHash("0x9e4f703fef5b6f50b15aa63abb8af24ab3baeb531ccb09e8ed0ff6fc729cb604"),
							common.HexToHash("0xad8d94ed4896893456302797f39eb4a435946e21a27ba320c6febfb8d003a716"),
						},
					},
					common.HexToAddress("0xff3a09c21b54912bb7a727e6060b051d7bcc602d"): {
						Amount: decimal.RequireFromString("20833333333333333333333"),
						Proofs: []common.Hash{
							common.HexToHash("0x230828c8f1b86ec898ebebf65dee1e3bec7c4ccdcc76eaa5e0d0bdb93cefcb77"),
							common.HexToHash("0xde927066c9b118f9e372964ecba8e831d813a195961cdeee781bde1f227285f1"),
							common.HexToHash("0x9e4f703fef5b6f50b15aa63abb8af24ab3baeb531ccb09e8ed0ff6fc729cb604"),
							common.HexToHash("0xad8d94ed4896893456302797f39eb4a435946e21a27ba320c6febfb8d003a716"),
						},
					},
				},
			},
		},
	}

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("TestCase[%d]", i+1), func(t *testing.T) {
			rewardMerkleTree, err := NewRewardMerkleTree(testCase.rewardMap)
			assert.Nil(t, err)

			err = rewardMerkleTree.GenerateMerkleProofs()
			assert.Nil(t, err)

			assert.Equal(t, testCase.expected.MerkleRoot.String(), rewardMerkleTree.MerkleProofs.MerkleRoot.String())
			for account, reward := range rewardMerkleTree.MerkleProofs.Rewards {
				expectedReward, ok := testCase.expected.Rewards[account]
				assert.True(t, ok)

				assert.Equal(t, expectedReward.Amount.String(), reward.Amount.String())

				for i, proof := range reward.Proofs {
					assert.Equal(t, expectedReward.Proofs[i].String(), proof.String())
				}
			}
		})
	}
}
