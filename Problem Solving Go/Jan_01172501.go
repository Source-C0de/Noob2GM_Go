func doesValidArrayExist(derived []int) bool {
    xOr := 0;

    for i:=0;i<len(derived);i++ {
        xOr ^= derived[i];
    }

    return xOr == 0;
}