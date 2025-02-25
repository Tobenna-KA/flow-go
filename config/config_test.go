package config

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/require"

	"github.com/onflow/flow-go/utils/unittest"
)

// TestBindPFlags ensures configuration is bound to the pflag set as expected and configuration values are overridden when set with CLI flags.
func TestBindPFlags(t *testing.T) {
	t.Run("should override config values when any flag is set", func(t *testing.T) {
		c := defaultConfig(t)
		flags := testFlagSet(c)
		err := flags.Set("networking-connection-pruning", "false")
		require.NoError(t, err)
		require.NoError(t, flags.Parse(nil))

		configFileUsed, err := BindPFlags(c, flags)
		require.NoError(t, err)
		require.False(t, configFileUsed)
		require.False(t, c.NetworkConfig.NetworkConnectionPruning)
	})
	t.Run("should return an error if flags are not parsed", func(t *testing.T) {
		c := defaultConfig(t)
		flags := testFlagSet(c)
		configFileUsed, err := BindPFlags(&FlowConfig{}, flags)
		require.False(t, configFileUsed)
		require.Error(t, err)
		require.True(t, errors.Is(err, errPflagsNotParsed))
	})
}

// TestDefaultConfig ensures the default Flow config is created and returned without errors.
func TestDefaultConfig(t *testing.T) {
	c := defaultConfig(t)
	require.Equalf(t, "./default-config.yml", c.ConfigFile, "expected default config file to be used")
	require.NoErrorf(t, c.Validate(), "unexpected error encountered validating default config")
	unittest.IdentifierFixture()
}

// TestFlowConfig_Validate ensures the Flow validate returns the expected number of validator.ValidationErrors when incorrect
// fields are set.
func TestFlowConfig_Validate(t *testing.T) {
	c := defaultConfig(t)
	// set invalid config values
	c.NetworkConfig.UnicastRateLimitersConfig.MessageRateLimit = -100
	c.NetworkConfig.UnicastRateLimitersConfig.BandwidthRateLimit = -100
	err := c.Validate()
	require.Error(t, err)
	errs, ok := errors.Unwrap(err).(validator.ValidationErrors)
	require.True(t, ok)
	require.Len(t, errs, 2)
}

// TestUnmarshall_UnsetFields ensures that if the config store has any missing config values an error is returned when the config is decoded into a Flow config.
func TestUnmarshall_UnsetFields(t *testing.T) {
	conf = viper.New()
	c := &FlowConfig{}
	err := Unmarshall(c)
	require.True(t, strings.Contains(err.Error(), "has unset fields"))
}

// Test_overrideConfigFile ensures configuration values can be overridden via the --config-file flag.
func Test_overrideConfigFile(t *testing.T) {
	t.Run("should override the default config if --config-file is set", func(t *testing.T) {
		file, err := os.CreateTemp("", "config-*.yml")
		require.NoError(t, err)
		defer os.Remove(file.Name())

		var data = fmt.Sprintf(`config-file: "%s"
network-config:
 networking-connection-pruning: false
`, file.Name())
		_, err = file.Write([]byte(data))
		require.NoError(t, err)
		c := defaultConfig(t)
		flags := testFlagSet(c)
		err = flags.Set(configFileFlagName, file.Name())

		require.NoError(t, err)
		overridden, err := overrideConfigFile(flags)
		require.NoError(t, err)
		require.True(t, overridden)

		// ensure config values overridden with values from our inline config
		require.Equal(t, conf.GetString(configFileFlagName), file.Name())
		require.False(t, conf.GetBool("networking-connection-pruning"))
	})
	t.Run("should return an error for missing --config file", func(t *testing.T) {
		c := defaultConfig(t)
		flags := testFlagSet(c)
		err := flags.Set(configFileFlagName, "./missing-config.yml")
		require.NoError(t, err)
		overridden, err := overrideConfigFile(flags)
		require.Error(t, err)
		require.False(t, overridden)
	})
	t.Run("should not attempt to override config if --config-file is not set", func(t *testing.T) {
		c := defaultConfig(t)
		flags := testFlagSet(c)
		overridden, err := overrideConfigFile(flags)
		require.NoError(t, err)
		require.False(t, overridden)
	})
	t.Run("should return an error for file types other than .yml", func(t *testing.T) {
		file, err := os.CreateTemp("", "config-*.json")
		require.NoError(t, err)
		defer os.Remove(file.Name())
		c := defaultConfig(t)
		flags := testFlagSet(c)
		err = flags.Set(configFileFlagName, file.Name())
		require.NoError(t, err)
		overridden, err := overrideConfigFile(flags)
		require.Error(t, err)
		require.False(t, overridden)
	})
}

// defaultConfig resets the config store gets the default Flow config.
func defaultConfig(t *testing.T) *FlowConfig {
	initialize()
	c, err := DefaultConfig()
	require.NoError(t, err)
	return c
}

func testFlagSet(c *FlowConfig) *pflag.FlagSet {
	flags := pflag.NewFlagSet("test", pflag.PanicOnError)
	// initialize default flags
	InitializePFlagSet(flags, c)
	return flags
}
