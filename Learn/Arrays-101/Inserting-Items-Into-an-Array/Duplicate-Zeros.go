package Inserting_Items_Into_an_Array

/*
func duplicateZeros(arr []int)  {
    for i:= 0; i < len(arr); i++ {
        if arr[i] == 0 {
            copy(arr[i+1:], arr[i:])
            i++
        }
    }
}
*/

func duplicateZeros(arr []int) {
	zeros := 0
	// Count zeros
	for _, number := range arr {
		if number == 0 {
			zeros++
		}
	}
	// [1 0 2]
	for i := len(arr) - 1; i >= 0; i-- {
		if arr[i] == 0 {
			if i+zeros < len(arr) {
				arr[i+zeros] = 0
			}
			zeros--
		}
		if i+zeros < len(arr) {
			arr[i+zeros] = arr[i]
		}
	}
}
