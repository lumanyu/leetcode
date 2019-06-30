//import "strconv"

const (
    S_BEGIN = iota
    S_SIGNED_BIT
    S_NUMBER
    S_END
)

func is_effective_number(c byte) bool{
    //fmt.Printf("%c\n", c)
    if c>='0' && c<='9' {
        return true
    }
    return false
}

func myAtoi(str string) int {
    INT_MAX := int(^uint32(0) >> 1)
	INT_MIN := ^INT_MAX
    
    var state int = S_BEGIN
    var signed_bit int32 = 1
    var sum int32 = 0 
    //fmt.Println(str)
    for i:=0; i<len(str); i++ {
        //fmt.Println(i, str[i]-'0')
        switch state {
        case S_BEGIN:
            if str[i] == '-' {
                signed_bit = -1
                state = S_NUMBER
                continue
            } else if str[i] == '+' {
                //signed_bit = 1
                state = S_NUMBER
                continue
            } else if str[i] == ' ' {
                continue
            } else if is_effective_number(str[i]) == false {
                return 0
            } else {
                state = S_NUMBER
            }
            fallthrough
        case S_NUMBER:
            if is_effective_number(str[i]) == true {
                //fmt.Println(i, strconv.Atoi(string(str[i]))
                n := str[i] - '0'
                last_sum := sum
                //sum *= 10
                fmt.Printf("%X\n", signed_bit*sum)
                if (int(sum)*int(signed_bit)*10 > INT_MAX) {
                    return INT_MAX
                } else if (int(sum)*int(signed_bit)*10<=INT_MIN) {
                    return INT_MIN
                }
                sum *= int32(10)
                fmt.Printf("%X\n", signed_bit*sum)
                sum += int32(n)
                if (signed_bit == 1 && sum<last_sum) {
                    return int(0x7FFFFFFF)
                } else if (signed_bit == -1 && int32(signed_bit*sum)>int32(signed_bit*last_sum)) {
                    //return 0x80000000
                    //return int(int32(0xFFFFFFFF))
                    fmt.Println(INT_MIN)
                    return INT_MIN
                }
                //fmt.Println(sum)
                //fmt.Println(signed_bit*sum, signed_bit*last_sum, n)
            } else {
                state = S_END
                break
            }
        case S_END:
        default:
            break
        }
    }
    return int(sum * signed_bit)
}
