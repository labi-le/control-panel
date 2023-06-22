package response

import (
	"github.com/gofiber/fiber/v2"
	"github.com/labi-le/control-panel/pkg/log"
	"net/http"
)

type Error struct {
	Error string `json:"error"`
}

type Success struct {
	Message string `json:"message"`
}

// Reply is a helper for responses
type Reply struct {
	l log.Logger
}

func New(l log.Logger) *Reply {
	return &Reply{l: l}
}

func request(ctx *fiber.Ctx, l log.Logger, status int, data interface{}) error {
	var (
		rStruct interface{}
		msg     string
	)

	//nolint:gocritic,gosimple //dn
	switch data.(type) {
	case error:
		rStruct = &Error{data.(error).Error()}
		msg = data.(error).Error()

	case Error:
		rStruct = &data
		msg = data.(Error).Error

	case Success:
		rStruct = &data
		msg = data.(Success).Message

	case any:
		rStruct = &data
		msg = ctx.String()
	}

	l.Debug(msg, " | ", http.StatusText(status))

	return ctx.
		Status(status).
		JSON(&rStruct)
}

func (r *Reply) BadRequest(ctx *fiber.Ctx, err error) error {
	return request(ctx, r.l, http.StatusBadRequest, err)
}

func (r *Reply) InternalServerError(ctx *fiber.Ctx, err error) error {
	return request(ctx, r.l, http.StatusInternalServerError, err)
}

func (r *Reply) OK(ctx *fiber.Ctx, data any) error {
	return request(ctx, r.l, http.StatusOK, data)
}

func (r *Reply) Created(ctx *fiber.Ctx, data any) error {
	return request(ctx, r.l, http.StatusCreated, data)
}

func (r *Reply) NoContent(ctx *fiber.Ctx, err error) error {
	return request(ctx, r.l, http.StatusNoContent, err)
}

func (r *Reply) NotFound(ctx *fiber.Ctx, err error) error {
	return request(ctx, r.l, http.StatusNotFound, err)
}

func (r *Reply) Unauthorized(ctx *fiber.Ctx, err error) error {
	return request(ctx, r.l, http.StatusUnauthorized, err)
}

func (r *Reply) Forbidden(ctx *fiber.Ctx, err error) error {
	return request(ctx, r.l, http.StatusForbidden, err)
}

func (r *Reply) Conflict(ctx *fiber.Ctx, err error) error {
	return request(ctx, r.l, http.StatusConflict, err)
}

func (r *Reply) UnprocessableEntity(ctx *fiber.Ctx, err error) error {
	return request(ctx, r.l, http.StatusUnprocessableEntity, err)
}

func (r *Reply) TooManyRequests(ctx *fiber.Ctx, err error) error {
	return request(ctx, r.l, http.StatusTooManyRequests, err)
}

func (r *Reply) ServiceUnavailable(ctx *fiber.Ctx, err error) error {
	return request(ctx, r.l, http.StatusServiceUnavailable, err)
}
