package i18n

import (
	"os"
	"strings"
	"sync"
)

// Language represents supported languages
type Language string

const (
	LangEN Language = "en"
	LangZH Language = "zh"
)

var (
	currentLang Language
	once        sync.Once
	mu          sync.RWMutex
)

// T translates a given key to the current language
func T(key string) string {
	once.Do(initLanguage)

	mu.RLock()
	defer mu.RUnlock()

	var translations map[string]string
	switch currentLang {
	case LangZH:
		translations = zhTranslations
	default:
		translations = enTranslations
	}

	if text, exists := translations[key]; exists {
		return text
	}

	// Fallback to English if key not found in current language
	if currentLang != LangEN {
		if text, exists := enTranslations[key]; exists {
			return text
		}
	}

	// Return key if translation not found
	return key
}

// SetLanguage manually sets the language
func SetLanguage(lang Language) {
	mu.Lock()
	defer mu.Unlock()
	currentLang = lang
}

// GetCurrentLanguage returns the current language
func GetCurrentLanguage() Language {
	mu.RLock()
	defer mu.RUnlock()
	return currentLang
}

// initLanguage detects and sets the language based on system locale
func initLanguage() {
	lang := detectSystemLanguage()
	mu.Lock()
	currentLang = lang
	mu.Unlock()
}

// detectSystemLanguage detects system language from environment variables
func detectSystemLanguage() Language {
	// Check environment variables
	for _, env := range []string{"LC_ALL", "LC_MESSAGES", "LANG"} {
		if locale := os.Getenv(env); locale != "" {
			if isChineseLocale(locale) {
				return LangZH
			}
		}
	}

	// Default to English
	return LangEN
}

// isChineseLocale checks if the locale indicates Chinese language
func isChineseLocale(locale string) bool {
	locale = strings.ToLower(locale)
	return strings.HasPrefix(locale, "zh") ||
		strings.Contains(locale, "chinese") ||
		strings.Contains(locale, "china")
}
