package logger

import (
	"context"

	"github.com/jae2274/goutils/llog"
	"github.com/jae2274/goutils/terr"
)

type AppLogger struct {
	apiLogger    *apiLogger
	stdoutLogger *llog.StdoutLLogger
}

func NewAppLogger(ctx context.Context, postUrl string) (*AppLogger, error) {
	apiLogger := &apiLogger{postUrl}
	stdoutLogger := &llog.StdoutLLogger{}

	firstLog := llog.Msg("AppLogger created").Build(ctx)

	err := apiLogger.Log(firstLog)
	if err != nil {
		return nil, terr.Wrap(err)
	}

	err = stdoutLogger.Log(firstLog)
	if err != nil {
		return nil, terr.Wrap(err)
	}

	return &AppLogger{
		apiLogger:    apiLogger,
		stdoutLogger: stdoutLogger,
	}, nil
}

func (al *AppLogger) Log(lg *llog.LLog) error {
	go func() { //apiLogger.Log()는 비동기로 실행, apiLogger.Log()가 실패해도 stdoutLogger.Log()는 실행
		err := al.apiLogger.Log(lg)
		if err != nil {
			al.stdoutLogger.Log(
				llog.Msg("failed to log to api").Data("error", err.Error()).Data("logContent", lg).Build(context.Background()),
			)
		}
	}()

	return al.stdoutLogger.Log(lg)
}
