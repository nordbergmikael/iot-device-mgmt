package api

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/diwise/iot-device-mgmt/internal/pkg/application"
	"github.com/diwise/iot-device-mgmt/internal/pkg/presentation/api/auth"
	"github.com/diwise/service-chassis/pkg/infrastructure/o11y"
	"github.com/diwise/service-chassis/pkg/infrastructure/o11y/tracing"
	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog"
	"go.opentelemetry.io/otel"
)

var tracer = otel.Tracer("iot-device-mgmt/api")

func RegisterHandlers(log zerolog.Logger, router *chi.Mux, policies io.Reader, app application.DeviceManagement) *chi.Mux {

	router.Get("/health", NewHealthHandler(log, app))

	router.Route("/api/v0", func(r chi.Router) {
		r.Route("/devices", func(r chi.Router) {
			r.Group(func(r chi.Router) {

				// Handle valid / invalid tokens.
				authenticator, err := auth.NewAuthenticator(context.Background(), log, policies)
				if err != nil {
					log.Fatal().Err(err).Msg("failed to create api authenticator")
				}
				r.Use(authenticator)

				r.Get("/", queryDevicesHandler(log, app))
				r.Post("/", createDeviceHandler(log, app))
				r.Get("/{id}", retrieveDeviceHandler(log, app))
				r.Patch("/{id}", patchDeviceHandler(log, app))
			})
		})

		r.Get("/environments", listEnvironments(log, app))
	})

	return router
}

func NewHealthHandler(log zerolog.Logger, app application.DeviceManagement) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	}
}

func listEnvironments(log zerolog.Logger, app application.DeviceManagement) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var err error

		ctx, span := tracer.Start(r.Context(), "list-environments")
		defer func() { tracing.RecordAnyErrorAndEndSpan(err, span) }()
		_, ctx, requestLogger := o11y.AddTraceIDToLoggerAndStoreInContext(span, log, ctx)

		env, err := app.ListEnvironments(ctx)
		if err != nil {
			requestLogger.Error().Err(err).Msg("unable to list environments")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		bytes, err := json.Marshal(&env)
		if err != nil {
			requestLogger.Error().Err(err).Msg("unable to marshal environments to json")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(bytes)
	}
}

func createDeviceHandler(log zerolog.Logger, app application.DeviceManagement) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var err error

		ctx, span := tracer.Start(r.Context(), "create-device")
		defer func() { tracing.RecordAnyErrorAndEndSpan(err, span) }()
		_, ctx, requestLogger := o11y.AddTraceIDToLoggerAndStoreInContext(span, log, ctx)

		body, err := io.ReadAll(r.Body)
		if err != nil {
			requestLogger.Error().Err(err).Msg("unable to read body")
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		var d application.Device
		err = json.Unmarshal(body, &d)
		if err != nil {
			requestLogger.Error().Err(err).Msg("unable to unmarshal body")
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		err = app.CreateDevice(ctx, d)
		if err != nil {
			requestLogger.Error().Err(err).Msg("unable to create device")
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
	}
}

func patchDeviceHandler(log zerolog.Logger, app application.DeviceManagement) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var err error

		ctx, span := tracer.Start(r.Context(), "patch-device")
		defer func() { tracing.RecordAnyErrorAndEndSpan(err, span) }()
		_, ctx, requestLogger := o11y.AddTraceIDToLoggerAndStoreInContext(span, log, ctx)

		deviceID := chi.URLParam(r, "id")

		b, err := io.ReadAll(r.Body)
		if err != nil {
			requestLogger.Error().Err(err).Msg("unable to read body")
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		var fields map[string]interface{}
		err = json.Unmarshal(b, &fields)
		if err != nil {
			requestLogger.Error().Err(err).Msg("unable to unmarshal body into map")
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		_, err = app.UpdateDevice(ctx, deviceID, fields)
		if err != nil {
			requestLogger.Error().Err(err).Msg("unable to update device")
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
	}
}

func queryDevicesHandler(log zerolog.Logger, app application.DeviceManagement) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var err error

		ctx, span := tracer.Start(r.Context(), "query-devices")
		defer func() { tracing.RecordAnyErrorAndEndSpan(err, span) }()
		_, ctx, requestLogger := o11y.AddTraceIDToLoggerAndStoreInContext(span, log, ctx)

		deviceArray := []application.Device{}

		devEUI := r.URL.Query().Get("devEUI")
		if devEUI == "" {
			devices, err := app.ListAllDevices(ctx)
			if err != nil {
				requestLogger.Error().Err(err).Msg("unable to fetch all devices")
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			deviceArray = append(deviceArray, devices...)
			requestLogger.Info().Msgf("returning information about %d devices", len(devices))
		} else {
			device, err := app.GetDeviceFromEUI(ctx, devEUI)
			if err != nil {
				requestLogger.Error().Err(err).Msg("device not found")
				w.WriteHeader(http.StatusNotFound)
				return
			}
			deviceArray = append(deviceArray, device)
			requestLogger.Info().Msgf("returning information about device %s", device.DeviceId)
		}

		bytes, err := json.Marshal(&deviceArray)
		if err != nil {
			requestLogger.Error().Err(err).Msg("unable to marshal device to json")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(bytes)
	}
}

func retrieveDeviceHandler(log zerolog.Logger, app application.DeviceManagement) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var err error

		ctx, span := tracer.Start(r.Context(), "get-device")
		defer func() { tracing.RecordAnyErrorAndEndSpan(err, span) }()
		_, ctx, requestLogger := o11y.AddTraceIDToLoggerAndStoreInContext(span, log, ctx)

		deviceID := chi.URLParam(r, "id")
		device, err := app.GetDevice(ctx, deviceID)
		if err != nil {
			requestLogger.Error().Err(err).Msg("device not found")
			w.WriteHeader(http.StatusNotFound)
			return
		}

		bytes, err := json.Marshal(device)
		if err != nil {
			requestLogger.Error().Err(err).Msg("unable to marshal device to json")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		requestLogger.Info().Msgf("returning information about device %s (%s)", device.DeviceId, deviceID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(bytes)
	}
}
