package kooperzap_test

import (
	"testing"

	"github.com/spotahome/kooper/v2/controller"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	kooperzap "github.com/thomas-tacquet/kooper-zap"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"go.uber.org/zap/zaptest/observer"
)

// TestKooperZap is a simple test that creates zap logger and then we call directly
// the logger function to write a message and check if this message has been saved to zap.
//
// Log testing is done under zap library this test is essentially here to ensure
// this package still matches the kooper logger interface.
func TestKooperZap(t *testing.T) {
	t.Parallel()

	// Use zap observer to ensure the passed logger works
	observedZapCore, recordedLogs := observer.New(zapcore.InfoLevel)
	observedLogger := zap.New(observedZapCore)

	// Create the controller with custom configuration.
	kooperCfg := controller.Config{
		Name:   "example",
		Logger: kooperzap.NewKooperLogger(observedLogger),
	}

	const logSample = "it works"

	kooperCfg.Logger.Infof(logSample)

	require.Equal(t, 1, recordedLogs.Len())

	firstLog := recordedLogs.All()[0]

	assert.Equal(t, logSample, firstLog.Message)
}
