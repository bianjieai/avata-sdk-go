package tests

import (
	"avata-sdk-go/models"
	"testing"
)

//创建nft类别
func TestCreateNFTClass(t *testing.T) {
	client := GetClient()
	params := &models.CreateNFTClassReq{
		Name:        "TestCreateNFTClass",
		Owner:       "iaa1h8kfpd5wjva9glmu4rk9cnlrcu5m2yhej4n3hf",
		OperationID: "TestCreateNFTClass",
	}
	result, err := client.NFT.CreateNFTClass(params)
	if err != nil {
		t.Log(err)
		return
	}
	t.Logf("%+v \n", result)
}

//查询nft类别
func TestQueryNFTClass(t *testing.T) {
	client := GetClient()
	params := &models.QueryNFTClassReq{}
	//加查询条件 不加条件 path:=""即可
	path := models.QueryNFTClass + "?" + "name=TestCreateNFTClass"
	//path := ""
	result, err := client.NFT.QueryNFTClass(params, path)
	if err != nil {
		t.Log(err)
		return
	}
	t.Logf("%+v \n", result)
}

//查询nft类别详情
func TestQueryNFTClassById(t *testing.T) {
	client := GetClient()
	id := "avata06trsneoyah2jna7lkezkv16ib7"
	result, err := client.NFT.QueryNFTClassById(id)
	if err != nil {
		t.Log(err)
		return
	}
	t.Logf("%+v \n", result)
}

//转让nft类别
func TestTransfersNFClass(t *testing.T) {
	client := GetClient()
	class_id := "avata06trsneoyah2jna7lkezkv16ib7"
	owner := "iaa1h8kfpd5wjva9glmu4rk9cnlrcu5m2yhej4n3hf"
	params := &models.TransfersNFClassReq{
		Recipient:   "iaa1kdrjdj6ptxfg3z8s9hj6crkuztmnjnzq4rujc5",
		OperationID: "TestTransfersNFClass",
	}
	result, err := client.NFT.TransfersNFClass(params, class_id, owner)
	if err != nil {
		t.Log(err)
		return
	}
	t.Logf("%+v \n", result)
}

//发行nft
func TestCreateNFT(t *testing.T) {
	client := GetClient()
	class_id := "avata06trsneoyah2jna7lkezkv16ib7"
	params := &models.CreateNFTReq{
		Name:        "TestCreateNFT",
		OperationID: "TestCreateNFT",
	}
	result, err := client.NFT.CreateNFT(params, class_id)
	if err != nil {
		t.Log(err)
		return
	}
	t.Logf("%+v \n", result)
}

//转让nft
func TestTransfersNFT(t *testing.T) {
	client := GetClient()
	class_id := "avata06trsneoyah2jna7lkezkv16ib7"
	owner := "iaa1h8kfpd5wjva9glmu4rk9cnlrcu5m2yhej4n3hf"
	nft_id := "avatagccrcjdywgljzmpkv2fa6vxtwpo"
	params := &models.TransfersNFTReq{
		Recipient:   "iaa1kdrjdj6ptxfg3z8s9hj6crkuztmnjnzq4rujc5",
		OperationID: "TestTransfersNFT",
	}
	result, err := client.NFT.TransfersNFT(params, class_id, owner, nft_id)
	if err != nil {
		t.Log(err)
		return
	}
	t.Logf("%+v \n", result)
}

//编辑nft
func TestEditorNFT(t *testing.T) {
	client := GetClient()
	class_id := "avata06trsneoyah2jna7lkezkv16ib7"
	owner := "iaa1h8kfpd5wjva9glmu4rk9cnlrcu5m2yhej4n3hf"
	nft_id := "avatagccrcjdywgljzmpkv2fa6vxtwpo"
	params := &models.EditorNFTReq{
		Name:        "EditorNFT",
		OperationID: "EditorNFT",
	}
	result, err := client.NFT.EditorNFT(params, class_id, owner, nft_id)
	if err != nil {
		t.Log(err)
		return
	}
	t.Logf("%+v \n", result)
}

//销毁nft
func TestDeleteNFT(t *testing.T) {
	client := GetClient()
	class_id := "avata06trsneoyah2jna7lkezkv16ib7"
	owner := "iaa1kdrjdj6ptxfg3z8s9hj6crkuztmnjnzq4rujc5"
	nft_id := "avatagccrcjdywgljzmpkv2fa6vxtwpo"
	params := &models.DeleteNFTReq{
		OperationID: "DeleteNFT",
	}
	result, err := client.NFT.DeleteNFT(params, class_id, owner, nft_id)
	if err != nil {
		t.Log(err)
		return
	}
	t.Logf("%+v \n", result)
}

//批量发行nft
func TestCreateNFTBatch(t *testing.T) {
	client := GetClient()
	class_id := "avata06trsneoyah2jna7lkezkv16ib7"
	params := &models.CreateNFTBatchReq{
		Name: "TestCreateNFTBatch1",
		Recipients: []struct {
			Amount    int    "json:\"amount\""
			Recipient string "json:\"recipient\""
		}{
			{Amount: 1, Recipient: "iaa1h8kfpd5wjva9glmu4rk9cnlrcu5m2yhej4n3hf"},
			{Amount: 1, Recipient: "iaa1kdrjdj6ptxfg3z8s9hj6crkuztmnjnzq4rujc5"},
		},
		OperationID: "TestCreateNFTBatch1",
	}
	result, err := client.NFT.CreateNFTBatch(params, class_id)
	if err != nil {
		t.Log(err)
		return
	}
	t.Logf("%+v \n", result)
}

//批量转让nft
func TestTransfersNFTBatch(t *testing.T) {
	client := GetClient()
	owner := "iaa1kdrjdj6ptxfg3z8s9hj6crkuztmnjnzq4rujc5"
	params := &models.TransfersNFTBatchReq{
		Data: []struct {
			NFTs []struct {
				ClassID string "json:\"class_id\""
				NFTID   string "json:\"nft_id\""
			} "json:\"nfts\""
			Recipient string "json:\"recipient\""
		}{
			{Recipient: "iaa1jta8kw63zd557y6qv02cllmespewg3uvtj4ucy", NFTs: []struct {
				ClassID string "json:\"class_id\""
				NFTID   string "json:\"nft_id\""
			}{
				{ClassID: "avata06trsneoyah2jna7lkezkv16ib7", NFTID: "avatabjmxoqjs1igipui4dbfbcak8r2o"},
				{ClassID: "avata06trsneoyah2jna7lkezkv16ib7", NFTID: "avata3rhx00bapvhoi32ms3f7ju7rahz"},
			}},
		},
		OperationID: "TestTransfersNFTBatch",
	}
	result, err := client.NFT.TransfersNFTBatch(params, owner)
	if err != nil {
		t.Log(err)
		return
	}
	t.Logf("%+v \n", result)
}

//批量编辑nft
func TestEditorNFTBatch(t *testing.T) {
	client := GetClient()
	owner := "iaa1jta8kw63zd557y6qv02cllmespewg3uvtj4ucy"
	params := &models.EditorNFTBatchReq{
		NFTs: []struct {
			ClassID string "json:\"class_id\""
			NFTID   string "json:\"nft_id\""
			Name    string "json:\"name\""
			Uri     string "json:\"uri\""
			Data    string "json:\"data\""
		}{
			{ClassID: "avata06trsneoyah2jna7lkezkv16ib7", NFTID: "avata3rhx00bapvhoi32ms3f7ju7rahz", Name: "TestEditorNFTBatch"},
			{ClassID: "avata06trsneoyah2jna7lkezkv16ib7", NFTID: "avatabjmxoqjs1igipui4dbfbcak8r2o", Name: "TestEditorNFTBatch"},
		},
		OperationID: "TestEditorNFTBatch",
	}
	result, err := client.NFT.EditorNFTBatch(params, owner)
	if err != nil {
		t.Log(err)
		return
	}
	t.Logf("%+v \n", result)
}

//批量销毁nft
func TestDeleteNFTBatch(t *testing.T) {
	client := GetClient()
	owner := "iaa1jta8kw63zd557y6qv02cllmespewg3uvtj4ucy"
	params := &models.DeleteNFTBatchReq{
		NFTs: []struct {
			ClassID string "json:\"class_id\""
			NFTID   string "json:\"nft_id\""
		}{
			{ClassID: "avata06trsneoyah2jna7lkezkv16ib7", NFTID: "avata3rhx00bapvhoi32ms3f7ju7rahz"},
			{ClassID: "avata06trsneoyah2jna7lkezkv16ib7", NFTID: "avatabjmxoqjs1igipui4dbfbcak8r2o"},
		},
		OperationID: "TestDeleteNFTBatch",
	}
	result, err := client.NFT.DeleteNFTBatch(params, owner)
	if err != nil {
		t.Log(err)
		return
	}
	t.Logf("%+v \n", result)
}

//查询nft
func TestQueryNFT(t *testing.T) {
	client := GetClient()
	params := &models.QueryNFTReq{}
	path := models.QueryNFT + "?name=TestCreateNFTBatch1"
	result, err := client.NFT.QueryNFT(params, path)
	if err != nil {
		t.Log(err)
		return
	}
	t.Logf("%+v \n", result)
}

//查询nft详情
func TestQueryNFTById(t *testing.T) {
	client := GetClient()
	class_id := "avata06trsneoyah2jna7lkezkv16ib7"
	nft_id := "avatazanb01zho24jjul8lq8zcpciphz"
	result, err := client.NFT.QueryNFTById(class_id, nft_id)
	if err != nil {
		t.Log(err)
		return
	}
	t.Logf("%+v \n", result)
}

//查询nft操作记录
func TestQueryNFTHistory(t *testing.T) {
	client := GetClient()
	class_id := "avata06trsneoyah2jna7lkezkv16ib7"
	nft_id := "avatagccrcjdywgljzmpkv2fa6vxtwpo"
	params := &models.QueryNFTHistoryReq{}
	path := ""
	result, err := client.NFT.QueryNFTHistory(params, path, class_id, nft_id)
	if err != nil {
		t.Log(err)
		return
	}
	t.Logf("%+v \n", result)
}
