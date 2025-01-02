// 2559. Count Vowel Strings in Ranges
package main

func vowelCheck(words string) bool{
    vowelString := "aeiou";
    for i:=0;i<len(vowelString);i++{
        if vowelString[i] == words[0] {
            for j:=0;j<len(vowelString); j++{
                if vowelString[j]==words[len(words)-1] {
                    return true;
                }
            }

        }
    }
    return false;
}
func vowelStrings(words []string, queries [][]int) []int {
    n := len(words)
	sum := make([]int, n)
	ans := []int{}

    if vowelCheck(words[0]){
        sum[0] = 1;
    }else{
        sum[0] = 0;
    }
    for i:=1; i<n; i++ {
        if vowelCheck(words[i]) == true {
            sum[i]=sum[i-1]+1;
        }else{
            sum[i]=sum[i-1];
        }
    }
    for i:=0; i<len(queries); i++ {
        l:= queries[i][0];
        r:= queries[i][1];
        if l==r {
            if(vowelCheck(words[l])){
                ans = append(ans,1);
            }else{
                ans = append(ans,0);
            }
        }else{
            if l==0 {
                ans = append(ans,sum[r]);
            }else{
                ans = append(ans,sum[r]-sum[l-1]);
            }
            
        }
    }
    return ans;
}
