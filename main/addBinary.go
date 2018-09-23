package main

import "fmt"

/**
"10100000100100110110010000010101111011011001101110111111111101000000101111001110001111100001101"
"110101001011101110001111100110001010100001101011101010000011011011001011101111001100000011011110011"


Given two binary strings, return their sum (also a binary string).

The input strings are both non-empty and contains only characters 1 or 0.

Example 1:

Input: a = "11", b = "1"
Output: "100"
Example 2:

Input: a = "1010", b = "1011"
Output: "10101"

 */

func addBinaryMy(a string, b string) string {
	i := len(a) - 1
	j := len(b) - 1
	var s string
	var c uint8
	for {
		if i >= 0 || j >= 0 {
			c /= 2
			if (i >= 0) {
				c += a[i] - '0';
				i--;
			}
			if (j >= 0) {
				c += b[j] - '0';
				j--;
			}
			if c%2 == 0 {
				s = "0" + s
			} else {
				s = "1" + s
			}

		} else {
			break
		}
	}
	if c > 1 {
		s = "1" + s
	}
	return s
}

func addBinary(a string, b string) string {

	IfAdd := 48

	aToByte := []byte(a)
	bToByte := []byte(b)

	aLength := len(aToByte)
	bLength := len(bToByte)

	theSameLength := 0

	var restPart []byte

	if aLength > bLength {

		restPart = aToByte[:aLength-bLength]
		aToByte = aToByte[aLength-bLength:]
		theSameLength = bLength

	} else if aLength < bLength {
		restPart = bToByte[:bLength-aLength]
		bToByte = bToByte[bLength-aLength:]
		theSameLength = aLength
	} else {
		theSameLength = aLength
	}

	for i := theSameLength - 1; i >= 0; i-- {
		switch {
		case aToByte[i] == 48 && bToByte[i] == 48 && IfAdd == 49:
			aToByte[i] = 49
			IfAdd = 48
		case aToByte[i] == 48 && bToByte[i] == 49 && IfAdd == 49:
			aToByte[i] = 48
			IfAdd = 49
		case aToByte[i] == 49 && bToByte[i] == 48 && IfAdd == 49:
			aToByte[i] = 48
			IfAdd = 49
		case aToByte[i] == 49 && bToByte[i] == 49 && IfAdd == 49:
			aToByte[i] = 49
			IfAdd = 49
		case aToByte[i] == 48 && bToByte[i] == 48 && IfAdd == 48:
			aToByte[i] = 48
			IfAdd = 48
		case aToByte[i] == 48 && bToByte[i] == 49 && IfAdd == 48:
			aToByte[i] = 49
			IfAdd = 48
		case aToByte[i] == 49 && bToByte[i] == 48 && IfAdd == 48:
			aToByte[i] = 49
			IfAdd = 48
		case aToByte[i] == 49 && bToByte[i] == 49 && IfAdd == 48:
			aToByte[i] = 48
			IfAdd = 49
		}
	}

	for i := len(restPart) - 1; i >= 0; i-- {
		switch {
		case restPart[i] == 48 && IfAdd == 48:
			restPart[i] = 48
			IfAdd = 48
		case restPart[i] == 48 && IfAdd == 49:
			restPart[i] = 49
			IfAdd = 48
		case restPart[i] == 49 && IfAdd == 48:
			restPart[i] = 49
			IfAdd = 48
		case restPart[i] == 49 && IfAdd == 49:
			restPart[i] = 48
			IfAdd = 49
		}
	}

	restPart = append(restPart, aToByte...)

	if IfAdd == 49 {
		restPart = append([]byte{49}, restPart...)
		return string(restPart)
	}

	return string(restPart)

}

func main() {
	fmt.Println(addBinary("0", "1"))
}
