package httputils

import (
	"drunklish/internal/service/word/dto"
	"encoding/csv"
	"errors"
	"fmt"
	"os"
)

var (
	ErrOpenFile = errors.New("fail with open file")
	ErrReadAll  = errors.New("fail with read all")
)

func TryCSV(fileName string) ([][]string, error) {
	file, err := os.OpenFile(fileName, os.O_RDONLY, os.ModePerm)

	if err != nil {
		return nil, fmt.Errorf("%w", ErrOpenFile)
	}

	defer file.Close()

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("%w", ErrReadAll)
	}

	return records, nil
}

func Parse(records [][]string) (dto.RequestFromUpload, error) {
	var response dto.RequestFromUpload
	var resWord dto.ResponseWord

	for i := 0; i < len(records); i++ {
		resWord.Word = records[i][0]
		resWord.Translate = records[i][1]
		response.Words.Words = append(response.Words.Words, resWord)
	}

	//for i := range records {
	//	response.Words.Words[i].Word = records[i][0]
	//	response.Words.Words[i].Translate = records[i][1]
	//}

	return response, nil
}
