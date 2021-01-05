package ethereum

import (
	"log"
	"os"

	"github.com/ethereum/go-ethereum/ethclient"
)

var client *ethclient.Client

// Client get client
func Client() *ethclient.Client {
	if client == nil {
		rpcURL := os.Getenv("ETHEREUM_RPC_URL")

		if len(rpcURL) == 0 {
			log.Fatalln("env ETHEREUM_RPC_URL can't be blank")
		}

		var err error
		client, err = ethclient.Dial(rpcURL)
		if err != nil {
			log.Fatalf("failed to connect geth node: %s", rpcURL)
		}
	}

	return client
}
