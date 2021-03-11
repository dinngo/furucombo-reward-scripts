# Furucombo COMBO Reward Script

This script calculates COMBO reward distribution.

[![Travis](https://travis-ci.com/dinngodev/furucombo-reward-scripts.svg?branch=master)](https://travis-ci.com/dinngodev/furucombo-reward-scripts)

## Announcement
* Transaction mining program started from 3:00 AM (UTC) Jan 15, 2021: [medium](https://medium.com/furucombo/announcing-furucombo-transaction-mining-program-33381f393230)
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

* bonus reward

  ```sh
  $ go run main.go bonus --help
  $ go run main.go bonus -c={CONFIG_PATH}
  ```

## Reward
* Tx mining reward COMBO [medium](https://medium.com/furucombo/announcing-furucombo-transaction-mining-program-33381f393230)
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

#### reward distribution is decided by trading rank, staking amount and staking duration
* `trading_rank` - higher trading volume in a round gives higher trading rank, e.g., user A/B/C's trading volume are 10/100/1000 USD, their trading rank are `1/2/3`.
* `staking_amount` - the amount of token in a staking contract, e.g., user A/B/C stakes `100/50/20` COMBO in the staking contract
* `staking_duration` - the duration of token in a staking contract, e.g., user A/B/C's staking duration is `46500` blocks ~= 1 week.
* Note that this script tracks the staking amount and duration every block, so higher staking amount for longer duration gives higher reward.

#### Let `share = (trading_rank * staking_amount * staking_duration)`
* A/B/C's `share` are `4,650,000/4,650,000/2,790,000`
* A/B/C's `weight` ~= `38%/38%/24%` by normalizing `share` into [0, 1]

#### Therefore, the reward distribution of 93,750 COMBO for a staking pool in a round
* A/B/C's `reward` ~= `35625/35625/22500` COMBO

### What's the difference between COMBO and COMBO/ETH LP pool?
* COMBO Pool: Every traded users no matter staked or not would share the rewards in this pool. If you made trades and have staked 0 COMBO, we will treat you as staked 250* COMBO when calculate the rewards; if you made trades and have staked 1000 COMBO, we will treat you as staked 1250 COMBO when calculate the rewards. Besides, there’s a trading volume cap of $1000 USD*. Which means that users with $1000 or $1000+ weekly trading volume will share the same trading rank.
* COMBO/ETH Uniswap V2 Liquidity Pool: Only people who staked could share the rewards in this pool. Besides, there’s a trading volume cap of $1 million USD*. Which means that users with $1 million or $1 million+ weekly trading volume will share the same trading rank.
* *The numbers are updated after the Feb 22–24 community vote.
