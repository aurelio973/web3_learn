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
