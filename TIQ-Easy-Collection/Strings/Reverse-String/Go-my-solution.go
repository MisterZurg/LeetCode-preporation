package Reverse_String

func reverseString(s []byte)  {
	for ind:=0; ind < len(s) / 2 ; ind++ {
		tmp := s[ind]
		s[ind] = s[len(s) - 1 - ind]
		s[len(s) - 1 - ind] = tmp
	}
}
