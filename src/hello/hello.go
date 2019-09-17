package main
import "fmt"
func quickSort(arr []int, left, right int) {
    if left >= right{
        return
    }
    start, end := left, right
    cur := arr[left]
    left++
    for left < right {
        for left < end && arr[left] <= cur {
            left++
        }
        for right > start && arr[right] >= cur {
            right--
        }
        if left < right{
            arr[left], arr[right] = arr[right], arr[left]
        }
	}
	if arr[right] < arr[start] {
		arr[start], arr[right] = arr[right], arr[start]
	}
    quickSort(arr, start, right - 1)
    quickSort(arr, right + 1, end)
}

func main(){
    test := []int{1,2,3,4,5,6,7}
    quickSort(test, 0, len(test) - 1)
    for _,j := range test{
        fmt.Println(j)
    }
}