func maxSubArray(nums []int) int {
    currSum := 0
    maxSubArray := (-1) * 10000 * len(nums)
    for _, v := range nums {
        currSum += v
        maxSubArray =  max(currSum, maxSubArray)
        if currSum < 0 {
            currSum = 0
        }
    }
    return maxSubArray
}


// currSum is stale if negative
// currSum is a possible candidate if positive

// therefore, 
// if currSum is not negative, expand
// if currSum negative stop.
