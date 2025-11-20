package config

import (
	"testing"
)

func Test_WithExistingConfig(t *testing.T) {
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
		t.Errorf("[CanBypassPrompt] Expected %q, got %q", expected, result)
	}

	// Cleanup
	config = nil
}

func Test_NoConfig(t *testing.T) {
	// Reset global config state before test
	resultInstallDir := GetInstallDir()
	expectedInstallDir := "/Users/obrutus/.prepare-code"

	resultBypassPrompt := CanBypassPrompt()
	expectedBypassPrompt := true

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
