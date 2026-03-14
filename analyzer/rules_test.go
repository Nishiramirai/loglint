package analyzer

import "testing"

func TestIsLowerStart(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"starting server", true},
		{"Starting server", false},
		{"123 server started", true},
		{"", true},
	}
	for _, tt := range tests {
		if result := isLowerStart(tt.input); result != tt.expected {
			t.Errorf("isLowerStart(%q) = %v; want %v", tt.input, result, tt.expected)
		}
	}
}

func TestIsEnglishOnly(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"starting server", true},
		{"запуск сервера", false},
		{"server 123", true},
		{"server error: ошибка", false},
	}
	for _, tt := range tests {
		if result := isEnglishOnly(tt.input); result != tt.expected {
			t.Errorf("isEnglishOnly(%q) = %v; want %v", tt.input, result, tt.expected)
		}
	}
}

func TestHasSpecialCharsOrEmojis(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"server started", false},
		{"connection failed!!!", true},
		{"server started 🚀", true},
		{"warning: error", true},
	}
	for _, tt := range tests {
		if result := hasSpecialCharsOrEmojis(tt.input); result != tt.expected {
			t.Errorf("hasSpecialCharsOrEmojis(%q) = %v; want %v", tt.input, result, tt.expected)
		}
	}
}

func TestHasSensitiveData(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"user authenticated successfully", false},
		{"user password: 123", true},
		{"api_key=secret", true},
		{"token: jwt", true},
		{"token validated", false},
	}
	for _, tt := range tests {
		if result := hasSensitiveData(tt.input); result != tt.expected {
			t.Errorf("hasSensitiveData(%q) = %v; want %v", tt.input, result, tt.expected)
		}
	}
}
