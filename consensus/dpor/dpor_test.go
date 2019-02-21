// Copyright 2018 The cpchain authors
// This file is part of the cpchain library.
//
// The cpchain library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The cpchain library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the cpchain library. If not, see <http://www.gnu.org/licenses/>.

package dpor

import (
	"math/big"
	"testing"

	"bitbucket.org/cpchain/chain/configs"
	"bitbucket.org/cpchain/chain/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
)

var (
	addr1 = common.HexToAddress("0xef3dd127de235f15ffb4fc0d71469d1339df6465")
	addr2 = common.HexToAddress("0xc05302acebd0730e3a18a058d7d1cb1204c4a092")
	addr3 = common.HexToAddress("0xe94b7b6c5a0e526a4d97f9768ad6097bde25c62a")
	addr4 = common.HexToAddress("0x3333333333333333333333333333333333333333")

	validator1 = common.HexToAddress("0x7b2f052a372951d02798853e39ee56c895109992")
	validator2 = common.HexToAddress("0x2f0176cc3a8617b6ddea6a501028fa4c6fc25ca1")
	validator3 = common.HexToAddress("0xe4d51117832e84f1d082e9fc12439b771a57e7b2")
	validator4 = common.HexToAddress("0x32bd7c33bb5060a85f361caf20c0bda9075c5d51")
)

func getProposerAddress() []common.Address {
	proposers := make([]common.Address, 3)
	proposers[0] = addr1
	proposers[1] = addr2
	proposers[2] = addr3
	return proposers
}

func getValidatorAddress() []common.Address {
	validators := make([]common.Address, 4)
	validators[0] = validator1
	validators[1] = validator2
	validators[2] = validator3
	validators[3] = validator4
	return validators
}

func getCandidates() []common.Address {
	return getProposerAddress()
}

func recents() map[uint64]common.Address {
	signers := make(map[uint64]common.Address)
	signers[0] = addr1
	signers[1] = addr2
	return signers
}

func Test_percentagePBFT(t *testing.T) {
	type args struct {
		n uint
		N uint
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"3*0>21*2", args{0, 21}, false},
		{"3*14>21*2", args{14, 21}, false},
		{"3*15>21*2", args{15, 21}, true},
		{"3*21>21*2", args{21, 21}, true},
		{"3*1000>21*2", args{1000, 21}, true},
	}
	dporUtil := &defaultDporUtil{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := dporUtil.percentagePBFT(tt.args.n, tt.args.N); got != tt.want {
				t.Errorf("percentagePBFT() = %v, want %v", got, tt.want)
			}
		})
	}
}

func newHeader() *types.Header {
	return &types.Header{
		ParentHash:   common.HexToHash("0x83cafc574e1f51ba9dc0568fc617a08ea2429fb384059c972f13b19fa1c8dd55"),
		Coinbase:     common.HexToAddress("0x8888f1F195AFa192CfeE860698584c030f4c9dB1"),
		StateRoot:    common.HexToHash("0xef1552a40b7165c3cd773806b9e0c165b75356e0314bf0706f279c729f51e017"),
		TxsRoot:      common.HexToHash("0x5fe50b260da6308036625b850b5d6ced6d0a9f814c0688bc91ffb7b7a3a54b67"),
		ReceiptsRoot: common.HexToHash("0xbc37d79753ad738a6dac4921e57392f145d8887476de3f783dfa7edae9283e52"),
		Number:       big.NewInt(1),
		GasLimit:     uint64(3141592),
		GasUsed:      uint64(21000),
		Time:         big.NewInt(1426516743),
		Extra:        []byte("0x0000000000000000000000000000000000000000000000000000000000000000095e7baea6a6c7c4c2dfeb977efac326af552d87e94b7b6c5a0e526a4d97f9768ad6097bde25c62ac05302acebd0730e3a18a058d7d1cb1204c4a092ef3dd127de235f15ffb4fc0d71469d1339df64650000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"),
	}
}

func TestNew(t *testing.T) {
	dpor := New(&configs.DporConfig{}, &fakeDb{}, nil)
	assert.NotNil(t, dpor)
}
