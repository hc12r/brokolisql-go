package services

import (
	"regexp"
	"strings"
	"unicode"
)

type Normalizer struct {
	MaxLength        int
	PreserveCase     bool
	ReplaceSpaces    bool
	SpaceReplacement string
}

func NewNormalizer() *Normalizer {
	return &Normalizer{
		MaxLength:        64,
		PreserveCase:     false,
		ReplaceSpaces:    true,
		SpaceReplacement: "_",
	}
}

func (n *Normalizer) NormalizeColumnName(name string) string {
	name = strings.TrimSpace(name)

	// Handle spaces first if replacement is enabled
	if n.ReplaceSpaces {
		name = strings.ReplaceAll(name, " ", n.SpaceReplacement)
	}

	var result strings.Builder

	// Replace special characters with the space replacement character
	// but preserve spaces if ReplaceSpaces is false
	re := regexp.MustCompile(`[^a-zA-Z0-9_ ]`)
	if n.ReplaceSpaces {
		name = re.ReplaceAllString(name, n.SpaceReplacement)
	} else {
		name = re.ReplaceAllString(name, "_")
	}

	// Ensure the name starts with a letter or underscore
	if len(name) > 0 && !unicode.IsLetter(rune(name[0])) && name[0] != '_' {
		result.WriteRune('_')
	}

	result.WriteString(name)

	// Apply case transformation
	if !n.PreserveCase {
		name = strings.ToUpper(result.String())
	} else {
		name = result.String()
	}

	// Truncate if longer than max length
	if n.MaxLength > 0 && len(name) > n.MaxLength {
		name = name[:n.MaxLength]
	}

	// Handle empty names
	if name == "" {
		name = "COLUMN"
	}

	return name
}

func (n *Normalizer) NormalizeColumnNames(names []string) []string {
	normalized := make([]string, len(names))

	// Track used names and their counts
	usedNames := make(map[string]int)

	for i, name := range names {
		normalizedName := n.NormalizeColumnName(name)

		// Handle duplicate names by adding a suffix
		if count, exists := usedNames[normalizedName]; exists {
			// For duplicates, add _0, _1, etc.
			normalizedName = normalizedName + "_" + string(rune('0'+count-1))
			usedNames[normalizedName] = 1
			usedNames[normalizedName[:len(normalizedName)-2]]++
		} else {
			usedNames[normalizedName] = 1
		}

		normalized[i] = normalizedName
	}

	return normalized
}
