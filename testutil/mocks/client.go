// Code generated by MockGen. DO NOT EDIT.
// Source: client/interface.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	types "github.com/babylonchain/babylon/x/btccheckpoint/types"
	types0 "github.com/babylonchain/babylon/x/btclightclient/types"
	types1 "github.com/babylonchain/babylon/x/checkpointing/types"
	types2 "github.com/babylonchain/babylon/x/epoching/types"
	config "github.com/babylonchain/rpc-client/config"
	chainhash "github.com/btcsuite/btcd/chaincfg/chainhash"
	wire "github.com/btcsuite/btcd/wire"
	types3 "github.com/cosmos/cosmos-sdk/types"
	types4 "github.com/cosmos/cosmos-sdk/x/staking/types"
	gomock "github.com/golang/mock/gomock"
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

// BlsPublicKeyList mocks base method.
func (m *MockBabylonClient) BlsPublicKeyList(epochNumber uint64) ([]*types1.ValidatorWithBlsKey, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BlsPublicKeyList", epochNumber)
	ret0, _ := ret[0].([]*types1.ValidatorWithBlsKey)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BlsPublicKeyList indicates an expected call of BlsPublicKeyList.
func (mr *MockBabylonClientMockRecorder) BlsPublicKeyList(epochNumber interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BlsPublicKeyList", reflect.TypeOf((*MockBabylonClient)(nil).BlsPublicKeyList), epochNumber)
}

// GetAddr mocks base method.
func (m *MockBabylonClient) GetAddr() (types3.AccAddress, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAddr")
	ret0, _ := ret[0].(types3.AccAddress)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAddr indicates an expected call of GetAddr.
func (mr *MockBabylonClientMockRecorder) GetAddr() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAddr", reflect.TypeOf((*MockBabylonClient)(nil).GetAddr))
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

// InsertBTCSpvProof mocks base method.
func (m *MockBabylonClient) InsertBTCSpvProof(msg *types.MsgInsertBTCSpvProof) (*types3.TxResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertBTCSpvProof", msg)
	ret0, _ := ret[0].(*types3.TxResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InsertBTCSpvProof indicates an expected call of InsertBTCSpvProof.
func (mr *MockBabylonClientMockRecorder) InsertBTCSpvProof(msg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertBTCSpvProof", reflect.TypeOf((*MockBabylonClient)(nil).InsertBTCSpvProof), msg)
}

// InsertHeader mocks base method.
func (m *MockBabylonClient) InsertHeader(msg *types0.MsgInsertHeader) (*types3.TxResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertHeader", msg)
	ret0, _ := ret[0].(*types3.TxResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InsertHeader indicates an expected call of InsertHeader.
func (mr *MockBabylonClientMockRecorder) InsertHeader(msg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertHeader", reflect.TypeOf((*MockBabylonClient)(nil).InsertHeader), msg)
}

// InsertHeaders mocks base method.
func (m *MockBabylonClient) InsertHeaders(msgs []*types0.MsgInsertHeader) (*types3.TxResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertHeaders", msgs)
	ret0, _ := ret[0].(*types3.TxResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InsertHeaders indicates an expected call of InsertHeaders.
func (mr *MockBabylonClientMockRecorder) InsertHeaders(msgs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertHeaders", reflect.TypeOf((*MockBabylonClient)(nil).InsertHeaders), msgs)
}

// MustGetAddr mocks base method.
func (m *MockBabylonClient) MustGetAddr() types3.AccAddress {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MustGetAddr")
	ret0, _ := ret[0].(types3.AccAddress)
	return ret0
}

// MustGetAddr indicates an expected call of MustGetAddr.
func (mr *MockBabylonClientMockRecorder) MustGetAddr() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MustGetAddr", reflect.TypeOf((*MockBabylonClient)(nil).MustGetAddr))
}

// MustInsertBTCSpvProof mocks base method.
func (m *MockBabylonClient) MustInsertBTCSpvProof(msg *types.MsgInsertBTCSpvProof) *types3.TxResponse {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MustInsertBTCSpvProof", msg)
	ret0, _ := ret[0].(*types3.TxResponse)
	return ret0
}

// MustInsertBTCSpvProof indicates an expected call of MustInsertBTCSpvProof.
func (mr *MockBabylonClientMockRecorder) MustInsertBTCSpvProof(msg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MustInsertBTCSpvProof", reflect.TypeOf((*MockBabylonClient)(nil).MustInsertBTCSpvProof), msg)
}

// MustQueryBTCCheckpointParams mocks base method.
func (m *MockBabylonClient) MustQueryBTCCheckpointParams() *types.Params {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MustQueryBTCCheckpointParams")
	ret0, _ := ret[0].(*types.Params)
	return ret0
}

// MustQueryBTCCheckpointParams indicates an expected call of MustQueryBTCCheckpointParams.
func (mr *MockBabylonClientMockRecorder) MustQueryBTCCheckpointParams() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MustQueryBTCCheckpointParams", reflect.TypeOf((*MockBabylonClient)(nil).MustQueryBTCCheckpointParams))
}

// QueryBTCCheckpointParams mocks base method.
func (m *MockBabylonClient) QueryBTCCheckpointParams() (*types.Params, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryBTCCheckpointParams")
	ret0, _ := ret[0].(*types.Params)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryBTCCheckpointParams indicates an expected call of QueryBTCCheckpointParams.
func (mr *MockBabylonClientMockRecorder) QueryBTCCheckpointParams() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryBTCCheckpointParams", reflect.TypeOf((*MockBabylonClient)(nil).QueryBTCCheckpointParams))
}

// QueryBTCLightclientParams mocks base method.
func (m *MockBabylonClient) QueryBTCLightclientParams() (*types0.Params, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryBTCLightclientParams")
	ret0, _ := ret[0].(*types0.Params)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryBTCLightclientParams indicates an expected call of QueryBTCLightclientParams.
func (mr *MockBabylonClientMockRecorder) QueryBTCLightclientParams() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryBTCLightclientParams", reflect.TypeOf((*MockBabylonClient)(nil).QueryBTCLightclientParams))
}

// QueryBaseHeader mocks base method.
func (m *MockBabylonClient) QueryBaseHeader() (*wire.BlockHeader, uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryBaseHeader")
	ret0, _ := ret[0].(*wire.BlockHeader)
	ret1, _ := ret[1].(uint64)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// QueryBaseHeader indicates an expected call of QueryBaseHeader.
func (mr *MockBabylonClientMockRecorder) QueryBaseHeader() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryBaseHeader", reflect.TypeOf((*MockBabylonClient)(nil).QueryBaseHeader))
}

// QueryContainsBlock mocks base method.
func (m *MockBabylonClient) QueryContainsBlock(blockHash *chainhash.Hash) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryContainsBlock", blockHash)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryContainsBlock indicates an expected call of QueryContainsBlock.
func (mr *MockBabylonClientMockRecorder) QueryContainsBlock(blockHash interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryContainsBlock", reflect.TypeOf((*MockBabylonClient)(nil).QueryContainsBlock), blockHash)
}

// QueryCurrentEpoch mocks base method.
func (m *MockBabylonClient) QueryCurrentEpoch() (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryCurrentEpoch")
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryCurrentEpoch indicates an expected call of QueryCurrentEpoch.
func (mr *MockBabylonClientMockRecorder) QueryCurrentEpoch() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryCurrentEpoch", reflect.TypeOf((*MockBabylonClient)(nil).QueryCurrentEpoch))
}

// QueryEpochingParams mocks base method.
func (m *MockBabylonClient) QueryEpochingParams() (*types2.Params, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryEpochingParams")
	ret0, _ := ret[0].(*types2.Params)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryEpochingParams indicates an expected call of QueryEpochingParams.
func (mr *MockBabylonClientMockRecorder) QueryEpochingParams() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryEpochingParams", reflect.TypeOf((*MockBabylonClient)(nil).QueryEpochingParams))
}

// QueryFinishedEpochBtcHeight mocks base method.
func (m *MockBabylonClient) QueryFinishedEpochBtcHeight(epochNum uint64) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryFinishedEpochBtcHeight", epochNum)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryFinishedEpochBtcHeight indicates an expected call of QueryFinishedEpochBtcHeight.
func (mr *MockBabylonClientMockRecorder) QueryFinishedEpochBtcHeight(epochNum interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryFinishedEpochBtcHeight", reflect.TypeOf((*MockBabylonClient)(nil).QueryFinishedEpochBtcHeight), epochNum)
}

// QueryHeaderChainTip mocks base method.
func (m *MockBabylonClient) QueryHeaderChainTip() (*chainhash.Hash, uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryHeaderChainTip")
	ret0, _ := ret[0].(*chainhash.Hash)
	ret1, _ := ret[1].(uint64)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// QueryHeaderChainTip indicates an expected call of QueryHeaderChainTip.
func (mr *MockBabylonClientMockRecorder) QueryHeaderChainTip() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryHeaderChainTip", reflect.TypeOf((*MockBabylonClient)(nil).QueryHeaderChainTip))
}

// QueryRawCheckpoint mocks base method.
func (m *MockBabylonClient) QueryRawCheckpoint(epochNumber uint64) (*types1.RawCheckpointWithMeta, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryRawCheckpoint", epochNumber)
	ret0, _ := ret[0].(*types1.RawCheckpointWithMeta)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryRawCheckpoint indicates an expected call of QueryRawCheckpoint.
func (mr *MockBabylonClientMockRecorder) QueryRawCheckpoint(epochNumber interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryRawCheckpoint", reflect.TypeOf((*MockBabylonClient)(nil).QueryRawCheckpoint), epochNumber)
}

// QueryRawCheckpointList mocks base method.
func (m *MockBabylonClient) QueryRawCheckpointList(status types1.CheckpointStatus) ([]*types1.RawCheckpointWithMeta, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryRawCheckpointList", status)
	ret0, _ := ret[0].([]*types1.RawCheckpointWithMeta)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryRawCheckpointList indicates an expected call of QueryRawCheckpointList.
func (mr *MockBabylonClientMockRecorder) QueryRawCheckpointList(status interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryRawCheckpointList", reflect.TypeOf((*MockBabylonClient)(nil).QueryRawCheckpointList), status)
}

// QueryReportedCheckpointBtcHeight mocks base method.
func (m *MockBabylonClient) QueryReportedCheckpointBtcHeight(hash []byte) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryReportedCheckpointBtcHeight", hash)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryReportedCheckpointBtcHeight indicates an expected call of QueryReportedCheckpointBtcHeight.
func (mr *MockBabylonClientMockRecorder) QueryReportedCheckpointBtcHeight(hash interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryReportedCheckpointBtcHeight", reflect.TypeOf((*MockBabylonClient)(nil).QueryReportedCheckpointBtcHeight), hash)
}

// QueryStakingParams mocks base method.
func (m *MockBabylonClient) QueryStakingParams() (*types4.Params, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryStakingParams")
	ret0, _ := ret[0].(*types4.Params)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryStakingParams indicates an expected call of QueryStakingParams.
func (mr *MockBabylonClientMockRecorder) QueryStakingParams() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryStakingParams", reflect.TypeOf((*MockBabylonClient)(nil).QueryStakingParams))
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
