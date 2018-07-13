// toll
package keypair

import (
	"bytes"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"errors"

	"github.com/bumoproject/bumo-sdk-go/src/3rd/base58"
	"github.com/bumoproject/bumo-sdk-go/src/3rd/ed25519/edwards25519"
	"github.com/bumoproject/bumo-sdk-go/src/3rd/secureRandom"
)

const (
	PublicKeySize  = 32
	PrivateKeySize = 64
	SignatureSize  = 64
)

//Create
func Create() (publicKey string, privateKey string, address string, err error) {

	var ranbuf [32]byte
	ranstr, err := secureRandom.New(32)
	if err != nil {
		return "", "", "", err
	}
	ranbyte := []byte(ranstr)
	for i := range ranbuf {
		ranbuf[i] = ranbyte[i]
	}
	public, _, err := GenerateKey(ranbuf)
	if err != nil {
		return "", "", "", err
	}

	publicKey, err = encodePublicKey(public)
	if err != nil {
		return "", "", "", err
	}
	privateKey, err = encodePrivateKey(&ranbuf)
	if err != nil {
		return "", "", "", err
	}
	address, err = encodeAddress(public)
	if err != nil {
		return "", "", "", err
	}

	return publicKey, privateKey, address, nil

}

//The private key gets the public key
func GetEncPublicKey(privateKey string) (publicKey string, err error) {
	if CheckPrivateKey(privateKey) == false {
		return "", errors.New("privateKey error")
	}
	PrivateKey, err := DecodePrivateKey(privateKey)
	if err != nil {
		return "", err
	}

	PublicKey, _, err := GenerateKey(*PrivateKey)
	if err != nil {
		return "", err
	}

	return encodePublicKey(PublicKey)
}

//The public key gets the address
func GetEncAddress(publicKey string) (address string, err error) {
	if CheckPublicKey(publicKey) == false {
		return "", errors.New("publicKey error")
	}
	PublicKey, err := DecodePublicKey(publicKey)
	if err != nil {
		return "", err
	}

	return encodeAddress(PublicKey)
}

//Verify the public key
func CheckPublicKey(PublicKey string) bool {
	if PublicKey == "" {
		return false
	}
	var epub []byte
	var ret bool
	var err error
	epub, err = hex.DecodeString(PublicKey)
	if err != nil {
		return false
	}

	if len(epub) != 38 || epub[0] != 0xb0 || epub[1] != 1 {
		return false
	}

	var hash1, hash2 []byte

	dpub := epub[:34]
	h1 := sha256.New()
	h1.Write([]byte(dpub))
	hash1 = h1.Sum(nil)
	h2 := sha256.New()
	h2.Write([]byte(hash1))
	hash2 = h2.Sum(nil)
	if !(hash2[0] == epub[34] && hash2[1] == epub[35] && hash2[2] == epub[36] && hash2[3] == epub[37]) {
		return false
	}

	ret = true
	return ret
}

//Verify the private key
func CheckPrivateKey(SprivateKey string) bool {
	var ret bool

	if SprivateKey == "" {
		return false
	}

	epriv, err := base58.Decode(SprivateKey)
	if err != nil {
		return false
	}

	if !(len(epriv) == 41 && epriv[0] == 0xDA && epriv[1] == 0x37 && epriv[2] == 0x9F && epriv[3] == 1) {
		return false
	}

	if !(epriv[36] == 0x00) {
		return false
	}

	var hash1, hash2 []byte

	dpriv := epriv[:37]

	h1 := sha256.New()
	h1.Write([]byte(dpriv))
	hash1 = h1.Sum(nil)

	h2 := sha256.New()
	h2.Write([]byte(hash1))
	hash2 = h2.Sum(nil)

	if !(hash2[0] == epriv[37] && hash2[1] == epriv[38] && hash2[2] == epriv[39] && hash2[3] == epriv[40]) {
		return false
	}

	ret = true
	return ret

}

//Verify the address key
func CheckAddress(Saddress string) bool {
	if Saddress == "" {
		return false
	}
	var addre []byte
	var ret bool
	var err error
	addre, err = base58.Decode(Saddress)
	if err != nil {
		return false
	}

	if !(addre[0] == 0X01 && addre[1] == 0X56) {

		return false
	} else if !(addre[2] == 1) {
		return false
	}
	var hash1, hash2 []byte

	daddr := addre[:23]

	h1 := sha256.New()
	h1.Write([]byte(daddr))
	hash1 = h1.Sum(nil)

	h2 := sha256.New()
	h2.Write([]byte(hash1))
	hash2 = h2.Sum(nil)

	if !(hash2[0] == addre[23] && hash2[1] == addre[24] && hash2[2] == addre[25] && hash2[3] == addre[26]) {
		return false
	}
	ret = true

	return ret

}

//Generate Key
func GenerateKey(ranbuf [32]byte) (publicKey *[PublicKeySize]byte, privateKey *[PrivateKeySize]byte, err error) {
	if len(ranbuf) == 0 {
		return nil, nil, errors.New("buf is nil")
	}

	privateKey = new([64]byte)
	publicKey = new([32]byte)
	copy(privateKey[:32], ranbuf[:])

	h := sha512.New()
	h.Write(privateKey[:32])
	digest := h.Sum(nil)

	digest[0] &= 248
	digest[31] &= 127
	digest[31] |= 64

	var A edwards25519.ExtendedGroupElement
	var hBytes [32]byte
	copy(hBytes[:], digest)
	edwards25519.GeScalarMultBase(&A, &hBytes)
	A.ToBytes(publicKey)

	copy(privateKey[32:], publicKey[:])
	return
}

func bytesCombine(pBytes ...[]byte) []byte {
	len := len(pBytes)
	s := make([][]byte, len)
	for index := 0; index < len; index++ {
		s[index] = pBytes[index]
	}
	sep := []byte("")
	return bytes.Join(s, sep)
}

//Encoding public Key
func encodePublicKey(publicKey *[32]byte) (GpublicKey string, err error) {
	if publicKey == nil {
		return "", errors.New("encode publicKey is error")
	}
	var ppblic [32]byte = *publicKey
	var str_result []byte
	var hash1, hash2 []byte
	str_result = append(str_result, 0xb0)
	str_result = append(str_result, 1)
	str_result = bytesCombine(str_result, ppblic[:32])
	h1 := sha256.New()
	h1.Write([]byte(str_result))
	hash1 = h1.Sum(nil)
	h2 := sha256.New()
	h2.Write([]byte(hash1))
	hash2 = h2.Sum(nil)
	str_result = bytesCombine(str_result, hash2[:4])

	str := hex.EncodeToString(str_result)
	return str, nil
}

//Encoding private Key
func encodePrivateKey(privateKey *[32]byte) (GPrivateKey string, err error) {
	if privateKey == nil {
		return "", errors.New("encode privateKey is error")
	}
	var ppriv [32]byte = *privateKey
	var str_result []byte
	var hash1, hash2 []byte
	str_result = append(str_result, 0xDA)
	str_result = append(str_result, 0x37)
	str_result = append(str_result, 0x9F)
	str_result = append(str_result, 1)

	str_result = bytesCombine(str_result, ppriv[:])
	str_result = append(str_result, 0x00)

	h1 := sha256.New()
	h1.Write([]byte(str_result))
	hash1 = h1.Sum(nil)

	h2 := sha256.New()
	h2.Write([]byte(hash1))
	hash2 = h2.Sum(nil)

	str_result = bytesCombine(str_result, hash2[:4])
	var result string
	result = base58.Encode(str_result)

	return result, nil
}

//Encoding address
func encodeAddress(publicKey *[32]byte) (GAccoun string, err error) {
	if publicKey == nil {
		return "", errors.New("encode publicKey is error")
	}
	var ppbilc [32]byte = *publicKey

	var str_result []byte
	var hash1, hash2, pubSha []byte
	str_result = append(str_result, 0X01)
	str_result = append(str_result, 0X56)
	str_result = append(str_result, 1)
	var ppbilc1 []byte
	ppbilc1 = ppbilc[:]

	ShaPub := sha256.New()
	ShaPub.Write([]byte(ppbilc1))
	pubSha = ShaPub.Sum(nil)

	str_result = bytesCombine(str_result, pubSha[12:32])

	h1 := sha256.New()
	h1.Write([]byte(str_result))
	hash1 = h1.Sum(nil)

	h2 := sha256.New()
	h2.Write([]byte(hash1))
	hash2 = h2.Sum(nil)

	str_result = bytesCombine(str_result, hash2[:4])
	var result string
	result = base58.Encode(str_result)

	return result, nil

}

//Decode public Key
func DecodePublicKey(spublicKey string) (PublicKey *[32]byte, err error) {
	if spublicKey == "" {
		return nil, errors.New("decode publicKey error :publicKey is nil")
	}
	ispub := CheckPublicKey(spublicKey)
	if !(ispub) {
		return nil, errors.New("check publicKey error")
	}
	var epub []byte
	epub, err = hex.DecodeString(spublicKey)
	if err != nil {
		return nil, err
	}
	var dpub [32]byte
	copy(dpub[:], epub[2:34])
	return &dpub, nil

}

//Decode private Key
func DecodePrivateKey(sprivateKey string) (PrivateKey *[32]byte, err error) {
	if sprivateKey == "" {
		return nil, errors.New("decode privateKey error :privateKey is nil")
	}
	ispri := CheckPrivateKey(sprivateKey)
	if !(ispri) {
		return nil, errors.New("check privateKey error")
	}

	epriv, err := base58.Decode(sprivateKey)
	if err != nil {
		return nil, err
	}

	var dpriv [32]byte
	copy(dpriv[:], epriv[4:36])

	return &dpriv, nil
}
