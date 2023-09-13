package utils

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

/**
* abi：调用合约生成data
* node：域名
* value：自定义数据
* method：abi:setABI 、 addr:setAddr 、 text:setText
 */
func ABICoding(node, value, method string, id int64) (string, error) {
	// 获取 abi
	abi, err := getResolveAbi()
	if err != nil {
		return "", err
	}
	byte32Node, err := getBytes32Node(node)
	if err != nil {
		return "", err
	}
	// 生成调用合约 hex data
	packData, err := abi.Pack(method, []interface{}{
		byte32Node,
		big.NewInt(id),
		[]byte(value),
	}...)
	if err != nil {
		return "", err
	}
	return hexutil.Encode(packData), nil
}

/**
* abi：查询合约生成data
* node：域名
* method：abi:ABI 、 addr:addr0 、 text:text
 */
func QueryABICoding(node, method string, id int64) (string, error) {
	// 获取 abi
	abi, err := getResolveAbi()
	if err != nil {
		return "", err
	}
	byte32Node, err := getBytes32Node(node)
	if err != nil {
		return "", err
	}
	//查询合约 hex data
	packData, err := abi.Pack(method, []interface{}{
		byte32Node,
		big.NewInt(id),
	}...)
	if err != nil {
		return "", err
	}
	return hexutil.Encode(packData), nil
}

/**
* abi：解析
* resultHex：查询合约的结果
* method：abi:ABI 、 addr:addr0 、 text:text
 */
func ABIResolver(resultHex, method string) ([]interface{}, error) {
	abi, err := getResolveAbi()
	if err != nil {
		return nil, err
	}
	resultBytes, err := hexutil.Decode(resultHex)
	if err != nil {
		return nil, err
	}
	resultData, err := abi.Methods[method].Outputs.Unpack(resultBytes)
	if err != nil {
		return nil, err
	}
	return resultData, nil
}

func hexStrToByte32(str string) ([32]byte, error) {
	var bytes32 [32]byte
	hexBytes, err := hexutil.Decode(str)
	if err != nil {
		return bytes32, err
	}
	copy(bytes32[:], hexBytes)
	return bytes32, nil
}

func getBytes32Node(node string) ([32]byte, error) {
	byte32Node, err := hexStrToByte32(node) // 域名 key
	if err != nil {
		return [32]byte{}, err
	}
	return byte32Node, nil
}

func getResolveAbi() (abi.ABI, error) {
	resolveAbi, err := abi.JSON(strings.NewReader(UtilsMetaData.ABI)) // 初始化 abi
	if err != nil {
		return abi.ABI{}, err
	}
	return resolveAbi, nil
}
