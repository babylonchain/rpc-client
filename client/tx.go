package client

import (
	"context"

	btcctypes "github.com/babylonchain/babylon/x/btccheckpoint/types"
	btclctypes "github.com/babylonchain/babylon/x/btclightclient/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// We do not expose ctx in our client calls, which means:
// - we do not support cancellation of submitting messages
// - the only timeout is the block inclusion timeout i.e block-timeout
// TODO: To properly support cancellation we need to expose ctx in our client calls
func (c *Client) InsertBTCSpvProof(msg *btcctypes.MsgInsertBTCSpvProof) (*sdk.TxResponse, error) {
	// generate context
	ctx := context.Background()
	return c.ChainClient.SendMsg(ctx, msg, "")
}

func (c *Client) InsertHeaders(msg *btclctypes.MsgInsertHeaders) (*sdk.TxResponse, error) {
	// generate context
	ctx := context.Background()
	return c.ChainClient.SendMsg(ctx, msg, "")
}

// TODO: implement necessary message invocations here
// - MsgInconsistencyEvidence
// - MsgStallingEvidence
