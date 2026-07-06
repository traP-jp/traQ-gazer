package main

import (
	"fmt"
	"log/slog"
	"os"
	"traQ-gazer/message"
	"traQ-gazer/oapi"
	"traQ-gazer/repo"

	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
)

func main() {
	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stdout, nil)))

	instance := echo.New()
	instance.Use(requestLogger())

	server := oapi.Server{}

	apiInstance := instance.Group("/api")
	oapi.RegisterHandlers(apiInstance, server)

	// まとめて賢くルーティングするのは厳しそうなので
	instance.Static("/", "dist")
	instance.File("/words", "dist/index.html")
	instance.File("/words/add", "dist/index.html")

	err := repo.SetUp()
	if err != nil {
		slog.Error("failed to set up", "err_type", fmt.Sprintf("%T", err))
	}

	messagePoller := message.NewMessagePoller()
	go messagePoller.Run()

	if err := instance.Start(":8080"); err != nil {
		slog.Error("server stopped", "err", err)
		os.Exit(1)
	}
}

func requestLogger() echo.MiddlewareFunc {
	return middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		HandleError:     true,
		LogLatency:      true,
		LogMethod:       true,
		LogRoutePath:    true,
		LogStatus:       true,
		LogResponseSize: true,
		LogValuesFunc: func(c *echo.Context, v middleware.RequestLoggerValues) error {
			level := slog.LevelInfo
			if v.Status >= 500 {
				level = slog.LevelError
			} else if v.Status >= 400 {
				level = slog.LevelWarn
			}

			route := v.RoutePath
			if route == "" {
				route = "unmatched"
			}

			attrs := []slog.Attr{
				slog.String("method", v.Method),
				slog.String("route", route),
				slog.Int("status", v.Status),
				slog.Duration("latency", v.Latency),
				slog.Int64("bytes_out", v.ResponseSize),
			}
			if v.Error != nil {
				attrs = append(attrs, slog.Bool("handler_error", true))
				attrs = append(attrs, slog.String("error_type", fmt.Sprintf("%T", v.Error)))
			}

			slog.LogAttrs(c.Request().Context(), level, "http request", attrs...)
			return nil
		},
	})
}
