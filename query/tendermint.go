package query

import (
	"strings"

	coretypes "github.com/tendermint/tendermint/rpc/core/types"
)

// GetBlock returns the tendermint block at a specific height
func (c *QueryClient) GetBlock(height int64) (*coretypes.ResultBlock, error) {
	ctx, cancel := c.getQueryContext()
	defer cancel()

	return c.RPCClient.Block(ctx, &height)
}

// BlockSearch searches for blocks satisfying the query
func (c *QueryClient) BlockSearch(query string, page *int, perPage *int, orderBy string) (*coretypes.ResultBlockSearch, error) {
	ctx, cancel := c.getQueryContext()
	defer cancel()

	return c.RPCClient.BlockSearch(ctx, query, page, perPage, orderBy)
}

// TxSearch searches for transactions satisfying the events specified on the events list
func (c *QueryClient) TxSearch(events []string, prove bool, page *int, perPage *int, orderBy string) (*coretypes.ResultTxSearch, error) {
	ctx, cancel := c.getQueryContext()
	defer cancel()

	return c.RPCClient.TxSearch(ctx, strings.Join(events, " AND "), prove, page, perPage, orderBy)
}

// GetTx returns the transaction with the specified hash
func (c *QueryClient) GetTx(hash []byte) (*coretypes.ResultTx, error) {
	ctx, cancel := c.getQueryContext()
	defer cancel()

	return c.RPCClient.Tx(ctx, hash, false)
}
