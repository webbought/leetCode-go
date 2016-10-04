// Given an array of integers, return indices of the two numbers such that they add up to a specific target.

// You may assume that each input would have exactly one solution.

/**
 * Example :
 * 
 * Given nums = [2, 7, 11, 15], target = 9,
 * Because nums[0] + nums[1] = 2 + 7 = 9,
 * return [0, 1].
 * 
 */

func twoSum(nums []int, target int) []int {
    pair := make(map[int]int)
    result := make([]int, 2)
    for i := 0; i < len(nums); i++ {
        value, ok := pair[nums[i]]
        if ok {
            result[0] = value
            result[1] = i
        } else {
            pair[target - nums[i]] = i
        }
    }
    return result
}