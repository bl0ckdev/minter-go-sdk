package transaction_test

import (
	"fmt"
	"github.com/MinterTeam/minter-go-sdk/transaction"
	"testing"
)

func TestTransactionRedeemCheck_Sign(t *testing.T) {
	data := transaction.NewRedeemCheckData().
		MustSetProof("da021d4f84728e0d3d312a18ec84c21768e0caa12a53cb0a1452771f72b0d1a91770ae139fd6c23bcf8cec50f5f2e733eabb8482cf29ee540e56c6639aac469600").
		MustSetRawCheck("Mcf89b01830f423f8a4d4e5400000000000000843b9aca00b8419b3beac2c6ad88a8bd54d24912754bb820e58345731cb1b9bc0885ee74f9e50a58a80aa990a29c98b05541b266af99d3825bb1e5ed4e540c6e2f7c9b40af9ecc011ca00f7ba6d0aa47d74274b960fba02be03158d0374b978dcaa5f56fc7cf1754f821a019a829a3b7bba2fc290f5c96e469851a3876376d6a6a4df937327b3a5e9e8297")

	tx, _ := transaction.NewBuilder(transaction.TestNetChainID).NewTransaction(data)
	signedTx, _ := tx.SetNonce(1).SetGasPrice(1).SetGasCoin("MNT").Sign("07bc17abdcee8b971bb8723e36fe9d2523306d5ab2d683631693238e0f9df142")
	signedTxEncode, _ := signedTx.Encode()
	fmt.Println(signedTxEncode)
	// Output:
	// 0xf8910102018a4d4e540000000000000005b7f68a535550455220544553548a5350525445535400000089056bc75e2d631000008a021e19e0c9bab24000000a893635c9adc5dea00000808001b845f8431ba07bf9c6916aabaac7fb34811b42350c0dbcfc6228cf2ce9b927254d01f9e0ec66a0039ea86546a950cd717544d9b19c30a5248cfeb0f93060145144b5bb511a4218

}
