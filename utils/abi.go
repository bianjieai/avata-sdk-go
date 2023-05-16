package utils

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

//调用合约 的 abi编码
func ABICoding(node string, id int64, value string) (string, error) {
	// 获取 abi
	abi, err := GetResolveAbi()
	if err != nil {
		return "", err
	}
	byte32Node, err := GetBytes32Node(node)
	if err != nil {
		return "", err
	}
	// 生成调用合约 hex data
	packData, err := abi.Pack("setABI", []interface{}{
		byte32Node,
		big.NewInt(id),
		[]byte(value),
	}...)
	if err != nil {
		return "", err
	}
	return hexutil.Encode(packData), nil
}

//解析合约的 abi编码
func QueryABICoding(node string, id int64) (string, error) {
	// 获取 abi
	abi, err := GetResolveAbi()
	if err != nil {
		return "", err
	}
	byte32Node, err := GetBytes32Node(node)
	if err != nil {
		return "", err
	}
	//查询合约 hex data
	packData, err := abi.Pack("ABI", []interface{}{
		byte32Node,
		big.NewInt(id),
	}...)
	if err != nil {
		return "", err
	}
	return hexutil.Encode(packData), nil
}

//abi解析
func ABIResolver(resultHex string) ([]interface{}, error) {
	abi, err := GetResolveAbi()
	if err != nil {
		return nil, err
	}
	resultBytes, err := hexutil.Decode(resultHex)
	if err != nil {
		return nil, err
	}
	resultData, err := abi.Methods["ABI"].Outputs.Unpack(resultBytes)
	if err != nil {
		return nil, err
	}
	return resultData, nil
}

func HexStrToByte32(str string) ([32]byte, error) {
	var bytes32 [32]byte
	hexBytes, err := hexutil.Decode(str)
	if err != nil {
		return bytes32, err
	}
	copy(bytes32[:], hexBytes)
	return bytes32, nil
}

func GetBytes32Node(node string) ([32]byte, error) {
	byte32Node, err := HexStrToByte32(node) // 域名 key
	if err != nil {
		return [32]byte{}, err
	}
	return byte32Node, nil
}

func GetResolveAbi() (abi.ABI, error) {
	resolveAbi, err := abi.JSON(strings.NewReader(UtilsMetaData.ABI)) // 初始化 abi
	if err != nil {
		return abi.ABI{}, err
	}
	return resolveAbi, nil
}
