package middlewares

import (
	"net/http"

	"github.com/iamNilotpal/openpulse/business/sys/validate"
	"github.com/iamNilotpal/openpulse/business/web/errors"
	"github.com/iamNilotpal/openpulse/foundation/web"
	"go.uber.org/zap"
)

func ErrorResponder(log *zap.SugaredLogger) middleware {
	responder := func(handler handler) http.HandlerFunc {
		h := func(w http.ResponseWriter, r *http.Request) {
			err := handler(w, r)

			var statusCode int
			var errResp web.APIError

			switch {
			case validate.IsFieldErrors(err):
				fieldErrors := validate.GetFieldErrors(err)
				errResp = web.APIError{
					Fields:    fieldErrors.Fields(),
					ErrorCode: errors.CodeToString(errors.InvalidInput),
					Message:   http.StatusText(http.StatusBadRequest),
				}
				statusCode = http.StatusBadRequest

			case errors.IsRequestError(err):
				reqErr := errors.GetRequestError(err)
				errResp = web.APIError{
					Message:   reqErr.Error(),
					ErrorCode: errors.CodeToString(reqErr.Code),
				}
				statusCode = reqErr.Status

			default:
				errResp = web.APIError{
					ErrorCode: errors.CodeToString(errors.InternalServerError),
					Message:   http.StatusText(http.StatusInternalServerError),
				}
				statusCode = http.StatusInternalServerError
			}

			if err = web.Error(w, statusCode, errResp); err != nil {
				log.Infow("API response sending error", "error", err)
			}
		}

		return h
	}

	return responder
}
