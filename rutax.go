package rutax

import (
	"errors"
	"fmt"
	"regexp"
)

const (
	personalLen = 12 // Length of TIN(Tax Identification Number) for individuals or sole proprietors.
	legalLen    = 10 // Length of TIN(Tax Identification Number) for legal entities.
)

// ID represents a Tax Identification Number (TIN) with information about its type.
type ID struct {
	Num     string `json:"num,omitempty"`      // TIN number.
	IsLegal bool   `json:"is_legal,omitempty"` // Whether it belongs to a legal entity.
}

var (
	// Regular expression for TIN format validation.
	taxIDexp = regexp.MustCompile(`^\d{10}(\d{2})?$`)

	// Weight coefficients for checksum calculation.
	weights = []int32{3, 7, 2, 4, 10, 3, 5, 9, 4, 6, 8}

	// Validation errors.
	ErrIDIncorrect    = errors.New("incorrect TIN format")
	ErrChecksumFailed = errors.New("checksum error")
)

// ParseID parses and validates a TIN.
func ParseID(taxID string) (ID, error) {
	// Basic format check.
	if !taxIDexp.MatchString(taxID) {
		return ID{}, fmt.Errorf("%w: format mismatch", ErrIDIncorrect)
	}

	runes := []rune(taxID)
	length := len(runes)

	id := ID{Num: taxID}

	switch length {
	case personalLen:
		// First checksum digit verification (11th position).
		sum := checksum(runes[:length-2], weights[1:])
		if runes[length-2]-'0' != sum {
			return ID{}, fmt.Errorf("%w: first checksum mismatch", ErrChecksumFailed)
		}

		// Second checksum digit verification (12th position).
		sum = checksum(runes[:length-1], weights)
		if runes[length-1]-'0' != sum {
			return ID{}, fmt.Errorf("%w: second checksum mismatch", ErrChecksumFailed)
		}

	case legalLen:
		// Checksum digit verification for legal entities (10th position).
		sum := checksum(runes[:length-1], weights[2:])
		if runes[length-1]-'0' != sum {
			return ID{}, fmt.Errorf("%w: checksum mismatch", ErrChecksumFailed)
		}
		id.IsLegal = true
	}

	return id, nil
}

// checksum calculates the checksum for given runes and weight coefficients.
func checksum(runes []rune, weights []int32) int32 {
	var sum int32
	for pos, char := range runes {
		sum += (char - '0') * weights[pos]
	}

	return sum % 11 % 10
}
