package qglyph

// Glyph is a sample exported struct for demonstration.
type Glyph struct {
	Emotion   string
	Intensity float64
}

// NewGlyph creates a new Glyph.
func NewGlyph(emotion string, intensity float64) *Glyph {
	return &Glyph{Emotion: emotion, Intensity: intensity}
}
