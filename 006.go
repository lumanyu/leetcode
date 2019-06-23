import "container/list"

type Stack struct {
 list *list.List
}

func NewStack() *Stack {
 list := list.New()
 return &Stack{list}
}

func (stack *Stack) Push(value interface{}) {
 stack.list.PushBack(value)
}

func (stack *Stack) Pop() interface{} {
 e := stack.list.Back()
 if e != nil {
     stack.list.Remove(e)
     return e.Value
 }
 return nil
}

func (stack *Stack) Peak() interface{} {
 e := stack.list.Back()
 if e != nil {
     return e.Value
 }

 return nil
}

func (stack *Stack) Len() int {
 return stack.list.Len()
}

func (stack *Stack) Empty() bool {
 return stack.list.Len() == 0
}

func transfer(i int, numRows int) (int, int) {
    tmp_i := i%(numRows*2-2) 
    tmp_j := (int)( (float64)(i+1) / (float64)(numRows*2-2) )
    var new_i,new_j int
    if tmp_i < numRows {
        new_i = tmp_i
        new_j = tmp_j*(numRows-1)
    } else {
        new_i = (numRows-1) - (tmp_i - (numRows-1))
        new_j = tmp_j*(numRows-1) + (tmp_i - numRows - 1)
    }
    fmt.Println(i," ", tmp_i, tmp_j, " ", new_i, new_j)
    return new_i,new_j
}

func convert(s string, numRows int) string {
    if numRows==1 {
        return s
    }
    //var s_arr[numRows] *Stack 
    //var queue_arr [5][]byte
    queue_arr := make([][]byte, numRows)
    //for i:=0; i<numRows; i++ {
    //    //s_arr[i] = NewStack()
    //    //queue_arr[i] = make([]char, 0)
    //    
    //    queue_arr[i] = make([]byte, 0)
    //}
    
    for i:=0; i< len(s); i++ {
        new_i, _ := transfer(i, numRows)
        queue_arr[new_i] = append(queue_arr[new_i], s[i])
    }

    var s_ret string
    for i:=0; i<numRows; i++ {
        for len(queue_arr[i]) != 0 {
            //s_ret.append(queue_arr[i][0])
            
            s_ret +=string(queue_arr[i][0])
            queue_arr[i] = queue_arr[i][1:]
        }
    }
    return s_ret
}
