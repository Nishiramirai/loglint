package analyzer

import (
	"go/ast"
	"go/token"
	"strings"
	"unicode"
	"unicode/utf8"

	"golang.org/x/tools/go/analysis"
)

func run(pass *analysis.Pass) (any, error) {
	for _, file := range pass.Files {
		ast.Inspect(file, func(n ast.Node) bool {
			callExpr, ok := n.(*ast.CallExpr)
			if !ok {
				return true
			}

			sel, ok := callExpr.Fun.(*ast.SelectorExpr)
			if !ok {
				return true
			}

			obj := pass.TypesInfo.Uses[sel.Sel]
			if obj == nil || obj.Pkg() == nil {
				return true
			}

			pkgPath := obj.Pkg().Path()

			isSlog := pkgPath == "log/slog"
			isZap := strings.HasPrefix(pkgPath, "go.uber.org/zap")

			if !isSlog && !isZap {
				return true
			}

			methodName := sel.Sel.Name
			switch methodName {
			case "Info", "Error", "Warn", "Debug", "Fatal":
				// Это вызов логгера
			default:
				return true
			}

			for _, arg := range callExpr.Args {
				ast.Inspect(arg, func(argNode ast.Node) bool {
					lit, ok := argNode.(*ast.BasicLit)
					if !ok || lit.Kind != token.STRING {
						return true
					}

					msg := strings.Trim(lit.Value, `"`)

					if len(msg) == 0 {
						return true
					}

					if !isLowerStart(msg) {
						pass.Reportf(lit.Pos(), "log message must start with a lowercase")
					}

					if !isEnglishOnly(msg) {
						pass.Reportf(lit.Pos(), "log message must be in English only")
					}

					if hasSpecialCharsOrEmojis(msg) {
						pass.Reportf(lit.Pos(), "log message must not contain special characters")
					}

					if hasSensitiveData(msg) {
						pass.Reportf(lit.Pos(), "log message must not contain sensitive data")
					}

					return true
				})
			}

			return true
		})
	}

	return nil, nil
}

func isLowerStart(s string) bool {
	r, _ := utf8.DecodeRuneInString(s)

	if unicode.IsLetter(r) && !unicode.IsLower(r) {
		return false
	}

	return true
}

func isEnglishOnly(s string) bool {
	for _, r := range s {
		if unicode.IsLetter(r) && !unicode.Is(unicode.Latin, r) {
			return false
		}
	}

	return true
}

func hasSpecialCharsOrEmojis(s string) bool {
	for _, r := range s {
		if !unicode.IsLetter(r) && !unicode.IsDigit(r) && !unicode.IsSpace(r) {
			return true
		}
	}

	return false
}

var sensitiveWords = []string{"password", "api_key", "token", "secret"}

func hasSensitiveData(s string) bool {
	lowerStr := strings.ToLower(s)
	for _, word := range sensitiveWords {
		if strings.Contains(lowerStr, word) {
			return true
		}
	}

	return false
}
