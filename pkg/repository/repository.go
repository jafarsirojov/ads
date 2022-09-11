package repository

import (
	"ads/internal/contract"
	"ads/internal/interfaces"
	"ads/pkg/config"
	"ads/pkg/db"
	"context"
	"github.com/jackc/pgx/v4"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

var Module = fx.Provide(NewRepo)

type repo struct {
	logger *zap.Logger
	config *config.Config
	db     db.Querier
}

type Params struct {
	fx.In
	Logger *zap.Logger
	Config *config.Config
	DB     db.Querier
}

func NewRepo(params Params) interfaces.AdRepo {
	return &repo{
		logger: params.Logger,
		config: params.Config,
		db:     params.DB,
	}
}

func (r *repo) Add(ctx context.Context, ad contract.Ad) (contract.Ad, error) {

	//txx, err := r.db.Begin(ctx)
	err := r.db.QueryRow(ctx,
		`INSERT INTO ad(
				title, 
				description, 
				price
			) VALUES(
				$1, $2, $3
			) RETURNING
				id, 
				title, 
				description, 
				price, 
				created_at
				;`,
		ad.Title,
		ad.Description,
		ad.Price,
	).Scan(
		&ad.ID,
		&ad.Title,
		&ad.Description,
		&ad.Price,
		&ad.CreatedAt,
	)

	if err != nil {
		r.logger.Error("pkg.repository.Add r.db.QueryRow",
			zap.Any("ad", ad), zap.Error(err))
		return ad, err
	}

	return ad, nil
}

func (r *repo) GetList(ctx context.Context) (ads []contract.Ad, err error) {
	rows, err := r.db.Query(ctx,
		`SELECT 
			id, 
			title, 
			description, 
			price, 
			created_at
		FROM ad
		WHERE 1=1 `)
	if err != nil {
		r.logger.Error("pkg.repository.Add r.db.Query", zap.Error(err))
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var ad contract.Ad
		err = rows.Scan(
			&ad.ID,
			&ad.Title,
			&ad.Description,
			&ad.Price,
			&ad.CreatedAt,
		)
		if err != nil {
			r.logger.Error("pkg.repository.Add rows.Scan", zap.Error(err))
			return nil, err
		}

		ads = append(ads, ad)
	}

	if len(ads) == 0 {
		return nil, contract.ErrNotFound
	}

	return ads, nil
}

func (r *repo) GetByID(ctx context.Context, id int) (ad contract.Ad, err error) {
	err = r.db.QueryRow(ctx,
		`SELECT 
			id, 
			title, 
			description, 
			price, 
			created_at
		FROM ad
		WHERE id = $1`,
		id).
		Scan(
			&ad.ID,
			&ad.Title,
			&ad.Description,
			&ad.Price,
			&ad.CreatedAt,
		)

	if err != nil {
		if err == pgx.ErrNoRows {
			return ad, contract.ErrNotFound
		}
		r.logger.Error("pkg.repository.GetByID r.db.QueryRow",
			zap.Int("id", id), zap.Error(err))
		return ad, err
	}

	return ad, nil
}
