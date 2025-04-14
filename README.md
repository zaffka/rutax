# Пакет валидации Российского ИНН (TaxID)

[![Go Reference](https://pkg.go.dev/badge/github.com/zaffka/rutax.svg)](https://pkg.go.dev/github.com/zaffka/rutax)
[![Go Report Card](https://goreportcard.com/badge/github.com/zaffka/rutax)](https://goreportcard.com/report/github.com/zaffka/rutax)
[![Tests](https://github.com/zaffka/rutax/actions/workflows/tests.yaml/badge.svg)](https://github.com/zaffka/rutax/actions/workflows/tests.yaml)

Пакет `rutax` предоставляет функциональность для валидации и парсинга ИНН (Идентификационного Номера Налогоплательщика, TaxID) Российской Федерации.

## Особенности

- Парсинг из строки и валидация ИНН физических и юридических лиц
- Проверка контрольных сумм согласно официальному алгоритму


## Установка

```bash
go get github.com/zaffka/rutax@latest
```

## Использование
#### `rutax.ParseID(innStr string) (rutax.ID, error)`

Основная функция пакета:
1. Проверяет соответствие строки формату ИНН
1. Проверяет контрольные суммы в ключевых позициях разбираемой строки
1. Возвращает структуру с данными ИНН

Возвращаемая структура `ID` содержит:
- `Num` - валидированный номер ИНН
- `IsLegal` - флаг принадлежности к юридическому лицу


Пример:
```go
package main

import (
	"fmt"
	"github.com/zaffka/rutax"
)

func main() {
	// Парсинг ИНН
	id, err := rutax.ParseID("7710140679")
	if err != nil {
		fmt.Println("Ошибка:", err)

		return
	}

	fmt.Printf("ИНН: %s, Юр.лицо: %v\n", id.Num, id.IsLegal)
}
```

### Ошибки

Пакет возвращает следующие типы ошибок:
- `ErrIDIncorrect` - ошибка формата ИНН (длина, символы чисел)
- `ErrChecksumFailed` - ошибка контрольной суммы


## Лицензия

MIT