func xorAllNums(nums1 []int, nums2 []int) int {
    x1,x2 := 0,0;
    result := 0

   for _, num := range nums1 {
        x1 ^= num
    }
    for _, num := range nums2 {
        x2 ^= num
    }
    if len(nums1)%2 == 1 {
        result ^= x2;
    }
    if len(nums2)%2 == 1 {
        result ^= x1
    }
    return result;
}