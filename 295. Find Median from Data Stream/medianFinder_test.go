package _95__Find_Median_from_Data_Stream

import "testing"

func TestMedianFinder(t *testing.T) {

	results := []float64{
		1.5,
		2.0,
		3.0,
	}

	obj := Constructor()
	obj.AddNum(1)
	obj.AddNum(2)
	i := 0
	if ret := obj.FindMedian(); ret != results[i] {
		t.Fatalf("case %d fails: %v\n", i, ret)
	}
	obj.AddNum(3)
	i++
	if ret := obj.FindMedian(); ret != results[i] {
		t.Fatalf("case %d fails: %v\n", i, ret)
	}
	obj.AddNum(4)
	obj.AddNum(5)
	i++
	if ret := obj.FindMedian(); ret != results[i] {
		t.Fatalf("case %d fails: %v\n", i, ret)
	}
}
