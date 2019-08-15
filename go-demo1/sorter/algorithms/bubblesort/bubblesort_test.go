package bubblesort

import "testing"

func TestBubbleSort1(t *testing.T){
	values :=[]int{1,3,2}
	BubbleSort(values)
	if values[0] != 1||values[1]!=2||values[2]!=3{
		t.Error("BubbleSort failed got ",values," Expected 1 2 3")
	}
}