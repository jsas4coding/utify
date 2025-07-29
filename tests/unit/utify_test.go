package unit

import (
	"testing"

	"github.com/jsas4coding/utify"
)

func defaultOpts() *utify.Options {
	return utify.OptionsDefault()
}

func TestBasicOutputFunctions(t *testing.T) {
	tests := []struct {
		name string
		fn   func(string, *utify.Options)
	}{
		{"Success", utify.Success},
		{"Error", utify.Error},
		{"Warning", utify.Warning},
		{"Info", utify.Info},
		{"Debug", utify.Debug},
		{"Critical", utify.Critical},
		{"Delete", utify.Delete},
		{"Update", utify.Update},
		{"Install", utify.Install},
		{"Upgrade", utify.Upgrade},
		{"Edit", utify.Edit},
		{"New", utify.New},
		{"Download", utify.Download},
		{"Upload", utify.Upload},
		{"Sync", utify.Sync},
		{"Search", utify.Search},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.fn("Test message", defaultOpts())
		})
	}
}

func TestFormattedOutputFunctions(t *testing.T) {
	tests := []struct {
		name string
		fn   func(string, *utify.Options, ...any)
	}{
		{"Successf", utify.Successf},
		{"Errorf", utify.Errorf},
		{"Warningf", utify.Warningf},
		{"Infof", utify.Infof},
		{"Debugf", utify.Debugf},
		{"Criticalf", utify.Criticalf},
		{"Deletef", utify.Deletef},
		{"Updatef", utify.Updatef},
		{"Installf", utify.Installf},
		{"Upgradef", utify.Upgradef},
		{"Editf", utify.Editf},
		{"Newf", utify.Newf},
		{"Downloadf", utify.Downloadf},
		{"Uploadf", utify.Uploadf},
		{"Syncf", utify.Syncf},
		{"Searchf", utify.Searchf},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.fn("Formatted %s", defaultOpts(), "message")
		})
	}
}

func TestGetFunctions(t *testing.T) {
	tests := []struct {
		name string
		fn   func(string, *utify.Options) (string, error)
	}{
		{"GetSuccess", utify.GetSuccess},
		{"GetError", utify.GetError},
		{"GetWarning", utify.GetWarning},
		{"GetInfo", utify.GetInfo},
		{"GetDebug", utify.GetDebug},
		{"GetCritical", utify.GetCritical},
		{"GetDelete", utify.GetDelete},
		{"GetUpdate", utify.GetUpdate},
		{"GetInstall", utify.GetInstall},
		{"GetUpgrade", utify.GetUpgrade},
		{"GetEdit", utify.GetEdit},
		{"GetNew", utify.GetNew},
		{"GetDownload", utify.GetDownload},
		{"GetUpload", utify.GetUpload},
		{"GetSync", utify.GetSync},
		{"GetSearch", utify.GetSearch},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.fn("Test message", defaultOpts())
			if err != nil && err != utify.ErrSilent {
				t.Errorf("%s returned unexpected error: %v", tt.name, err)
			}
		})
	}
}

func TestLogOnlyFunctions(t *testing.T) {
	tests := []struct {
		name string
		fn   func(string)
	}{
		{"LogSuccess", utify.LogSuccess},
		{"LogError", utify.LogError},
		{"LogWarning", utify.LogWarning},
		{"LogInfo", utify.LogInfo},
		{"LogDebug", utify.LogDebug},
		{"LogCritical", utify.LogCritical},
		{"LogDelete", utify.LogDelete},
		{"LogUpdate", utify.LogUpdate},
		{"LogInstall", utify.LogInstall},
		{"LogUpgrade", utify.LogUpgrade},
		{"LogEdit", utify.LogEdit},
		{"LogNew", utify.LogNew},
		{"LogDownload", utify.LogDownload},
		{"LogUpload", utify.LogUpload},
		{"LogSync", utify.LogSync},
		{"LogSearch", utify.LogSearch},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.fn("Log-only message")
		})
	}
}
