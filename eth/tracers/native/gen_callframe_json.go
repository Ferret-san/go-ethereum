// Code generated by github.com/fjl/gencodec. DO NOT EDIT.

package native

import (
	"encoding/json"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/vm"
)

var _ = (*callFrameMarshaling)(nil)

// MarshalJSON marshals as JSON.
func (c callFrame) MarshalJSON() ([]byte, error) {
	type callFrame0 struct {
		BeforeEVMTransfers *[]arbitrumTransfer `json:"beforeEVMTransfers,omitempty"`
		AfterEVMTransfers  *[]arbitrumTransfer `json:"afterEVMTransfers,omitempty"`
		BalanceChanges     *[]balanceChange    `json:"balanceChanges,omitempty"`
		Type               vm.OpCode           `json:"-"`
		From               common.Address      `json:"from"`
		Gas                hexutil.Uint64      `json:"gas"`
		GasUsed            hexutil.Uint64      `json:"gasUsed"`
		To                 *common.Address     `json:"to,omitempty" rlp:"optional"`
		Input              hexutil.Bytes       `json:"input" rlp:"optional"`
		Output             hexutil.Bytes       `json:"output,omitempty" rlp:"optional"`
		Error              string              `json:"error,omitempty" rlp:"optional"`
		RevertReason       string              `json:"revertReason,omitempty"`
		Calls              []callFrame         `json:"calls,omitempty" rlp:"optional"`
		Logs               []callLog           `json:"logs,omitempty" rlp:"optional"`
		Value              *hexutil.Big        `json:"value,omitempty" rlp:"optional"`
		TypeString         string              `json:"type"`
	}
	var enc callFrame0
	enc.BeforeEVMTransfers = c.BeforeEVMTransfers
	enc.AfterEVMTransfers = c.AfterEVMTransfers
	enc.BalanceChanges = c.BalanceChanges
	enc.Type = c.Type
	enc.From = c.From
	enc.Gas = hexutil.Uint64(c.Gas)
	enc.GasUsed = hexutil.Uint64(c.GasUsed)
	enc.To = c.To
	enc.Input = c.Input
	enc.Output = c.Output
	enc.Error = c.Error
	enc.RevertReason = c.RevertReason
	enc.Calls = c.Calls
	enc.Logs = c.Logs
	enc.Value = (*hexutil.Big)(c.Value)
	enc.TypeString = c.TypeString()
	return json.Marshal(&enc)
}

// UnmarshalJSON unmarshals from JSON.
func (c *callFrame) UnmarshalJSON(input []byte) error {
	type callFrame0 struct {
		BeforeEVMTransfers *[]arbitrumTransfer `json:"beforeEVMTransfers,omitempty"`
		AfterEVMTransfers  *[]arbitrumTransfer `json:"afterEVMTransfers,omitempty"`
		BalanceChanges     *[]balanceChange    `json:"balanceChanges,omitempty"`
		Type               *vm.OpCode          `json:"-"`
		From               *common.Address     `json:"from"`
		Gas                *hexutil.Uint64     `json:"gas"`
		GasUsed            *hexutil.Uint64     `json:"gasUsed"`
		To                 *common.Address     `json:"to,omitempty" rlp:"optional"`
		Input              *hexutil.Bytes      `json:"input" rlp:"optional"`
		Output             *hexutil.Bytes      `json:"output,omitempty" rlp:"optional"`
		Error              *string             `json:"error,omitempty" rlp:"optional"`
		RevertReason       *string             `json:"revertReason,omitempty"`
		Calls              []callFrame         `json:"calls,omitempty" rlp:"optional"`
		Logs               []callLog           `json:"logs,omitempty" rlp:"optional"`
		Value              *hexutil.Big        `json:"value,omitempty" rlp:"optional"`
	}
	var dec callFrame0
	if err := json.Unmarshal(input, &dec); err != nil {
		return err
	}
	if dec.BeforeEVMTransfers != nil {
		c.BeforeEVMTransfers = dec.BeforeEVMTransfers
	}
	if dec.AfterEVMTransfers != nil {
		c.AfterEVMTransfers = dec.AfterEVMTransfers
	}
	if dec.BalanceChanges != nil {
		c.BalanceChanges = dec.BalanceChanges
	}
	if dec.Type != nil {
		c.Type = *dec.Type
	}
	if dec.From != nil {
		c.From = *dec.From
	}
	if dec.Gas != nil {
		c.Gas = uint64(*dec.Gas)
	}
	if dec.GasUsed != nil {
		c.GasUsed = uint64(*dec.GasUsed)
	}
	if dec.To != nil {
		c.To = dec.To
	}
	if dec.Input != nil {
		c.Input = *dec.Input
	}
	if dec.Output != nil {
		c.Output = *dec.Output
	}
	if dec.Error != nil {
		c.Error = *dec.Error
	}
	if dec.RevertReason != nil {
		c.RevertReason = *dec.RevertReason
	}
	if dec.Calls != nil {
		c.Calls = dec.Calls
	}
	if dec.Logs != nil {
		c.Logs = dec.Logs
	}
	if dec.Value != nil {
		c.Value = (*big.Int)(dec.Value)
	}
	return nil
}
