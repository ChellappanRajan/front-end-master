package main 

import "testing"

func TestBinart(t *testing.T){
	
	nums := []int{1,2,3,4}
	target := 0
	result := LinearSearch(nums,target); 

	if(result != 0){
	  t.Errorf("Failed")
	}

}
