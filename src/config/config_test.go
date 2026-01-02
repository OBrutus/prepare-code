package config

import (
	"os"
	"testing"
)

func setupCI(t *testing.T) {
	originalCI := os.Getenv("CI")
	os.Unsetenv("CI")

	// t.Cleanup automatically runs after the test finishes
	t.Cleanup(func() {
		if originalCI != "" {
			os.Setenv("CI", originalCI)
		}
	})
}

func Test_WithExistingConfig(t *testing.T) {
	setupCI()

	// Reset global config state before test
	config = &Config{
		InstallDir:   "/custom/install/path",
		BypassPrompt: true,
	}

	result := GetInstallDir()
	expected := "/custom/install/path"

	if result != expected {
		t.Errorf("[GetInstallDir] Expected %q, got %q", expected, result)
	}

	resultBypass := CanBypassPrompt()
	expectedBypass := true

	if resultBypass != expectedBypass {
		t.Errorf("[CanBypassPrompt] Expected %t, got %t", expectedBypass, resultBypass)
	}

	// Cleanup
	config = nil
}

func Test_NoConfig(t *testing.T) {
	setupCI(t)

	// Reset global config state before test
	config = nil

	resultInstallDir := GetInstallDir()
	expectedInstallDir := getOldInstallDirPath()

	resultBypassPrompt := CanBypassPrompt()
	expectedBypassPrompt := false // Default when config file doesn't exist

	if resultInstallDir != expectedInstallDir {
		t.Errorf("[InstallDir] "+
			"Expected %q, got %q", expectedInstallDir, resultInstallDir,
		)
	}

	if resultBypassPrompt != expectedBypassPrompt {
		t.Errorf("[CanBypassPrompt] "+
			"Expected %t, got %t", expectedBypassPrompt, resultBypassPrompt,
		)
	}

	// Cleanup
	config = nil
}
