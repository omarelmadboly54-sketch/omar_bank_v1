package token

import "time"
/*
NOTE:i made all the folder of token according to section 21
*/
type Maker interface {
	CreateToken(username string, duration time.Duration)(string,error)

	VerifyToken(token string)(*Payload,error)
}