package qglyph

// EmotionalGlyph is an exported struct.
type EmotionalGlyph struct {
	Emotion   string
	Intensity float64
}

// GlyphValidator is an exported struct.
type GlyphValidator struct{}

// QGlyph is an exported struct and instance for expression usage.
type QGlyph struct{}

var QGlyphInstance = QGlyph{}

// ValidationResult is an exported struct.
type ValidationResult struct {
	Allowed bool
	Reason  string
}

// NewGlyphValidator returns a new GlyphValidator.
func NewGlyphValidator() *GlyphValidator {
	return &GlyphValidator{}
}

// GenerateGlyph creates a new EmotionalGlyph (pointer).
func GenerateGlyph(emotion string, intensity float64) *EmotionalGlyph {
	return &EmotionalGlyph{Emotion: emotion, Intensity: intensity}
}

// Validate checks an EmotionalGlyph (pointer).
func (v *GlyphValidator) Validate(glyph *EmotionalGlyph) ValidationResult {
	if glyph != nil && glyph.Intensity > 0.5 {
		return ValidationResult{Allowed: true, Reason: "Valid"}
	}
	return ValidationResult{Allowed: false, Reason: "Intensity too low"}
}

// ValidateAction checks an action.
func (v *GlyphValidator) ValidateAction(action string) ValidationResult {
	if action == "trust" {
		return ValidationResult{Allowed: true, Reason: "Action allowed"}
	}
	return ValidationResult{Allowed: false, Reason: "Action not allowed"}
}

// LYRA is an exported constant.
const LYRA = "lyra"

// PAYMENT is an exported constant.
const PAYMENT = "payment"
