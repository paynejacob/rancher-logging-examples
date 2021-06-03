package main

import (
	"fmt"
	golorem "github.com/drhodes/golorem"
	"github.com/kelseyhightower/envconfig"
	"log"
	"math/rand"
	"os"
	"text/template"
	"time"
)

type config struct {
	LogFormat string `default:"[{{.Timestamp}}] @{{.Level}} {{.Message}} {{.Number}} ::{{.Decimal}}"`

	Frequency    time.Duration `default:"1s"`
	ConstantRate bool          `default:"true"`
}

type logLevel string

const (
	DEBUG    logLevel = "DEBUG"
	INFO     logLevel = "INFO"
	WARNING  logLevel = "WARNING"
	ERROR    logLevel = "ERROR"
	CRITICAL logLevel = "CRITICAL"
)

func main() {
	logLevels := []logLevel{DEBUG, INFO, WARNING, ERROR, CRITICAL}
	var logCount int64
	var delay time.Duration
	var config = &config{}

	r := rand.New(rand.NewSource(time.Now().Unix()))

	if err := envconfig.Process("", config); err != nil {
		log.Fatalf("failed to parse configuration from environment: %v", err)
	}

	for {
		levelIndex := r.Intn(len(logLevels))

		_ = template.Must(template.New("").Parse(config.LogFormat+"\n")).Execute(os.Stdout, map[string]string{
			"Timestamp": time.Now().Format(time.RFC3339),
			"Level":     string(logLevels[levelIndex]),
			"Message":   golorem.Sentence(5, 10),
			"Number":    fmt.Sprintf("%d", logCount),
			"Decimal":   fmt.Sprintf("%f", r.Float64()),
		})

		if config.ConstantRate {
			delay, _ = time.ParseDuration(fmt.Sprintf("%ds", int(config.Frequency.Seconds())))
		} else {
			delay, _ = time.ParseDuration(fmt.Sprintf("%ds", r.Intn(int(config.Frequency.Seconds()))))
		}

		logCount++

		time.Sleep(delay)
	}
}
