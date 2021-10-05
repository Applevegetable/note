package main

import (
	"fmt"
	//"internal/syscall/windows"
	"strings"
	//"math"
)
func countWord(text string) map[string]int{
		words:=strings.Fields(strings.ToLower(text))//返回的是一个string类型的切片
		frequnecy:= make(map[string]int)
		for _,word := range words{
			word=strings.Trim(word,`,."- `)//输入是string类型的
			frequnecy[word]++

		}
		return frequnecy


}
func main() {

    //tempture := []float64{		-28.0,32.0,-31.0,-29.0,-23.0,-29.0,-28.0,-33.0,	}

	/*实现数据分组
	groups:=make(map[float64][]float64)//创建一个映射，值的类型是切片
	for _,t := range tempture{
		g:=math.Trunc(t/10)*10
		groups[g]=append(groups[g], g)
	}
	for g,temtempture := range groups{
		fmt.Printf("%v :%v\n",g,temtempture)
	}
	*/

	/*映射用为集合set
	set:=make(map[float64]bool)
	for _,t:=range tempture{
		set[t]=true
	}
	if set[-28.0]{
		fmt.Println("set number")
	}
	fmt.Println(set)
	*/
	word:=`As far as eye could reach he saw nothing but the stems of the great plants about him receding in the violet shade, and far overhead the multiple transparency of huge leaves filtering the sunshine to the solemn splendour of twilight in which he walked. Whenever he felt able he ran again; the ground continued soft and springy, covered with the same resilient weed which was the first thing his hands had touched in Malacandra. Once or twice a small red creature scuttled across his path, but otherwise there seemed to be no life stirring in the wood; nothing to fear—except the fact of wandering unprovisioned and alone in a forest of unknown vegetation thousands or millions of miles beyond the reach or knowledge of man.`
	///frequnecy:=make(map[string]int)
	frequnecy:=countWord(word)
	for word,count:=range frequnecy{
		if count>1 {
			fmt.Printf("%d %v\n",count,word)
		}
	}
	

}
