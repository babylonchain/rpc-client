package query

import (
	"github.com/babylonchain/rpc-client/client"
	coretypes "github.com/tendermint/tendermint/rpc/core/types"
	"strings"
)

// GetBlock returns the tendermint block at a specific height
func GetBlock(c *client.Client, height int64) (*coretypes.ResultBlock, error) {
	q := c.GetDefaultQuery()
	ctx, cancel := q.GetQueryContext()
	defer cancel()

	return q.Client.RPCClient.Block(ctx, &height)
}

// TxSearch searches for transactions satisfying the events specified on the events list
func TxSearch(c *client.Client, events []string, prove bool, page *int, perPage *int, orderBy string) (*coretypes.ResultTxSearch, error) {
	q := c.GetDefaultQuery()
	ctx, cancel := q.GetQueryContext()
	defer cancel()

	return q.Client.RPCClient.TxSearch(ctx, strings.Join(events, " AND "), prove, page, perPage, orderBy)
}

// GetTx returns the transaction with the specified hash
func GetTx(c *client.Client, hash []byte) (*coretypes.ResultTx, error) {
	q := c.GetDefaultQuery()
	ctx, cancel := q.GetQueryContext()
	defer cancel()

	return q.Client.RPCClient.Tx(ctx, hash, false)
}
