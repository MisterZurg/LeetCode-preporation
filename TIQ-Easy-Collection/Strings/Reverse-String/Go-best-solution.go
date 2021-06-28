package Reverse_String

func reverseStringBest(s []byte)  {
	reversed := make([]byte, 0)
	for i := len(s)-1; i >= 0; i-- {
		reversed = append(reversed,s[i])
	}

	for i, val := range reversed {
		s[i] = val
	}
}
