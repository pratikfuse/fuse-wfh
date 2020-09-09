package sheet

import (
	//sheets	"google.golang.org/api/sheets/v4"
	"context"
	"google.golang.org/api/sheets/v4"
)

type sheet struct {
	service *sheets.Service
}

func (s *sheet) GetInstance() error {
	ctx := context.Background()

	sheetService, err := sheets.NewService(ctx)

	if err != nil {
		return nil
	}
	s.service = sheetService

	return nil
}
