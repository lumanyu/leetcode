func findMedianSortedArrays(num1 []int, num2 []int) float64 {
    if len(num1) == 1 && len(num2) == 1 {
        return (float64)(num1[0] + num2[0]) / float64(2)
    }
    
    i := 0
    var order [2]int
    for i=0;i+1<len(num1);i++ {
        if num1[i] > num1[i+1] {
            order[0] = 1
        } else if num1[i] < num1[i+1]  {
            order[0] = 2
        }
    }
    
    for i=0;i+1<len(num2);i++ {
        if num2[i] > num2[i+1] {
            order[1] = 1
        } else if num2[i] < num2[i+1]  {
            order[1] = 2
        }
    }
    
    var mid int = (len(num1) + len(num2))/2
    
    var even int = 0
    if (len(num1)+len(num2))%2 == 0 {
        even = 1
    }
    
    var final_result float64 = 0
    
    
    //get bu ma
    i = 0
    j := 0
    k := 0
    m := 0
    n := 0
    var select_value float64 = 0
    
    for ; i< len(num1) && j<len(num2) && k<=mid; k=k+1 {
        
        if order[0] == 1 {
            m = len(num1) - i -1
        } else if order[0] == 2 {
            m = i
        } else {
            m = i
        }

        if order[1] == 1 {
            n = len(num2) - j - 1
        } else if order[1] == 2 {
            n = j
        } else {
            n = j
        }
        
        
        if num1[m] < num2[n] {
            select_value = (float64)(num1[m])
            i++
        } else {
            select_value = (float64)(num2[n])
            j++

        }
        
        if even == 0 && k == mid {
            return select_value      
        } else if even == 1 && (k == mid-1) {
            final_result += select_value
        } else if even == 1 && (k==mid) {
            final_result += select_value
            final_result /= 2
            return final_result
        }
        

    }
    
    for ;i<len(num1) && k<=mid; k++ {
        
        if order[0] == 1 {
            m = len(num1) - i -1
        } else if order[0] == 2 {
            m = i
        } else {
            m = i
        }
        select_value = (float64)(num1[m])
        i++
        
        if even == 0 && k == mid {
            return select_value
        } else if even == 1 && (k == mid-1) {
            final_result += select_value
        } else if even == 1 && (k==mid) {
            final_result += select_value
            final_result /= 2
            return final_result
        }
    }

    for ;j<len(num2) && k<=mid; k++ {
        
        if order[0] == 1 {
            n = len(num2) - j -1
        } else if order[0] == 2 {
            n = j
        } else {
            n = j
        }
        select_value = (float64)(num2[n])
        j++
        
        if even == 0 && k == mid {
            return select_value
        } else if even == 1 && (k == mid-1) {
            final_result += select_value
        } else if even == 1 && (k==mid) {
            final_result += select_value
            final_result /= 2
            return final_result
        }
    }
    return -1
}
