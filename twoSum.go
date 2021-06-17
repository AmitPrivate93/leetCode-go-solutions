package main
import (
    "fmt"
)

func twoSum(nums []int, target int) []int {
    m:=make(map[int]int)
    for index1, value := range nums {
        if index2, found := m[target-value]; found{
            return []int{index2,index1}
        }
        m[value] = index1
    }
    return nil
}
func main(){
	fmt.Println(twoSum([]int{1,3,2,4,9,7,1},7))
}
