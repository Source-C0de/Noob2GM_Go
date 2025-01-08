package main

func checkPrefix(s1 string, s2 string) bool{

    for i:=0;i<len(s1);i++ {
        if s1[i]!=s2[i] {
            return false;
        };
    }
    return true;
}
func checkSuffix(s1 string, s2 string) bool{
    i,j := len(s1)-1, len(s2)-1;
    for i>=0 {
        if s1[i]!=s2[j] {
            return false;
        }
        i--;
        j--;
    }
    return true;
}
func countPrefixSuffixPairs(words []string) int {
    cnt := 0;
    for i:=0;i<len(words)-1;i++{
        s1:= words[i];
        for j:=i+1;j<len(words);j++{
            s2:= words[j];
            fmt.Println(len(s1),len(s2));
            if len(s1)>len(s2){
                continue;
            }
            if checkPrefix(s1,s2)==true && checkSuffix(s1,s2) == true {
                cnt++;
            }
        }
    }
    return cnt;
}