package main
import (
	"fmt"
	"strings"
)

func main(){
		message:="zenmehuish s "
		keyword:="GOLANG"
		keyIndex:=0
		cipherText:=""
		
		message=strings.ToUpper(strings.Replace(message," ","",-1))
		keyword=strings.ToUpper(strings.Replace(keyword," ","",-1))

		for i:=0;i<len(message);i++ {
			c:=message[i]
			if c>='A'&& c<='Z'{
				c-='A'
				k:=keyword[keyIndex]-'A'
				//直接转换为rune/byte模式，计算基础编码值，之后加入“A'将其转变为大写字母
				c=(c+k)%26+'A'
				keyIndex++
				keyIndex%=len(keyword)

			}
			cipherText+=string(c)
			

		}
		fmt.Println(cipherText)


}