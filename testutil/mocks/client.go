// Code generated by MockGen. DO NOT EDIT.
// Source: client/interface.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	types "github.com/babylonchain/babylon/x/btccheckpoint/types"
	types0 "github.com/babylonchain/babylon/x/btclightclient/types"
	types1 "github.com/babylonchain/babylon/x/checkpointing/types"
	types2 "github.com/babylonchain/babylon/x/epoching/types"
	types3 "github.com/babylonchain/babylon/x/monitor/types"
	types4 "github.com/babylonchain/babylon/x/zoneconcierge/types"
	config "github.com/babylonchain/rpc-client/config"
	chainhash "github.com/btcsuite/btcd/chaincfg/chainhash"
	types5 "github.com/cosmos/cosmos-sdk/types"
	query "github.com/cosmos/cosmos-sdk/types/query"
	types6 "github.com/cosmos/cosmos-sdk/x/staking/types"
	gomock "github.com/golang/mock/gomock"
	coretypes "github.com/cometbft/cometbft/rpc/core/types"
)

// MockBabylonClient is a mock of BabylonClient interface.
type MockBabylonClient struct {
	ctrl     *gomock.Controller
	recorder *MockBabylonClientMockRecorder
}

// MockBabylonClientMockRecorder is the mock recorder for MockBabylonClient.
type MockBabylonClientMockRecorder struct {
	mock *MockBabylonClient
}

// NewMockBabylonClient creates a new mock instance.
func NewMockBabylonClient(ctrl *gomock.Controller) *MockBabylonClient {
	mock := &MockBabylonClient{ctrl: ctrl}
	mock.recorder = &MockBabylonClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBabylonClient) EXPECT() *MockBabylonClientMockRecorder {
	return m.recorder
}

// BTCBaseHeader mocks base method.
func (m *MockBabylonClient) BTCBaseHeader() (*types0.QueryBaseHeaderResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BTCBaseHeader")
	ret0, _ := ret[0].(*types0.QueryBaseHeaderResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BTCBaseHeader indicates an expected call of BTCBaseHeader.
func (mr *MockBabylonClientMockRecorder) BTCBaseHeader() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BTCBaseHeader", reflect.TypeOf((*MockBabylonClient)(nil).BTCBaseHeader))
}

// BTCCheckpointParams mocks base method.
func (m *MockBabylonClient) BTCCheckpointParams() (*types.QueryParamsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BTCCheckpointParams")
	ret0, _ := ret[0].(*types.QueryParamsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BTCCheckpointParams indicates an expected call of BTCCheckpointParams.
func (mr *MockBabylonClientMockRecorder) BTCCheckpointParams() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BTCCheckpointParams", reflect.TypeOf((*MockBabylonClient)(nil).BTCCheckpointParams))
}

// BTCHeaderChainTip mocks base method.
func (m *MockBabylonClient) BTCHeaderChainTip() (*types0.QueryTipResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BTCHeaderChainTip")
	ret0, _ := ret[0].(*types0.QueryTipResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BTCHeaderChainTip indicates an expected call of BTCHeaderChainTip.
func (mr *MockBabylonClientMockRecorder) BTCHeaderChainTip() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BTCHeaderChainTip", reflect.TypeOf((*MockBabylonClient)(nil).BTCHeaderChainTip))
}

// BTCMainChain mocks base method.
func (m *MockBabylonClient) BTCMainChain(pagination *query.PageRequest) (*types0.QueryMainChainResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BTCMainChain", pagination)
	ret0, _ := ret[0].(*types0.QueryMainChainResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BTCMainChain indicates an expected call of BTCMainChain.
func (mr *MockBabylonClientMockRecorder) BTCMainChain(pagination interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BTCMainChain", reflect.TypeOf((*MockBabylonClient)(nil).BTCMainChain), pagination)
}

// BTCCheckpointInfo mocks base method.
func (m *MockBabylonClient) BTCCheckpointInfo(epochNumber uint64) (*types.QueryBtcCheckpointInfoResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BTCPositionAtEpoch", epochNumber)
	ret0, _ := ret[0].(*types.QueryBtcCheckpointInfoResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BTCCheckpointInfo indicates an expected call of BTCCheckpointInfo.
func (mr *MockBabylonClientMockRecorder) BTCCheckpointInfo(epochNumber interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BTCCheckpointInfo", reflect.TypeOf((*MockBabylonClient)(nil).BTCCheckpointInfo), epochNumber)
}

// BTCCheckpointsInfo mocks base method.
func (m *MockBabylonClient) BTCCheckpointsInfo(pagination *query.PageRequest) (*types.QueryBtcCheckpointsInfoResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BTCCheckpointsInfo", pagination)
	ret0, _ := ret[0].(*types.QueryBtcCheckpointsInfoResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BTCCheckpointsInfo indicates an expected call of BTCCheckpointsInfo.
func (mr *MockBabylonClientMockRecorder) BTCCheckpointsInfo(pagination interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BTCCheckpointsInfo", reflect.TypeOf((*MockBabylonClient)(nil).BTCCheckpointsInfo), pagination)
}

// BlsPublicKeyList mocks base method.
func (m *MockBabylonClient) BlsPublicKeyList(epochNumber uint64, pagination *query.PageRequest) (*types1.QueryBlsPublicKeyListResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BlsPublicKeyList", epochNumber, pagination)
	ret0, _ := ret[0].(*types1.QueryBlsPublicKeyListResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BlsPublicKeyList indicates an expected call of BlsPublicKeyList.
func (mr *MockBabylonClientMockRecorder) BlsPublicKeyList(epochNumber, pagination interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BlsPublicKeyList", reflect.TypeOf((*MockBabylonClient)(nil).BlsPublicKeyList), epochNumber, pagination)
}

// ConnectedChainEpochInfo mocks base method.
func (m *MockBabylonClient) ConnectedChainEpochInfo(chainID string, epochNum uint64) (*types4.QueryEpochChainInfoResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConnectedChainEpochInfo", chainID, epochNum)
	ret0, _ := ret[0].(*types4.QueryEpochChainInfoResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ConnectedChainEpochInfo indicates an expected call of ConnectedChainEpochInfo.
func (mr *MockBabylonClientMockRecorder) ConnectedChainEpochInfo(chainID, epochNum interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConnectedChainEpochInfo", reflect.TypeOf((*MockBabylonClient)(nil).ConnectedChainEpochInfo), chainID, epochNum)
}

// ConnectedChainHeaders mocks base method.
func (m *MockBabylonClient) ConnectedChainHeaders(chainID string, pagination *query.PageRequest) (*types4.QueryListHeadersResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConnectedChainHeaders", chainID, pagination)
	ret0, _ := ret[0].(*types4.QueryListHeadersResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ConnectedChainHeaders indicates an expected call of ConnectedChainHeaders.
func (mr *MockBabylonClientMockRecorder) ConnectedChainHeaders(chainID, pagination interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConnectedChainHeaders", reflect.TypeOf((*MockBabylonClient)(nil).ConnectedChainHeaders), chainID, pagination)
}

// ConnectedChainInfo mocks base method.
func (m *MockBabylonClient) ConnectedChainInfo(chainID string) (*types4.QueryChainInfoResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConnectedChainInfo", chainID)
	ret0, _ := ret[0].(*types4.QueryChainInfoResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ConnectedChainInfo indicates an expected call of ConnectedChainInfo.
func (mr *MockBabylonClientMockRecorder) ConnectedChainInfo(chainID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConnectedChainInfo", reflect.TypeOf((*MockBabylonClient)(nil).ConnectedChainInfo), chainID)
}

// ConnectedChainList mocks base method.
func (m *MockBabylonClient) ConnectedChainList() (*types4.QueryChainListResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConnectedChainList")
	ret0, _ := ret[0].(*types4.QueryChainListResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ConnectedChainList indicates an expected call of ConnectedChainList.
func (mr *MockBabylonClientMockRecorder) ConnectedChainList() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConnectedChainList", reflect.TypeOf((*MockBabylonClient)(nil).ConnectedChainList))
}

// ContainsBTCBlock mocks base method.
func (m *MockBabylonClient) ContainsBTCBlock(blockHash *chainhash.Hash) (*types0.QueryContainsBytesResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ContainsBTCBlock", blockHash)
	ret0, _ := ret[0].(*types0.QueryContainsBytesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ContainsBTCBlock indicates an expected call of ContainsBTCBlock.
func (mr *MockBabylonClientMockRecorder) ContainsBTCBlock(blockHash interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ContainsBTCBlock", reflect.TypeOf((*MockBabylonClient)(nil).ContainsBTCBlock), blockHash)
}

// CurrentEpoch mocks base method.
func (m *MockBabylonClient) CurrentEpoch() (*types2.QueryCurrentEpochResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CurrentEpoch")
	ret0, _ := ret[0].(*types2.QueryCurrentEpochResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CurrentEpoch indicates an expected call of CurrentEpoch.
func (mr *MockBabylonClientMockRecorder) CurrentEpoch() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CurrentEpoch", reflect.TypeOf((*MockBabylonClient)(nil).CurrentEpoch))
}

// DelegationLifecycle mocks base method.
func (m *MockBabylonClient) DelegationLifecycle(delegator string) (*types2.QueryDelegationLifecycleResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DelegationLifecycle", delegator)
	ret0, _ := ret[0].(*types2.QueryDelegationLifecycleResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DelegationLifecycle indicates an expected call of DelegationLifecycle.
func (mr *MockBabylonClientMockRecorder) DelegationLifecycle(delegator interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DelegationLifecycle", reflect.TypeOf((*MockBabylonClient)(nil).DelegationLifecycle), delegator)
}

// EndedEpochBTCHeight mocks base method.
func (m *MockBabylonClient) EndedEpochBTCHeight(epochNum uint64) (*types3.QueryEndedEpochBtcHeightResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EndedEpochBTCHeight", epochNum)
	ret0, _ := ret[0].(*types3.QueryEndedEpochBtcHeightResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EndedEpochBTCHeight indicates an expected call of EndedEpochBTCHeight.
func (mr *MockBabylonClientMockRecorder) EndedEpochBTCHeight(epochNum interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EndedEpochBTCHeight", reflect.TypeOf((*MockBabylonClient)(nil).EndedEpochBTCHeight), epochNum)
}

// EpochStatusCount mocks base method.
func (m *MockBabylonClient) EpochStatusCount(epochCount uint64) (*types1.QueryRecentEpochStatusCountResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EpochStatusCount", epochCount)
	ret0, _ := ret[0].(*types1.QueryRecentEpochStatusCountResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EpochStatusCount indicates an expected call of EpochStatusCount.
func (mr *MockBabylonClientMockRecorder) EpochStatusCount(epochCount interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EpochStatusCount", reflect.TypeOf((*MockBabylonClient)(nil).EpochStatusCount), epochCount)
}

// EpochingParams mocks base method.
func (m *MockBabylonClient) EpochingParams() (*types2.QueryParamsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EpochingParams")
	ret0, _ := ret[0].(*types2.QueryParamsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EpochingParams indicates an expected call of EpochingParams.
func (mr *MockBabylonClientMockRecorder) EpochingParams() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EpochingParams", reflect.TypeOf((*MockBabylonClient)(nil).EpochingParams))
}

// EpochsInfo mocks base method.
func (m *MockBabylonClient) EpochsInfo(pagination *query.PageRequest) (*types2.QueryEpochsInfoResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EpochsInfo", pagination)
	ret0, _ := ret[0].(*types2.QueryEpochsInfoResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EpochsInfo indicates an expected call of EpochsInfo.
func (mr *MockBabylonClientMockRecorder) EpochsInfo(pagination interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EpochsInfo", reflect.TypeOf((*MockBabylonClient)(nil).EpochsInfo), pagination)
}

// EpochsInfoForEpochRange mocks base method.
func (m *MockBabylonClient) EpochsInfoForEpochRange(startEpoch, endEpoch uint64) (*types2.QueryEpochsInfoResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EpochsInfoForEpochRange", startEpoch, endEpoch)
	ret0, _ := ret[0].(*types2.QueryEpochsInfoResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EpochsInfoForEpochRange indicates an expected call of EpochsInfoForEpochRange.
func (mr *MockBabylonClientMockRecorder) EpochsInfoForEpochRange(startEpoch, endEpoch interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EpochsInfoForEpochRange", reflect.TypeOf((*MockBabylonClient)(nil).EpochsInfoForEpochRange), startEpoch, endEpoch)
}

// FinalizedConnectedChainInfo mocks base method.
func (m *MockBabylonClient) FinalizedConnectedChainInfo(chainID string) (*types4.QueryFinalizedChainInfoResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FinalizedConnectedChainInfo", chainID)
	ret0, _ := ret[0].(*types4.QueryFinalizedChainInfoResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FinalizedConnectedChainInfo indicates an expected call of FinalizedConnectedChainInfo.
func (mr *MockBabylonClientMockRecorder) FinalizedConnectedChainInfo(chainID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FinalizedConnectedChainInfo", reflect.TypeOf((*MockBabylonClient)(nil).FinalizedConnectedChainInfo), chainID)
}

// GetAddr mocks base method.
func (m *MockBabylonClient) GetAddr() (types5.AccAddress, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAddr")
	ret0, _ := ret[0].(types5.AccAddress)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAddr indicates an expected call of GetAddr.
func (mr *MockBabylonClientMockRecorder) GetAddr() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAddr", reflect.TypeOf((*MockBabylonClient)(nil).GetAddr))
}

// GetBlock mocks base method.
func (m *MockBabylonClient) GetBlock(height int64) (*coretypes.ResultBlock, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBlock", height)
	ret0, _ := ret[0].(*coretypes.ResultBlock)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBlock indicates an expected call of GetBlock.
func (mr *MockBabylonClientMockRecorder) GetBlock(height interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBlock", reflect.TypeOf((*MockBabylonClient)(nil).GetBlock), height)
}

// GetConfig mocks base method.
func (m *MockBabylonClient) GetConfig() *config.BabylonConfig {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetConfig")
	ret0, _ := ret[0].(*config.BabylonConfig)
	return ret0
}

// GetConfig indicates an expected call of GetConfig.
func (mr *MockBabylonClientMockRecorder) GetConfig() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetConfig", reflect.TypeOf((*MockBabylonClient)(nil).GetConfig))
}

// GetTagIdx mocks base method.
func (m *MockBabylonClient) GetTagIdx() uint8 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTagIdx")
	ret0, _ := ret[0].(uint8)
	return ret0
}

// GetTagIdx indicates an expected call of GetTagIdx.
func (mr *MockBabylonClientMockRecorder) GetTagIdx() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTagIdx", reflect.TypeOf((*MockBabylonClient)(nil).GetTagIdx))
}

// GetTx mocks base method.
func (m *MockBabylonClient) GetTx(hash []byte) (*coretypes.ResultTx, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTx", hash)
	ret0, _ := ret[0].(*coretypes.ResultTx)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTx indicates an expected call of GetTx.
func (mr *MockBabylonClientMockRecorder) GetTx(hash interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTx", reflect.TypeOf((*MockBabylonClient)(nil).GetTx), hash)
}

// InsertBTCSpvProof mocks base method.
func (m *MockBabylonClient) InsertBTCSpvProof(msg *types.MsgInsertBTCSpvProof) (*types5.TxResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertBTCSpvProof", msg)
	ret0, _ := ret[0].(*types5.TxResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InsertBTCSpvProof indicates an expected call of InsertBTCSpvProof.
func (mr *MockBabylonClientMockRecorder) InsertBTCSpvProof(msg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertBTCSpvProof", reflect.TypeOf((*MockBabylonClient)(nil).InsertBTCSpvProof), msg)
}

// InsertHeader mocks base method.
func (m *MockBabylonClient) InsertHeader(msg *types0.MsgInsertHeader) (*types5.TxResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertHeader", msg)
	ret0, _ := ret[0].(*types5.TxResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InsertHeader indicates an expected call of InsertHeader.
func (mr *MockBabylonClientMockRecorder) InsertHeader(msg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertHeader", reflect.TypeOf((*MockBabylonClient)(nil).InsertHeader), msg)
}

// InsertHeaders mocks base method.
func (m *MockBabylonClient) InsertHeaders(msgs []*types0.MsgInsertHeader) (*types5.TxResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertHeaders", msgs)
	ret0, _ := ret[0].(*types5.TxResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InsertHeaders indicates an expected call of InsertHeaders.
func (mr *MockBabylonClientMockRecorder) InsertHeaders(msgs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertHeaders", reflect.TypeOf((*MockBabylonClient)(nil).InsertHeaders), msgs)
}

// LatestEpochFromStatus mocks base method.
func (m *MockBabylonClient) LatestEpochFromStatus(status types1.CheckpointStatus) (*types1.QueryLastCheckpointWithStatusResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LatestEpochFromStatus", status)
	ret0, _ := ret[0].(*types1.QueryLastCheckpointWithStatusResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LatestEpochFromStatus indicates an expected call of LatestEpochFromStatus.
func (mr *MockBabylonClientMockRecorder) LatestEpochFromStatus(status interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LatestEpochFromStatus", reflect.TypeOf((*MockBabylonClient)(nil).LatestEpochFromStatus), status)
}

// LatestEpochMsgs mocks base method.
func (m *MockBabylonClient) LatestEpochMsgs(endEpoch, epochCount uint64, pagination *query.PageRequest) (*types2.QueryLatestEpochMsgsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LatestEpochMsgs", endEpoch, epochCount, pagination)
	ret0, _ := ret[0].(*types2.QueryLatestEpochMsgsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LatestEpochMsgs indicates an expected call of LatestEpochMsgs.
func (mr *MockBabylonClientMockRecorder) LatestEpochMsgs(endEpoch, epochCount, pagination interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LatestEpochMsgs", reflect.TypeOf((*MockBabylonClient)(nil).LatestEpochMsgs), endEpoch, epochCount, pagination)
}

// MustGetAddr mocks base method.
func (m *MockBabylonClient) MustGetAddr() types5.AccAddress {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MustGetAddr")
	ret0, _ := ret[0].(types5.AccAddress)
	return ret0
}

// MustGetAddr indicates an expected call of MustGetAddr.
func (mr *MockBabylonClientMockRecorder) MustGetAddr() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MustGetAddr", reflect.TypeOf((*MockBabylonClient)(nil).MustGetAddr))
}

// QueryBTCCheckpoint mocks base method.
func (m *MockBabylonClient) QueryBTCCheckpoint(f func(context.Context, types.QueryClient) error) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryBTCCheckpoint", f)
	ret0, _ := ret[0].(error)
	return ret0
}

// QueryBTCCheckpoint indicates an expected call of QueryBTCCheckpoint.
func (mr *MockBabylonClientMockRecorder) QueryBTCCheckpoint(f interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryBTCCheckpoint", reflect.TypeOf((*MockBabylonClient)(nil).QueryBTCCheckpoint), f)
}

// QueryBTCLightclient mocks base method.
func (m *MockBabylonClient) QueryBTCLightclient(f func(context.Context, types0.QueryClient) error) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryBTCLightclient", f)
	ret0, _ := ret[0].(error)
	return ret0
}

// QueryBTCLightclient indicates an expected call of QueryBTCLightclient.
func (mr *MockBabylonClientMockRecorder) QueryBTCLightclient(f interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryBTCLightclient", reflect.TypeOf((*MockBabylonClient)(nil).QueryBTCLightclient), f)
}

// QueryCheckpointing mocks base method.
func (m *MockBabylonClient) QueryCheckpointing(f func(context.Context, types1.QueryClient) error) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryCheckpointing", f)
	ret0, _ := ret[0].(error)
	return ret0
}

// QueryCheckpointing indicates an expected call of QueryCheckpointing.
func (mr *MockBabylonClientMockRecorder) QueryCheckpointing(f interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryCheckpointing", reflect.TypeOf((*MockBabylonClient)(nil).QueryCheckpointing), f)
}

// QueryEpoching mocks base method.
func (m *MockBabylonClient) QueryEpoching(f func(context.Context, types2.QueryClient) error) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryEpoching", f)
	ret0, _ := ret[0].(error)
	return ret0
}

// QueryEpoching indicates an expected call of QueryEpoching.
func (mr *MockBabylonClientMockRecorder) QueryEpoching(f interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryEpoching", reflect.TypeOf((*MockBabylonClient)(nil).QueryEpoching), f)
}

// QueryMonitor mocks base method.
func (m *MockBabylonClient) QueryMonitor(f func(context.Context, types3.QueryClient) error) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryMonitor", f)
	ret0, _ := ret[0].(error)
	return ret0
}

// QueryMonitor indicates an expected call of QueryMonitor.
func (mr *MockBabylonClientMockRecorder) QueryMonitor(f interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryMonitor", reflect.TypeOf((*MockBabylonClient)(nil).QueryMonitor), f)
}

// QueryStaking mocks base method.
func (m *MockBabylonClient) QueryStaking(f func(context.Context, types6.QueryClient) error) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryStaking", f)
	ret0, _ := ret[0].(error)
	return ret0
}

// QueryStaking indicates an expected call of QueryStaking.
func (mr *MockBabylonClientMockRecorder) QueryStaking(f interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryStaking", reflect.TypeOf((*MockBabylonClient)(nil).QueryStaking), f)
}

// QueryZoneConcierge mocks base method.
func (m *MockBabylonClient) QueryZoneConcierge(f func(context.Context, types4.QueryClient) error) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryZoneConcierge", f)
	ret0, _ := ret[0].(error)
	return ret0
}

// QueryZoneConcierge indicates an expected call of QueryZoneConcierge.
func (mr *MockBabylonClientMockRecorder) QueryZoneConcierge(f interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryZoneConcierge", reflect.TypeOf((*MockBabylonClient)(nil).QueryZoneConcierge), f)
}

// RawCheckpoint mocks base method.
func (m *MockBabylonClient) RawCheckpoint(epochNumber uint64) (*types1.QueryRawCheckpointResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RawCheckpoint", epochNumber)
	ret0, _ := ret[0].(*types1.QueryRawCheckpointResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RawCheckpoint indicates an expected call of RawCheckpoint.
func (mr *MockBabylonClientMockRecorder) RawCheckpoint(epochNumber interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RawCheckpoint", reflect.TypeOf((*MockBabylonClient)(nil).RawCheckpoint), epochNumber)
}

// RawCheckpointList mocks base method.
func (m *MockBabylonClient) RawCheckpointList(status types1.CheckpointStatus, pagination *query.PageRequest) (*types1.QueryRawCheckpointListResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RawCheckpointList", status, pagination)
	ret0, _ := ret[0].(*types1.QueryRawCheckpointListResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RawCheckpointList indicates an expected call of RawCheckpointList.
func (mr *MockBabylonClientMockRecorder) RawCheckpointList(status, pagination interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RawCheckpointList", reflect.TypeOf((*MockBabylonClient)(nil).RawCheckpointList), status, pagination)
}

// ReportedCheckpointBTCHeight mocks base method.
func (m *MockBabylonClient) ReportedCheckpointBTCHeight(hashStr string) (*types3.QueryReportedCheckpointBtcHeightResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReportedCheckpointBTCHeight", hashStr)
	ret0, _ := ret[0].(*types3.QueryReportedCheckpointBtcHeightResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReportedCheckpointBTCHeight indicates an expected call of ReportedCheckpointBTCHeight.
func (mr *MockBabylonClientMockRecorder) ReportedCheckpointBTCHeight(hashStr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReportedCheckpointBTCHeight", reflect.TypeOf((*MockBabylonClient)(nil).ReportedCheckpointBTCHeight), hashStr)
}

// StakingParams mocks base method.
func (m *MockBabylonClient) StakingParams() (*types6.QueryParamsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StakingParams")
	ret0, _ := ret[0].(*types6.QueryParamsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StakingParams indicates an expected call of StakingParams.
func (mr *MockBabylonClientMockRecorder) StakingParams() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StakingParams", reflect.TypeOf((*MockBabylonClient)(nil).StakingParams))
}

// Stop mocks base method.
func (m *MockBabylonClient) Stop() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Stop")
}

// Stop indicates an expected call of Stop.
func (mr *MockBabylonClientMockRecorder) Stop() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockBabylonClient)(nil).Stop))
}

// TxSearch mocks base method.
func (m *MockBabylonClient) TxSearch(events []string, prove bool, page, perPage *int, orderBy string) (*coretypes.ResultTxSearch, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TxSearch", events, prove, page, perPage, orderBy)
	ret0, _ := ret[0].(*coretypes.ResultTxSearch)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TxSearch indicates an expected call of TxSearch.
func (mr *MockBabylonClientMockRecorder) TxSearch(events, prove, page, perPage, orderBy interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TxSearch", reflect.TypeOf((*MockBabylonClient)(nil).TxSearch), events, prove, page, perPage, orderBy)
}
