package tests

import (
	"testing"

	"github.com/bianjieai/avata-sdk-go/models"
)

//tx_hash=0x9666364e2f932c49bbabef252ce0776be6d7373597f100487dbc292149258d3b
// 创建 NFT 类别示例
func TestCreateNFTClass(t *testing.T) {
	params := &models.CreateNFTClassReq{
		Name:        "v2_CreateNFTClass",
		Symbol:      "v2_class",
		Owner:       "0x1A6f8Ed0d40Fcb915e59444AA096241146108D82",
		OperationID: "v2_TestCreateNFTClass",
	}
	result, err := client.NFT.CreateNFTClass(params)
	if err != nil {
		t.Log(err)
		return
	}
	t.Logf("%+v \n", result)
}

// 查询 NFT 类别示例
func TestQueryNFTClasses(t *testing.T) {
	params := &models.QueryNFTClassesReq{
		Name: "v2_CreateNFTClass",
	}
	result, err := client.NFT.QueryNFTClasses(params)
	if err != nil {
		t.Log(err)
		return
	}
	t.Logf("%+v \n", result)
}

// 查询 NFT 类别详情示例
func TestQueryNFTClass(t *testing.T) {
	id := "0x43eddaefa2cb9098b70a5d83cab59363ac920f0e"
	result, err := client.NFT.QueryNFTClass(id)
	if err != nil {
		t.Log(err)
		return
	}
	t.Logf("%+v \n", result)
}

// 转让 NFT 类别示例
func TestTransferNFTClass(t *testing.T) {

	classId := "0x43eddaefa2cb9098b70a5d83cab59363ac920f0e"
	owner := "0x1A6f8Ed0d40Fcb915e59444AA096241146108D82"
	params := &models.TransferNFClassReq{
		Recipient:   "0x9c3d37463fCA8Cd2cec3548a63f1910ec2Cb0BCe",
		OperationID: "v2_TestTransferNFTClass",
	}
	result, err := client.NFT.TransferNFTClass(params, classId, owner)
	if err != nil {
		t.Log(err)
		return
	}
	t.Logf("%+v \n", result)
}

// 发行 NFT 示例
func TestMintNFT(t *testing.T) {
	classId := "0x43eddaefa2cb9098b70a5d83cab59363ac920f0e"
	params := &models.MintNFTReq{
		Uri:         "http://123321test.com",
		Recipient:   "0x9c3d37463fCA8Cd2cec3548a63f1910ec2Cb0BCe",
		OperationID: "v2_TestMintNFT",
	}
	result, err := client.NFT.MintNFT(params, classId)
	if err != nil {
		t.Log(err)
		return
	}
	t.Logf("%+v \n", result)
}

// 转让 NFT 示例
func TestTransferNFT(t *testing.T) {
	classId := "0x43eddaefa2cb9098b70a5d83cab59363ac920f0e"
	owner := "0x9c3d37463fCA8Cd2cec3548a63f1910ec2Cb0BCe"
	nftId := "1"
	params := &models.TransferNFTReq{
		Recipient:   "0x1A6f8Ed0d40Fcb915e59444AA096241146108D82",
		OperationID: "v2_TestTransferNFT",
	}
	result, err := client.NFT.TransferNFT(params, classId, owner, nftId)
	if err != nil {
		t.Log(err)
		return
	}
	t.Logf("%+v \n", result)
}

// 编辑 NFT 示例
func TestEditNFT(t *testing.T) {
	classId := "0x43eddaefa2cb9098b70a5d83cab59363ac920f0e"
	owner := "0x1A6f8Ed0d40Fcb915e59444AA096241146108D82"
	nftId := "1"
	params := &models.EditNFTReq{
		Uri:         "http://123321test.com",
		OperationID: "v2_TestEditNFT",
	}
	result, err := client.NFT.EditNFT(params, classId, owner, nftId)
	if err != nil {
		t.Log(err)
		return
	}
	t.Logf("%+v \n", result)
}

// 销毁 NFT 示例
func TestBurnNFT(t *testing.T) {

	classId := "0x43eddaefa2cb9098b70a5d83cab59363ac920f0e"
	owner := "0x1A6f8Ed0d40Fcb915e59444AA096241146108D82"
	nftId := "1"
	params := &models.BurnNFTReq{
		OperationID: "BurnNFT",
	}
	result, err := client.NFT.BurnNFT(params, classId, owner, nftId)
	if err != nil {
		t.Log(err)
		return
	}
	t.Logf("%+v \n", result)
}

// 查询 NFT 示例
func TestQueryNFTs(t *testing.T) {

	params := &models.QueryNFTsReq{
		ClassID: "0x43eddaefa2cb9098b70a5d83cab59363ac920f0e",
		ID:      "1",
	}
	result, err := client.NFT.QueryNFTs(params)
	if err != nil {
		t.Log(err)
		return
	}
	t.Logf("%+v \n", result)
}

// 查询 NFT 详情示例
func TestQueryNFT(t *testing.T) {
	classId := "0x43eddaefa2cb9098b70a5d83cab59363ac920f0e"
	nftId := "1"
	result, err := client.NFT.QueryNFT(classId, nftId)
	if err != nil {
		t.Log(err)
		return
	}
	t.Logf("%+v \n", result)
}

// 查询 NFT 操作记录示例
func TestQueryNFTHistory(t *testing.T) {

	classId := "0x43edDaEfa2cb9098b70a5d83cAb59363Ac920f0E"
	nftId := "1"
	params := &models.QueryNFTHistoryReq{}
	result, err := client.NFT.QueryNFTHistory(params, classId, nftId)
	if err != nil {
		t.Log(err)
		return
	}
	t.Logf("%+v \n", result)
}
