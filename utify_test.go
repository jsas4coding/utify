package utify

import (
	"os"
	"path/filepath"
	"testing"
)

func defaultOpts() *Options {
	return OptionsDefault()
}

func TestBasicOutputFunctions(t *testing.T) {
	tests := []struct {
		name string
		fn   func(string, *Options)
	}{
		{"Success", Success},
		{"Error", Error},
		{"Warning", Warning},
		{"Info", Info},
		{"Debug", Debug},
		{"Critical", Critical},
		{"Delete", Delete},
		{"Update", Update},
		{"Install", Install},
		{"Upgrade", Upgrade},
		{"Edit", Edit},
		{"New", New},
		{"Download", Download},
		{"Upload", Upload},
		{"Sync", Sync},
		{"Search", Search},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt := tt
			tt.fn("Test message", defaultOpts())
		})
	}
}

func TestFormattedOutputFunctions(t *testing.T) {
	tests := []struct {
		name string
		fn   func(string, *Options, ...any)
	}{
		{"Successf", Successf},
		{"Errorf", Errorf},
		{"Warningf", Warningf},
		{"Infof", Infof},
		{"Debugf", Debugf},
		{"Criticalf", Criticalf},
		{"Deletef", Deletef},
		{"Updatef", Updatef},
		{"Installf", Installf},
		{"Upgradef", Upgradef},
		{"Editf", Editf},
		{"Newf", Newf},
		{"Downloadf", Downloadf},
		{"Uploadf", Uploadf},
		{"Syncf", Syncf},
		{"Searchf", Searchf},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt := tt
			tt.fn("Formatted %s", defaultOpts(), "message")
		})
	}
}

func TestGetFunctions(t *testing.T) {
	tests := []struct {
		name string
		fn   func(string, *Options) (string, error)
	}{
		{"GetSuccess", GetSuccess},
		{"GetError", GetError},
		{"GetWarning", GetWarning},
		{"GetInfo", GetInfo},
		{"GetDebug", GetDebug},
		{"GetCritical", GetCritical},
		{"GetDelete", GetDelete},
		{"GetUpdate", GetUpdate},
		{"GetInstall", GetInstall},
		{"GetUpgrade", GetUpgrade},
		{"GetEdit", GetEdit},
		{"GetNew", GetNew},
		{"GetDownload", GetDownload},
		{"GetUpload", GetUpload},
		{"GetSync", GetSync},
		{"GetSearch", GetSearch},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.fn("Test message", defaultOpts())
			if err != nil && err != ErrSilent {
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
		{"LogSuccess", LogSuccess},
		{"LogError", LogError},
		{"LogWarning", LogWarning},
		{"LogInfo", LogInfo},
		{"LogDebug", LogDebug},
		{"LogCritical", LogCritical},
		{"LogDelete", LogDelete},
		{"LogUpdate", LogUpdate},
		{"LogInstall", LogInstall},
		{"LogUpgrade", LogUpgrade},
		{"LogEdit", LogEdit},
		{"LogNew", LogNew},
		{"LogDownload", LogDownload},
		{"LogUpload", LogUpload},
		{"LogSync", LogSync},
		{"LogSearch", LogSearch},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt := tt
			tt.fn("Log-only message")
		})
	}
}

func TestConfigFunctions(t *testing.T) {
	t.Run("SetColorTable", func(t *testing.T) {
		SetColorTable(map[string]string{
			"Red":  "#FF0000",
			"Blue": "#0000FF",
		})
	})

	t.Run("SetLogTargetAndGet", func(t *testing.T) {
		tmpfile := filepath.Join(os.TempDir(), "utify_test.log")
		err := SetLogTarget(tmpfile)
		if err != nil {
			t.Fatalf("SetLogTarget failed: %v", err)
		}
		got := GetLogTarget()
		if got != tmpfile {
			t.Errorf("expected log target %s, got %s", tmpfile, got)
		}
		CloseLogger()
	})

	t.Run("SetLoggingEnabled", func(t *testing.T) {
		SetLoggingEnabled(true)
		if !IsLoggingEnabled() {
			t.Error("expected logging to be enabled")
		}
		SetLoggingEnabled(false)
		if IsLoggingEnabled() {
			t.Error("expected logging to be disabled")
		}
	})

	t.Run("ForceIconsModes", func(t *testing.T) {
		ForceNerdFont()
		ForceRegularIcons()
		DisableIcons()
		_ = IsNerdFontDetected()
		_ = GetIconType()
	})
}

func TestGetFormattedFunctions(t *testing.T) {
	tests := []struct {
		name string
		fn   func(string, *Options, ...any) (string, error)
	}{
		{"GetSuccessf", GetSuccessf},
		{"GetErrorf", GetErrorf},
		{"GetWarningf", GetWarningf},
		{"GetInfof", GetInfof},
		{"GetDebugf", GetDebugf},
		{"GetCriticalf", GetCriticalf},
		{"GetDeletef", GetDeletef},
		{"GetUpdatef", GetUpdatef},
		{"GetInstallf", GetInstallf},
		{"GetUpgradef", GetUpgradef},
		{"GetEditf", GetEditf},
		{"GetNewf", GetNewf},
		{"GetDownloadf", GetDownloadf},
		{"GetUploadf", GetUploadf},
		{"GetSyncf", GetSyncf},
		{"GetSearchf", GetSearchf},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.fn("Formatted %s", defaultOpts(), "message")
			if err != nil && err != ErrSilent {
				t.Errorf("%s returned unexpected error: %v", tt.name, err)
			}
		})
	}
}

func TestLogFormattedFunctions(t *testing.T) {
	tests := []struct {
		name string
		fn   func(string, ...any)
	}{
		{"LogSuccessf", LogSuccessf},
		{"LogErrorf", LogErrorf},
		{"LogWarningf", LogWarningf},
		{"LogInfof", LogInfof},
		{"LogDebugf", LogDebugf},
		{"LogCriticalf", LogCriticalf},
		{"LogDeletef", LogDeletef},
		{"LogUpdatef", LogUpdatef},
		{"LogInstallf", LogInstallf},
		{"LogUpgradef", LogUpgradef},
		{"LogEditf", LogEditf},
		{"LogNewf", LogNewf},
		{"LogDownloadf", LogDownloadf},
		{"LogUploadf", LogUploadf},
		{"LogSyncf", LogSyncf},
		{"LogSearchf", LogSearchf},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.fn("Formatted %s", "message")
		})
	}
}
