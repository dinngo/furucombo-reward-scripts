# Furucombo COMBO Reward Script

This script calculates COMBO reward distribution.

[![Travis](https://travis-ci.com/dinngodev/furucombo-reward-scripts.svg?branch=master)](https://travis-ci.com/dinngodev/furucombo-reward-scripts)
[![Execute stakingv2 reward calculation](https://github.com/dinngodev/furucombo-reward-scripts/actions/workflows/stakingv2.yml/badge.svg)](https://github.com/dinngodev/furucombo-reward-scripts/actions/workflows/stakingv2.yml)

## Setup
* Configure json file under /config

```sh
$ cp .env.sample .env
$ vim .env
ETHEREUM_RPC_URL={YOUR_ETHEREUM_RPC_URL}
ETHERSCAN_API_KEY={YOUR_KEY}
```

* ETHEREUM_RPC_URL is an ethereum archive node url, e.g., <https://mainnet.infura.io/v3/{YOUR_KEY}>
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

* usage reward

  ```sh
  $ go run main.go usage --help
  $ go run main.go usage -c={CONFIG_PATH}
  ```

## Reward
### Mining COMBO
* Season 1 [Medium](https://medium.com/furucombo/announcing-furucombo-transaction-mining-program-33381f393230)
* Season 2 [Medium](https://medium.com/furucombo/announcing-combo-mining-season-2-e0c20e586c47)
* Season 3 [Medium](https://medium.com/furucombo/combo-mining-season-3-5e5f248923b2)

| COMBO pool | COMBO/ETH UNIV2 Pool |
| ---------: | -------------------: |
| [0](/rewards/staking/0/0x7c46eFAe8632A0c0e1C25718bae91b6b62D9A16E/rewards.json)   | [0](/rewards/staking/0/0x78d742F43Ce72B3D7bDBB2147c252F7a8bab3de4/rewards.json)   |
| [1](/rewards/staking/1/0x7c46eFAe8632A0c0e1C25718bae91b6b62D9A16E/rewards.json)   | [1](/rewards/staking/1/0x78d742F43Ce72B3D7bDBB2147c252F7a8bab3de4/rewards.json)   |
| [2](/rewards/staking/2/0x7c46eFAe8632A0c0e1C25718bae91b6b62D9A16E/rewards.json)   | [2](/rewards/staking/2/0x78d742F43Ce72B3D7bDBB2147c252F7a8bab3de4/rewards.json)   |
| [3](/rewards/staking/3/0x7c46eFAe8632A0c0e1C25718bae91b6b62D9A16E/rewards.json)   | [3](/rewards/staking/3/0x78d742F43Ce72B3D7bDBB2147c252F7a8bab3de4/rewards.json)   |
| [4](/rewards/staking/4/0x7c46eFAe8632A0c0e1C25718bae91b6b62D9A16E/rewards.json)   | [4](/rewards/staking/4/0x78d742F43Ce72B3D7bDBB2147c252F7a8bab3de4/rewards.json)   |
| [5](/rewards/staking/5/0x7c46eFAe8632A0c0e1C25718bae91b6b62D9A16E/rewards.json)   | [5](/rewards/staking/5/0x78d742F43Ce72B3D7bDBB2147c252F7a8bab3de4/rewards.json)   |
| [6](/rewards/staking/6/0x7c46eFAe8632A0c0e1C25718bae91b6b62D9A16E/rewards.json)   | [6](/rewards/staking/6/0x78d742F43Ce72B3D7bDBB2147c252F7a8bab3de4/rewards.json)   |
| [7](/rewards/staking/7/0x7c46eFAe8632A0c0e1C25718bae91b6b62D9A16E/rewards.json)   | [7](/rewards/staking/7/0x78d742F43Ce72B3D7bDBB2147c252F7a8bab3de4/rewards.json)   |
| [8](/rewards/staking/8/0x7c46eFAe8632A0c0e1C25718bae91b6b62D9A16E/rewards.json)   | [8](/rewards/staking/8/0x78d742F43Ce72B3D7bDBB2147c252F7a8bab3de4/rewards.json)   |
| [9](/rewards/staking/9/0x7c46eFAe8632A0c0e1C25718bae91b6b62D9A16E/rewards.json)   | [9](/rewards/staking/9/0x78d742F43Ce72B3D7bDBB2147c252F7a8bab3de4/rewards.json)   |
| [10](/rewards/staking/10/0x7c46eFAe8632A0c0e1C25718bae91b6b62D9A16E/rewards.json) | [10](/rewards/staking/10/0x78d742F43Ce72B3D7bDBB2147c252F7a8bab3de4/rewards.json) |
| [11](/rewards/staking/11/0x7c46eFAe8632A0c0e1C25718bae91b6b62D9A16E/rewards.json) | [11](/rewards/staking/11/0x78d742F43Ce72B3D7bDBB2147c252F7a8bab3de4/rewards.json) |
| [12](/rewards/staking/12/0x7c46eFAe8632A0c0e1C25718bae91b6b62D9A16E/rewards.json) | [12](/rewards/staking/12/0x78d742F43Ce72B3D7bDBB2147c252F7a8bab3de4/rewards.json) |
| [13](/rewards/staking/13/0x7c46eFAe8632A0c0e1C25718bae91b6b62D9A16E/rewards.json) | [13](/rewards/staking/13/0x78d742F43Ce72B3D7bDBB2147c252F7a8bab3de4/rewards.json) |
| [14](/rewards/staking/14/0x7c46eFAe8632A0c0e1C25718bae91b6b62D9A16E/rewards.json) | [14](/rewards/staking/14/0x78d742F43Ce72B3D7bDBB2147c252F7a8bab3de4/rewards.json) |
| [15](/rewards/staking/15/0x7c46eFAe8632A0c0e1C25718bae91b6b62D9A16E/rewards.json) | [15](/rewards/staking/15/0x78d742F43Ce72B3D7bDBB2147c252F7a8bab3de4/rewards.json) |

### Bonus COMBO
| COMBO pool | COMBO/ETH UNIV2 Pool |
| ---------: | -------------------: |
| [1inch](/rewards/bonus/0/rewards.json) |  [1inch](/rewards/bonus/1/rewards.json) |
| [Curve](/rewards/bonus/2/rewards.json) |  [Curve](/rewards/bonus/3/rewards.json) |
| [Aave](/rewards/bonus/4/rewards.json) |  [Aave](/rewards/bonus/5/rewards.json) |


### Retroactive COMBO
* [Medium](https://medium.com/furucombo/first-furucombo-grant-7b1e48175c99)

| Retroactive |
| ----------: |
| [0](/rewards/retroactive/0/rewards.json) |
| [1](/rewards/retroactive/1/rewards.json) |
| [2](/rewards/retroactive/2/rewards.json) |
| [3](/rewards/retroactive/3/rewards.json) |
| [4](/rewards/retroactive/4/rewards.json) |
| [5](/rewards/retroactive/5/rewards.json) |
| [6](/rewards/retroactive/6/rewards.json) |
| [7](/rewards/retroactive/7/rewards.json) |
| [8](/rewards/retroactive/8/rewards.json) |

### Uniswap UNI
* [Medium](https://medium.com/furucombo/uni-decision-has-been-made-distribution-to-community-253a51e742dc)

| Uniswap UNI |
| ----------: |
| [0](/rewards/uni_distribution/0/rewards.json) |
