package main

// import "fmt"

// 1.只出现一次的数字
// 方法1：for循环结合if和map
func singleNumber(nums []int) int {
    count:=make(map[int]int)
    for _,j := range(nums){
        count[j]++
    }
    for j,times :=range(count){
        if times == 1{
        return j
        }
    }
    return 0
}

// 方法2：for循环结合异或运算
func singleNumber(nums []int) int {
    result:=0
    for _,j :=range(nums){
        result^=j
    }
    return result
}

// 方法3：排序+分情况判断
func singleNumber(nums []int) int {
    sort.Ints(nums)
    // 只有一个元素时
    if len(nums)==1{
        return nums[0]
    }

    // 单个数字在开头
    if nums[0]!=nums[1]{
        return nums[0]
    }
    // 单个数字在结尾
    if nums[len(nums)-1]!=nums[len(nums)-2]{
        return nums[len(nums)-1]
    }
    // 单个数字在中间
    for i :=1;i<len(nums)-1;i++{
        if nums[i]!=nums[i-1] && nums[i]!=nums[i+1]{
            return nums[i]
        }
    }
    return 0
}

//2.回文数
// 方法1：翻转判断
func isPalindrome(x int) bool {
    if x<0 ||(x!=0 && x%10==0){
        return false
    }
    original :=x
    reverse:=0
    for x!=0{
        reverse = reverse*10+x%10
        x/=10
    }
    return original == reverse
}

// 方法2：转换成字符串判断
func isPalindrome(x int) bool {
    s:=fmt.Sprintf("%d",x)
    for i:=0;i<len(s)/2;i++{
        if s[i]!=s[len(s)-1-i]{
            return false
        }
    }
    return true
}

//3.有效的括号
// 方法1：栈 switch判断
func isValid(s string) bool {
    left:=[]rune{}
    right:=map[rune]rune{')':'(',']':'[','}':'{'} 
    for _,i :=range s {
        switch i {
            case '(','[','{':
                left = append(left,i)
            case ')',']','}':
                if len(left)==0|| left[len(left)-1]!=right[i]{
                    return false
                }
        left = left[0:len(left)-1]
        }
    }   
    return len(left)==0
}

// 方法2：栈 if/else判断
func isValid(s string) bool {
    left:=[]rune{}
    for _,i :=range s {
        if i =='('||i == '['|| i== '{'{
            left = append(left,i)
        } else if i == ')' {
            if len(left) == 0 || left[len(left)-1]!='('{
                return false
            }
            left = left[0:len(left)-1]
        } else if i == ']' {
            if len(left) == 0 || left[len(left)-1]!='['{
                return false
            }
            left = left[0:len(left)-1]
        } else if i == '}' {
            if len(left) == 0 || left[len(left)-1]!='{'{
                return false
            }
            left = left[0:len(left)-1]
        }
    
    }
    return len(left)==0
}

// 4.最长公共前缀
// 方法1：for循环纵向比较
func longestCommonPrefix(strs []string) string {
    if len(strs) == 0{
        return ""
    }
    for i :=0;i<len(strs[0]);i++ {
        s:=strs[0][i]
        for j :=1;j<len(strs);j++ {
            if i >=len(strs[j]) || strs[j][i]!= s {
                return strs[0][:i]
            }
        }
    }
    return strs[0]
}

// 方法2：for循环横向比较
func longestCommonPrefix(strs []string) string {
    if len(strs) == 0{
        return ""
    }
    pre := strs[0]
    for i :=1;i<len(strs);i++ {
        var j int
        for j =0;j<len(pre) && j < len(strs[i]);j++ {
            if pre[j]!= strs[i][j] {
                break
            }
        }
        pre = pre[0:j]
        if pre =="" {
            return ""
        }
    }
    return pre
}

// 5.加一
func plusOne(digits []int) []int {
    for i :=len(digits)-1;i>=0;i-- {
        if digits[i] < 9 {
            digits[i]++
            return digits
        }
        digits[i] = 0
    }
    return append([]int{1},digits...)
}

//6.删除有序数组中的重复项
// 方法1：嵌套循环查找
func removeDuplicates(nums []int) int {
    n:=len(nums)
    if len(nums) == 0 {
        return 0
    }
    for i:=0;i<n-1;i++ {
        for j:=i+1;j<n; {
            if nums[i] == nums[j] {
                for k:=j;k<n-1;k++ {
                    nums[k] = nums[k+1]
                }
                n--
            } else {
                j++
            }

        }
    }
    return n
}

// 方法2：
func removeDuplicates(nums []int) int {
    if len(nums) == 0 {
        return 0
    }
    i:= 0
    for j :=1;j<len(nums);j++ {
        if nums[i]!= nums[j]{
            i++
            nums[i] = nums[j]
        }
    }
    return i+1
}

//7.合并区间
// 方法1：单循环比较 边遍历边合并
func merge(intervals [][]int) [][]int {
    if len(intervals) ==0 {
        return intervals
    }
    // 区间排序
    sort.Slice(intervals,func(i,j int)bool {
        return intervals[i][0]<intervals[j][0]
    })
    result := [][]int{intervals[0]}
    for i :=1;i<len(intervals);i++ {
        last :=result[len(result)-1]
        start :=intervals[i]

            if start[0]<=last[1] {
                if start[1]>last[1]{
                    last[1]=start[1]
                }
            } else {
                result = append(result,start)
            }
    }
    return result
}

// 方法2：双循环比较
func merge(intervals [][]int) [][]int {
    if len(intervals) ==0 {
        return intervals
    }
    // 区间排序
    sort.Slice(intervals,func(i,j int)bool {
        return intervals[i][0]<intervals[j][0]
    })
    var result [][]int
    i := 0

    for i < len(intervals) {
        left := intervals[i][0]
        right :=intervals[i][1]

        j:= i+1
        for j <len(intervals) && intervals[j][0] <= right {
            if intervals[j][1] > right {
                right = intervals[j][1]
            }
            j++
        }
        result = append(result,[]int{left,right})
        i = j
    }
    return result
}

// 8.两数之和
//方法1：嵌套循环
func twoSum(nums []int, target int) []int {
    for i:=0;i<len(nums);i++ {
        for j := i+1;j<len(nums);j++ {
            if nums[i] + nums [j] ==target{
                return []int{i,j}
            }
        }
    }
    return []int{}
}
//方法2：map
func twoSum(nums []int, target int) []int {
    numMap:=make(map[int]int)

    for i,num :=range nums {
        s :=target-num
        if index,exists := numMap[s]; exists {
            return []int{index,i}
        }
        numMap[num] = i
    }
    return nil
}

