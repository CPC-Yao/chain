// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package reward

import (
	"math/big"
	"strings"

	cpchain "bitbucket.org/cpchain/chain"
	"bitbucket.org/cpchain/chain/accounts/abi"
	"bitbucket.org/cpchain/chain/accounts/abi/bind"
	"bitbucket.org/cpchain/chain/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/event"
)

// RewardABI is the input ABI used to generate the binding from.
const RewardABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_period\",\"type\":\"uint256\"}],\"name\":\"setPeriod\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"isContract\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"bonusPool\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"isEnode\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"refundAll\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"getFreeBalanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"nextRound\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"nextRoundStartTime\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"setNextRoundStartTime\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_addr\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"refundDeposit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"criteria\",\"type\":\"uint256\"}],\"name\":\"setCriteria\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getTotalAmount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"investors\",\"outputs\":[{\"name\":\"freeDeposit\",\"type\":\"uint256\"},{\"name\":\"lockedDeposit\",\"type\":\"uint256\"},{\"name\":\"returned\",\"type\":\"uint256\"},{\"name\":\"toRenew\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"setBonusPool\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"disableContract\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"wantRenew\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"getLockedBalanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"newRaise\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"quitRenew\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getTotalLockedAmount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getInvestors\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"startNewRound\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"basicCriteria\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"locked\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"submitDeposit\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"period\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"isToRenew\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"getTotalBalanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"OwnerSetBonusPool\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"who\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"SubmitDeposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"who\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"WithdrawDeposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"who\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"JoinEnodes\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"who\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"RefundDeposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"round\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"lock\",\"type\":\"bool\"},{\"indexed\":false,\"name\":\"_bonusPool\",\"type\":\"uint256\"}],\"name\":\"NewRaise\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"who\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"DepositInsufficient\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_addr\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_iscontinue\",\"type\":\"bool\"}],\"name\":\"ContinuedInvest\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"FundBonusPool\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"num\",\"type\":\"uint256\"}],\"name\":\"RefundAll\",\"type\":\"event\"}]"

// RewardBin is the compiled bytecode used for deploying new contracts.
const RewardBin = `0x60806040526000805460a060020a60ff0219167401000000000000000000000000000000000000000017815569043c33c1937564800000600155600281905560038190556006556276a70060075534801561005957600080fd5b5060008054600160a060020a031916331790556113948061007b6000396000f3006080604052600436106101745763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416630f3a9f6581146101bf57806316279055146101d95780632693ee801461020e5780632a0e96c2146102355780632e1a7d4d1461025657806338e771ab1461026e578063414d11e61461028357806347e40553146102a45780634aeb31eb146102b95780634d5b03f3146102ce5780635215d054146102e65780635f2c01ea1461030a57806365ac4341146103225780636f7bc9be146103375780637c0fa59514610380578063894ba833146103885780638ca2ca7f1461039d5780638d0dcb5f146103b25780638f83bc7f146103d35780638ffb6376146103e857806391d2b32e146103fd578063b2f5a54c14610412578063bd85948c14610477578063c0209d1f1461048c578063cf309012146104a1578063e9843a3e146104b6578063ef78d4fd146104be578063f509dd4a146104d3578063fdfefa7f146104f4575b600254610187903463ffffffff61051516565b6002556040805134815290517f71030773066b852afef8d0f98dbfdaec8e9a62f2f5533916ec7dfa15a0edc1f29181900360200190a1005b3480156101cb57600080fd5b506101d760043561052f565b005b3480156101e557600080fd5b506101fa600160a060020a036004351661054b565b604080519115158252519081900360200190f35b34801561021a57600080fd5b50610223610553565b60408051918252519081900360200190f35b34801561024157600080fd5b506101fa600160a060020a0360043516610559565b34801561026257600080fd5b506101d7600435610572565b34801561027a57600080fd5b506101d761062b565b34801561028f57600080fd5b50610223600160a060020a03600435166106fc565b3480156102b057600080fd5b50610223610717565b3480156102c557600080fd5b5061022361071d565b3480156102da57600080fd5b506101d7600435610723565b3480156102f257600080fd5b506101d7600160a060020a036004351660243561073f565b34801561031657600080fd5b506101d7600435610834565b34801561032e57600080fd5b50610223610850565b34801561034357600080fd5b50610358600160a060020a03600435166108ac565b6040805194855260208501939093528383019190915215156060830152519081900360800190f35b6101d76108d6565b34801561039457600080fd5b506101d7610939565b3480156103a957600080fd5b506101d76109fa565b3480156103be57600080fd5b50610223600160a060020a0360043516610a30565b3480156103df57600080fd5b506101d7610a4e565b3480156103f457600080fd5b506101d7610ad6565b34801561040957600080fd5b50610223610b09565b34801561041e57600080fd5b50610427610b70565b60408051602080825283518183015283519192839290830191858101910280838360005b8381101561046357818101518382015260200161044b565b505050509050019250505060405180910390f35b34801561048357600080fd5b506101d7610b81565b34801561049857600080fd5b50610223610d82565b3480156104ad57600080fd5b506101fa610d88565b6101d7610d98565b3480156104ca57600080fd5b50610223610eec565b3480156104df57600080fd5b506101fa600160a060020a0360043516610ef2565b34801561050057600080fd5b50610223600160a060020a0360043516610f13565b60008282018381101561052457fe5b8091505b5092915050565b600054600160a060020a0316331461054657600080fd5b600755565b6000903b1190565b60025481565b600061056c60048363ffffffff610f4b16565b92915050565b3360009081526008602052604090205481111561058e57600080fd5b336000908152600860205260409020546105ae908263ffffffff610f6a16565b33600081815260086020526040808220939093559151909183156108fc02918491818181858888f193505050501580156105ec573d6000803e3d6000fd5b50604080513381526020810183905281517f195ddc41d185a27fe901831dcad44dd85716c95be78b1d71aa42393697966d40929181900390910190a150565b600080548190600160a060020a0316331461064557600080fd5b505060055460005b818110156106c557600580546106bd91908390811061066857fe5b600091825260208220015460058054600160a060020a03909216926008929091908690811061069357fe5b6000918252602080832090910154600160a060020a0316835282019290925260400190205461073f565b60010161064d565b6040805183815290517fa70a53972d6afb0fc38bd683cd5955faa5fa55e6629744a51e7a2aaa0ecc4e049181900360200190a15050565b600160a060020a031660009081526008602052604090205490565b60035481565b60065481565b600054600160a060020a0316331461073a57600080fd5b600655565b600054600160a060020a0316331461075657600080fd5b600160a060020a03821660009081526008602052604090205481111561077b57600080fd5b600160a060020a0382166000908152600860205260409020546107a4908263ffffffff610f6a16565b600160a060020a038316600081815260086020526040808220939093559151909183156108fc02918491818181858888f193505050501580156107eb573d6000803e3d6000fd5b5060408051600160a060020a03841681526020810183905281517f23285480e7a48c23c9ee70f743b41d58594f15a194b45ad805149ba14f8316d2929181900390910190a15050565b600054600160a060020a0316331461084b57600080fd5b600155565b600080805b6005548110156108a6576005805461089c9161088f918490811061087557fe5b600091825260209091200154600160a060020a0316610f13565b839063ffffffff61051516565b9150600101610855565b50919050565b60086020526000908152604090208054600182015460028301546003909301549192909160ff1684565b600054600160a060020a031633146108ed57600080fd5b600254610900903463ffffffff61051516565b600281905560408051918252517f162a21b4a3cda9776fd151508bc1f4fac3ceaed4b487c9aaa7888c95484502cf9181900360200190a1565b60008054600160a060020a0316331461095157600080fd5b610959610a4e565b5060005b6005548110156109be5760006008600060046001018481548110151561097f57fe5b600091825260208083209190910154600160a060020a031683528201929092526040019020600301805460ff191691151591909117905560010161095d565b60016006556109cb610b81565b6109d361062b565b506000805474ff0000000000000000000000000000000000000000191660a060020a179055565b60005460a060020a900460ff1615610a1157600080fd5b336000908152600860205260409020600301805460ff19166001179055565b600160a060020a031660009081526008602052604090206001015490565b600054600160a060020a03163314610a6557600080fd5b6000805474ff00000000000000000000000000000000000000001916908190556003546002546040805192835260a060020a90930460ff16151560208301528183015290517fc8ea7d3c44e48dda18a813373040ce0eda7c908ad2cd30b53310d9b4b30012149181900360600190a1565b60005460a060020a900460ff1615610aed57600080fd5b336000908152600860205260409020600301805460ff19169055565b600080805b6005548110156108a657610b6660086000600460010184815481101515610b3157fe5b6000918252602080832090910154600160a060020a03168352820192909252604001902060010154839063ffffffff61051516565b9150600101610b0e565b6060610b7c6004610f7c565b905090565b60008054819081908190600160a060020a03163314610b9f57600080fd5b600654421015610c1057604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f746865206e65787420726f756e64206e6f742073746172740000000000000000604482015290519081900360640190fd5b60006003541115610c2357610c23610fe2565b600354610c3790600163ffffffff61051516565b60035560075442016201517f1901600655600093505b600554841015610d58576005805485908110610c6557fe5b6000918252602080832090910154600160a060020a0316808352600890915260409091209093509150610c9783610f13565b9050600154811015610d105760028201548254610cb99163ffffffff61051516565b8255600182015415610cc757fe5b60408051600160a060020a03851681526020810183905281517f9873c485f5a9e0be9a918f4d6ad5b64912fcb8352006b316a63427b1f408e824929181900390910190a1610d46565b81546001830154610d269163ffffffff61051516565b6001808401919091556002830154835560038301805460ff191690911790555b60006002830155600190930192610c4d565b50506000805474ff0000000000000000000000000000000000000000191660a060020a1790555050565b60015481565b60005460a060020a900460ff1681565b60005460a060020a900460ff1615610daf57600080fd5b610db83361054b565b15610e4a57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f706c65617365206e6f742075736520636f6e74726163742063616c6c2074686960448201527f732066756e6374696f6e00000000000000000000000000000000000000000000606482015290519081900360840190fd5b60003311610e5757600080fd5b610e6033610559565b1515610e7957610e7760043363ffffffff6112a616565b505b33600090815260086020526040902054610e99903463ffffffff61051516565b33600081815260086020908152604091829020939093558051918252349282019290925281517f78d81951b78dad84771f88d35b4c93a632e1ed2da8706bbc7d8e465110686830929181900390910190a1565b60075481565b600160a060020a031660009081526008602052604090206003015460ff1690565b600160a060020a03811660009081526008602052604081208054600190910154610f43828263ffffffff61051516565b949350505050565b600160a060020a03166000908152602091909152604090205460ff1690565b600082821115610f7657fe5b50900390565b606081600101805480602002602001604051908101604052809291908181526020018280548015610fd657602002820191906000526020600020905b8154600160a060020a03168152600190910190602001808311610fb8575b50505050509050919050565b600080600080610ff0610b09565b9350831515610ffe576112a0565b5060005b6005548110156112a057600580546008916000918490811061102057fe5b6000918252602080832090910154600160a060020a03168352820192909252604001902060010154600254909350611070908590611064908663ffffffff61132616565b9063ffffffff61135116565b600254909250611086908363ffffffff610f6a16565b600255600580546110d7918491600891600091869081106110a357fe5b6000918252602080832090910154600160a060020a031683528201929092526040019020600201549063ffffffff61051516565b60058054600891600091859081106110eb57fe5b6000918252602080832090910154600160a060020a031683528201929092526040018120600201919091556005805460089291908490811061112957fe5b6000918252602080832090910154600160a060020a0316835282019290925260400190206003015460ff1615156111ef5761117483600860006004600101858154811015156110a357fe5b600580546008916000918590811061118857fe5b6000918252602080832090910154600160a060020a03168352820192909252604001812060020191909155600580546008918391859081106111c657fe5b6000918252602080832090910154600160a060020a031683528201929092526040019020600101555b600580547f2772659b237083773d3a2874ab3591def1a8625215ae057bde8fc4ef3dee729091908390811061122057fe5b600091825260208220015460058054600160a060020a03909216926008929091908690811061124b57fe5b6000918252602080832090910154600160a060020a03908116845283820194909452604092830190912060030154825194909316845260ff9092161515918301919091528051918290030190a1600101611002565b50505050565b600160a060020a03811660009081526020839052604081205460ff16156112cf5750600061056c565b50600160a060020a03166000818152602083815260408220805460ff1916600190811790915593840180548086018255908352912001805473ffffffffffffffffffffffffffffffffffffffff1916909117905590565b6000808315156113395760009150610528565b5082820282848281151561134957fe5b041461052457fe5b600080828481151561135f57fe5b049493505050505600a165627a7a72305820d844eae0637a6f7aa749302accfaa61e8f50f4dadbe9c36b39968b65a53315250029`

// DeployReward deploys a new cpchain contract, binding an instance of Reward to it.
func DeployReward(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Reward, error) {
	parsed, err := abi.JSON(strings.NewReader(RewardABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(RewardBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Reward{RewardCaller: RewardCaller{contract: contract}, RewardTransactor: RewardTransactor{contract: contract}, RewardFilterer: RewardFilterer{contract: contract}}, nil
}

// Reward is an auto generated Go binding around an cpchain contract.
type Reward struct {
	RewardCaller     // Read-only binding to the contract
	RewardTransactor // Write-only binding to the contract
	RewardFilterer   // Log filterer for contract events
}

// RewardCaller is an auto generated read-only Go binding around an cpchain contract.
type RewardCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RewardTransactor is an auto generated write-only Go binding around an cpchain contract.
type RewardTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RewardFilterer is an auto generated log filtering Go binding around an cpchain contract events.
type RewardFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RewardSession is an auto generated Go binding around an cpchain contract,
// with pre-set call and transact options.
type RewardSession struct {
	Contract     *Reward           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RewardCallerSession is an auto generated read-only Go binding around an cpchain contract,
// with pre-set call options.
type RewardCallerSession struct {
	Contract *RewardCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// RewardTransactorSession is an auto generated write-only Go binding around an cpchain contract,
// with pre-set transact options.
type RewardTransactorSession struct {
	Contract     *RewardTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RewardRaw is an auto generated low-level Go binding around an cpchain contract.
type RewardRaw struct {
	Contract *Reward // Generic contract binding to access the raw methods on
}

// RewardCallerRaw is an auto generated low-level read-only Go binding around an cpchain contract.
type RewardCallerRaw struct {
	Contract *RewardCaller // Generic read-only contract binding to access the raw methods on
}

// RewardTransactorRaw is an auto generated low-level write-only Go binding around an cpchain contract.
type RewardTransactorRaw struct {
	Contract *RewardTransactor // Generic write-only contract binding to access the raw methods on
}

// NewReward creates a new instance of Reward, bound to a specific deployed contract.
func NewReward(address common.Address, backend bind.ContractBackend) (*Reward, error) {
	contract, err := bindReward(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Reward{RewardCaller: RewardCaller{contract: contract}, RewardTransactor: RewardTransactor{contract: contract}, RewardFilterer: RewardFilterer{contract: contract}}, nil
}

// NewRewardCaller creates a new read-only instance of Reward, bound to a specific deployed contract.
func NewRewardCaller(address common.Address, caller bind.ContractCaller) (*RewardCaller, error) {
	contract, err := bindReward(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RewardCaller{contract: contract}, nil
}

// NewRewardTransactor creates a new write-only instance of Reward, bound to a specific deployed contract.
func NewRewardTransactor(address common.Address, transactor bind.ContractTransactor) (*RewardTransactor, error) {
	contract, err := bindReward(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RewardTransactor{contract: contract}, nil
}

// NewRewardFilterer creates a new log filterer instance of Reward, bound to a specific deployed contract.
func NewRewardFilterer(address common.Address, filterer bind.ContractFilterer) (*RewardFilterer, error) {
	contract, err := bindReward(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RewardFilterer{contract: contract}, nil
}

// bindReward binds a generic wrapper to an already deployed contract.
func bindReward(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RewardABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Reward *RewardRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Reward.Contract.RewardCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Reward *RewardRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Reward.Contract.RewardTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Reward *RewardRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Reward.Contract.RewardTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Reward *RewardCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Reward.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Reward *RewardTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Reward.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Reward *RewardTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Reward.Contract.contract.Transact(opts, method, params...)
}

// BasicCriteria is a free data retrieval call binding the contract method 0xc0209d1f.
//
// Solidity: function basicCriteria() constant returns(uint256)
func (_Reward *RewardCaller) BasicCriteria(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Reward.contract.Call(opts, out, "basicCriteria")
	return *ret0, err
}

// BasicCriteria is a free data retrieval call binding the contract method 0xc0209d1f.
//
// Solidity: function basicCriteria() constant returns(uint256)
func (_Reward *RewardSession) BasicCriteria() (*big.Int, error) {
	return _Reward.Contract.BasicCriteria(&_Reward.CallOpts)
}

// BasicCriteria is a free data retrieval call binding the contract method 0xc0209d1f.
//
// Solidity: function basicCriteria() constant returns(uint256)
func (_Reward *RewardCallerSession) BasicCriteria() (*big.Int, error) {
	return _Reward.Contract.BasicCriteria(&_Reward.CallOpts)
}

// BonusPool is a free data retrieval call binding the contract method 0x2693ee80.
//
// Solidity: function bonusPool() constant returns(uint256)
func (_Reward *RewardCaller) BonusPool(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Reward.contract.Call(opts, out, "bonusPool")
	return *ret0, err
}

// BonusPool is a free data retrieval call binding the contract method 0x2693ee80.
//
// Solidity: function bonusPool() constant returns(uint256)
func (_Reward *RewardSession) BonusPool() (*big.Int, error) {
	return _Reward.Contract.BonusPool(&_Reward.CallOpts)
}

// BonusPool is a free data retrieval call binding the contract method 0x2693ee80.
//
// Solidity: function bonusPool() constant returns(uint256)
func (_Reward *RewardCallerSession) BonusPool() (*big.Int, error) {
	return _Reward.Contract.BonusPool(&_Reward.CallOpts)
}

// GetFreeBalanceOf is a free data retrieval call binding the contract method 0x414d11e6.
//
// Solidity: function getFreeBalanceOf(_addr address) constant returns(uint256)
func (_Reward *RewardCaller) GetFreeBalanceOf(opts *bind.CallOpts, _addr common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Reward.contract.Call(opts, out, "getFreeBalanceOf", _addr)
	return *ret0, err
}

// GetFreeBalanceOf is a free data retrieval call binding the contract method 0x414d11e6.
//
// Solidity: function getFreeBalanceOf(_addr address) constant returns(uint256)
func (_Reward *RewardSession) GetFreeBalanceOf(_addr common.Address) (*big.Int, error) {
	return _Reward.Contract.GetFreeBalanceOf(&_Reward.CallOpts, _addr)
}

// GetFreeBalanceOf is a free data retrieval call binding the contract method 0x414d11e6.
//
// Solidity: function getFreeBalanceOf(_addr address) constant returns(uint256)
func (_Reward *RewardCallerSession) GetFreeBalanceOf(_addr common.Address) (*big.Int, error) {
	return _Reward.Contract.GetFreeBalanceOf(&_Reward.CallOpts, _addr)
}

// GetInvestors is a free data retrieval call binding the contract method 0xb2f5a54c.
//
// Solidity: function getInvestors() constant returns(address[])
func (_Reward *RewardCaller) GetInvestors(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _Reward.contract.Call(opts, out, "getInvestors")
	return *ret0, err
}

// GetInvestors is a free data retrieval call binding the contract method 0xb2f5a54c.
//
// Solidity: function getInvestors() constant returns(address[])
func (_Reward *RewardSession) GetInvestors() ([]common.Address, error) {
	return _Reward.Contract.GetInvestors(&_Reward.CallOpts)
}

// GetInvestors is a free data retrieval call binding the contract method 0xb2f5a54c.
//
// Solidity: function getInvestors() constant returns(address[])
func (_Reward *RewardCallerSession) GetInvestors() ([]common.Address, error) {
	return _Reward.Contract.GetInvestors(&_Reward.CallOpts)
}

// GetLockedBalanceOf is a free data retrieval call binding the contract method 0x8d0dcb5f.
//
// Solidity: function getLockedBalanceOf(_addr address) constant returns(uint256)
func (_Reward *RewardCaller) GetLockedBalanceOf(opts *bind.CallOpts, _addr common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Reward.contract.Call(opts, out, "getLockedBalanceOf", _addr)
	return *ret0, err
}

// GetLockedBalanceOf is a free data retrieval call binding the contract method 0x8d0dcb5f.
//
// Solidity: function getLockedBalanceOf(_addr address) constant returns(uint256)
func (_Reward *RewardSession) GetLockedBalanceOf(_addr common.Address) (*big.Int, error) {
	return _Reward.Contract.GetLockedBalanceOf(&_Reward.CallOpts, _addr)
}

// GetLockedBalanceOf is a free data retrieval call binding the contract method 0x8d0dcb5f.
//
// Solidity: function getLockedBalanceOf(_addr address) constant returns(uint256)
func (_Reward *RewardCallerSession) GetLockedBalanceOf(_addr common.Address) (*big.Int, error) {
	return _Reward.Contract.GetLockedBalanceOf(&_Reward.CallOpts, _addr)
}

// GetTotalAmount is a free data retrieval call binding the contract method 0x65ac4341.
//
// Solidity: function getTotalAmount() constant returns(uint256)
func (_Reward *RewardCaller) GetTotalAmount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Reward.contract.Call(opts, out, "getTotalAmount")
	return *ret0, err
}

// GetTotalAmount is a free data retrieval call binding the contract method 0x65ac4341.
//
// Solidity: function getTotalAmount() constant returns(uint256)
func (_Reward *RewardSession) GetTotalAmount() (*big.Int, error) {
	return _Reward.Contract.GetTotalAmount(&_Reward.CallOpts)
}

// GetTotalAmount is a free data retrieval call binding the contract method 0x65ac4341.
//
// Solidity: function getTotalAmount() constant returns(uint256)
func (_Reward *RewardCallerSession) GetTotalAmount() (*big.Int, error) {
	return _Reward.Contract.GetTotalAmount(&_Reward.CallOpts)
}

// GetTotalBalanceOf is a free data retrieval call binding the contract method 0xfdfefa7f.
//
// Solidity: function getTotalBalanceOf(_addr address) constant returns(uint256)
func (_Reward *RewardCaller) GetTotalBalanceOf(opts *bind.CallOpts, _addr common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Reward.contract.Call(opts, out, "getTotalBalanceOf", _addr)
	return *ret0, err
}

// GetTotalBalanceOf is a free data retrieval call binding the contract method 0xfdfefa7f.
//
// Solidity: function getTotalBalanceOf(_addr address) constant returns(uint256)
func (_Reward *RewardSession) GetTotalBalanceOf(_addr common.Address) (*big.Int, error) {
	return _Reward.Contract.GetTotalBalanceOf(&_Reward.CallOpts, _addr)
}

// GetTotalBalanceOf is a free data retrieval call binding the contract method 0xfdfefa7f.
//
// Solidity: function getTotalBalanceOf(_addr address) constant returns(uint256)
func (_Reward *RewardCallerSession) GetTotalBalanceOf(_addr common.Address) (*big.Int, error) {
	return _Reward.Contract.GetTotalBalanceOf(&_Reward.CallOpts, _addr)
}

// GetTotalLockedAmount is a free data retrieval call binding the contract method 0x91d2b32e.
//
// Solidity: function getTotalLockedAmount() constant returns(uint256)
func (_Reward *RewardCaller) GetTotalLockedAmount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Reward.contract.Call(opts, out, "getTotalLockedAmount")
	return *ret0, err
}

// GetTotalLockedAmount is a free data retrieval call binding the contract method 0x91d2b32e.
//
// Solidity: function getTotalLockedAmount() constant returns(uint256)
func (_Reward *RewardSession) GetTotalLockedAmount() (*big.Int, error) {
	return _Reward.Contract.GetTotalLockedAmount(&_Reward.CallOpts)
}

// GetTotalLockedAmount is a free data retrieval call binding the contract method 0x91d2b32e.
//
// Solidity: function getTotalLockedAmount() constant returns(uint256)
func (_Reward *RewardCallerSession) GetTotalLockedAmount() (*big.Int, error) {
	return _Reward.Contract.GetTotalLockedAmount(&_Reward.CallOpts)
}

// Investors is a free data retrieval call binding the contract method 0x6f7bc9be.
//
// Solidity: function investors( address) constant returns(freeDeposit uint256, lockedDeposit uint256, returned uint256, toRenew bool)
func (_Reward *RewardCaller) Investors(opts *bind.CallOpts, arg0 common.Address) (struct {
	FreeDeposit   *big.Int
	LockedDeposit *big.Int
	Returned      *big.Int
	ToRenew       bool
}, error) {
	ret := new(struct {
		FreeDeposit   *big.Int
		LockedDeposit *big.Int
		Returned      *big.Int
		ToRenew       bool
	})
	out := ret
	err := _Reward.contract.Call(opts, out, "investors", arg0)
	return *ret, err
}

// Investors is a free data retrieval call binding the contract method 0x6f7bc9be.
//
// Solidity: function investors( address) constant returns(freeDeposit uint256, lockedDeposit uint256, returned uint256, toRenew bool)
func (_Reward *RewardSession) Investors(arg0 common.Address) (struct {
	FreeDeposit   *big.Int
	LockedDeposit *big.Int
	Returned      *big.Int
	ToRenew       bool
}, error) {
	return _Reward.Contract.Investors(&_Reward.CallOpts, arg0)
}

// Investors is a free data retrieval call binding the contract method 0x6f7bc9be.
//
// Solidity: function investors( address) constant returns(freeDeposit uint256, lockedDeposit uint256, returned uint256, toRenew bool)
func (_Reward *RewardCallerSession) Investors(arg0 common.Address) (struct {
	FreeDeposit   *big.Int
	LockedDeposit *big.Int
	Returned      *big.Int
	ToRenew       bool
}, error) {
	return _Reward.Contract.Investors(&_Reward.CallOpts, arg0)
}

// IsContract is a free data retrieval call binding the contract method 0x16279055.
//
// Solidity: function isContract(addr address) constant returns(bool)
func (_Reward *RewardCaller) IsContract(opts *bind.CallOpts, addr common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Reward.contract.Call(opts, out, "isContract", addr)
	return *ret0, err
}

// IsContract is a free data retrieval call binding the contract method 0x16279055.
//
// Solidity: function isContract(addr address) constant returns(bool)
func (_Reward *RewardSession) IsContract(addr common.Address) (bool, error) {
	return _Reward.Contract.IsContract(&_Reward.CallOpts, addr)
}

// IsContract is a free data retrieval call binding the contract method 0x16279055.
//
// Solidity: function isContract(addr address) constant returns(bool)
func (_Reward *RewardCallerSession) IsContract(addr common.Address) (bool, error) {
	return _Reward.Contract.IsContract(&_Reward.CallOpts, addr)
}

// IsEnode is a free data retrieval call binding the contract method 0x2a0e96c2.
//
// Solidity: function isEnode(_addr address) constant returns(bool)
func (_Reward *RewardCaller) IsEnode(opts *bind.CallOpts, _addr common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Reward.contract.Call(opts, out, "isEnode", _addr)
	return *ret0, err
}

// IsEnode is a free data retrieval call binding the contract method 0x2a0e96c2.
//
// Solidity: function isEnode(_addr address) constant returns(bool)
func (_Reward *RewardSession) IsEnode(_addr common.Address) (bool, error) {
	return _Reward.Contract.IsEnode(&_Reward.CallOpts, _addr)
}

// IsEnode is a free data retrieval call binding the contract method 0x2a0e96c2.
//
// Solidity: function isEnode(_addr address) constant returns(bool)
func (_Reward *RewardCallerSession) IsEnode(_addr common.Address) (bool, error) {
	return _Reward.Contract.IsEnode(&_Reward.CallOpts, _addr)
}

// IsToRenew is a free data retrieval call binding the contract method 0xf509dd4a.
//
// Solidity: function isToRenew(_addr address) constant returns(bool)
func (_Reward *RewardCaller) IsToRenew(opts *bind.CallOpts, _addr common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Reward.contract.Call(opts, out, "isToRenew", _addr)
	return *ret0, err
}

// IsToRenew is a free data retrieval call binding the contract method 0xf509dd4a.
//
// Solidity: function isToRenew(_addr address) constant returns(bool)
func (_Reward *RewardSession) IsToRenew(_addr common.Address) (bool, error) {
	return _Reward.Contract.IsToRenew(&_Reward.CallOpts, _addr)
}

// IsToRenew is a free data retrieval call binding the contract method 0xf509dd4a.
//
// Solidity: function isToRenew(_addr address) constant returns(bool)
func (_Reward *RewardCallerSession) IsToRenew(_addr common.Address) (bool, error) {
	return _Reward.Contract.IsToRenew(&_Reward.CallOpts, _addr)
}

// Locked is a free data retrieval call binding the contract method 0xcf309012.
//
// Solidity: function locked() constant returns(bool)
func (_Reward *RewardCaller) Locked(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Reward.contract.Call(opts, out, "locked")
	return *ret0, err
}

// Locked is a free data retrieval call binding the contract method 0xcf309012.
//
// Solidity: function locked() constant returns(bool)
func (_Reward *RewardSession) Locked() (bool, error) {
	return _Reward.Contract.Locked(&_Reward.CallOpts)
}

// Locked is a free data retrieval call binding the contract method 0xcf309012.
//
// Solidity: function locked() constant returns(bool)
func (_Reward *RewardCallerSession) Locked() (bool, error) {
	return _Reward.Contract.Locked(&_Reward.CallOpts)
}

// NextRound is a free data retrieval call binding the contract method 0x47e40553.
//
// Solidity: function nextRound() constant returns(uint256)
func (_Reward *RewardCaller) NextRound(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Reward.contract.Call(opts, out, "nextRound")
	return *ret0, err
}

// NextRound is a free data retrieval call binding the contract method 0x47e40553.
//
// Solidity: function nextRound() constant returns(uint256)
func (_Reward *RewardSession) NextRound() (*big.Int, error) {
	return _Reward.Contract.NextRound(&_Reward.CallOpts)
}

// NextRound is a free data retrieval call binding the contract method 0x47e40553.
//
// Solidity: function nextRound() constant returns(uint256)
func (_Reward *RewardCallerSession) NextRound() (*big.Int, error) {
	return _Reward.Contract.NextRound(&_Reward.CallOpts)
}

// NextRoundStartTime is a free data retrieval call binding the contract method 0x4aeb31eb.
//
// Solidity: function nextRoundStartTime() constant returns(uint256)
func (_Reward *RewardCaller) NextRoundStartTime(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Reward.contract.Call(opts, out, "nextRoundStartTime")
	return *ret0, err
}

// NextRoundStartTime is a free data retrieval call binding the contract method 0x4aeb31eb.
//
// Solidity: function nextRoundStartTime() constant returns(uint256)
func (_Reward *RewardSession) NextRoundStartTime() (*big.Int, error) {
	return _Reward.Contract.NextRoundStartTime(&_Reward.CallOpts)
}

// NextRoundStartTime is a free data retrieval call binding the contract method 0x4aeb31eb.
//
// Solidity: function nextRoundStartTime() constant returns(uint256)
func (_Reward *RewardCallerSession) NextRoundStartTime() (*big.Int, error) {
	return _Reward.Contract.NextRoundStartTime(&_Reward.CallOpts)
}

// Period is a free data retrieval call binding the contract method 0xef78d4fd.
//
// Solidity: function period() constant returns(uint256)
func (_Reward *RewardCaller) Period(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Reward.contract.Call(opts, out, "period")
	return *ret0, err
}

// Period is a free data retrieval call binding the contract method 0xef78d4fd.
//
// Solidity: function period() constant returns(uint256)
func (_Reward *RewardSession) Period() (*big.Int, error) {
	return _Reward.Contract.Period(&_Reward.CallOpts)
}

// Period is a free data retrieval call binding the contract method 0xef78d4fd.
//
// Solidity: function period() constant returns(uint256)
func (_Reward *RewardCallerSession) Period() (*big.Int, error) {
	return _Reward.Contract.Period(&_Reward.CallOpts)
}

// DisableContract is a paid mutator transaction binding the contract method 0x894ba833.
//
// Solidity: function disableContract() returns()
func (_Reward *RewardTransactor) DisableContract(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Reward.contract.Transact(opts, "disableContract")
}

// DisableContract is a paid mutator transaction binding the contract method 0x894ba833.
//
// Solidity: function disableContract() returns()
func (_Reward *RewardSession) DisableContract() (*types.Transaction, error) {
	return _Reward.Contract.DisableContract(&_Reward.TransactOpts)
}

// DisableContract is a paid mutator transaction binding the contract method 0x894ba833.
//
// Solidity: function disableContract() returns()
func (_Reward *RewardTransactorSession) DisableContract() (*types.Transaction, error) {
	return _Reward.Contract.DisableContract(&_Reward.TransactOpts)
}

// NewRaise is a paid mutator transaction binding the contract method 0x8f83bc7f.
//
// Solidity: function newRaise() returns()
func (_Reward *RewardTransactor) NewRaise(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Reward.contract.Transact(opts, "newRaise")
}

// NewRaise is a paid mutator transaction binding the contract method 0x8f83bc7f.
//
// Solidity: function newRaise() returns()
func (_Reward *RewardSession) NewRaise() (*types.Transaction, error) {
	return _Reward.Contract.NewRaise(&_Reward.TransactOpts)
}

// NewRaise is a paid mutator transaction binding the contract method 0x8f83bc7f.
//
// Solidity: function newRaise() returns()
func (_Reward *RewardTransactorSession) NewRaise() (*types.Transaction, error) {
	return _Reward.Contract.NewRaise(&_Reward.TransactOpts)
}

// QuitRenew is a paid mutator transaction binding the contract method 0x8ffb6376.
//
// Solidity: function quitRenew() returns()
func (_Reward *RewardTransactor) QuitRenew(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Reward.contract.Transact(opts, "quitRenew")
}

// QuitRenew is a paid mutator transaction binding the contract method 0x8ffb6376.
//
// Solidity: function quitRenew() returns()
func (_Reward *RewardSession) QuitRenew() (*types.Transaction, error) {
	return _Reward.Contract.QuitRenew(&_Reward.TransactOpts)
}

// QuitRenew is a paid mutator transaction binding the contract method 0x8ffb6376.
//
// Solidity: function quitRenew() returns()
func (_Reward *RewardTransactorSession) QuitRenew() (*types.Transaction, error) {
	return _Reward.Contract.QuitRenew(&_Reward.TransactOpts)
}

// RefundAll is a paid mutator transaction binding the contract method 0x38e771ab.
//
// Solidity: function refundAll() returns()
func (_Reward *RewardTransactor) RefundAll(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Reward.contract.Transact(opts, "refundAll")
}

// RefundAll is a paid mutator transaction binding the contract method 0x38e771ab.
//
// Solidity: function refundAll() returns()
func (_Reward *RewardSession) RefundAll() (*types.Transaction, error) {
	return _Reward.Contract.RefundAll(&_Reward.TransactOpts)
}

// RefundAll is a paid mutator transaction binding the contract method 0x38e771ab.
//
// Solidity: function refundAll() returns()
func (_Reward *RewardTransactorSession) RefundAll() (*types.Transaction, error) {
	return _Reward.Contract.RefundAll(&_Reward.TransactOpts)
}

// RefundDeposit is a paid mutator transaction binding the contract method 0x5215d054.
//
// Solidity: function refundDeposit(_addr address, _value uint256) returns()
func (_Reward *RewardTransactor) RefundDeposit(opts *bind.TransactOpts, _addr common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Reward.contract.Transact(opts, "refundDeposit", _addr, _value)
}

// RefundDeposit is a paid mutator transaction binding the contract method 0x5215d054.
//
// Solidity: function refundDeposit(_addr address, _value uint256) returns()
func (_Reward *RewardSession) RefundDeposit(_addr common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Reward.Contract.RefundDeposit(&_Reward.TransactOpts, _addr, _value)
}

// RefundDeposit is a paid mutator transaction binding the contract method 0x5215d054.
//
// Solidity: function refundDeposit(_addr address, _value uint256) returns()
func (_Reward *RewardTransactorSession) RefundDeposit(_addr common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Reward.Contract.RefundDeposit(&_Reward.TransactOpts, _addr, _value)
}

// SetBonusPool is a paid mutator transaction binding the contract method 0x7c0fa595.
//
// Solidity: function setBonusPool() returns()
func (_Reward *RewardTransactor) SetBonusPool(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Reward.contract.Transact(opts, "setBonusPool")
}

// SetBonusPool is a paid mutator transaction binding the contract method 0x7c0fa595.
//
// Solidity: function setBonusPool() returns()
func (_Reward *RewardSession) SetBonusPool() (*types.Transaction, error) {
	return _Reward.Contract.SetBonusPool(&_Reward.TransactOpts)
}

// SetBonusPool is a paid mutator transaction binding the contract method 0x7c0fa595.
//
// Solidity: function setBonusPool() returns()
func (_Reward *RewardTransactorSession) SetBonusPool() (*types.Transaction, error) {
	return _Reward.Contract.SetBonusPool(&_Reward.TransactOpts)
}

// SetCriteria is a paid mutator transaction binding the contract method 0x5f2c01ea.
//
// Solidity: function setCriteria(criteria uint256) returns()
func (_Reward *RewardTransactor) SetCriteria(opts *bind.TransactOpts, criteria *big.Int) (*types.Transaction, error) {
	return _Reward.contract.Transact(opts, "setCriteria", criteria)
}

// SetCriteria is a paid mutator transaction binding the contract method 0x5f2c01ea.
//
// Solidity: function setCriteria(criteria uint256) returns()
func (_Reward *RewardSession) SetCriteria(criteria *big.Int) (*types.Transaction, error) {
	return _Reward.Contract.SetCriteria(&_Reward.TransactOpts, criteria)
}

// SetCriteria is a paid mutator transaction binding the contract method 0x5f2c01ea.
//
// Solidity: function setCriteria(criteria uint256) returns()
func (_Reward *RewardTransactorSession) SetCriteria(criteria *big.Int) (*types.Transaction, error) {
	return _Reward.Contract.SetCriteria(&_Reward.TransactOpts, criteria)
}

// SetNextRoundStartTime is a paid mutator transaction binding the contract method 0x4d5b03f3.
//
// Solidity: function setNextRoundStartTime(time uint256) returns()
func (_Reward *RewardTransactor) SetNextRoundStartTime(opts *bind.TransactOpts, time *big.Int) (*types.Transaction, error) {
	return _Reward.contract.Transact(opts, "setNextRoundStartTime", time)
}

// SetNextRoundStartTime is a paid mutator transaction binding the contract method 0x4d5b03f3.
//
// Solidity: function setNextRoundStartTime(time uint256) returns()
func (_Reward *RewardSession) SetNextRoundStartTime(time *big.Int) (*types.Transaction, error) {
	return _Reward.Contract.SetNextRoundStartTime(&_Reward.TransactOpts, time)
}

// SetNextRoundStartTime is a paid mutator transaction binding the contract method 0x4d5b03f3.
//
// Solidity: function setNextRoundStartTime(time uint256) returns()
func (_Reward *RewardTransactorSession) SetNextRoundStartTime(time *big.Int) (*types.Transaction, error) {
	return _Reward.Contract.SetNextRoundStartTime(&_Reward.TransactOpts, time)
}

// SetPeriod is a paid mutator transaction binding the contract method 0x0f3a9f65.
//
// Solidity: function setPeriod(_period uint256) returns()
func (_Reward *RewardTransactor) SetPeriod(opts *bind.TransactOpts, _period *big.Int) (*types.Transaction, error) {
	return _Reward.contract.Transact(opts, "setPeriod", _period)
}

// SetPeriod is a paid mutator transaction binding the contract method 0x0f3a9f65.
//
// Solidity: function setPeriod(_period uint256) returns()
func (_Reward *RewardSession) SetPeriod(_period *big.Int) (*types.Transaction, error) {
	return _Reward.Contract.SetPeriod(&_Reward.TransactOpts, _period)
}

// SetPeriod is a paid mutator transaction binding the contract method 0x0f3a9f65.
//
// Solidity: function setPeriod(_period uint256) returns()
func (_Reward *RewardTransactorSession) SetPeriod(_period *big.Int) (*types.Transaction, error) {
	return _Reward.Contract.SetPeriod(&_Reward.TransactOpts, _period)
}

// StartNewRound is a paid mutator transaction binding the contract method 0xbd85948c.
//
// Solidity: function startNewRound() returns()
func (_Reward *RewardTransactor) StartNewRound(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Reward.contract.Transact(opts, "startNewRound")
}

// StartNewRound is a paid mutator transaction binding the contract method 0xbd85948c.
//
// Solidity: function startNewRound() returns()
func (_Reward *RewardSession) StartNewRound() (*types.Transaction, error) {
	return _Reward.Contract.StartNewRound(&_Reward.TransactOpts)
}

// StartNewRound is a paid mutator transaction binding the contract method 0xbd85948c.
//
// Solidity: function startNewRound() returns()
func (_Reward *RewardTransactorSession) StartNewRound() (*types.Transaction, error) {
	return _Reward.Contract.StartNewRound(&_Reward.TransactOpts)
}

// SubmitDeposit is a paid mutator transaction binding the contract method 0xe9843a3e.
//
// Solidity: function submitDeposit() returns()
func (_Reward *RewardTransactor) SubmitDeposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Reward.contract.Transact(opts, "submitDeposit")
}

// SubmitDeposit is a paid mutator transaction binding the contract method 0xe9843a3e.
//
// Solidity: function submitDeposit() returns()
func (_Reward *RewardSession) SubmitDeposit() (*types.Transaction, error) {
	return _Reward.Contract.SubmitDeposit(&_Reward.TransactOpts)
}

// SubmitDeposit is a paid mutator transaction binding the contract method 0xe9843a3e.
//
// Solidity: function submitDeposit() returns()
func (_Reward *RewardTransactorSession) SubmitDeposit() (*types.Transaction, error) {
	return _Reward.Contract.SubmitDeposit(&_Reward.TransactOpts)
}

// WantRenew is a paid mutator transaction binding the contract method 0x8ca2ca7f.
//
// Solidity: function wantRenew() returns()
func (_Reward *RewardTransactor) WantRenew(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Reward.contract.Transact(opts, "wantRenew")
}

// WantRenew is a paid mutator transaction binding the contract method 0x8ca2ca7f.
//
// Solidity: function wantRenew() returns()
func (_Reward *RewardSession) WantRenew() (*types.Transaction, error) {
	return _Reward.Contract.WantRenew(&_Reward.TransactOpts)
}

// WantRenew is a paid mutator transaction binding the contract method 0x8ca2ca7f.
//
// Solidity: function wantRenew() returns()
func (_Reward *RewardTransactorSession) WantRenew() (*types.Transaction, error) {
	return _Reward.Contract.WantRenew(&_Reward.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(_value uint256) returns()
func (_Reward *RewardTransactor) Withdraw(opts *bind.TransactOpts, _value *big.Int) (*types.Transaction, error) {
	return _Reward.contract.Transact(opts, "withdraw", _value)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(_value uint256) returns()
func (_Reward *RewardSession) Withdraw(_value *big.Int) (*types.Transaction, error) {
	return _Reward.Contract.Withdraw(&_Reward.TransactOpts, _value)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(_value uint256) returns()
func (_Reward *RewardTransactorSession) Withdraw(_value *big.Int) (*types.Transaction, error) {
	return _Reward.Contract.Withdraw(&_Reward.TransactOpts, _value)
}

// RewardContinuedInvestIterator is returned from FilterContinuedInvest and is used to iterate over the raw logs and unpacked data for ContinuedInvest events raised by the Reward contract.
type RewardContinuedInvestIterator struct {
	Event *RewardContinuedInvest // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log       // Log channel receiving the found contract events
	sub  cpchain.Subscription // Subscription for errors, completion and termination
	done bool                 // Whether the subscription completed delivering logs
	fail error                // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RewardContinuedInvestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RewardContinuedInvest)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RewardContinuedInvest)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RewardContinuedInvestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RewardContinuedInvestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RewardContinuedInvest represents a ContinuedInvest event raised by the Reward contract.
type RewardContinuedInvest struct {
	Addr       common.Address
	Iscontinue bool
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterContinuedInvest is a free log retrieval operation binding the contract event 0x2772659b237083773d3a2874ab3591def1a8625215ae057bde8fc4ef3dee7290.
//
// Solidity: e ContinuedInvest(_addr address, _iscontinue bool)
func (_Reward *RewardFilterer) FilterContinuedInvest(opts *bind.FilterOpts) (*RewardContinuedInvestIterator, error) {

	logs, sub, err := _Reward.contract.FilterLogs(opts, "ContinuedInvest")
	if err != nil {
		return nil, err
	}
	return &RewardContinuedInvestIterator{contract: _Reward.contract, event: "ContinuedInvest", logs: logs, sub: sub}, nil
}

// WatchContinuedInvest is a free log subscription operation binding the contract event 0x2772659b237083773d3a2874ab3591def1a8625215ae057bde8fc4ef3dee7290.
//
// Solidity: e ContinuedInvest(_addr address, _iscontinue bool)
func (_Reward *RewardFilterer) WatchContinuedInvest(opts *bind.WatchOpts, sink chan<- *RewardContinuedInvest) (event.Subscription, error) {

	logs, sub, err := _Reward.contract.WatchLogs(opts, "ContinuedInvest")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RewardContinuedInvest)
				if err := _Reward.contract.UnpackLog(event, "ContinuedInvest", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// RewardDepositInsufficientIterator is returned from FilterDepositInsufficient and is used to iterate over the raw logs and unpacked data for DepositInsufficient events raised by the Reward contract.
type RewardDepositInsufficientIterator struct {
	Event *RewardDepositInsufficient // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log       // Log channel receiving the found contract events
	sub  cpchain.Subscription // Subscription for errors, completion and termination
	done bool                 // Whether the subscription completed delivering logs
	fail error                // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RewardDepositInsufficientIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RewardDepositInsufficient)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RewardDepositInsufficient)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RewardDepositInsufficientIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RewardDepositInsufficientIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RewardDepositInsufficient represents a DepositInsufficient event raised by the Reward contract.
type RewardDepositInsufficient struct {
	Who   common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterDepositInsufficient is a free log retrieval operation binding the contract event 0x9873c485f5a9e0be9a918f4d6ad5b64912fcb8352006b316a63427b1f408e824.
//
// Solidity: e DepositInsufficient(who address, value uint256)
func (_Reward *RewardFilterer) FilterDepositInsufficient(opts *bind.FilterOpts) (*RewardDepositInsufficientIterator, error) {

	logs, sub, err := _Reward.contract.FilterLogs(opts, "DepositInsufficient")
	if err != nil {
		return nil, err
	}
	return &RewardDepositInsufficientIterator{contract: _Reward.contract, event: "DepositInsufficient", logs: logs, sub: sub}, nil
}

// WatchDepositInsufficient is a free log subscription operation binding the contract event 0x9873c485f5a9e0be9a918f4d6ad5b64912fcb8352006b316a63427b1f408e824.
//
// Solidity: e DepositInsufficient(who address, value uint256)
func (_Reward *RewardFilterer) WatchDepositInsufficient(opts *bind.WatchOpts, sink chan<- *RewardDepositInsufficient) (event.Subscription, error) {

	logs, sub, err := _Reward.contract.WatchLogs(opts, "DepositInsufficient")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RewardDepositInsufficient)
				if err := _Reward.contract.UnpackLog(event, "DepositInsufficient", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// RewardFundBonusPoolIterator is returned from FilterFundBonusPool and is used to iterate over the raw logs and unpacked data for FundBonusPool events raised by the Reward contract.
type RewardFundBonusPoolIterator struct {
	Event *RewardFundBonusPool // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log       // Log channel receiving the found contract events
	sub  cpchain.Subscription // Subscription for errors, completion and termination
	done bool                 // Whether the subscription completed delivering logs
	fail error                // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RewardFundBonusPoolIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RewardFundBonusPool)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RewardFundBonusPool)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RewardFundBonusPoolIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RewardFundBonusPoolIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RewardFundBonusPool represents a FundBonusPool event raised by the Reward contract.
type RewardFundBonusPool struct {
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterFundBonusPool is a free log retrieval operation binding the contract event 0x71030773066b852afef8d0f98dbfdaec8e9a62f2f5533916ec7dfa15a0edc1f2.
//
// Solidity: e FundBonusPool(value uint256)
func (_Reward *RewardFilterer) FilterFundBonusPool(opts *bind.FilterOpts) (*RewardFundBonusPoolIterator, error) {

	logs, sub, err := _Reward.contract.FilterLogs(opts, "FundBonusPool")
	if err != nil {
		return nil, err
	}
	return &RewardFundBonusPoolIterator{contract: _Reward.contract, event: "FundBonusPool", logs: logs, sub: sub}, nil
}

// WatchFundBonusPool is a free log subscription operation binding the contract event 0x71030773066b852afef8d0f98dbfdaec8e9a62f2f5533916ec7dfa15a0edc1f2.
//
// Solidity: e FundBonusPool(value uint256)
func (_Reward *RewardFilterer) WatchFundBonusPool(opts *bind.WatchOpts, sink chan<- *RewardFundBonusPool) (event.Subscription, error) {

	logs, sub, err := _Reward.contract.WatchLogs(opts, "FundBonusPool")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RewardFundBonusPool)
				if err := _Reward.contract.UnpackLog(event, "FundBonusPool", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// RewardJoinEnodesIterator is returned from FilterJoinEnodes and is used to iterate over the raw logs and unpacked data for JoinEnodes events raised by the Reward contract.
type RewardJoinEnodesIterator struct {
	Event *RewardJoinEnodes // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log       // Log channel receiving the found contract events
	sub  cpchain.Subscription // Subscription for errors, completion and termination
	done bool                 // Whether the subscription completed delivering logs
	fail error                // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RewardJoinEnodesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RewardJoinEnodes)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RewardJoinEnodes)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RewardJoinEnodesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RewardJoinEnodesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RewardJoinEnodes represents a JoinEnodes event raised by the Reward contract.
type RewardJoinEnodes struct {
	Who   common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterJoinEnodes is a free log retrieval operation binding the contract event 0x36743ba59d8493fafba069f65f5f2af34c863a3f1576a9bb9f85713cf3fabffd.
//
// Solidity: e JoinEnodes(who address, value uint256)
func (_Reward *RewardFilterer) FilterJoinEnodes(opts *bind.FilterOpts) (*RewardJoinEnodesIterator, error) {

	logs, sub, err := _Reward.contract.FilterLogs(opts, "JoinEnodes")
	if err != nil {
		return nil, err
	}
	return &RewardJoinEnodesIterator{contract: _Reward.contract, event: "JoinEnodes", logs: logs, sub: sub}, nil
}

// WatchJoinEnodes is a free log subscription operation binding the contract event 0x36743ba59d8493fafba069f65f5f2af34c863a3f1576a9bb9f85713cf3fabffd.
//
// Solidity: e JoinEnodes(who address, value uint256)
func (_Reward *RewardFilterer) WatchJoinEnodes(opts *bind.WatchOpts, sink chan<- *RewardJoinEnodes) (event.Subscription, error) {

	logs, sub, err := _Reward.contract.WatchLogs(opts, "JoinEnodes")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RewardJoinEnodes)
				if err := _Reward.contract.UnpackLog(event, "JoinEnodes", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// RewardNewRaiseIterator is returned from FilterNewRaise and is used to iterate over the raw logs and unpacked data for NewRaise events raised by the Reward contract.
type RewardNewRaiseIterator struct {
	Event *RewardNewRaise // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log       // Log channel receiving the found contract events
	sub  cpchain.Subscription // Subscription for errors, completion and termination
	done bool                 // Whether the subscription completed delivering logs
	fail error                // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RewardNewRaiseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RewardNewRaise)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RewardNewRaise)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RewardNewRaiseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RewardNewRaiseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RewardNewRaise represents a NewRaise event raised by the Reward contract.
type RewardNewRaise struct {
	Round     *big.Int
	Lock      bool
	BonusPool *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterNewRaise is a free log retrieval operation binding the contract event 0xc8ea7d3c44e48dda18a813373040ce0eda7c908ad2cd30b53310d9b4b3001214.
//
// Solidity: e NewRaise(round uint256, lock bool, _bonusPool uint256)
func (_Reward *RewardFilterer) FilterNewRaise(opts *bind.FilterOpts) (*RewardNewRaiseIterator, error) {

	logs, sub, err := _Reward.contract.FilterLogs(opts, "NewRaise")
	if err != nil {
		return nil, err
	}
	return &RewardNewRaiseIterator{contract: _Reward.contract, event: "NewRaise", logs: logs, sub: sub}, nil
}

// WatchNewRaise is a free log subscription operation binding the contract event 0xc8ea7d3c44e48dda18a813373040ce0eda7c908ad2cd30b53310d9b4b3001214.
//
// Solidity: e NewRaise(round uint256, lock bool, _bonusPool uint256)
func (_Reward *RewardFilterer) WatchNewRaise(opts *bind.WatchOpts, sink chan<- *RewardNewRaise) (event.Subscription, error) {

	logs, sub, err := _Reward.contract.WatchLogs(opts, "NewRaise")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RewardNewRaise)
				if err := _Reward.contract.UnpackLog(event, "NewRaise", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// RewardOwnerSetBonusPoolIterator is returned from FilterOwnerSetBonusPool and is used to iterate over the raw logs and unpacked data for OwnerSetBonusPool events raised by the Reward contract.
type RewardOwnerSetBonusPoolIterator struct {
	Event *RewardOwnerSetBonusPool // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log       // Log channel receiving the found contract events
	sub  cpchain.Subscription // Subscription for errors, completion and termination
	done bool                 // Whether the subscription completed delivering logs
	fail error                // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RewardOwnerSetBonusPoolIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RewardOwnerSetBonusPool)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RewardOwnerSetBonusPool)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RewardOwnerSetBonusPoolIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RewardOwnerSetBonusPoolIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RewardOwnerSetBonusPool represents a OwnerSetBonusPool event raised by the Reward contract.
type RewardOwnerSetBonusPool struct {
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterOwnerSetBonusPool is a free log retrieval operation binding the contract event 0x162a21b4a3cda9776fd151508bc1f4fac3ceaed4b487c9aaa7888c95484502cf.
//
// Solidity: e OwnerSetBonusPool(value uint256)
func (_Reward *RewardFilterer) FilterOwnerSetBonusPool(opts *bind.FilterOpts) (*RewardOwnerSetBonusPoolIterator, error) {

	logs, sub, err := _Reward.contract.FilterLogs(opts, "OwnerSetBonusPool")
	if err != nil {
		return nil, err
	}
	return &RewardOwnerSetBonusPoolIterator{contract: _Reward.contract, event: "OwnerSetBonusPool", logs: logs, sub: sub}, nil
}

// WatchOwnerSetBonusPool is a free log subscription operation binding the contract event 0x162a21b4a3cda9776fd151508bc1f4fac3ceaed4b487c9aaa7888c95484502cf.
//
// Solidity: e OwnerSetBonusPool(value uint256)
func (_Reward *RewardFilterer) WatchOwnerSetBonusPool(opts *bind.WatchOpts, sink chan<- *RewardOwnerSetBonusPool) (event.Subscription, error) {

	logs, sub, err := _Reward.contract.WatchLogs(opts, "OwnerSetBonusPool")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RewardOwnerSetBonusPool)
				if err := _Reward.contract.UnpackLog(event, "OwnerSetBonusPool", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// RewardRefundAllIterator is returned from FilterRefundAll and is used to iterate over the raw logs and unpacked data for RefundAll events raised by the Reward contract.
type RewardRefundAllIterator struct {
	Event *RewardRefundAll // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log       // Log channel receiving the found contract events
	sub  cpchain.Subscription // Subscription for errors, completion and termination
	done bool                 // Whether the subscription completed delivering logs
	fail error                // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RewardRefundAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RewardRefundAll)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RewardRefundAll)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RewardRefundAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RewardRefundAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RewardRefundAll represents a RefundAll event raised by the Reward contract.
type RewardRefundAll struct {
	Num *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterRefundAll is a free log retrieval operation binding the contract event 0xa70a53972d6afb0fc38bd683cd5955faa5fa55e6629744a51e7a2aaa0ecc4e04.
//
// Solidity: e RefundAll(num uint256)
func (_Reward *RewardFilterer) FilterRefundAll(opts *bind.FilterOpts) (*RewardRefundAllIterator, error) {

	logs, sub, err := _Reward.contract.FilterLogs(opts, "RefundAll")
	if err != nil {
		return nil, err
	}
	return &RewardRefundAllIterator{contract: _Reward.contract, event: "RefundAll", logs: logs, sub: sub}, nil
}

// WatchRefundAll is a free log subscription operation binding the contract event 0xa70a53972d6afb0fc38bd683cd5955faa5fa55e6629744a51e7a2aaa0ecc4e04.
//
// Solidity: e RefundAll(num uint256)
func (_Reward *RewardFilterer) WatchRefundAll(opts *bind.WatchOpts, sink chan<- *RewardRefundAll) (event.Subscription, error) {

	logs, sub, err := _Reward.contract.WatchLogs(opts, "RefundAll")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RewardRefundAll)
				if err := _Reward.contract.UnpackLog(event, "RefundAll", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// RewardRefundDepositIterator is returned from FilterRefundDeposit and is used to iterate over the raw logs and unpacked data for RefundDeposit events raised by the Reward contract.
type RewardRefundDepositIterator struct {
	Event *RewardRefundDeposit // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log       // Log channel receiving the found contract events
	sub  cpchain.Subscription // Subscription for errors, completion and termination
	done bool                 // Whether the subscription completed delivering logs
	fail error                // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RewardRefundDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RewardRefundDeposit)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RewardRefundDeposit)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RewardRefundDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RewardRefundDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RewardRefundDeposit represents a RefundDeposit event raised by the Reward contract.
type RewardRefundDeposit struct {
	Who   common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterRefundDeposit is a free log retrieval operation binding the contract event 0x23285480e7a48c23c9ee70f743b41d58594f15a194b45ad805149ba14f8316d2.
//
// Solidity: e RefundDeposit(who address, value uint256)
func (_Reward *RewardFilterer) FilterRefundDeposit(opts *bind.FilterOpts) (*RewardRefundDepositIterator, error) {

	logs, sub, err := _Reward.contract.FilterLogs(opts, "RefundDeposit")
	if err != nil {
		return nil, err
	}
	return &RewardRefundDepositIterator{contract: _Reward.contract, event: "RefundDeposit", logs: logs, sub: sub}, nil
}

// WatchRefundDeposit is a free log subscription operation binding the contract event 0x23285480e7a48c23c9ee70f743b41d58594f15a194b45ad805149ba14f8316d2.
//
// Solidity: e RefundDeposit(who address, value uint256)
func (_Reward *RewardFilterer) WatchRefundDeposit(opts *bind.WatchOpts, sink chan<- *RewardRefundDeposit) (event.Subscription, error) {

	logs, sub, err := _Reward.contract.WatchLogs(opts, "RefundDeposit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RewardRefundDeposit)
				if err := _Reward.contract.UnpackLog(event, "RefundDeposit", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// RewardSubmitDepositIterator is returned from FilterSubmitDeposit and is used to iterate over the raw logs and unpacked data for SubmitDeposit events raised by the Reward contract.
type RewardSubmitDepositIterator struct {
	Event *RewardSubmitDeposit // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log       // Log channel receiving the found contract events
	sub  cpchain.Subscription // Subscription for errors, completion and termination
	done bool                 // Whether the subscription completed delivering logs
	fail error                // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RewardSubmitDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RewardSubmitDeposit)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RewardSubmitDeposit)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RewardSubmitDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RewardSubmitDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RewardSubmitDeposit represents a SubmitDeposit event raised by the Reward contract.
type RewardSubmitDeposit struct {
	Who   common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterSubmitDeposit is a free log retrieval operation binding the contract event 0x78d81951b78dad84771f88d35b4c93a632e1ed2da8706bbc7d8e465110686830.
//
// Solidity: e SubmitDeposit(who address, value uint256)
func (_Reward *RewardFilterer) FilterSubmitDeposit(opts *bind.FilterOpts) (*RewardSubmitDepositIterator, error) {

	logs, sub, err := _Reward.contract.FilterLogs(opts, "SubmitDeposit")
	if err != nil {
		return nil, err
	}
	return &RewardSubmitDepositIterator{contract: _Reward.contract, event: "SubmitDeposit", logs: logs, sub: sub}, nil
}

// WatchSubmitDeposit is a free log subscription operation binding the contract event 0x78d81951b78dad84771f88d35b4c93a632e1ed2da8706bbc7d8e465110686830.
//
// Solidity: e SubmitDeposit(who address, value uint256)
func (_Reward *RewardFilterer) WatchSubmitDeposit(opts *bind.WatchOpts, sink chan<- *RewardSubmitDeposit) (event.Subscription, error) {

	logs, sub, err := _Reward.contract.WatchLogs(opts, "SubmitDeposit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RewardSubmitDeposit)
				if err := _Reward.contract.UnpackLog(event, "SubmitDeposit", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// RewardWithdrawDepositIterator is returned from FilterWithdrawDeposit and is used to iterate over the raw logs and unpacked data for WithdrawDeposit events raised by the Reward contract.
type RewardWithdrawDepositIterator struct {
	Event *RewardWithdrawDeposit // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log       // Log channel receiving the found contract events
	sub  cpchain.Subscription // Subscription for errors, completion and termination
	done bool                 // Whether the subscription completed delivering logs
	fail error                // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RewardWithdrawDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RewardWithdrawDeposit)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RewardWithdrawDeposit)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RewardWithdrawDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RewardWithdrawDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RewardWithdrawDeposit represents a WithdrawDeposit event raised by the Reward contract.
type RewardWithdrawDeposit struct {
	Who   common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterWithdrawDeposit is a free log retrieval operation binding the contract event 0x195ddc41d185a27fe901831dcad44dd85716c95be78b1d71aa42393697966d40.
//
// Solidity: e WithdrawDeposit(who address, value uint256)
func (_Reward *RewardFilterer) FilterWithdrawDeposit(opts *bind.FilterOpts) (*RewardWithdrawDepositIterator, error) {

	logs, sub, err := _Reward.contract.FilterLogs(opts, "WithdrawDeposit")
	if err != nil {
		return nil, err
	}
	return &RewardWithdrawDepositIterator{contract: _Reward.contract, event: "WithdrawDeposit", logs: logs, sub: sub}, nil
}

// WatchWithdrawDeposit is a free log subscription operation binding the contract event 0x195ddc41d185a27fe901831dcad44dd85716c95be78b1d71aa42393697966d40.
//
// Solidity: e WithdrawDeposit(who address, value uint256)
func (_Reward *RewardFilterer) WatchWithdrawDeposit(opts *bind.WatchOpts, sink chan<- *RewardWithdrawDeposit) (event.Subscription, error) {

	logs, sub, err := _Reward.contract.WatchLogs(opts, "WithdrawDeposit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RewardWithdrawDeposit)
				if err := _Reward.contract.UnpackLog(event, "WithdrawDeposit", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// SafeMathABI is the input ABI used to generate the binding from.
const SafeMathABI = "[]"

// SafeMathBin is the compiled bytecode used for deploying new contracts.
const SafeMathBin = `0x604c602c600b82828239805160001a60731460008114601c57601e565bfe5b5030600052607381538281f30073000000000000000000000000000000000000000030146080604052600080fd00a165627a7a7230582045c9149d657c1500044f71c78b102d8f1ec69984a0af0f9988d3b1efa9d6acba0029`

// DeploySafeMath deploys a new cpchain contract, binding an instance of SafeMath to it.
func DeploySafeMath(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SafeMath, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SafeMathBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// SafeMath is an auto generated Go binding around an cpchain contract.
type SafeMath struct {
	SafeMathCaller     // Read-only binding to the contract
	SafeMathTransactor // Write-only binding to the contract
	SafeMathFilterer   // Log filterer for contract events
}

// SafeMathCaller is an auto generated read-only Go binding around an cpchain contract.
type SafeMathCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathTransactor is an auto generated write-only Go binding around an cpchain contract.
type SafeMathTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathFilterer is an auto generated log filtering Go binding around an cpchain contract events.
type SafeMathFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathSession is an auto generated Go binding around an cpchain contract,
// with pre-set call and transact options.
type SafeMathSession struct {
	Contract     *SafeMath         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SafeMathCallerSession is an auto generated read-only Go binding around an cpchain contract,
// with pre-set call options.
type SafeMathCallerSession struct {
	Contract *SafeMathCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// SafeMathTransactorSession is an auto generated write-only Go binding around an cpchain contract,
// with pre-set transact options.
type SafeMathTransactorSession struct {
	Contract     *SafeMathTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// SafeMathRaw is an auto generated low-level Go binding around an cpchain contract.
type SafeMathRaw struct {
	Contract *SafeMath // Generic contract binding to access the raw methods on
}

// SafeMathCallerRaw is an auto generated low-level read-only Go binding around an cpchain contract.
type SafeMathCallerRaw struct {
	Contract *SafeMathCaller // Generic read-only contract binding to access the raw methods on
}

// SafeMathTransactorRaw is an auto generated low-level write-only Go binding around an cpchain contract.
type SafeMathTransactorRaw struct {
	Contract *SafeMathTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSafeMath creates a new instance of SafeMath, bound to a specific deployed contract.
func NewSafeMath(address common.Address, backend bind.ContractBackend) (*SafeMath, error) {
	contract, err := bindSafeMath(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// NewSafeMathCaller creates a new read-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathCaller(address common.Address, caller bind.ContractCaller) (*SafeMathCaller, error) {
	contract, err := bindSafeMath(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathCaller{contract: contract}, nil
}

// NewSafeMathTransactor creates a new write-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathTransactor(address common.Address, transactor bind.ContractTransactor) (*SafeMathTransactor, error) {
	contract, err := bindSafeMath(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathTransactor{contract: contract}, nil
}

// NewSafeMathFilterer creates a new log filterer instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathFilterer(address common.Address, filterer bind.ContractFilterer) (*SafeMathFilterer, error) {
	contract, err := bindSafeMath(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SafeMathFilterer{contract: contract}, nil
}

// bindSafeMath binds a generic wrapper to an already deployed contract.
func bindSafeMath(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.SafeMathCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transact(opts, method, params...)
}

// SetABI is the input ABI used to generate the binding from.
const SetABI = "[]"

// SetBin is the compiled bytecode used for deploying new contracts.
const SetBin = `0x604c602c600b82828239805160001a60731460008114601c57601e565bfe5b5030600052607381538281f30073000000000000000000000000000000000000000030146080604052600080fd00a165627a7a723058200b6ed157c7502b1d00516e602e8b58d21646becbed02a3371dfe548453a415b90029`

// DeploySet deploys a new cpchain contract, binding an instance of Set to it.
func DeploySet(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Set, error) {
	parsed, err := abi.JSON(strings.NewReader(SetABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SetBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Set{SetCaller: SetCaller{contract: contract}, SetTransactor: SetTransactor{contract: contract}, SetFilterer: SetFilterer{contract: contract}}, nil
}

// Set is an auto generated Go binding around an cpchain contract.
type Set struct {
	SetCaller     // Read-only binding to the contract
	SetTransactor // Write-only binding to the contract
	SetFilterer   // Log filterer for contract events
}

// SetCaller is an auto generated read-only Go binding around an cpchain contract.
type SetCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SetTransactor is an auto generated write-only Go binding around an cpchain contract.
type SetTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SetFilterer is an auto generated log filtering Go binding around an cpchain contract events.
type SetFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SetSession is an auto generated Go binding around an cpchain contract,
// with pre-set call and transact options.
type SetSession struct {
	Contract     *Set              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SetCallerSession is an auto generated read-only Go binding around an cpchain contract,
// with pre-set call options.
type SetCallerSession struct {
	Contract *SetCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// SetTransactorSession is an auto generated write-only Go binding around an cpchain contract,
// with pre-set transact options.
type SetTransactorSession struct {
	Contract     *SetTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SetRaw is an auto generated low-level Go binding around an cpchain contract.
type SetRaw struct {
	Contract *Set // Generic contract binding to access the raw methods on
}

// SetCallerRaw is an auto generated low-level read-only Go binding around an cpchain contract.
type SetCallerRaw struct {
	Contract *SetCaller // Generic read-only contract binding to access the raw methods on
}

// SetTransactorRaw is an auto generated low-level write-only Go binding around an cpchain contract.
type SetTransactorRaw struct {
	Contract *SetTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSet creates a new instance of Set, bound to a specific deployed contract.
func NewSet(address common.Address, backend bind.ContractBackend) (*Set, error) {
	contract, err := bindSet(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Set{SetCaller: SetCaller{contract: contract}, SetTransactor: SetTransactor{contract: contract}, SetFilterer: SetFilterer{contract: contract}}, nil
}

// NewSetCaller creates a new read-only instance of Set, bound to a specific deployed contract.
func NewSetCaller(address common.Address, caller bind.ContractCaller) (*SetCaller, error) {
	contract, err := bindSet(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SetCaller{contract: contract}, nil
}

// NewSetTransactor creates a new write-only instance of Set, bound to a specific deployed contract.
func NewSetTransactor(address common.Address, transactor bind.ContractTransactor) (*SetTransactor, error) {
	contract, err := bindSet(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SetTransactor{contract: contract}, nil
}

// NewSetFilterer creates a new log filterer instance of Set, bound to a specific deployed contract.
func NewSetFilterer(address common.Address, filterer bind.ContractFilterer) (*SetFilterer, error) {
	contract, err := bindSet(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SetFilterer{contract: contract}, nil
}

// bindSet binds a generic wrapper to an already deployed contract.
func bindSet(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SetABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Set *SetRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Set.Contract.SetCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Set *SetRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Set.Contract.SetTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Set *SetRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Set.Contract.SetTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Set *SetCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Set.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Set *SetTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Set.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Set *SetTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Set.Contract.contract.Transact(opts, method, params...)
}
