package sort

func SelectSort(nums []int){
	for i:=0;i<len(nums);i++{
		for j:=i+1;j<len(nums);j++{
			if nums[i]>nums[j]{
				nums[i],nums[j] = nums[j],nums[i]
			}
		}
	}
}

func BubbleSort(nums []int){
	for i:=0;i<len(nums);i++{
		for j:=0;j<len(nums)-i-1;j++{
			if nums[j]<nums[j+1]{
				nums[j],nums[j+1] = nums[j+1],nums[j]
			}
		}
	}
}

func InsertSort(nums []int){
	//todo
}