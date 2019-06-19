func basic_condition(i int, j int, len int, style int) int{
    switch style {
        default:
        //if i+j<len {
        if j<len {
            return 1
        }
        break
        case 1:
        if (i-j)>=0 && (i+j)<len {
            return 1
        } else {
            break
        }
        case 2:
        case 3:
        if (i-j+1)>=0 && (i+j)<len {
           return 1
        } else {
            break
        }
    }
    return 0
}

func bound_test(index int, len int) int{
    if index>=0 && index<len {
        return 1
    } else {
        return 0
    }
}

func longestPalindrome(s string) string {
    if len(s)==0 {
        return ""
    }
    var repeat_start = 0
    var repeat_end = 0
    var repeat_len = 0
    for i:=0; i<len(s); i++ {
        
        var tmp_repeat_start = 0
        var tmp_repeat_end = 0
        var tmp_repeat_len = 0
        
        var style = 0
        var j = 1
        for ; 1 == basic_condition(i, j, len(s), style); j++ {
            
            var match_style1 = 0
            var match_style2 = 0
            if bound_test(i+j, len(s))==1 && bound_test(i-j, len(s))==1 && s[i+j] == s[i-j] && (style==0 || style==1 || style==3) {
                match_style1 = 1
                fmt.Printf("loop1, i:%d,j:%d\n",i,j)
                //continue
            }
            if bound_test(i+j, len(s))==1 && bound_test(i-j+1, len(s))==1 && s[i+j] == s[i-j+1] && (style==0 || style==2 || style==3) {
                match_style2 = 1
                fmt.Printf("loop2, i:%d,j:%d\n",i,j)
                //continue
            //} else if style == 0 {
            //    fmt.Printf("loop3, i:%d,j:%d\n",i,j)
            //    continue
            }
            if match_style1 == 1 && match_style2 == 1 {
                fmt.Printf("loop3, i:%d,j:%d\n",i,j)
                style =3
                continue
            } else if match_style1==1 {
                fmt.Printf("loop6, i:%d,j:%d\n",i,j)
                style = 1
                continue 
            } else if match_style2 ==1 {
                fmt.Printf("loop7, i:%d,j:%d\n",i,j)
                style = 2
                continue
            } else {
                //fmt.Printf("loop4, i:%d,j:%d\n",i,j)
                j--
                break
            }
        }
        if style==0 {
            continue
        } else if style==3 {
            
            fmt.Printf("loop here, i:%d,j:%d\n",i,j)
            return s
        }
        
        //if (style==1 && i-j<0) || (style==2 && i+j+1<0) || i+j >= len(s) {  
        if 0 == basic_condition(i, j, len(s), style) {
            j--
        }

        if j>0 {
            switch style {
                case 1:
                    tmp_repeat_start = i-j
                case 2:
                    tmp_repeat_start = i-j+1
                default:
                    continue
            }
            tmp_repeat_end = i+j
            tmp_repeat_len = tmp_repeat_end-tmp_repeat_start+1
            
            if tmp_repeat_len > repeat_len {
                repeat_start = tmp_repeat_start
                repeat_end = tmp_repeat_end
                repeat_len = tmp_repeat_len
                
                
        
                fmt.Printf("i:%d,j:%d\n",i,j)
        
                fmt.Printf("start:%d,end:%d, len:%d\n",repeat_start,repeat_end,repeat_len)
            }
            
        }
        
    }
    
    fmt.Printf("%d,%d\n",repeat_start,repeat_end)
    sp:=s[repeat_start:repeat_end+1]
    return sp
}
