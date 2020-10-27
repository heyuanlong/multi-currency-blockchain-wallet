package contracts

import (
	"bytes"
	"math/big"
	"testing"

	"upex-wallet/wallet-base/newbitx/misclib/utils"
)

func TestMethodID(t *testing.T) {
	c := WalletSimpleCaller{
		MethodName: "sendMultiSig",
	}
	if !bytes.Equal(c.MethodID(), []byte{0x39, 0x12, 0x52, 0x15}) {
		t.Error("method id matchs")
	}

}

func TestUnpackValues(t *testing.T) {
	var (
		expectAddress = "0x0Ec6816971513071E6b766B7BEd84Cc9A74D84a0"
		expectAmount  = big.NewInt(2418247420000000000)
	)
	c := WalletSimpleCaller{
		MethodName: "sendMultiSig",
	}
	input := "0000000000000000000000000ec6816971513071e6b766b7bed84cc9a74d84a0000000000000000000000000000000000000000000000000218f5737501ad80000000000000000000000000000000000000000000000000000000000000000c0000000000000000000000000000000000000000000000000000000005b1257290000000000000000000000000000000000000000000000000000000000000b610000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000041a8674cb427294f17cdb316438d04605c7be9d3f50530b8b0f7329d3faedf647a3a7520a41cc15eb36372eeebcca971daa29ffe1956a88dac3498808982382edb1c00000000000000000000000000000000000000000000000000000000000000"
	// input = "000000000000000000000000c16b542ff490e01fcc0dc58a60e1efdc3e357ca600000000000000000000000000000000000000000000000002c855df12d4f00000000000000000000000000000000000000000000000000000000000000000c0000000000000000000000000000000000000000000000000000000005b1006270000000000000000000000000000000000000000000000000000000000000af80000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000041e7fbd635f58f52dc86ceff50cf1fb73b7de97c80b8d4c0358b127b4f11dbd87032800296572bc191ef98b2feac4987b1a540c16109e4242a1da602f755bed2de1b00000000000000000000000000000000000000000000000000000000000000"
	// input = "0000000000000000000000000a8ae62c245b98ed5d1a44430706fccbcecb8e4400000000000000000000000000000000000000000000000002b2f444bd39fc0000000000000000000000000000000000000000000000000000000000000000c0000000000000000000000000000000000000000000000000000000005b0562ec0000000000000000000000000000000000000000000000000000000000000972000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000004152fe4ba176ef35c0b30591ff6e17aeb38d6307496d5bcaed4cf671d232b835410187a1ea1cb12f218542e70db2c88d7b6f7b1b82187e6cc82b7f4ea7028fabb51c00000000000000000000000000000000000000000000000000000000000000"
	address, amount, err := c.UnpackParams(utils.HexToBytes(input))
	if err != nil {
		t.Errorf("unpack values error %+v", err)
	}
	if address.String() != expectAddress {
		t.Errorf("unpack values error, expect %s, give %v", expectAddress, address.String())
	}
	if amount.Cmp(expectAmount) != 0 {
		t.Errorf("unpack values error, expect %s, give %v", expectAmount, amount)
	}
}
