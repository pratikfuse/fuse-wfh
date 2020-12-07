package sheet

import (
	"context"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

type Sheet struct {
	Service *sheets.Service
}

func GetInstance(apiKey string) (Sheet, error) {
	sheetService := Sheet{}
	ctx := context.Background()
	s, err := sheets.NewService(ctx, option.WithAPIKey(apiKey))

	if err != nil {
		return sheetService, err
	}
	sheetService.Service = s
	return sheetService, nil
}
