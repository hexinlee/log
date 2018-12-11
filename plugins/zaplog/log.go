package zaplog

import (
	"os"
	"time"
	"runtime"
	"strings"
	"github.com/hexinlee/log/conf"
	"github.com/hexinlee/log/fileout"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Log struct {
	logger *zap.Logger
}

//var Log *zap.Logger //全局日志

func parseLevel(lvl string) zapcore.Level {
	switch strings.ToLower(lvl) {
	case "panic", "dpanic":
		return zapcore.PanicLevel
	case "fatal":
		return zapcore.FatalLevel
	case "error":
		return zapcore.ErrorLevel
	case "warn", "warning":
		return zapcore.WarnLevel
	case "info":
		return zapcore.InfoLevel
	case "debug":
		return zapcore.DebugLevel
	default:
		return zapcore.DebugLevel
	}
}

//创建日志
func New(opts ...conf.Option) *Log {
	o := &conf.Options{
		LogPath:    conf.LogPath,
		LogName:    conf.LogName,
		LogLevel:   conf.LogLevel,
		MaxSize:    conf.MaxSize,
		MaxAge:     conf.MaxAge,
		Stacktrace: conf.Stacktrace,
		IsStdOut:   conf.IsStdOut,
		//ProjectName: conf.ProjectName,
	}
	for _, opt := range opts {
		opt(o)
	}
	writers := []zapcore.WriteSyncer{fileout.NewRollingFile(o.LogPath, o.LogName, o.MaxSize, o.MaxAge)}
	if o.IsStdOut == "yes" {
		writers = append(writers, os.Stdout)
	}
	logger := newZapLogger(parseLevel(o.LogLevel), parseLevel(o.Stacktrace), zapcore.NewMultiWriteSyncer(writers...))
	zap.RedirectStdLog(logger)
	//logger = logger.With(zap.String("project", o.ProjectName)) //加上项目名称
	return &Log{logger: logger}
}

// CallerEncoder will add caller to log. format is "filename:lineNum:funcName", e.g:"zaplog/zaplog_test.go:15:zaplog.TestNewLogger"
func CallerEncoder(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(strings.Join([]string{caller.TrimmedPath(), runtime.FuncForPC(caller.PC).Name()}, ":"))
	//_, file, line, ok := runtime.Caller(calldepth)
}

func newZapLogger(level, stacktrace zapcore.Level, output zapcore.WriteSyncer) *zap.Logger {
	encCfg := zapcore.EncoderConfig{
		TimeKey:    "@",
		LevelKey:   "lev",
		NameKey:    "app",
		CallerKey:  "func",
		MessageKey: "msg",
		//StacktraceKey: "stacktrace",
		LineEnding: zapcore.DefaultLineEnding,
		//EncodeCaller:   zapcore.ShortCallerEncoder,
		EncodeCaller:   CallerEncoder,
		EncodeDuration: zapcore.NanosDurationEncoder,
		//EncodeTime: zapcore.ISO8601TimeEncoder,
		EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(t.Format("2006-01-02 15:04:05"))
		},
		EncodeLevel: zapcore.LowercaseLevelEncoder,
	}

	var encoder zapcore.Encoder
	dyn := zap.NewAtomicLevel()
	//encCfg.EncodeLevel = zapcore.LowercaseLevelEncoder
	//encoder = zapcore.NewJSONEncoder(encCfg) // zapcore.NewConsoleEncoder(encCfg)
	dyn.SetLevel(level)
	encoder = zapcore.NewJSONEncoder(encCfg)

	return zap.New(zapcore.NewCore(encoder, output, dyn), zap.AddCaller(), zap.AddStacktrace(stacktrace), zap.AddCallerSkip(2))
}
