package handler

import (
	"ads/internal/ad"
	"ads/internal/contract"
	"ads/internal/responses"
	"ads/pkg/reply"
	"encoding/json"
	"github.com/gorilla/mux"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"net/http"
	"strconv"
)

var Module = fx.Provide(New)

type Handler interface {
	Add(http.ResponseWriter, *http.Request)
	GetList(http.ResponseWriter, *http.Request)
	GetByID(http.ResponseWriter, *http.Request)
}

type Params struct {
	fx.In
	Logger     *zap.Logger
	AdsService ad.AdsService
}

type handler struct {
	logger     *zap.Logger
	AdsService ad.AdsService
}

func New(params Params) Handler {
	return &handler{
		logger:     params.Logger,
		AdsService: params.AdsService,
	}
}

func (h *handler) Add(w http.ResponseWriter, r *http.Request) {

	var (
		request  contract.Ad
		response contract.Response
	)

	defer reply.Json(w, http.StatusOK, &response)

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		h.logger.Error("cmd.app.handler.Add json.NewDecoder(r.Body).Decode(&request)",
			zap.Any("request", request), zap.Any("r.Body", r.Body), zap.Error(err))
		response = responses.BadRequest
		return
	}

	if len(request.Title) > 200 {
		h.logger.Error("cmd.app.handler.Add request title > 200",
			zap.Any("request", request), zap.Int("len", len(request.Title)), zap.Error(err))
		response = responses.BadRequest
		return
	}

	if len(request.Description) > 1000 {
		h.logger.Error("cmd.app.handler.Add request description > 1000",
			zap.Any("request", request), zap.Int("len", len(request.Title)), zap.Error(err))
		response = responses.BadRequest
		return
	}

	if len(request.LinksToPhotos) > 3 || len(request.LinksToPhotos) == 0 {
		h.logger.Error("cmd.app.handler.Add request LinksToPhotos incorrect",
			zap.Any("request", request), zap.Int("count", len(request.LinksToPhotos)), zap.Error(err))
		response = responses.BadRequest
		return
	}

	newAd, err := h.AdsService.Add(r.Context(), request)
	if err != nil {
		h.logger.Error("cmd.app.handler.Add h.AdsService.Add(r.Context(), request)",
			zap.Any("request", request), zap.Error(err))
		response = responses.InternalErr
		return
	}

	response = responses.Success
	response.Payload = newAd
}

func (h *handler) GetList(w http.ResponseWriter, r *http.Request) {

	var (
		offsetStr   = r.Header.Get("offset")
		sortByPrice = r.Header.Get("sortByPrice")
		sortByDate  = r.Header.Get("sortByDate")

		response contract.Response
		offset   int
		err      error
	)

	defer reply.Json(w, http.StatusOK, &response)

	if offset, err = strconv.Atoi(offsetStr); err != nil || offset < 0 {
		offset = 0
	}

	list, err := h.AdsService.GetList(r.Context(), offset, sortByPrice, sortByDate)
	if err != nil {
		if err == contract.ErrNotFound {
			response = responses.NotFound
			return
		}

		h.logger.Error("cmd.app.handler.Add h.AdsService.GetList", zap.Error(err))
		response = responses.InternalErr
		return
	}

	response = responses.Success
	response.Payload = list
}

func (h *handler) GetByID(w http.ResponseWriter, r *http.Request) {
	var (
		idStr    = mux.Vars(r)["id"]
		response contract.Response

		id  int
		err error
	)

	defer reply.Json(w, http.StatusOK, &response)

	if id, err = strconv.Atoi(idStr); err != nil || id == 0 {
		h.logger.Error("cmd.app.handler.GetByID incorrect id in params",
			zap.String("idStr", idStr), zap.Error(err))
		response = responses.BadRequest
		return
	}

	payload, err := h.AdsService.GetByID(r.Context(), id)
	if err != nil {
		h.logger.Error("cmd.app.handler.GetByID h.AdsService.GetByID",
			zap.Int("id", id), zap.Error(err))
		response = responses.BadRequest
		return
	}

	response = responses.Success
	response.Payload = payload
}
