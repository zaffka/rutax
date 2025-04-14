package rutax

import (
	"errors"
	"fmt"
	"regexp"
)

const (
	personalLen = 12 // Длина ИНН для физического лица или ИП.
	legalLen    = 10 // Длина ИНН для юридического лица.
)

// ID представляет ИНН с информацией о его типе.
type ID struct {
	Num     string `json:"num,omitempty"`      // Номер ИНН.
	IsLegal bool   `json:"is_legal,omitempty"` // Является ли юридическим лицом.
}

var (
	// Регулярное выражение для проверки формата ИНН.
	taxIDexp = regexp.MustCompile(`^\d{10}(\d{2})?$`)

	// Весовые коэффициенты для расчета контрольных сумм.
	weights = []int32{3, 7, 2, 4, 10, 3, 5, 9, 4, 6, 8}

	// Ошибки валидации.
	ErrIDIncorrect    = errors.New("некорректный формат ИНН")
	ErrChecksumFailed = errors.New("ошибка контрольной суммы")
)

// ParseID парсит и валидирует ИНН.
func ParseID(taxID string) (ID, error) {
	// Проверка базового формата.
	if !taxIDexp.MatchString(taxID) {
		return ID{}, fmt.Errorf("%w: не соответствует формату", ErrIDIncorrect)
	}

	runes := []rune(taxID)
	length := len(runes)

	id := ID{Num: taxID}

	switch length {
	case personalLen:
		// Проверка первой контрольной цифры (11-я позиция).
		sum := checksum(runes[:length-2], weights[1:])
		if runes[length-2]-'0' != sum {
			return ID{}, fmt.Errorf("%w: первая контрольная сумма не совпадает", ErrChecksumFailed)
		}

		// Проверка второй контрольной цифры (12-я позиция).
		sum = checksum(runes[:length-1], weights)
		if runes[length-1]-'0' != sum {
			return ID{}, fmt.Errorf("%w: вторая контрольная сумма не совпадает", ErrChecksumFailed)
		}

	case legalLen:
		// Проверка контрольной цифры для юр.лица (10-я позиция).
		sum := checksum(runes[:length-1], weights[2:])
		if runes[length-1]-'0' != sum {
			return ID{}, fmt.Errorf("%w: контрольная сумма не совпадает", ErrChecksumFailed)
		}
		id.IsLegal = true
	}

	return id, nil
}

// checksum вычисляет контрольную сумму для указанных рун и весовых коэффициентов.
func checksum(runes []rune, weights []int32) int32 {
	var sum int32
	for pos, char := range runes {
		sum += (char - '0') * weights[pos]
	}

	return sum % 11 % 10
}
