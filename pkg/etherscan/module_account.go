package etherscan

// AccountTxs  ?module=account&action=txlist
func (c *Client) AccountTxs(params Params) ([]NormalTx, error) {
	url := c.NewURL("account", "txlist", params)
	return c.accountTxs(url, params)
}

// AccountTxsPolygon  ?module=account&action=txlist
func (c *Client) AccountTxsPolygon(params Params) ([]NormalTx, error) {
	url := c.NewURLPolygon("account", "txlist", params)
	return c.accountTxs(url, params)
}

func (c *Client) accountTxs(url string, params Params) ([]NormalTx, error) {

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
	return c.accountInternalTxs(url, params)
}

// AccountInternalTxs  ?module=account&action=txlistinternal
func (c *Client) AccountInternalTxsPolygon(params Params) ([]InternalTx, error) {
	url := c.NewURLPolygon("account", "txlistinternal", params)
	return c.accountInternalTxs(url, params)
}

func (c *Client) accountInternalTxs(url string, params Params) ([]InternalTx, error) {
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
	return c.accountTokenTxs(url, params)
}

// AccountTokenTxsPolygon  ?module=account&action=tokentx
func (c *Client) AccountTokenTxsPolygon(params Params) ([]ERC20Transfer, error) {
	url := c.NewURLPolygon("account", "tokentx", params)
	return c.accountTokenTxs(url, params)
}

func (c *Client) accountTokenTxs(url string, params Params) ([]ERC20Transfer, error) {
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
