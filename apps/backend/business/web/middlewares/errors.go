package middlewares

import (
	"encoding/json"
	stdErrors "errors"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/iamNilotpal/openpulse/business/sys/validate"
	"github.com/iamNilotpal/openpulse/business/web/errors"
	"github.com/iamNilotpal/openpulse/foundation/web"
	"go.uber.org/zap"
)

func ErrorResponder(log *zap.SugaredLogger) middleware {
	responder := func(handler handler) http.HandlerFunc {
		h := func(w http.ResponseWriter, r *http.Request) {
			err := handler(w, r)
			if err == nil {
				return
			}

			var statusCode int
			var errResp *web.APIError
			var syntaxError *json.SyntaxError

			switch {
			case stdErrors.As(err, &syntaxError):
				statusCode = 400
				errResp = web.NewAPIError(
					fmt.Sprintf("Request body contains badly-formed JSON (at position %d).", syntaxError.Offset),
					errors.FromErrorCode(errors.InvalidInput),
					nil,
				)

			case stdErrors.Is(err, io.ErrUnexpectedEOF):
				statusCode = 400
				errResp = web.NewAPIError(
					"Request body contains badly-formed JSON.",
					errors.FromErrorCode(errors.InvalidInput),
					nil,
				)

			case strings.HasPrefix(err.Error(), "json: unknown field "):
				statusCode = 400
				fieldName := strings.TrimPrefix(err.Error(), "json: unknown field ")
				msg := fmt.Sprintf("Request body contains unknown field %s.", fieldName)
				errResp = web.NewAPIError(msg, errors.FromErrorCode(errors.UnknownField), nil)

			case stdErrors.Is(err, io.EOF):
				statusCode = 400
				errResp = web.NewAPIError(
					"Request body must not be empty.", errors.FromErrorCode(errors.MissingRequiredFields), nil,
				)

			case validate.IsFieldErrors(err):
				fieldErrors := validate.GetFieldErrors(err)

				statusCode = http.StatusBadRequest
				errResp = web.NewAPIError(
					http.StatusText(http.StatusBadRequest),
					errors.FromErrorCode(errors.InvalidInput),
					fieldErrors.Fields(),
				)

			case errors.IsRequestError(err):
				reqErr := errors.GetRequestError(err)

				statusCode = reqErr.Status
				errResp = web.NewAPIError(
					reqErr.Error(),
					errors.FromErrorCode(reqErr.Code),
					nil,
				)

			default:
				statusCode = http.StatusInternalServerError
				errResp = web.NewAPIError(
					http.StatusText(http.StatusInternalServerError),
					errors.FromErrorCode(errors.InternalServerError),
					nil,
				)
			}

			if err = web.Error(w, statusCode, errResp); err != nil {
				log.Infow("API response sending error", "error", err)
			}
		}

		return h
	}

	return responder
}
