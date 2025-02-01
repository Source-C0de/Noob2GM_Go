package main

import fmt

func isArraySpecial(nums []int) bool {
    for i:=1;i<len(nums);i++ {
        if nums[i]%2 == 0 && nums[i-1]%2==1 {
            continue;
        }
        if nums[i-1]%2 == 0 && nums[i]%2==1 {
            continue;
        }else{
            return false;
        }
        
    }
    return true;
}