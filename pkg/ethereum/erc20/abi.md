# ABI cmd

- ERC20 Contract

```sh
abigen --abi pkg/ethereum/erc20/abi/ERC20.json --pkg erc20 --out pkg/ethereum/erc20/erc20_contract.go --type ERC20Contract
```

- DSToken Contract

```sh
abigen --abi pkg/ethereum/erc20/abi/DSToken.json --pkg erc20 --out pkg/ethereum/erc20/dstoken_contract.go --type DSTokenContract
```
