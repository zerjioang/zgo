package strconv

import (
	"fmt"
	"testing"
)

func TestByteArrayToInt(t *testing.T) {
	// integer for convert
	num := int64(1354321354812)
	fmt.Println("Original number:", num)

	// integer to byte array
	byteArr := IntToByteArray(num)
	fmt.Println("Byte Array", byteArr)

	// byte array to integer again
	numAgain := ByteArrayToInt(byteArr)
	fmt.Println("Converted number:", numAgain)
}
