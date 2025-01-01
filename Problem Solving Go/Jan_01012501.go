// Maximum Score After Splitting a String 

package main

import "fmt"

// Brute-force solution TC- O(n2)
func maximumScore_BruteForce(s string) int {
	ma:= 0;
    for i:=0;i<len(s)-1;i++ {
        l := 0;
        r := 0;
        for il:=0;il<=i;il++ {
            if s[il] == '0'{
                l++;
            }
        }
        for iR:=i+1;iR<len(s);iR++ {
            if s[iR] == '1'{
                r++;
            }
        }
        fmt.Println(l,r);
        cnt := l+r;
        if (cnt>=ma){
            ma = cnt;
        }
    }
    return ma;

}


//Optimal Solution Sliding Windows

func maximumScore_Optimal(s string) int {
	return 0
}


func main()  {
	
}
