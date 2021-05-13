package ethereum

import (
	"log"
	"os"

	"github.com/ethereum/go-ethereum/ethclient"
)

var client *ethclient.Client
var clientMatic *ethclient.Client

// Client get client
func Client() *ethclient.Client {
	rpcURL := os.Getenv("ETHEREUM_RPC_URL")
	if len(rpcURL) == 0 {
		log.Fatalln("env ETHEREUM_RPC_URL can't be blank")
	}

	return dial(rpcURL, client)
}

func ClientMatic() *ethclient.Client {
	rpcURL := os.Getenv("MATIC_RPC_URL")
	if len(rpcURL) == 0 {
		log.Fatalln("env MATIC_RPC_URL can't be blank")
	}

	return dial(rpcURL, clientMatic)

}

func dial(rpcURL string, client *ethclient.Client) *ethclient.Client {
	if client == nil {
		var err error
		client, err = ethclient.Dial(rpcURL)
		if err != nil {
			log.Fatalf("failed to connect geth node: %s", rpcURL)
		}
	}

	return client
}
