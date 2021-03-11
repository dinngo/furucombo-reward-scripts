# Furucombo COMBO Reward Script

This script calculates COMBO reward distribution.

[![Travis](https://travis-ci.com/dinngodev/furucombo-reward-scripts.svg?branch=master)](https://travis-ci.com/dinngodev/furucombo-reward-scripts)

## Announcement
* Season 2 mining program started from 3:00 AM (UTC) Mar 11, 2021: [medium](https://medium.com/furucombo/announcing-combo-mining-season-2-e0c20e586c47)
* Retroactive COMBO was be distributed at 3:00 AM (UTC) Jan 15, 2021: [medium](https://medium.com/furucombo/first-furucombo-grant-7b1e48175c99)

## Setup
* Configure json file under /config

```sh
$ cp .env.sample .env
$ vim .env
ETHEREUM_RPC_URL={YOUR_ETHEREUM_RPC_URL}
ETHERSCAN_API_KEY={YOUR_KEY}
```

* ETHEREUM_RPC_URL is an ethereum node url, e.g., <https://mainnet.infura.io/v3/{YOUR_KEY}>
* ETHERSCAN_API_KEY is an etherscan api key from <https://etherscan.io/myapikey>

## Run

* one-off reward, e.g., retroactive

  ```sh
  $ go run main.go once --help
  $ go run main.go once -c={CONFIG_PATH}
  ```

* staking reward

  ```sh
  $ go run main.go staking --help
  $ go run main.go staking -c={CONFIG_PATH}
  ```

* staking reward for season 2

  ```sh
  $ go run main.go stakingv2 --help
  $ go run main.go stakingv2 -c={CONFIG_PATH}
  ```

* bonus reward

  ```sh
  $ go run main.go bonus --help
  $ go run main.go bonus -c={CONFIG_PATH}
  ```

## Reward
* Tx mining reward COMBO [season 1 medium](https://medium.com/furucombo/announcing-furucombo-transaction-mining-program-33381f393230) [season 2 medium](https://medium.com/furucombo/announcing-combo-mining-season-2-e0c20e586c47)
  * Round 0
    * COMBO Pool [Reward](/rewards/staking/0/0x7c46eFAe8632A0c0e1C25718bae91b6b62D9A16E/rewards.json)
    * COMBO/ETH UNIV2 Pool [Reward](/rewards/staking/0/0x78d742F43Ce72B3D7bDBB2147c252F7a8bab3de4/rewards.json)
  * Round 1
    * COMBO Pool [Reward](/rewards/staking/1/0x7c46eFAe8632A0c0e1C25718bae91b6b62D9A16E/rewards.json)
    * COMBO/ETH UNIV2 Pool [Reward](/rewards/staking/1/0x78d742F43Ce72B3D7bDBB2147c252F7a8bab3de4/rewards.json)
  * Round 2
    * COMBO Pool [Reward](/rewards/staking/2/0x7c46eFAe8632A0c0e1C25718bae91b6b62D9A16E/rewards.json)
    * COMBO/ETH UNIV2 Pool [Reward](/rewards/staking/2/0x78d742F43Ce72B3D7bDBB2147c252F7a8bab3de4/rewards.json)
  * Round 3
    * COMBO Pool [Reward](/rewards/staking/3/0x7c46eFAe8632A0c0e1C25718bae91b6b62D9A16E/rewards.json)
    * COMBO/ETH UNIV2 Pool [Reward](/rewards/staking/3/0x78d742F43Ce72B3D7bDBB2147c252F7a8bab3de4/rewards.json)
  * Round 4
    * COMBO Pool [Reward](/rewards/staking/4/0x7c46eFAe8632A0c0e1C25718bae91b6b62D9A16E/rewards.json)
    * COMBO/ETH UNIV2 Pool [Reward](/rewards/staking/4/0x78d742F43Ce72B3D7bDBB2147c252F7a8bab3de4/rewards.json)
  * Round 5
    * COMBO Pool [Reward](/rewards/staking/5/0x7c46eFAe8632A0c0e1C25718bae91b6b62D9A16E/rewards.json)
    * COMBO/ETH UNIV2 Pool [Reward](/rewards/staking/5/0x78d742F43Ce72B3D7bDBB2147c252F7a8bab3de4/rewards.json)
  * Round 6
    * COMBO Pool [Reward](/rewards/staking/6/0x7c46eFAe8632A0c0e1C25718bae91b6b62D9A16E/rewards.json)
    * COMBO/ETH UNIV2 Pool [Reward](/rewards/staking/6/0x78d742F43Ce72B3D7bDBB2147c252F7a8bab3de4/rewards.json)
  * Round 7
    * COMBO Pool [Reward](/rewards/staking/7/0x7c46eFAe8632A0c0e1C25718bae91b6b62D9A16E/rewards.json)
    * COMBO/ETH UNIV2 Pool [Reward](/rewards/staking/7/0x78d742F43Ce72B3D7bDBB2147c252F7a8bab3de4/rewards.json)
* Retroactive COMBO [medium](https://medium.com/furucombo/first-furucombo-grant-7b1e48175c99)
  * Round 0 [Reward](/rewards/retroactive/0/rewards.json)
  * Round 1 [Reward](/rewards/retroactive/1/rewards.json)
  * Round 2 [Reward](/rewards/retroactive/2/rewards.json)
  * Round 3 [Reward](/rewards/retroactive/3/rewards.json)
  * Round 4 [Reward](/rewards/retroactive/4/rewards.json)
  * Round 5 [Reward](/rewards/retroactive/5/rewards.json)
  * Round 6 [Reward](/rewards/retroactive/6/rewards.json)
  * Round 7 [Reward](/rewards/retroactive/7/rewards.json)
  * Round 8 [Reward](/rewards/retroactive/8/rewards.json)
* Bonus reward COMBO
  * 1inch
    * COMBO Pool [Reward](/rewards/bonus/0/rewards.json)
    * COMBO/ETH UNIV2 Pool [Reward](/rewards/bonus/1/rewards.json)
  * Curve
    * COMBO Pool [Reward](/rewards/bonus/2/rewards.json)
    * COMBO/ETH UNIV2 Pool [Reward](/rewards/bonus/3/rewards.json)
  * Aave
    * COMBO Pool [Reward](/rewards/bonus/3/rewards.json)
    * COMBO/ETH UNIV2 Pool [Reward](/rewards/bonus/4/rewards.json)
* UNI distribution [medium](https://medium.com/furucombo/uni-decision-has-been-made-distribution-to-community-253a51e742dc)
  * Round 0 [Reward](/rewards/uni_distribution/0/rewards.json)

## Calculations

### Reward

#### Reward is decided by staking amount and staking duration, and user must do at least one tx on furucombo.
* `staking_amount` - the amount of token in a staking contract, e.g., user A/B/C stakes `50/30/20` COMBO in the staking contract
* `staking_duration` - the duration of token in a staking contract, e.g., user A/B/C's staking duration is `46500` blocks ~= 1 week.
* Eligibility requirement: send out at least one transaction on Furucombo during each week’s program.
* Note that this script tracks the staking amount and duration every block, so higher staking amount for longer duration gives higher reward.

#### Let `staking_area = staking_amount * staking_duration`
* A/B/C's `staking_area` are `2,325,000/1,395,000/930,000`
* A/B/C's `weight` = `50%/30%/20%` by normalizing `staking_area` into [0, 1]

#### Therefore, the reward distribution of 46,875 COMBO for a staking pool
* A/B/C's `reward` = `23437.5/14062.5/9,375` COMBO

### What's the difference between COMBO and COMBO/ETH LP pool?
* COMBO Pool
  * Rewards: 46,875 COMBO
  * Competition rule: base amount = 250 COMBO
  * Every traded user, no matter if they stake or not, will share the rewards in this pool. If you have made trades but staked 0 COMBO, we will automatically act as if staked 250 COMBO when calculating the rewards; if you have made trades and have staked 1000 COMBO, we will act as if you have staked 1250 COMBO when calculating the rewards.
* COMBO/ETH Uniswap V2 Liquidity Pool
  * Rewards: 140,625 COMBO
  * Only people who stake LP tokens will be able to share the rewards in this pool.
