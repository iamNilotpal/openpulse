package middlewares

import (
	"net/http"

	"github.com/iamNilotpal/openpulse/business/sys/validate"
	v1 "github.com/iamNilotpal/openpulse/business/web/v1"
	"github.com/iamNilotpal/openpulse/foundation/web"
	"go.uber.org/zap"
)

func ErrorResponder(log *zap.SugaredLogger) middleware {
	h := func(handler handler) http.HandlerFunc {
		h := func(w http.ResponseWriter, r *http.Request) {
			err := handler(w, r)

			var statusCode int
			var errResp web.APIError

			switch {
			case validate.IsFieldErrors(err):
				fieldErrors := validate.GetFieldErrors(err)
				errResp = web.APIError{
					Fields:    fieldErrors.Fields(),
					ErrorCode: v1.ToString(v1.InvalidInputErrorCode),
					Message:   http.StatusText(http.StatusBadRequest),
				}
				statusCode = http.StatusBadRequest

			case v1.IsRequestError(err):
				reqErr := v1.GetRequestError(err)
				errResp = web.APIError{
					Message:   reqErr.Error(),
					ErrorCode: v1.ToString(reqErr.Code),
				}
				statusCode = reqErr.Status

			default:
				errResp = web.APIError{
					ErrorCode: v1.ToString(v1.InternalErrorCode),
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

	return h
}
