# Furucombo COMBO Reward Script

This script calculates COMBO reward distribution.

[![Travis](https://travis-ci.com/dinngodev/furucombo-reward-scripts.svg?branch=master)](https://travis-ci.com/dinngodev/furucombo-reward-scripts)

## Announcement
* Transaction mining program start from 3:00 AM (UTC) Jan 15, 2021: [medium](https://medium.com/furucombo/announcing-furucombo-transaction-mining-program-33381f393230)
* Retroactive COMBO will be distributed at 3:00 AM (UTC) Jan 15, 2021: [medium](https://medium.com/furucombo/first-furucombo-grant-7b1e48175c99)

## Setup

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

## Reward

* Retroactive COMBO [medium](https://medium.com/furucombo/first-furucombo-grant-7b1e48175c99)
  * Round 0 [reward](/rewards/retroactive/0/rewards.json)

* UNI distribution [medium](https://medium.com/furucombo/uni-decision-has-been-made-distribution-to-community-253a51e742dc)
  * Round 0 [reward](/rewards/uni_distribution/0/rewards.json)

## Calculations

### Staking Reward

#### Staking reward distribution is decided by trading rank, staking amount and staking duration

* `trading_rank` - higher trading volume in a round gives higher trading rank, e.g., user A/B/C's trading volume are 10/100/1000 USD, their trading rank are `1/2/3`.
* `staking_amount` - the amount of token in a staking contract, e.g., user A/B/C stakes `100/50/20` COMBO in the staking contract
* `staking_duration` - the duration of token in a staking contract, e.g., user A/B/C's staking duration is `46500` blocks ~= 1 week.
* Note that this script tracks the staking amount and duration every block, so higher staking amount for longer duration gives higher reward.

#### Let `share = (trading_rank * staking_amount * staking_duration)`

* A/B/C's `share` are `4,650,000/4,650,000/2,790,000`
* A/B/C's `weight` ~= `38%/38%/24%` by normalizing `share` into [0, 1]

#### Therefore, the reward distribution of 93,750 COMBO for a staking pool in a round

* A/B/C's `reward` ~= `35625/35625/22500` COMBO

#### Note that there is a base staking amount for each staking pool
* Pool 0 - COMBO Pool: this script sets `base_staking_amount` = 20 COMBO for each trading user, it means every trading user could receive COMBO reward by assuming everyone staked 20 COMBO in a round.
* Pool 1 - COMBO LP Pool: this script sets `base_staking_amount` = 0 COMBO_LP_TOKEN, it means users need to trade and staking LP token to receive COMBO reward.