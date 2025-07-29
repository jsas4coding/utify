package unit

import (
	"strings"
	"testing"

	"github.com/jsas4coding/utify"
	testutil "github.com/jsas4coding/utify/internal/tests"
)

func TestUtifyOptionsDefault(t *testing.T) {
	opts := utify.OptionsDefault()
	if opts == nil {
		t.Fatal("Expected OptionsDefault to return non-nil")
	}
	if opts.Bold || opts.Italic || opts.NoColor || opts.NoIcon || opts.NoStyle || opts.Exit {
		t.Error("Expected default options to have all flags false")
	}
}

func TestSetColorTableAndCloseLogger(t *testing.T) {
	// Just ensure no panic
	utify.SetColorTable(map[string]string{"test": "#FFF"})
	utify.CloseLogger()
}

func TestGetIconTypeChanges(t *testing.T) {
	original := utify.GetIconType()
	utify.DisableIcons()
	if utify.GetIconType() != 0 {
		t.Error("Expected GetIconType to return 0 after DisableIcons")
	}
	utify.ForceRegularIcons()
	if utify.GetIconType() != 1 {
		t.Error("Expected GetIconType to return 1 after ForceRegularIcons")
	}
	utify.ForceNerdFont()
	if utify.GetIconType() != 2 {
		t.Error("Expected GetIconType to return 2 after ForceNerdFont")
	}
	// Restore
	switch original {
	case 0:
		utify.DisableIcons()
	case 1:
		utify.ForceRegularIcons()
	case 2:
		utify.ForceNerdFont()
	}
}

func TestAllOutputFunctions(t *testing.T) {
	opts := utify.OptionsDefault()

	type testCase struct {
		name string
		fn   func()
	}
	msgs := []string{
		"Success", "Error", "Warning", "Info", "Debug", "Critical",
		"Delete", "Update", "Install", "Upgrade", "Edit", "New",
		"Download", "Upload", "Sync", "Search",
	}
	tests := []testCase{}

	// Add direct output
	for _, m := range msgs {
		tests = append(tests, testCase{
			name: "Direct" + m,
			fn: func() {
				switch m {
				case "Success":
					utify.Success("msg", opts)
				case "Error":
					utify.Error("msg", opts)
				case "Warning":
					utify.Warning("msg", opts)
				case "Info":
					utify.Info("msg", opts)
				case "Debug":
					utify.Debug("msg", opts)
				case "Critical":
					utify.Critical("msg", opts)
				case "Delete":
					utify.Delete("msg", opts)
				case "Update":
					utify.Update("msg", opts)
				case "Install":
					utify.Install("msg", opts)
				case "Upgrade":
					utify.Upgrade("msg", opts)
				case "Edit":
					utify.Edit("msg", opts)
				case "New":
					utify.New("msg", opts)
				case "Download":
					utify.Download("msg", opts)
				case "Upload":
					utify.Upload("msg", opts)
				case "Sync":
					utify.Sync("msg", opts)
				case "Search":
					utify.Search("msg", opts)
				}
			},
		})
	}

	// Add formatted output
	for _, m := range msgs {
		tests = append(tests, testCase{
			name: "Formatted" + m,
			fn: func() {
				switch m {
				case "Success":
					utify.Successf("%s", opts, "fmt")
				case "Error":
					utify.Errorf("%s", opts, "fmt")
				case "Warning":
					utify.Warningf("%s", opts, "fmt")
				case "Info":
					utify.Infof("%s", opts, "fmt")
				case "Debug":
					utify.Debugf("%s", opts, "fmt")
				case "Critical":
					utify.Criticalf("%s", opts, "fmt")
				case "Delete":
					utify.Deletef("%s", opts, "fmt")
				case "Update":
					utify.Updatef("%s", opts, "fmt")
				case "Install":
					utify.Installf("%s", opts, "fmt")
				case "Upgrade":
					utify.Upgradef("%s", opts, "fmt")
				case "Edit":
					utify.Editf("%s", opts, "fmt")
				case "New":
					utify.Newf("%s", opts, "fmt")
				case "Download":
					utify.Downloadf("%s", opts, "fmt")
				case "Upload":
					utify.Uploadf("%s", opts, "fmt")
				case "Sync":
					utify.Syncf("%s", opts, "fmt")
				case "Search":
					utify.Searchf("%s", opts, "fmt")
				}
			},
		})
	}

	// Add GetX and GetXf
	for _, m := range msgs {
		tests = append(tests, testCase{
			name: "Get" + m,
			fn: func() {
				var s string
				var e error
				switch m {
				case "Success":
					s, e = utify.GetSuccess("msg", opts)
				case "Error":
					s, e = utify.GetError("msg", opts)
				case "Warning":
					s, e = utify.GetWarning("msg", opts)
				case "Info":
					s, e = utify.GetInfo("msg", opts)
				case "Debug":
					s, e = utify.GetDebug("msg", opts)
				case "Critical":
					s, e = utify.GetCritical("msg", opts)
				case "Delete":
					s, e = utify.GetDelete("msg", opts)
				case "Update":
					s, e = utify.GetUpdate("msg", opts)
				case "Install":
					s, e = utify.GetInstall("msg", opts)
				case "Upgrade":
					s, e = utify.GetUpgrade("msg", opts)
				case "Edit":
					s, e = utify.GetEdit("msg", opts)
				case "New":
					s, e = utify.GetNew("msg", opts)
				case "Download":
					s, e = utify.GetDownload("msg", opts)
				case "Upload":
					s, e = utify.GetUpload("msg", opts)
				case "Sync":
					s, e = utify.GetSync("msg", opts)
				case "Search":
					s, e = utify.GetSearch("msg", opts)
				}
				if s == "" {
					t.Error("Expected non-empty string from Get function")
				}
				if e != nil && e != utify.ErrSilent {
					t.Errorf("Unexpected error: %v", e)
				}
			},
		})

		tests = append(tests, testCase{
			name: "GetFormatted" + m,
			fn: func() {
				var s string
				var e error
				switch m {
				case "Success":
					s, e = utify.GetSuccessf("%s", opts, "fmt")
				case "Error":
					s, e = utify.GetErrorf("%s", opts, "fmt")
				case "Warning":
					s, e = utify.GetWarningf("%s", opts, "fmt")
				case "Info":
					s, e = utify.GetInfof("%s", opts, "fmt")
				case "Debug":
					s, e = utify.GetDebugf("%s", opts, "fmt")
				case "Critical":
					s, e = utify.GetCriticalf("%s", opts, "fmt")
				case "Delete":
					s, e = utify.GetDeletef("%s", opts, "fmt")
				case "Update":
					s, e = utify.GetUpdatef("%s", opts, "fmt")
				case "Install":
					s, e = utify.GetInstallf("%s", opts, "fmt")
				case "Upgrade":
					s, e = utify.GetUpgradef("%s", opts, "fmt")
				case "Edit":
					s, e = utify.GetEditf("%s", opts, "fmt")
				case "New":
					s, e = utify.GetNewf("%s", opts, "fmt")
				case "Download":
					s, e = utify.GetDownloadf("%s", opts, "fmt")
				case "Upload":
					s, e = utify.GetUploadf("%s", opts, "fmt")
				case "Sync":
					s, e = utify.GetSyncf("%s", opts, "fmt")
				case "Search":
					s, e = utify.GetSearchf("%s", opts, "fmt")
				}
				if s == "" {
					t.Error("Expected non-empty string from Get formatted function")
				}
				if e != nil && e != utify.ErrSilent {
					t.Errorf("Unexpected error: %v", e)
				}
			},
		})
	}

	// Add LogX and LogXf
	for _, m := range msgs {
		tests = append(tests, testCase{
			name: "Log" + m,
			fn: func() {
				switch m {
				case "Success":
					utify.LogSuccess("msg")
				case "Error":
					utify.LogError("msg")
				case "Warning":
					utify.LogWarning("msg")
				case "Info":
					utify.LogInfo("msg")
				case "Debug":
					utify.LogDebug("msg")
				case "Critical":
					utify.LogCritical("msg")
				case "Delete":
					utify.LogDelete("msg")
				case "Update":
					utify.LogUpdate("msg")
				case "Install":
					utify.LogInstall("msg")
				case "Upgrade":
					utify.LogUpgrade("msg")
				case "Edit":
					utify.LogEdit("msg")
				case "New":
					utify.LogNew("msg")
				case "Download":
					utify.LogDownload("msg")
				case "Upload":
					utify.LogUpload("msg")
				case "Sync":
					utify.LogSync("msg")
				case "Search":
					utify.LogSearch("msg")
				}
			},
		})

		tests = append(tests, testCase{
			name: "LogFormatted" + m,
			fn: func() {
				switch m {
				case "Success":
					utify.LogSuccessf("%s", "fmt")
				case "Error":
					utify.LogErrorf("%s", "fmt")
				case "Warning":
					utify.LogWarningf("%s", "fmt")
				case "Info":
					utify.LogInfof("%s", "fmt")
				case "Debug":
					utify.LogDebugf("%s", "fmt")
				case "Critical":
					utify.LogCriticalf("%s", "fmt")
				case "Delete":
					utify.LogDeletef("%s", "fmt")
				case "Update":
					utify.LogUpdatef("%s", "fmt")
				case "Install":
					utify.LogInstallf("%s", "fmt")
				case "Upgrade":
					utify.LogUpgradef("%s", "fmt")
				case "Edit":
					utify.LogEditf("%s", "fmt")
				case "New":
					utify.LogNewf("%s", "fmt")
				case "Download":
					utify.LogDownloadf("%s", "fmt")
				case "Upload":
					utify.LogUploadf("%s", "fmt")
				case "Sync":
					utify.LogSyncf("%s", "fmt")
				case "Search":
					utify.LogSearchf("%s", "fmt")
				}
			},
		})
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			output := testutil.CaptureOutput(tc.fn)
			if strings.Contains(tc.name, "Log") {
				// Log functions should not output to stdout
				if output != "" {
					t.Errorf("Expected no stdout for %s, got: %q", tc.name, output)
				}
			}
		})
	}
}
