package utils

import (
	"crypto/rand"
	"math/big"
)

func GenerateBookingNumber() (*big.Int, error) {
	// Generate random 8-digit number (10000000 to 99999999)
	min := int64(10000000)
	max := int64(99999999)

	n, err := rand.Int(rand.Reader, big.NewInt(max-min+1))
	if err != nil {
		return big.NewInt(0), err
	}

	bookingNo := big.NewInt(n.Int64() + min)

	return bookingNo, nil
}
