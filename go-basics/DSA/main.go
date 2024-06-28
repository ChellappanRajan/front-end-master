package main
import "fmt"


func main(){

	myNums := []int{1,2,3,4,5}
	target := 9
fmt.Print(LinearSearch(myNums,target))

}

func LinearSearch(nums []int, target int) int{
	for i, num := range nums {
       if(num == target) {
			return i
	   }
	}
	return -1
}
