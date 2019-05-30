import "sort"

func pos(nums []int, t int, repeat int) int {
    for p, v := range(nums) {
        if v == t {
            repeat-=1
            if repeat==0 {
                return p
            }
        }
    }
    return -1
}

func twoSum(nums []int, target int) []int {
    ori_nums := make([]int, len(nums))
    copy(ori_nums, nums)
    sort.Ints(nums)
    fmt.Printf("%v\n", ori_nums);
    fmt.Printf("%v\n", nums);
    j :=len(nums)-1
    for i:=0; i<len(nums); i++ {
        for ; j>=0; j-- {
            if nums[i] + nums[j] > target {
                continue
            } else if nums[i] + nums[j] < target {
                break
            } else {
                fmt.Printf("%d, %d\n", nums[i], nums[j]);
                if nums[i] == nums[j] {
                    return []int{pos(ori_nums, nums[i], 1), pos(ori_nums, nums[j], 2)}
                } else {
                    return []int{pos(ori_nums, nums[i], 1), pos(ori_nums, nums[j], 1)}
                }
            }
        }
    }
    return []int{-1, -1}
}
