package tests

import (
	"testing"

	"github.com/bianjieai/avata-sdk-go/models"
)

// tx_hash=0x9666364e2f932c49bbabef252ce0776be6d7373597f100487dbc292149258d3b
// 以 EVM 方式创建 NFT 类别示例
func TestCreateNFTClass(t *testing.T) {
	params := &models.CreateNFTClassReq{
		Name:        "v3_CreateNFTClass1",
		Symbol:      "v3_class",
		Owner:       "0xfb74240135ebCf3bB56F1CDe680FB85bd36E71F9",
		OperationID: "v3_TestCreateNFTClass_sh1q",
	}
	result, err := client.NFT.CreateNFTClass(params)
	if err != nil {
		t.Log(err)
		return
	}
	t.Logf("%+v \n", result)
}

// 以原生方式创建 NFT 类别示例
func TestCreateNativeNFTClass(t *testing.T) {
	params := &models.CreateNativeNFTClassReq{
		Name:        "nativeNFTClasstest",
		Symbol:      "v3_class",
		Owner:       "iaa1ld6zgqf4a08nhdt0rn0xsract0fkuu0equagvz",
		OperationID: "v3_TestCreateNFTClass_go_sdk1",
	}
	result, err := client.NFT.CreateNativeNFTClass(params)
	if err != nil {
		t.Log(err)
		return
	}
	t.Logf("%+v \n", result)
}

// 以 EVM 方式查询 NFT 类别示例
func TestQueryNFTClasses(t *testing.T) {
	params := &models.QueryNFTClassesReq{}
	result, err := client.NFT.QueryNFTClasses(params)
	if err != nil {
		t.Log(err)
		return
	}
	t.Logf("%+v \n", result)
}

// 以原生方式查询 NFT 类别示例
func TestQueryNativeNFTClasses(t *testing.T) {
	params := &models.QueryNativeNFTClassesReq{}
	result, err := client.NFT.QueryNativeNFTClasses(params)
	if err != nil {
		t.Log(err)
		return
	}
	t.Logf("%+v \n", result)
}

// 以 EVM 方式查询 NFT 类别详情示例
func TestQueryNFTClass(t *testing.T) {
	id := "0x18c82844dA374B741Dfb433cce01241C7db49a98"
	result, err := client.NFT.QueryNFTClass(id)
	if err != nil {
		t.Log(err)
		return
	}
	t.Logf("%+v \n", result)
}

// 以原生方式查询 NFT 类别详情示例
func TestQueryNativeNFTClass(t *testing.T) {
	id := "lmhtestnftclass01"
	result, err := client.NFT.QueryNativeNFTClass(id)
	if err != nil {
		t.Log(err)
		return
	}
	t.Logf("%+v \n", result)
}

// 以 EVM 方式转让 NFT 类别示例
func TestTransferNFTClass(t *testing.T) {

	classId := "0xb195397d69A85edD12552182f360DC83d86Ee5d6"
	owner := "0xfb74240135ebCf3bB56F1CDe680FB85bd36E71F9"
	params := &models.TransferNFClassReq{
		OperationId: "testshihengtransfernftclass",
		Recipient:   "0x362C87e50Ef0d60AA9f827DCa0136FF6E7927398",
	}
	result, err := client.NFT.TransferNFTClass(params, classId, owner)
	if err != nil {
		t.Log(err)
		return
	}
	t.Logf("%+v \n", result)
}

// 以原生方式转让 NFT 类别示例
func TestTransferNativeNFTClass(t *testing.T) {

	classId := "lmhtestnftclass01"
	owner := "iaa123sl8s8wpqzxpuwvz5xj5ltc8p5apksdrfr5wc"
	params := &models.TransferNativeNFClassReq{
		Recipient:   "iaa1jjmwg5ah27aynuwt2phwa8sfvzh4lvvlelddxm",
		OperationId: OperationID,
	}
	result, err := client.NFT.TransferNativeNFTClass(params, classId, owner)
	if err != nil {
		t.Log(err)
		return
	}
	t.Logf("%+v \n", result)
}

// 以 EVM 方式发行 NFT 示例
func TestMintNFT(t *testing.T) {
	classId := "0xb195397d69A85edD12552182f360DC83d86Ee5d6"
	params := &models.MintNFTReq{
		Uri:         "http://123321shitest.com",
		Recipient:   "0x362C87e50Ef0d60AA9f827DCa0136FF6E7927398",
		OperationID: "v3_TestMintNFT11",
	}
	result, err := client.NFT.MintNFT(params, classId)
	if err != nil {
		t.Log(err)
		return
	}
	t.Logf("%+v \n", result)
}

// 以原生方式发行 NFT 示例
func TestMintNativeNFT(t *testing.T) {
	classId := "lmhtestnftclass01"
	params := &models.MintNativeNFTReq{
		OperationID: "v3_TestMintwNativseN14112311231",
		Recipient:   "iaa1sj0dsuntd464wgdsa6kjafd6xu0eu0pzzj8gkx",
	}
	result, err := client.NFT.MintNativeNFT(params, classId)
	if err != nil {
		t.Log(err)
		return
	}
	t.Logf("%+v \n", result)
}

// 以 EVM 方式转让 NFT 示例
func TestTransferNFT(t *testing.T) {
	classId := "0xb195397d69A85edD12552182f360DC83d86Ee5d6"
	owner := "0x362C87e50Ef0d60AA9f827DCa0136FF6E7927398"
	nftId := "2"
	params := &models.TransferNFTReq{
		Recipient:   "0xfb74240135ebCf3bB56F1CDe680FB85bd36E71F9",
		OperationID: "v3_TestTransferNFT11",
	}
	result, err := client.NFT.TransferNFT(params, classId, owner, nftId)
	if err != nil {
		t.Log(err)
		return
	}
	t.Logf("%+v \n", result)
}

// 以原生方式转让 NFT 示例
func TestTransferNativeNFT(t *testing.T) {
	classId := "lmhtestnftclass01"
	owner := "iaa1jjmwg5ah27aynuwt2phwa8sfvzh4lvvlelddxm"
	nftId := "1"
	params := &models.TransferNativeNFTReq{
		Recipient:   "iaa1jn94j6mxgmeg3eyqe6e55gnug6rqd9l5hn523p",
		OperationID: "v3_TestTransferNFT",
	}
	result, err := client.NFT.TransferNativeNFT(params, classId, owner, nftId)
	if err != nil {
		t.Log(err)
		return
	}
	t.Logf("%+v \n", result)
}

// 以 EVM 方式编辑 NFT 示例
func TestEditNFT(t *testing.T) {
	classId := "0xb195397d69A85edD12552182f360DC83d86Ee5d6"
	owner := "0xfb74240135ebCf3bB56F1CDe680FB85bd36E71F9"
	nftId := "1"
	params := &models.EditNFTReq{
		Uri:         "http://1233shitest.com",
		OperationID: "v3_TestEditNFT1111shiheng1",
	}
	result, err := client.NFT.EditNFT(params, classId, owner, nftId)
	if err != nil {
		t.Log(err)
		return
	}
	t.Logf("%+v \n", result)
}

// 以原生方式编辑 NFT 示例
func TestEditNativeNFT(t *testing.T) {
	classId := "lmhtestnftclass01"
	owner := "iaa1jn94j6mxgmeg3eyqe6e55gnug6rqd9l5hn523p" //iaa1jn94j6mxgmeg3eyqe6e55gnug6rqd9l5hn523p  --owner不对的情况
	nftId := "avataokjyn5fgkjebhat94onzuvnzbon"
	params := &models.EditNativeNFTReq{
		Uri:         "http://123321test.com",
		OperationID: "op828222223a1122",
	}
	result, err := client.NFT.EditNativeNFT(params, classId, owner, nftId)
	if err != nil {
		t.Log(err)
		return
	}
	t.Logf("%+v \n", result)
}

// 以 EVM 方式销毁 NFT 示例
func TestBurnNFT(t *testing.T) {

	classId := "lmhtestnftclass01"
	owner := "iaa1sj0dsuntd464wgdsa6kjafd6xu0eu0pzzj8gkx"
	nftId := "avatabczlynlejxabqxv9muk2i6objos"
	params := &models.BurnNFTReq{
		OperationID: "v3_Bu2rnNFT211231",
	}
	result, err := client.NFT.BurnNFT(params, classId, owner, nftId)
	if err != nil {
		t.Log(err)
		return
	}
	t.Logf("%+v \n", result)
}

// 以原生方式销毁 NFT 示例
func TestBurnNativeNFT(t *testing.T) {

	classId := "lmhtestnftclass01"
	owner := "iaa1sj0dsuntd464wgdsa6kjafd6xu0eu0pzzj8gkx"
	nftId := "avataqeerstm9hcwfiutxr6gs1ktizcg"
	params := &models.BurnNativeNFTReq{
		OperationID: "BurnNFT12sssre12112",
	}
	result, err := client.NFT.BurnNativeNFT(params, classId, owner, nftId)
	if err != nil {
		t.Log(err)
		return
	}
	t.Logf("%+v \n", result)
}

// 以 EVM 方式查询 NFT 示例
func TestQueryNFTs(t *testing.T) {

	params := &models.QueryNFTsReq{
		ClassID: "0xb195397d69A85edD12552182f360DC83d86Ee5d6",
	}
	result, err := client.NFT.QueryNFTs(params)
	if err != nil {
		t.Log(err)
		return
	}
	t.Logf("%+v \n", result)
}

// 以原生方式查询 NFT 示例
func TestQueryNativeNFTs(t *testing.T) {

	params := &models.QueryNativeNFTsReq{
		ClassID: "lmhtestnftclass01",
	}
	result, err := client.NFT.QueryNativeNFTs(params)
	if err != nil {
		t.Log(err)
		return
	}
	t.Logf("%+v \n", result)
}

// 以 EVM 方式查询 NFT 详情示例
func TestQueryNFT(t *testing.T) {
	classId := "0xb195397d69A85edD12552182f360DC83d86Ee5d6"
	nftId := "1"
	result, err := client.NFT.QueryNFT(classId, nftId)
	if err != nil {
		t.Log(err)
		return
	}
	t.Logf("%+v \n", result)
}

// 以原生方式查询 NFT 详情示例
func TestQueryNativeNFT(t *testing.T) {
	classId := "lmhtestnftclass01"
	nftId := "avatachfkcwwp1nfnitzy9gqhui2s7pk"
	result, err := client.NFT.QueryNativeNFT(classId, nftId)
	if err != nil {
		t.Log(err)
		return
	}
	t.Logf("%+v \n", result)
}

// 以 EVM 方式查询 NFT 操作记录示例
func TestQueryNFTHistory(t *testing.T) {

	classId := "0xb195397d69A85edD12552182f360DC83d86Ee5d6"
	nftId := "1"
	params := &models.QueryNFTHistoryReq{}
	result, err := client.NFT.QueryNFTHistory(params, classId, nftId)
	if err != nil {
		t.Log(err)
		return
	}
	t.Logf("%+v \n", result)
}

// 以原生方式查询 NFT 操作记录示例
func TestQueryNativeNFTHistory(t *testing.T) {

	classId := "lmhtestnftclass01"
	nftId := "avatachfkcwwp1nfnitzy9gqhui2s7pk"
	params := &models.QueryNativeNFTHistoryReq{
		CountTotal: "1",
	}
	result, err := client.NFT.QueryNativeNFTHistory(params, classId, nftId)
	if err != nil {
		t.Log(err)
		return
	}
	t.Logf("%+v \n", result)
}
