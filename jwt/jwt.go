package jwt

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

func init() {
	Secret = "123456"
}

type AccountToken struct {
	UserID   int    `json:"userID"`
	Username string `json:"username"`
}

// Header JWT请求头
type Header struct {
	Alg string `json:"alg"`
	Typ string `json:"typ"`
}

type Payload struct {
	Issuer     string // 签发者
	IssuedAt   string // 签发时间
	Expiration string // 过期时间
	Audience   string // 接收方
}

var Secret string

func Sign(jwtHeader Header, payload Payload) (res string, err error) {
	jwtH, err := json.Marshal(jwtHeader)
	if err != nil {
		return "", errors.New("jwt请求头异常")
	}
	jh := base64.RawURLEncoding.EncodeToString(jwtH)
	jwtP, err := json.Marshal(payload)
	if err != nil {
		return "", errors.New("jwt负载异常")
	}
	jp := base64.RawURLEncoding.EncodeToString(jwtP)
	jhp := fmt.Sprintf("%s.%s", jh, jp)
	hasher := hmac.New(sha256.New, []byte(Secret))
	_, err = hasher.Write([]byte(jhp))
	if err != nil {
		return "", err
	}
	js := base64.RawURLEncoding.EncodeToString(hasher.Sum(nil))
	return fmt.Sprintf("%s.%s.%s", jh, jp, js), nil
}

func Verify(token string) (err error) {
	arr := strings.Split(token, ".")
	// 验证token是否为三段
	if len(arr) != 3 {
		return errors.New("token 格式错误")
	}
	// 提取jwt Header
	jwtH, err := base64.RawURLEncoding.DecodeString(arr[0])
	if err != nil {
		return errors.New("解压jwt header失败")
	}
	var h Header
	json.Unmarshal(jwtH, &h)
	if h.Alg != "HS256" {
		return errors.New("无法识别加密算法")
	}
	data := strings.Join(arr[0:2], ".")
	hasher := hmac.New(sha256.New, []byte(Secret))
	_, err = hasher.Write([]byte(data))
	if err != nil {
		return err
	}
	// 提取签名
	sign, err := base64.RawURLEncoding.DecodeString(arr[2])
	// 签名正确
	if hmac.Equal(hasher.Sum(nil), sign) {
		return nil
	}
	// 签名错误
	return errors.New("非法签名")
}

func GetPayload(token string) (payload Payload, err error) {
	arr := strings.Split(token, ".")
	jwtP, err := base64.RawURLEncoding.DecodeString(arr[1])
	if err != nil {
		return Payload{}, err
	}
	json.Unmarshal(jwtP, &payload)
	return payload, nil
}
