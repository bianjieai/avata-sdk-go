package tests

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/bianjieai/avata-sdk-go/models"
)

// 创建 NFT 类别示例
func TestCreateNFTClass(t *testing.T) {
	params := &models.CreateNFTClassReq{
		Name:        "TestCreateNFTClass2",
		Owner:       "iaa1tu0gve9se3qgqkadn22d7ar74pal7vqt3yvna9",
		OperationID: "TestCreateNFTClass2",
	}
	result := client.NFT.CreateNFTClass(params)
	if result.Code != 0 {
		t.Log(result.Message)
		return
	}
	if result.Http.Code != http.StatusOK {
		t.Log(result.Error)
		return
	}
	t.Logf("%+v \n", result.Data)

	var txRes models.TxRes
	dataBytes, _ := json.Marshal(result)
	_ = json.Unmarshal(dataBytes, &txRes)

	t.Logf("%+v \n", txRes)
}

// 查询 NFT 类别示例
func TestQueryNFTClasses(t *testing.T) {
	params := &models.QueryNFTClassesReq{
		Name: "TestCreateNFTClass2",
	}
	result := client.NFT.QueryNFTClasses(params)
	if result.Code != 0 {
		t.Log(result.Message)
		return
	}
	if result.Http.Code != http.StatusOK {
		t.Log(result.Error)
		return
	}
	t.Logf("%+v \n", result.Data)

	var nftClassRes models.QueryNFTClassesRes
	dataBytes, _ := json.Marshal(result)
	_ = json.Unmarshal(dataBytes, &nftClassRes)

	t.Logf("%+v \n", nftClassRes)
}

// 查询 NFT 类别详情示例
func TestQueryNFTClass(t *testing.T) {
	
	id := "avatauuj0hj53thkyahiaitfmctsensn"
	result := client.NFT.QueryNFTClass(id)
	if result.Code != 0 {
		t.Log(result.Message)
		return
	}
	if result.Http.Code != http.StatusOK {
		t.Log(result.Error)
		return
	}
	t.Logf("%+v \n", result.Data)

	var nftClassByIdRes models.QueryNFTClassRes
	dataBytes, _ := json.Marshal(result)
	_ = json.Unmarshal(dataBytes, &nftClassByIdRes)

	t.Logf("%+v \n", nftClassByIdRes)
}

// 转让 NFT 类别示例
func TestTransferNFTClass(t *testing.T) {
	
	classId := "avatauuj0hj53thkyahiaitfmctsensn"
	owner := "iaa1tu0gve9se3qgqkadn22d7ar74pal7vqt3yvna9"
	params := &models.TransferNFClassReq{
		Recipient:   "iaa10ldfc2n60ngfpwxnm8qgy5y5hh3vmse6mk4y6v",
		OperationID: "TestTransferNFTClass",
	}
	result := client.NFT.TransferNFTClass(params, classId, owner)
	if result.Code != 0 {
		t.Log(result.Message)
		return
	}
	if result.Http.Code != http.StatusOK {
		t.Log(result.Error)
		return
	}
	t.Logf("%+v \n", result.Data)

	var txRes models.TxRes
	dataBytes, _ := json.Marshal(result)
	_ = json.Unmarshal(dataBytes, &txRes)

	t.Logf("%+v \n", txRes)
}

// 发行 NFT 示例
func TestMintNFT(t *testing.T) {
	
	tag := make(map[string]string)
	tag["nihaoaaa"] = "aaabbbcccddd"
	classId := "avatauuj0hj53thkyahiaitfmctsensn"
	params := &models.MintNFTReq{
		Name:        "TestCreateNFT",
		OperationID: "TestCreateNFT",
		Tag:         tag,
	}
	result := client.NFT.MintNFT(params, classId)
	if result.Code != 0 {
		t.Log(result.Message)
		return
	}
	if result.Http.Code != http.StatusOK {
		t.Log(result.Error)
		return
	}
	t.Logf("%+v \n", result.Data)

	var txRes models.TxRes
	dataBytes, _ := json.Marshal(result)
	_ = json.Unmarshal(dataBytes, &txRes)

	t.Logf("%+v \n", txRes)
}

// 转让 NFT 示例
func TestTransferNFT(t *testing.T) {
	
	classId := "avatauuj0hj53thkyahiaitfmctsensn"
	owner := "iaa10ldfc2n60ngfpwxnm8qgy5y5hh3vmse6mk4y6v"
	nftId := "avatarjqt2kiwlbbced5ieugj2h8cue3"
	params := &models.TransferNFTReq{
		Recipient:   "iaa1tu0gve9se3qgqkadn22d7ar74pal7vqt3yvna9",
		OperationID: "TestTransferNFT",
	}
	result := client.NFT.TransferNFT(params, classId, owner, nftId)
	if result.Code != 0 {
		t.Log(result.Message)
		return
	}
	if result.Http.Code != http.StatusOK {
		t.Log(result.Error)
		return
	}
	t.Logf("%+v \n", result.Data)

	var txRes models.TxRes
	dataBytes, _ := json.Marshal(result)
	_ = json.Unmarshal(dataBytes, &txRes)

	t.Logf("%+v \n", txRes)
}

// 编辑 NFT 示例
func TestEditNFT(t *testing.T) {
	
	classId := "avatauuj0hj53thkyahiaitfmctsensn"
	owner := "iaa1tu0gve9se3qgqkadn22d7ar74pal7vqt3yvna9"
	nftId := "avatarjqt2kiwlbbced5ieugj2h8cue3"
	params := &models.EditNFTReq{
		Name:        "EditNFT",
		OperationID: "EditNFT",
	}
	result := client.NFT.EditNFT(params, classId, owner, nftId)
	if result.Code != 0 {
		t.Log(result.Message)
		return
	}
	if result.Http.Code != http.StatusOK {
		t.Log(result.Error)
		return
	}
	t.Logf("%+v \n", result.Data)

	var txRes models.TxRes
	dataBytes, _ := json.Marshal(result)
	_ = json.Unmarshal(dataBytes, &txRes)

	t.Logf("%+v \n", txRes)
}

// 销毁 NFT 示例
func TestBurnNFT(t *testing.T) {
	
	classId := "avatauuj0hj53thkyahiaitfmctsensn"
	owner := "iaa1tu0gve9se3qgqkadn22d7ar74pal7vqt3yvna9"
	nftId := "avatarjqt2kiwlbbced5ieugj2h8cue3"
	params := &models.BurnNFTReq{
		OperationID: "BurnNFT",
	}
	result := client.NFT.BurnNFT(params, classId, owner, nftId)
	if result.Code != 0 {
		t.Log(result.Message)
		return
	}
	if result.Http.Code != http.StatusOK {
		t.Log(result.Error)
		return
	}
	t.Logf("%+v \n", result.Data)

	var txRes models.TxRes
	dataBytes, _ := json.Marshal(result)
	_ = json.Unmarshal(dataBytes, &txRes)

	t.Logf("%+v \n", txRes)
}

// 批量发行 NFT 示例
func TestBatchMintNFT(t *testing.T) {
	
	classId := "avatauuj0hj53thkyahiaitfmctsensn"
	var recipients []models.Recipients
	recipients = append(recipients, models.Recipients{Amount: 1, Recipient: "iaa1tu0gve9se3qgqkadn22d7ar74pal7vqt3yvna9"})
	params := &models.BatchMintNFTReq{
		Name:        "TestBatchCreateNFT1",
		Recipients:  recipients,
		OperationID: "TestBatchCreateNFT2",
	}
	result := client.NFT.BatchMintNFT(params, classId)
	if result.Code != 0 {
		t.Log(result.Message)
		return
	}
	if result.Http.Code != http.StatusOK {
		t.Log(result.Error)
		return
	}
	t.Logf("%+v \n", result.Data)

	var txRes models.TxRes
	dataBytes, _ := json.Marshal(result)
	_ = json.Unmarshal(dataBytes, &txRes)

	t.Logf("%+v \n", txRes)
}

// 批量转让 NFT 示例
func TestBatchTransferNFT(t *testing.T) {
	
	owner := "iaa1tu0gve9se3qgqkadn22d7ar74pal7vqt3yvna9"
	var nfts []models.BatchTransferNFTs
	nfts = append(nfts, models.BatchTransferNFTs{
		ClassID: "avatauuj0hj53thkyahiaitfmctsensn",
		NFTID:   "avataj5h7algcaibxiz5ipbi5o97kfqs",
	}, models.BatchTransferNFTs{
		ClassID: "avatauuj0hj53thkyahiaitfmctsensn",
		NFTID:   "avatarxamgw6sqwvclfqfnvkgkwkp6zf",
	})
	var data []models.BatchTransferNFTData
	data = append(data, models.BatchTransferNFTData{
		NFTs:      nil,
		Recipient: "iaa153uyr6ghtumt3lrtdwndplk4ggal9r6gm6953g",
	})
	params := &models.BatchTransferNFTReq{
		Data:        data,
		OperationID: "TestBatchTransferNFT",
	}
	result := client.NFT.BatchTransferNFT(params, owner)
	if result.Code != 0 {
		t.Log(result.Message)
		return
	}
	if result.Http.Code != http.StatusOK {
		t.Log(result.Error)
		return
	}
	t.Logf("%+v \n", result.Data)

	var txRes models.TxRes
	dataBytes, _ := json.Marshal(result)
	_ = json.Unmarshal(dataBytes, &txRes)

	t.Logf("%+v \n", txRes)
}

// 批量编辑 NFT 示例
func TestBatchEditNFT(t *testing.T) {
	
	owner := "iaa153uyr6ghtumt3lrtdwndplk4ggal9r6gm6953g"

	var nfts []models.BatchEditNfts
	nfts = append(nfts, models.BatchEditNfts{
		ClassID: "avatauuj0hj53thkyahiaitfmctsensn",
		NFTID:   "avataj5h7algcaibxiz5ipbi5o97kfqs",
		Name:    "TestBatchEditNFT",
	}, models.BatchEditNfts{
		ClassID: "avatauuj0hj53thkyahiaitfmctsensn",
		NFTID:   "avatarxamgw6sqwvclfqfnvkgkwkp6zf",
		Name:    "TestBatchEditNFT",
	})
	params := &models.BatchEditNFTReq{
		NFTs:        nfts,
		OperationID: OperationID,
	}
	result := client.NFT.BatchEditNFT(params, owner)
	if result.Code != 0 {
		t.Log(result.Message)
		return
	}
	if result.Http.Code != http.StatusOK {
		t.Log(result.Error)
		return
	}
	t.Logf("%+v \n", result.Data)

	var txRes models.TxRes
	dataBytes, _ := json.Marshal(result)
	_ = json.Unmarshal(dataBytes, &txRes)

	t.Logf("%+v \n", txRes)
}

// 批量销毁 NFT 示例
func TestBatchBurnNFT(t *testing.T) {
	owner := "iaa153uyr6ghtumt3lrtdwndplk4ggal9r6gm6953g"
	var nfts []models.NFTs
	nfts = append(nfts, models.NFTs{
		ClassID: "avatauuj0hj53thkyahiaitfmctsensn",
		NFTID:   "avataj5h7algcaibxiz5ipbi5o97kfqs",
	}, models.NFTs{
		ClassID: "avatauuj0hj53thkyahiaitfmctsensn",
		NFTID:   "avatarxamgw6sqwvclfqfnvkgkwkp6zf",
	})

	params := &models.BatchBurnNFTReq{
		NFTs:        nfts,
		OperationID: OperationID,
	}
	result := client.NFT.BatchBurnNFT(params, owner)
	if result.Code != 0 {
		t.Log(result.Message)
		return
	}
	if result.Http.Code != http.StatusOK {
		t.Log(result.Error)
		return
	}
	t.Logf("%+v \n", result.Data)

	var txRes models.TxRes
	dataBytes, _ := json.Marshal(result)
	_ = json.Unmarshal(dataBytes, &txRes)

	t.Logf("%+v \n", txRes)
}

// 查询 NFT 示例
func TestQueryNFTs(t *testing.T) {
	
	params := &models.QueryNFTsReq{}
	result := client.NFT.QueryNFTs(params)
	if result.Code != 0 {
		t.Log(result.Message)
		return
	}
	if result.Http.Code != http.StatusOK {
		t.Log(result.Error)
		return
	}
	t.Logf("%+v \n", result.Data)

	var nftRes models.QueryNFTsRes
	dataBytes, _ := json.Marshal(result)
	_ = json.Unmarshal(dataBytes, &nftRes)

	t.Logf("%+v \n", nftRes)
}

// 查询 NFT 详情示例
func TestQueryNFT(t *testing.T) {
	
	classId := "avatauuj0hj53thkyahiaitfmctsensn"
	nftId := "avataj5h7algcaibxiz5ipbi5o97kfqs"
	result := client.NFT.QueryNFT(classId, nftId)
	if result.Code != 0 {
		t.Log(result.Message)
		return
	}
	if result.Http.Code != http.StatusOK {
		t.Log(result.Error)
		return
	}
	t.Logf("%+v \n", result.Data)

	var nftByIdRes models.QueryNFTRes
	dataBytes, _ := json.Marshal(result)
	_ = json.Unmarshal(dataBytes, &nftByIdRes)

	t.Logf("%+v \n", nftByIdRes)
}

// 查询 NFT 操作记录示例
func TestQueryNFTHistory(t *testing.T) {
	
	classId := "avatauuj0hj53thkyahiaitfmctsensn"
	nftId := "avataj5h7algcaibxiz5ipbi5o97kfqs"
	params := &models.QueryNFTHistoryReq{}
	result := client.NFT.QueryNFTHistory(params, classId, nftId)
	if result.Code != 0 {
		t.Log(result.Message)
		return
	}
	if result.Http.Code != http.StatusOK {
		t.Log(result.Error)
		return
	}
	t.Logf("%+v \n", result.Data)

	var nftHistoryRes models.QueryNFTHistoryRes
	dataBytes, _ := json.Marshal(result)
	_ = json.Unmarshal(dataBytes, &nftHistoryRes)

	t.Logf("%+v \n", nftHistoryRes)
}
