package i18n

import (
	"testing"
	"time"
)

// BenchmarkTranslation tests the performance of the T() function
func BenchmarkTranslation(b *testing.B) {
	// Test key lookup performance
	b.Run("ExistingKey", func(b *testing.B) {
		for b.Loop() {
			_ = T("version.short")
		}
	})

	b.Run("NonExistingKey", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = T("non.existing.key")
		}
	})
}

// BenchmarkLanguageSwitch tests the performance of language switching
func BenchmarkLanguageSwitch(b *testing.B) {
	for i := 0; b.Loop(); i++ {
		if i%2 == 0 {
			SetLanguage(LangZH)
		} else {
			SetLanguage(LangEN)
		}
		_ = T("version.short")
	}
}

// TestTranslationPerformance measures actual performance
func TestTranslationPerformance(t *testing.T) {
	start := time.Now()

	// Perform 10000 translations
	for range 10000 {
		_ = T("version.short")
		_ = T("status.short")
		_ = T("push.all.short")
	}

	duration := time.Since(start)
	t.Logf("10000 translations took: %v", duration)

	// Should be less than 10ms for 10000 translations (very fast)
	if duration > 10*time.Millisecond {
		t.Errorf("Translation performance too slow: %v", duration)
	}
}

// TestLanguageDetection tests automatic language detection
func TestLanguageDetection(t *testing.T) {
	tests := []struct {
		locale   string
		expected Language
	}{
		{"zh_CN.UTF-8", LangZH},
		{"zh_TW.UTF-8", LangZH},
		{"zh", LangZH},
		{"en_US.UTF-8", LangEN},
		{"en_GB.UTF-8", LangEN},
		{"fr_FR.UTF-8", LangEN}, // Fallback to English
		{"", LangEN},            // Empty fallback to English
	}

	for _, tt := range tests {
		t.Run(tt.locale, func(t *testing.T) {
			result := LangEN // Default
			if isChineseLocale(tt.locale) {
				result = LangZH
			}

			if result != tt.expected {
				t.Errorf("detectSystemLanguage(%s) = %v, want %v", tt.locale, result, tt.expected)
			}
		})
	}
}

// TestTranslationContent tests that translations are complete
func TestTranslationContent(t *testing.T) {
	commonKeys := []string{
		"root.short",
		"version.short",
		"status.short",
		"push.all.short",
		"push.selected.short",
	}

	for _, key := range commonKeys {
		// Test English
		SetLanguage(LangEN)
		enText := T(key)
		if enText == key {
			t.Errorf("Missing English translation for key: %s", key)
		}

		// Test Chinese
		SetLanguage(LangZH)
		zhText := T(key)
		if zhText == key {
			t.Errorf("Missing Chinese translation for key: %s", key)
		}

		// Translations should be different (unless they're the same by design)
		if enText == zhText && key != "version.version" {
			t.Logf("Warning: English and Chinese translations are identical for key: %s", key)
		}
	}
}
