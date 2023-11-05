package main

import "go.uber.org/zap"

func NewLogger() (*zap.Logger, error) {
	cfg := zap.NewProductionConfig()
	cfg.OutputPaths = []string{
		"./out.log",
		"stdout",
	}
	return cfg.Build()
}

func main() {
	lg, err := NewLogger()
	if err != nil {
		panic(err)
	}
	sugar := lg.Sugar()
	defer sugar.Sync()

	//lg, _ := zap.NewProduction()
	//lg, _ = zap.NewDevelopment()
	//lg.Sync()
	//lg.Info("test log", zap.String("input", "123"), zap.Int("time", 1920))
	//sugar := lg.Sugar()
	sugar.Infow("1234567", "pwd", 123456, "name", "bob")
}