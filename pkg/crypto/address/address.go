package address

import (
	"crypto/sha1" // placeholder for RIPEMD-160
	"math/big"

	"github.com/devprosvn/VNPrider/pkg/crypto/hash"
)

const alphabet = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"

// base58Encode encodes bytes into a base58 string.
func base58Encode(b []byte) string {
	x := new(big.Int).SetBytes(b)
	base := big.NewInt(58)
	zero := big.NewInt(0)
	mod := new(big.Int)
	var encoded []byte
	for x.Cmp(zero) > 0 {
		x.QuoRem(x, base, mod)
		encoded = append([]byte{alphabet[mod.Int64()]}, encoded...)
	}
	// leading zeros
	for _, v := range b {
		if v == 0x00 {
			encoded = append([]byte{alphabet[0]}, encoded...)
		} else {
			break
		}
	}
	return string(encoded)
}

// GenerateAddress produces a Base58Check-encoded address.
func GenerateAddress(pubKey []byte, version byte) string {
	h1 := hash.ComputeSHA3(pubKey)
	// NOTE: RIPEMD-160 is not available; SHA1 is used as a placeholder
	r := sha1.New()
	r.Write(h1)
	payload := append([]byte{version}, r.Sum(nil)...)
	checksumFull := hash.ComputeSHA3(hash.ComputeSHA3(payload))
	addrBin := append(payload, checksumFull[:4]...)
	return base58Encode(addrBin)
}
