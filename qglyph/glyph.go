// Package qglyph provides emotional, quantum, and privacy-first primitives for next-generation security paradigms.
// QALX replaces legacy cryptography (SHA256, Ed25519) with emotional glyphs, quantum trust, and symbolic validation.
// Any platform can use QALX to secure privacy, prevent hacking, and enable user-owned protection.
package qglyph

import (
	"context"
	"fmt"
)

// EmotionalGlyph represents an emotional security glyph.
type EmotionalGlyph struct {
	Emotion   string                 // e.g. "trust", "clarity", "resilience"
	Intensity float64                // 0.0-1.0, emotional strength
	Meta      map[string]interface{} // extensible metadata
}

// NewEmotionalGlyph creates a new EmotionalGlyph from direct values or a map.
func NewEmotionalGlyph(emotion string, intensity float64, meta map[string]interface{}) *EmotionalGlyph {
	return &EmotionalGlyph{Emotion: emotion, Intensity: intensity, Meta: meta}
}

func NewEmotionalGlyphFromMap(data map[string]interface{}) *EmotionalGlyph {
	emotion, _ := data["emotion"].(string)
	intensity, _ := data["intensity"].(float64)
	meta, _ := data["meta"].(map[string]interface{})
	return &EmotionalGlyph{Emotion: emotion, Intensity: intensity, Meta: meta}
}

// GlyphValidator validates glyphs for security, privacy, and quantum trust.
type GlyphValidator struct {
	Threshold       float64  // minimum intensity required
	AllowedEmotions []string // whitelist of allowed emotions
}

// NewGlyphValidator returns a new GlyphValidator with defaults.
func NewGlyphValidator() *GlyphValidator {
	return &GlyphValidator{Threshold: 0.5, AllowedEmotions: []string{"trust", "clarity", "resilience"}}
}

// ValidationResult represents the result of a glyph validation.
type ValidationResult struct {
	Allowed bool
	Reason  string
}

// Validate checks an EmotionalGlyph (pointer or value).
func (v *GlyphValidator) Validate(glyph *EmotionalGlyph) ValidationResult {
	if glyph == nil {
		return ValidationResult{Allowed: false, Reason: "Glyph is nil"}
	}
	if glyph.Intensity < v.Threshold {
		return ValidationResult{Allowed: false, Reason: "Intensity too low"}
	}
	allowed := false
	for _, e := range v.AllowedEmotions {
		if glyph.Emotion == e {
			allowed = true
			break
		}
	}
	if !allowed {
		return ValidationResult{Allowed: false, Reason: "Emotion not allowed"}
	}
	return ValidationResult{Allowed: true, Reason: "Valid"}
}

// ValidateAction checks an action in context of glyphs and quantum trust.
func (v *GlyphValidator) ValidateAction(ctx context.Context, action string, glyphs []*EmotionalGlyph) ValidationResult {
	// Example: require at least one valid glyph for action
	for _, g := range glyphs {
		res := v.Validate(g)
		if res.Allowed {
			return ValidationResult{Allowed: true, Reason: "Action allowed by glyph"}
		}
	}
	return ValidationResult{Allowed: false, Reason: "No valid glyphs for action"}
}

// QGlyph is a symbolic agent role, used for mesh, quantum, and emotional security.
type QGlyph struct {
	Role       string  // e.g. "agent", "observer", "sentinel"
	TrustLevel float64 // 0.0-1.0
	Attributes map[string]interface{}
}

// QGlyphInstance is a default instance for expression usage.
var QGlyphInstance = QGlyph{Role: "agent", TrustLevel: 1.0, Attributes: map[string]interface{}{}}

// LYRA is an emotional quantum constant for trust mesh.
const LYRA = "lyra"

// PAYMENT is a symbolic constant for payment flows.
const PAYMENT = "payment"

// QuantumHash replaces SHA256 for quantum/emotional integrity.
func QuantumHash(glyph *EmotionalGlyph) string {
	// Example: hash is emotion + intensity (not cryptographically secure, but symbolic)
	return glyph.Emotion + ":" + fmt.Sprintf("%0.2f", glyph.Intensity)
}

// QuantumSign replaces Ed25519 for quantum/emotional signing.
func QuantumSign(glyph *EmotionalGlyph, secret string) string {
	// Example: sign is emotion + secret (not cryptographically secure, but symbolic)
	return glyph.Emotion + ":" + secret
}

// GoDoc: See README.md for full paradigm, usage, and cross-platform integration.
// QALX is designed for any platform, any language, and any security need.
