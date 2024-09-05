package service

import (
	"crypto/rsa"
	"encoding/base64"
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"
	"time"
)

func EncodeMM(password string) string {
	// 示例modulus和exponent的Base64值
	modulusBase64, exponentBase64 := GetModulusAndExpoent()
	// 将Base64转换为16进制
	modulusHex, err := base64ToHex(modulusBase64)
	if err != nil {
		return ""
	}
	exponentHex, err := base64ToHex(exponentBase64)
	if err != nil {
		return ""
	}

	// 配置RSA公钥
	publicKey, err := configureRSA(modulusHex, exponentHex)
	if err != nil {
		return ""
	}

	// 加密并转换为Base64
	encryptedBase64, err := encryptAndBase64(publicKey, password)
	if err != nil {
		return ""
	}

	return encryptedBase64
}

// base64ToHex 将Base64编码转换为16进制字符串
func base64ToHex(base64Str string) (string, error) {
	bytes, err := base64.StdEncoding.DecodeString(base64Str)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

// configureRSA 使用modulus和exponent配置RSA公钥
func configureRSA(modulusHex, exponentHex string) (*rsa.PublicKey, error) {
	// 将16进制字符串转换为大整数
	modulus := new(big.Int)
	modulus.SetString(modulusHex, 16)

	// 将16进制字符串转换为uint64
	exponentBytes, err := hex.DecodeString(exponentHex)
	if err != nil {
		return nil, err
	}
	exponentBytes = append(make([]byte, 8-len(exponentBytes)), exponentBytes...)
	exponent := binary.BigEndian.Uint64(exponentBytes)
	// 配置RSA公钥
	publicKey := &rsa.PublicKey{
		N: modulus,
		E: int(exponent),
	}

	return publicKey, nil
}

// encryptAndBase64 使用RSA公钥加密并将结果转换为Base64
func encryptAndBase64(publicKey *rsa.PublicKey, message string) (string, error) {
	encryptedBytes, err := rsa.EncryptPKCS1v15(nil, publicKey, []byte(message))
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(encryptedBytes), nil
}
func GetModulusAndExpoent() (string, string) {
	client := &http.Client{}
	url := fmt.Sprintf("https://grd.ccnu.edu.cn/yjsxt/xtgl/login_getPublicKey.html?time=%d", time.Now().Unix())
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", ""
	}
	resp, err := client.Do(req)
	if err != nil {
		return "", ""
	}
	defer resp.Body.Close()
	var reply struct {
		Modulus  string `json:"modulus"`
		Exponent string `json:"exponent"`
	}
	err = json.NewDecoder(resp.Body).Decode(&reply)
	if err != nil {
		return "", ""
	}
	return reply.Modulus, reply.Exponent
}
