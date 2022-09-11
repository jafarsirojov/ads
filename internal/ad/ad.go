package ad

import (
	"ads/internal/contract"
	"ads/internal/interfaces"
	"ads/pkg/utils"
	"context"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

var Module = fx.Provide(New)

type Params struct {
	fx.In
	Logger *zap.Logger
	AdRepo interfaces.AdRepo
}

type service struct {
	logger *zap.Logger
	adRepo interfaces.AdRepo
}

func New(params Params) AdsService {
	return &service{
		logger: params.Logger,
		adRepo: params.AdRepo,
	}
}

type AdsService interface {
	Add(ctx context.Context, ad contract.Ad) (contract.Ad, error)
	GetList(ctx context.Context, offset int, sortByPrice, sortByDate string) (ads []contract.AdFromList, err error)
	GetByID(ctx context.Context, id int) (contract.Ad, error)
}

func (s *service) Add(ctx context.Context, ad contract.Ad) (contract.Ad, error) {
	ad, err := s.adRepo.Add(ctx, ad)
	if err != nil {
		s.logger.Error("internal.ad.Add s.adRepo.Add", zap.Any("ad", ad), zap.Error(err))
		return ad, err
	}

	return ad, nil
}

func (s *service) GetList(ctx context.Context, offset int, sortByPrice,
	sortByDate string) (ads []contract.AdFromList, err error) {

	var (
		priceSort string
		dateSort  string
	)

	if utils.InArray(sortByPrice, sortType) {
		priceSort = sortByPrice
	}

	if utils.InArray(sortByDate, sortType) {
		dateSort = sortByDate
	}

	ads, err = s.adRepo.GetList(ctx, offset, priceSort, dateSort)
	if err != nil {
		s.logger.Error("internal.ad.GetList s.adRepo.GetList", zap.Error(err))
		return nil, err
	}

	return ads, nil
}

func (s *service) GetByID(ctx context.Context, id int) (ad contract.Ad, err error) {
	ad, err = s.adRepo.GetByID(ctx, id)
	if err != nil {
		s.logger.Error("internal.ad.GetByID s.adRepo.GetByID", zap.Error(err))
		return ad, err
	}

	return ad, nil
}

var sortType = []string{"asc", "desc"}
