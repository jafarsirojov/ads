package interfaces

import (
	"ads/internal/contract"
	"context"
)

type AdRepo interface {
	Add(ctx context.Context, ad contract.Ad) (contract.Ad, error)
	GetList(ctx context.Context, offset int, priceSort, dateSort string) ([]contract.AdFromList, error)
	GetByID(context.Context, int) (contract.Ad, error)
}
