package etherscan

// AccountTxs  ?module=account&action=txlist
func (c *Client) AccountTxs(params Params) ([]NormalTx, error) {
	url := c.NewURL("account", "txlist", params)
	resp, err := c.MakeReq(url)
	if err != nil {
		return nil, err
	}

	txs := make([]NormalTx, 0)
	if err := jsonex.Unmarshal(resp, &txs); err != nil {
		return nil, err
	}

	return txs, nil
}

// AccountInternalTxs  ?module=account&action=txlistinternal
func (c *Client) AccountInternalTxs(params Params) ([]InternalTx, error) {
	url := c.NewURL("account", "txlistinternal", params)
	resp, err := c.MakeReq(url)
	if err != nil {
		return nil, err
	}

	txs := make([]InternalTx, 0)
	if err := jsonex.Unmarshal(resp, &txs); err != nil {
		return nil, err
	}

	return txs, nil
}

// AccountTokenTxs  ?module=account&action=tokentx
func (c *Client) AccountTokenTxs(params Params) ([]ERC20Transfer, error) {
	url := c.NewURL("account", "tokentx", params)
	resp, err := c.MakeReq(url)
	if err != nil {
		return nil, err
	}

	txs := make([]ERC20Transfer, 0)
	if err := jsonex.Unmarshal(resp, &txs); err != nil {
		return nil, err
	}

	return txs, nil
}
