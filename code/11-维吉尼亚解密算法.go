package main

import (
	"fmt"
)

func main() {
		cipherText:= "CSOITEUIWUIZNSROCNKFD"
		keyword:= "GOLANG"
		message:=""
		keyIndex:=0

		for i:=0;i<len(cipherText);i++{
			c:=cipherText[i]-'A'
			k:=keyword[keyIndex]-'A'

			c=(c-k+26)%26+'A'
			
			message+= string(c)

			keyIndex++
			keyIndex%=len(keyword)
		}
		fmt.Println(message)

}