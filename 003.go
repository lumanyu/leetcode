func isRepeatStr(bs []byte, begin int, end int) int {
    var fixed_space []byte = make(int, 64*(end-begin+1))
    var j int =0
    for i in range(begin, end) {
        if fixed_space[bs[i]-'a'+j*64] == 1 {
            return 1
        } else {
            fixed_space[bs[i]-'a'+j*64] = 1
        }
    }
    return 0
}

//max-> ... -> 0
func lengthOfLongestSubstring(s string) int {
    var bs []byte = []byte(s)
    for i := 0; i<len(s); i++ {
        //var begin *byte = &bs[0]
        //var end *byte = &bs[len(s)-1-i]
        //var sential *byte = &bs[len(s)-1]
        var begin int = 0
        var end int = len(s)-1-i
        var sential int = len(s)-1
        var repeat int = 0
        for end < sential {
            if isRepeatStr(begin, end) {
                repeat = 1
                break
            }
            begin++
            end++
        }
        if repeat == 0 {
            return (end-begin+1)
        }
    }
}
