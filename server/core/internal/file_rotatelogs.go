package internal

import (
	"GinVueAdmin/server/global"
	rotateLogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap/zapcore"
	"os"
	"path"
	"time"
)

type fileRotateLogs struct{}

var FileRotateLogs = new(fileRotateLogs)

func (r *fileRotateLogs) GetWriteSyncer(level string) (zapcore.WriteSyncer, error) {
	fileWriter, err := rotateLogs.New(
		path.Join(global.GVA_CONFIG.Zap.Director, "%Y-%m-%d", level+".log"),
		rotateLogs.WithClock(rotateLogs.Local),
		rotateLogs.WithMaxAge(time.Duration(global.GVA_CONFIG.Zap.MaxAge)*24*time.Hour), // 日志留存时间
		rotateLogs.WithRotationTime(time.Hour*24),
	)
	if global.GVA_CONFIG.Zap.LogInConsole {
		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(fileWriter)), err
	}
	return zapcore.AddSync(fileWriter), err
}
