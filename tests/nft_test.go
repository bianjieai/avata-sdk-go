package tests

import (
	"avata-sdk-go/models"
	"testing"
)

//创建nft类别
func TestCreateNFTClass(t *testing.T) {
	client := GetClient()
	params := &models.CreateNFTClassReq{
		Name:        "TestCreateNFTClass2",
		Owner:       "iaa1tu0gve9se3qgqkadn22d7ar74pal7vqt3yvna9",
		OperationID: "TestCreateNFTClass2",
	}
	result := client.NFT.CreateNFTClass(params)
	t.Logf("%+v \n", result)
}

//查询nft类别
func TestQueryNFTClass(t *testing.T) {
	client := GetClient()
	params := &models.QueryNFTClassReq{
		Name: "TestCreateNFTClass2",
	}
	result := client.NFT.QueryNFTClass(params)
	t.Logf("%+v \n", result)
}

//查询nft类别详情
func TestQueryNFTClassById(t *testing.T) {
	client := GetClient()
	id := "avatauuj0hj53thkyahiaitfmctsensn"
	result := client.NFT.QueryNFTClassById(id)
	t.Logf("%+v \n", result)
}

//转让nft类别
func TestTransfersNFClass(t *testing.T) {
	client := GetClient()
	class_id := "avatauuj0hj53thkyahiaitfmctsensn"
	owner := "iaa1tu0gve9se3qgqkadn22d7ar74pal7vqt3yvna9"
	params := &models.TransfersNFClassReq{
		Recipient:   "iaa10ldfc2n60ngfpwxnm8qgy5y5hh3vmse6mk4y6v",
		OperationID: "TestTransfersNFClass",
	}
	result := client.NFT.TransfersNFClass(params, class_id, owner)
	t.Logf("%+v \n", result)
}

//发行nft
func TestCreateNFT(t *testing.T) {
	client := GetClient()
	tag := make(map[string]string)
	tag["nihaoaaa"] = "aaabbbcccddd"
	class_id := "avatauuj0hj53thkyahiaitfmctsensn"
	params := &models.CreateNFTReq{
		Name:        "TestCreateNFT",
		OperationID: "TestCreateNFT",
		Tag:         tag,
	}
	result := client.NFT.CreateNFT(params, class_id)
	t.Logf("%+v \n", result)
}

//转让nft
func TestTransfersNFT(t *testing.T) {
	client := GetClient()
	class_id := "avatauuj0hj53thkyahiaitfmctsensn"
	owner := "iaa10ldfc2n60ngfpwxnm8qgy5y5hh3vmse6mk4y6v"
	nft_id := "avatarjqt2kiwlbbced5ieugj2h8cue3"
	params := &models.TransfersNFTReq{
		Recipient:   "iaa1tu0gve9se3qgqkadn22d7ar74pal7vqt3yvna9",
		OperationID: "TestTransfersNFT",
	}
	result := client.NFT.TransfersNFT(params, class_id, owner, nft_id)
	t.Logf("%+v \n", result)
}

//编辑nft
func TestEditorNFT(t *testing.T) {
	client := GetClient()
	class_id := "avatauuj0hj53thkyahiaitfmctsensn"
	owner := "iaa1tu0gve9se3qgqkadn22d7ar74pal7vqt3yvna9"
	nft_id := "avatarjqt2kiwlbbced5ieugj2h8cue3"
	params := &models.EditorNFTReq{
		Name:        "EditorNFT",
		OperationID: "EditorNFT",
	}
	result := client.NFT.EditorNFT(params, class_id, owner, nft_id)
	t.Logf("%+v \n", result)
}

//销毁nft
func TestDeleteNFT(t *testing.T) {
	client := GetClient()
	class_id := "avatauuj0hj53thkyahiaitfmctsensn"
	owner := "iaa1tu0gve9se3qgqkadn22d7ar74pal7vqt3yvna9"
	nft_id := "avatarjqt2kiwlbbced5ieugj2h8cue3"
	params := &models.DeleteNFTReq{
		OperationID: "DeleteNFT",
	}
	result := client.NFT.DeleteNFT(params, class_id, owner, nft_id)
	t.Logf("%+v \n", result)
}

//批量发行nft
func TestCreateNFTBatch(t *testing.T) {
	client := GetClient()
	class_id := "avatauuj0hj53thkyahiaitfmctsensn"
	params := &models.CreateNFTBatchReq{
		Name: "TestCreateNFTBatch1",
		Recipients: []struct {
			Amount    int    "json:\"amount\""
			Recipient string "json:\"recipient\""
		}{
			{Amount: 1, Recipient: "iaa1tu0gve9se3qgqkadn22d7ar74pal7vqt3yvna9"},
			{Amount: 1, Recipient: "iaa10ldfc2n60ngfpwxnm8qgy5y5hh3vmse6mk4y6v"},
		},
		OperationID: "TestCreateNFTBatch2",
	}
	result := client.NFT.CreateNFTBatch(params, class_id)
	t.Logf("%+v \n", result)
}

//批量转让nft
func TestTransfersNFTBatch(t *testing.T) {
	client := GetClient()
	owner := "iaa1tu0gve9se3qgqkadn22d7ar74pal7vqt3yvna9"
	params := &models.TransfersNFTBatchReq{
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
		OperationID: "TestTransfersNFTBatch",
	}
	result := client.NFT.TransfersNFTBatch(params, owner)
	t.Logf("%+v \n", result)
}

//批量编辑nft
func TestEditorNFTBatch(t *testing.T) {
	client := GetClient()
	owner := "iaa153uyr6ghtumt3lrtdwndplk4ggal9r6gm6953g"
	params := &models.EditorNFTBatchReq{
		NFTs: []struct {
			ClassID string "json:\"class_id\""
			NFTID   string "json:\"nft_id\""
			Name    string "json:\"name\""
			Uri     string "json:\"uri\""
			Data    string "json:\"data\""
		}{
			{ClassID: "avatauuj0hj53thkyahiaitfmctsensn", NFTID: "avataj5h7algcaibxiz5ipbi5o97kfqs", Name: "TestEditorNFTBatch"},
			{ClassID: "avatauuj0hj53thkyahiaitfmctsensn", NFTID: "avatarxamgw6sqwvclfqfnvkgkwkp6zf", Name: "TestEditorNFTBatch"},
		},
		OperationID: "TestEditorNFTBatch",
	}
	result := client.NFT.EditorNFTBatch(params, owner)
	t.Logf("%+v \n", result)
}

//批量销毁nft
func TestDeleteNFTBatch(t *testing.T) {
	client := GetClient()
	owner := "iaa153uyr6ghtumt3lrtdwndplk4ggal9r6gm6953g"
	params := &models.DeleteNFTBatchReq{
		NFTs: []struct {
			ClassID string "json:\"class_id\""
			NFTID   string "json:\"nft_id\""
		}{
			{ClassID: "avatauuj0hj53thkyahiaitfmctsensn", NFTID: "avataj5h7algcaibxiz5ipbi5o97kfqs"},
			{ClassID: "avatauuj0hj53thkyahiaitfmctsensn", NFTID: "avatarxamgw6sqwvclfqfnvkgkwkp6zf"},
		},
		OperationID: "TestDeleteNFTBatch",
	}
	result := client.NFT.DeleteNFTBatch(params, owner)
	t.Logf("%+v \n", result)
}

//查询nft
func TestQueryNFT(t *testing.T) {
	client := GetClient()
	params := &models.QueryNFTReq{}
	result := client.NFT.QueryNFT(params)
	t.Logf("%+v \n", result)
}

//查询nft详情
func TestQueryNFTById(t *testing.T) {
	client := GetClient()
	class_id := "avata06trsneoyah2jna7lkezkv16ib7"
	nft_id := "avatazanb01zho24jjul8lq8zcpciphz"
	result := client.NFT.QueryNFTById(class_id, nft_id)
	t.Logf("%+v \n", result)
}

//查询nft操作记录
func TestQueryNFTHistory(t *testing.T) {
	client := GetClient()
	class_id := "avatauuj0hj53thkyahiaitfmctsensn"
	nft_id := "avataj5h7algcaibxiz5ipbi5o97kfqs"
	params := &models.QueryNFTHistoryReq{}
	result := client.NFT.QueryNFTHistory(params, class_id, nft_id)
	t.Logf("%+v \n", result)
}
