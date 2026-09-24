package util

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

/*
Note: i made this code according to section 6 to be used in the testing
*/
const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixMilli())
}

func RandomInt(min,max int64) int64{
	return  min+rand.Int63n(max-min+1)
}

func RandomString(n int)string{
	var sb strings.Builder
	k:=len(alphabet)

	for i:=0;i<n;i++{
		c:=alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

func RandomOwner()string{
	return RandomString(6)
}

func RandomMoney() int64{
	return RandomInt(0,1000)
}

func RandomCurrency() string{
	currencies:=[]string{USD, EUR, CAD}
	n:=len(currencies)
	return currencies[rand.Intn(n)]
}

/* i made this code according to section 17*/
func RandomEmail()string{
	return fmt.Sprintf("%s@email.com",RandomString(6))
}
