package main
import "fmt"

func checkPrefix(s string, pre string) bool{

    for i:=0;i<len(pre);i++ {
        if s[i]!= pre[i] {
            return false;
        }
    }
    return true;
}
func prefixCount(words []string, pref string) int {
    cnt := 0;
    for i:=0;i<len(words);i++ {
        if len(words[i])<len(pref) {
            continue;
        }else{
            if(checkPrefix(words[i],pref)==true){
                cnt++;
            }
        }
    }
    return cnt;
}

