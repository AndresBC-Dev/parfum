package middleware

import (
	"fmt"
	"log"
	"net/http"

	"github.com/fatih/color"
)

type statusRecorder struct {
	http.ResponseWriter
	status int
}

func (r *statusRecorder) WriteHeader(status int) {
	r.status = status
	r.ResponseWriter.WriteHeader(status)
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		recorder := &statusRecorder{ResponseWriter: w, status: http.StatusOK}
		next.ServeHTTP(recorder, r)

		// Colores para diferentes códigos de estado
		greenBg := color.New(color.BgGreen).SprintFunc()
		yellowBg := color.New(color.BgYellow).SprintFunc()
		redBg := color.New(color.BgRed).SprintFunc()

		// Colorear el código de estado según el valor
		var statusColor string
		switch {
		case recorder.status >= 200 && recorder.status < 300:
			statusColor = greenBg(fmt.Sprintf(" %d ", recorder.status))
		case recorder.status >= 400 && recorder.status < 500:
			statusColor = yellowBg(fmt.Sprintf(" %d ", recorder.status))
		case recorder.status >= 500:
			statusColor = redBg(fmt.Sprintf(" %d ", recorder.status))
		default:
			statusColor = fmt.Sprintf(" %d ", recorder.status)
		}

		log.Printf("%s %s %s |%s|", r.Method, r.RequestURI, r.RemoteAddr, statusColor)
	})
}
