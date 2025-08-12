package logx

import (
  "go.uber.org/zap"
  "gopkg.in/natefinch/lumberjack.v2"
)

func New() (*zap.Logger, error) {
  // TODO: zap core + lumberjack для ротируемых логов
  _ = lumberjack.Logger{} // чтобы зависимость подтянулась
  return zap.NewProduction()
}

