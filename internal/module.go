package internal

import (
	"ads/internal/ad"
	"go.uber.org/fx"
)

var Module = fx.Options(
	ad.Module,
)
