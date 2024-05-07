package fileops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func WriteFloatTofile(value float64, fileName string) {

	valueText := fmt.Sprint(value)
	os.WriteFile(fileName, []byte(valueText), 0644)
}

func ReadFloatFile(fileName string) (float64, error) {

	data, err := os.ReadFile(fileName)

	if err != nil {
		return 1000, errors.New("error al leer el archivo, no se encuentra o se llama diferente")
	}

	valueText := string(data)
	value, err := strconv.ParseFloat(valueText, 64)

	if err != nil {
		return 1000, errors.New("error al convertir el archivo")
	}

	return value, nil
}
