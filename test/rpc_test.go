package test

import (
	"encoding/json"
	"fmt"
	"github.com/JFJun/go-substrate-crypto/ss58"
	"github.com/coldwallet-group/bifrost-go/client"
	"testing"
)

func Test_GetBlockByNumber(t *testing.T) {
	c, err := client.New("wss://rpc.polkadot.io")
	if err != nil {
		t.Fatal(err)
	}
	c.SetPrefix(ss58.PolkadotPrefix)
	//expand.SetSerDeOptions(false)
	resp, err := c.GetBlockByNumber(3901804)
	if err != nil {
		t.Fatal(err)
	}
	d, _ := json.Marshal(resp)
	fmt.Println(string(d))
}

func Test_GetAccountInfo(t *testing.T) {
	c, err := client.New("wss://rpc.polkadot.io")
	if err != nil {
		t.Fatal(err)
	}
	c.SetPrefix(ss58.PolkadotPrefix)
	ai, err := c.GetAccountInfo("1exaAg2VJRQbyUBAeXcktChCAqjVP9TUxF3zo23R2T6EGdE")
	if err != nil {
		t.Fatal(err)
	}
	d, _ := json.Marshal(ai)
	fmt.Println(string(d))
	fmt.Println(ai.Data.Free.String())
}
