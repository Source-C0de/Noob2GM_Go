func findThePrefixCommonArray(A []int, B []int) []int {
    ans := []int{}
    cnt := 0
    a := make(map[int]bool) // To track elements in A

    for i := 0; i < len(A); i++ {
        a[A[i]] = true // Mark A[i] as seen
        cnt = 0

        // Compare elements of B up to index i
        for j := 0; j <= i; j++ {
            if a[B[j]] {
                cnt++
            }
        }
        ans = append(ans, cnt) // Add count to result
    }

    return ans
}