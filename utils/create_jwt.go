package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
)

type Header struct {
	Alg string    `json:"alg"` 
	Typ string    `json:"typ"` 
}

type Payload struct {
	Sub  int    `json:"sub"`
	Email string   `json:"email"` 
    Name string    `json:"name"`
}

func CreateJwt(secret string ,  data Payload) (string , error) {
	header := Header{
		Alg: "HS256",
		Typ: "JWT",
	}

	byteArrayHeader , err := json.Marshal(header)
    if err != nil {
		return "" , err
	}
    
	headerBase64 := base64UrlEncode(byteArrayHeader)

    byteArrayData , err := json.Marshal(data) 
	if err != nil {
		return "" , err
	}

    payloadBase64 := base64UrlEncode(byteArrayData)

	byteArraySecret := []byte(secret)

	fullmessage := headerBase64 + "." + payloadBase64 
    
	byteArrayMessage := []byte(fullmessage) 

	h:= hmac.New(sha256.New , byteArraySecret)
	h.Write(byteArrayMessage)

	signature := h.Sum(nil) 

	signatureBase64 := base64UrlEncode(signature)

	jwt := headerBase64 + "." + payloadBase64 + "." + signatureBase64 

	return  jwt , nil
}

func base64UrlEncode(data []byte) string {
	return base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(data)
}