package tests

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/bianjieai/avata-sdk-go/models"
)

// 创建 NFT 类别示例
func TestCreateNFTClass(t *testing.T) {
	client := GetClient()

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
func TestQueryNFTClass(t *testing.T) {
	client := GetClient()

	params := &models.QueryNFTClassReq{
		Name: "TestCreateNFTClass2",
	}
	result := client.NFT.QueryNFTClass(params)
	if result.Code != 0 {
		t.Log(result.Message)
		return
	}
	if result.Http.Code != http.StatusOK {
		t.Log(result.Error)
		return
	}
	t.Logf("%+v \n", result.Data)

	var nftClassRes models.QueryNFTClassRes
	dataBytes, _ := json.Marshal(result)
	_ = json.Unmarshal(dataBytes, &nftClassRes)

	t.Logf("%+v \n", nftClassRes)
}

// 查询 NFT 类别详情示例
func TestQueryNFTClassById(t *testing.T) {
	client := GetClient()
	id := "avatauuj0hj53thkyahiaitfmctsensn"
	result := client.NFT.QueryNFTClassById(id)
	if result.Code != 0 {
		t.Log(result.Message)
		return
	}
	if result.Http.Code != http.StatusOK {
		t.Log(result.Error)
		return
	}
	t.Logf("%+v \n", result.Data)

	var nftClassByIdRes models.QueryNFTClassByIdRes
	dataBytes, _ := json.Marshal(result)
	_ = json.Unmarshal(dataBytes, &nftClassByIdRes)

	t.Logf("%+v \n", nftClassByIdRes)
}

// 转让 NFT 类别示例
func TestTransfersNFClass(t *testing.T) {
	client := GetClient()
	classId := "avatauuj0hj53thkyahiaitfmctsensn"
	owner := "iaa1tu0gve9se3qgqkadn22d7ar74pal7vqt3yvna9"
	params := &models.TransfersNFClassReq{
		Recipient:   "iaa10ldfc2n60ngfpwxnm8qgy5y5hh3vmse6mk4y6v",
		OperationID: "TestTransfersNFClass",
	}
	result := client.NFT.TransfersNFClass(params, classId, owner)
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
func TestCreateNFT(t *testing.T) {
	client := GetClient()
	tag := make(map[string]string)
	tag["nihaoaaa"] = "aaabbbcccddd"
	classId := "avatauuj0hj53thkyahiaitfmctsensn"
	params := &models.CreateNFTReq{
		Name:        "TestCreateNFT",
		OperationID: "TestCreateNFT",
		Tag:         tag,
	}
	result := client.NFT.CreateNFT(params, classId)
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
	client := GetClient()
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
	client := GetClient()
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
func TestDeleteNFT(t *testing.T) {
	client := GetClient()
	classId := "avatauuj0hj53thkyahiaitfmctsensn"
	owner := "iaa1tu0gve9se3qgqkadn22d7ar74pal7vqt3yvna9"
	nftId := "avatarjqt2kiwlbbced5ieugj2h8cue3"
	params := &models.DeleteNFTReq{
		OperationID: "DeleteNFT",
	}
	result := client.NFT.DeleteNFT(params, classId, owner, nftId)
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
func TestBatchCreateNFT(t *testing.T) {
	client := GetClient()
	classId := "avatauuj0hj53thkyahiaitfmctsensn"
	params := &models.BatchCreateNFTReq{
		Name: "TestBatchCreateNFT1",
		Recipients: []struct {
			Amount    int    "json:\"amount\""
			Recipient string "json:\"recipient\""
		}{
			{Amount: 1, Recipient: "iaa1tu0gve9se3qgqkadn22d7ar74pal7vqt3yvna9"},
			{Amount: 1, Recipient: "iaa10ldfc2n60ngfpwxnm8qgy5y5hh3vmse6mk4y6v"},
		},
		OperationID: "TestBatchCreateNFT2",
	}
	result := client.NFT.BatchCreateNFT(params, classId)
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
	client := GetClient()
	owner := "iaa1tu0gve9se3qgqkadn22d7ar74pal7vqt3yvna9"
	params := &models.BatchTransferNFTReq{
		Data: []struct {
			NFTs []struct {
				ClassID string "json:\"class_id\""
				NFTID   string "json:\"nft_id\""
			} "json:\"nfts\""
			Recipient string "json:\"recipient\""
		}{
			{Recipient: "iaa153uyr6ghtumt3lrtdwndplk4ggal9r6gm6953g", NFTs: []struct {
				ClassID string "json:\"class_id\""
				NFTID   string "json:\"nft_id\""
			}{
				{ClassID: "avatauuj0hj53thkyahiaitfmctsensn", NFTID: "avataj5h7algcaibxiz5ipbi5o97kfqs"},
				{ClassID: "avatauuj0hj53thkyahiaitfmctsensn", NFTID: "avatarxamgw6sqwvclfqfnvkgkwkp6zf"},
			}},
		},
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
	client := GetClient()
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
func TestBatchDeleteNFT(t *testing.T) {
	client := GetClient()

	owner := "iaa153uyr6ghtumt3lrtdwndplk4ggal9r6gm6953g"
	var nfts []models.NFTs
	nfts = append(nfts, models.NFTs{
		ClassID: "avatauuj0hj53thkyahiaitfmctsensn",
		NFTID:   "avataj5h7algcaibxiz5ipbi5o97kfqs",
	}, models.NFTs{
		ClassID: "avatauuj0hj53thkyahiaitfmctsensn",
		NFTID:   "avatarxamgw6sqwvclfqfnvkgkwkp6zf",
	})

	params := &models.BatchDeleteNFTReq{
		NFTs:        nfts,
		OperationID: OperationID,
	}
	result := client.NFT.BatchDeleteNFT(params, owner)
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
func TestQueryNFT(t *testing.T) {
	client := GetClient()
	params := &models.QueryNFTReq{}
	result := client.NFT.QueryNFT(params)
	if result.Code != 0 {
		t.Log(result.Message)
		return
	}
	if result.Http.Code != http.StatusOK {
		t.Log(result.Error)
		return
	}
	t.Logf("%+v \n", result.Data)

	var nftRes models.QueryNFTRes
	dataBytes, _ := json.Marshal(result)
	_ = json.Unmarshal(dataBytes, &nftRes)

	t.Logf("%+v \n", nftRes)
}

// 查询 NFT 详情示例
func TestQueryNFTById(t *testing.T) {
	client := GetClient()
	classId := "avatauuj0hj53thkyahiaitfmctsensn"
	nftId := "avataj5h7algcaibxiz5ipbi5o97kfqs"
	result := client.NFT.QueryNFTById(classId, nftId)
	if result.Code != 0 {
		t.Log(result.Message)
		return
	}
	if result.Http.Code != http.StatusOK {
		t.Log(result.Error)
		return
	}
	t.Logf("%+v \n", result.Data)

	var nftByIdRes models.QueryNFTByIdRes
	dataBytes, _ := json.Marshal(result)
	_ = json.Unmarshal(dataBytes, &nftByIdRes)

	t.Logf("%+v \n", nftByIdRes)
}

// 查询 NFT 操作记录示例
func TestQueryNFTHistory(t *testing.T) {
	client := GetClient()
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
