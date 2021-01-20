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

## Reward

* Retroactive COMBO [medium](https://medium.com/furucombo/first-furucombo-grant-7b1e48175c99)
  * Round 0 [reward](/rewards/retroactive/0/rewards.json)
  * Round 1 [reward](/rewards/retroactive/1/rewards.json)
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

### Notices
* There is a base staking amount for each staking pool
  * COMBO Pool: Every traded users no matter staked or not would share the rewards in this pool. If you made trades and have staked 0 COMBO, the script will treat you as staked 20 COMBO when calculate the rewards; if you made trades and have staked 1000 COMBO, the script will treat you as staked 1020 COMBO when calculate the rewards.
  * COMBO/ETH Uniswap V2 Liquidity Pool: Only people who staked could share the rewards in this pool.

* Transactions from smart contracts will not be counted. For example, Argent and DappLogic will not be counted as the behavior of these wallets are quite unique. Quick check? If your transaction has the tag “Interacted With (To): Furucombo: Proxy vx.x.x” on etherscan, it will be counted.