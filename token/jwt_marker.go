package token

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

const minSecretKeySize = 32

type JWTMaker struct {
	secretKey string
}

func NEWJWTMaker(secretKey string) (Maker, error) {
	if len(secretKey) < minSecretKeySize {
		return nil, fmt.Errorf("invalid key size:must be at least %d characters",minSecretKeySize)
	}
	return &JWTMaker{secretKey},nil
}

func (maker *JWTMaker)CreateToken(username string, duration time.Duration)(string,error){
	payload,err:=NewPayload(username,duration)
	if err!=nil{
		return "",err
	}
	jwtToken:=jwt.NewWithClaims(jwt.SigningMethodHS256,payload)
	return jwtToken.SignedString([]byte(maker.secretKey))

}

func (maker *JWTMaker)VerifyToken(token string)(*Payload,error){
	KeyFunc:=func(token *jwt.Token)(interface{},error){
		_,ok:=token.Method.(*jwt.SigningMethodHMAC)
		if!ok{
			return nil,errInvalidToken
		}
		return []byte(maker.secretKey),nil

	}
	jwtToken,err:=jwt.ParseWithClaims(token,&Payload{},KeyFunc)
	if err!=nil{
		verr,ok:=err.(*jwt.ValidationError)
		if ok && errors.Is(verr.Inner,errExpiredToken){
			return nil,errExpiredToken
		}
		return nil,errInvalidToken
	}
	payload,ok:=jwtToken.Claims.(*Payload)
	if!ok{
		return nil,errInvalidToken
	}
	return payload,nil

}