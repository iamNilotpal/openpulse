package middlewares

import (
	"encoding/json"
	stdErrors "errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"syscall"

	"github.com/iamNilotpal/openpulse/business/web/errors"
	"github.com/iamNilotpal/openpulse/foundation/validate"
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

			if IsResponseWriteError(err) {
				log.Infow("API response sending error", "error", err)
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
					"Error while processing request.",
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

func IsResponseWriteError(err error) bool {
	// Ignore syscall.EPIPE and syscall.ECONNRESET errors which occurs
	// when a write operation happens on the http.ResponseWriter that
	// has simultaneously been disconnected by the client (TCP
	// connections is broken). For instance, when large amounts of
	// data is being written or streamed to the client.
	// https://blog.cloudflare.com/the-complete-guide-to-golang-net-http-timeouts/
	// https://gosamples.dev/broken-pipe/
	// https://gosamples.dev/connection-reset-by-peer/

	switch {
	case stdErrors.Is(err, syscall.EPIPE):

		// Usually, you get the broken pipe error when you write to the connection after the
		// RST (TCP RST Flag) is sent.
		// The broken pipe is a TCP/IP error occurring when you write to a stream where the
		// other end (the peer) has closed the underlying connection. The first write to the
		// closed connection causes the peer to reply with an RST packet indicating that the
		// connection should be terminated immediately. The second write to the socket that
		// has already received the RST causes the broken pipe error.
		return true

	case stdErrors.Is(err, syscall.ECONNRESET):

		// Usually, you get connection reset by peer error when you read from the
		// connection after the RST (TCP RST Flag) is sent.
		// The connection reset by peer is a TCP/IP error that occurs when the other end (peer)
		// has unexpectedly closed the connection. It happens when you send a packet from your
		// end, but the other end crashes and forcibly closes the connection with the RST
		// packet instead of the TCP FIN, which is used to close a connection under normal
		// circumstances.
		return true
	}

	return false
}
