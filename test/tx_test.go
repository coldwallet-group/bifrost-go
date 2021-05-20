package test

import (
	"encoding/json"
	"fmt"
	"github.com/JFJun/go-substrate-crypto/crypto"
	"github.com/JFJun/go-substrate-crypto/ss58"
	"github.com/coldwallet-group/bifrost-go/client"
	"github.com/coldwallet-group/bifrost-go/expand"
	"github.com/coldwallet-group/bifrost-go/tx"
	"github.com/coldwallet-group/bifrost-go/utils"
	"math/big"
	"testing"
)
func Test_acc(t *testing.T){

	t.Log(utils.AddressToPublicKey("5DXVzVZSAreGnbuHvbxZNQpADwPtn2Y4WDpw8gPb7vbvvuae"))
	t.Log(ss58.EncodeByPubHex("40abeafb6f99b8b7acd4c828a08f11a06e92b332ef8912ed1333993ce9e8d77d", []byte{0x2a}))
	//40abeafb6f99b8b7acd4c828a08f11a06e92b332ef8912ed1333993ce9e8d77d
	c, err := client.New("http://13.114.44.225:31933")
	if err != nil {
		t.Fatal(err)
	}
	b,err :=c.GetBlockByNumber(5302725)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(String(b))
	h,_ :=c.C.RPC.Chain.GetFinalizedHead()
	t.Log(h.Hex())
}
func Test_Tx2(t *testing.T) {
	// 1. 初始化rpc客户端
	c, err := client.New("http://13.114.44.225:31933")
	if err != nil {
		t.Fatal(err)
	}
	//2. 如果某些链（例如：chainX)的地址的字节前面需要0xff,则下面这个值设置为false
	//expand.SetSerDeOptions(false)
	from := "5GgFf8FPYaPM1MgKuK8J61u5EbjbY2gab6R1t3qRhFS3ytgZ"
	to := "5F4JVJ51EdSqWUMpW9V8DaaorPkAXmmw3kay2dQv8e6w9eiT"
	amount := uint64(100000000)
	//3. 获取from地址的nonce
	acc, err := c.GetAccountInfo(from)
	if err != nil {
		t.Fatal(err)
	}
	nonce := uint64(acc.Nonce)
	//4. 创建一个substrate交易，这个方法满足所有遵循substrate 的交易结构的链
	transaction := tx.NewSubstrateTransaction(from, nonce)
	//5. 初始化metadata的扩张结构
	ed, err := expand.NewMetadataExpand(c.Meta)
	if err != nil {
		t.Fatal(err)
	}
	//6. 初始化Balances.transfer的call方法
	call, err := ed.BalanceTransferCall(to, big.NewInt(int64(amount)))
	if err != nil {
		t.Fatal(err)
	}
	/*
		//Balances.transfer_keep_alive  call方法
		btkac,err:=ed.BalanceTransferKeepAliveCall(to,amount)
	*/

	/*
		toAmount:=make(map[string]uint64)
		toAmount[to] = amount
		//...
		//true: user Balances.transfer_keep_alive  false: Balances.transfer
		ubtc,err:=ed.UtilityBatchTxCall(toAmount,false)
	*/

	//7. 设置交易的必要参数
	transaction.SetGenesisHashAndBlockHash(c.GetGenesisHash(), c.GetGenesisHash()).
		SetSpecAndTxVersion(uint32(c.SpecVersion), uint32(c.TransactionVersion)).
		SetCall(call) //设置call
	//8. 签名交易
	sig, err := transaction.SignTransaction("", crypto.Sr25519Type)
	if err != nil {
		t.Fatal(err)
	}
	//9. 提交交易
	var result interface{}
	err = c.C.Client.Call(&result, "author_submitExtrinsic", sig)
	if err != nil {
		t.Fatal(err)
	}
	//10. txid
	txid := result.(string)
	fmt.Println(txid)
}
func String(d interface{})string{
	str,_ := json.Marshal(d)
	return string(str)
}