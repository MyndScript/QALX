package coherra

import (
	"math"

	"github.com/google/uuid"
)

// MinCoherence is the minimum allowed coherence for quantum metrics.
const (
	MinCoherence = 0.85
)

// QuantumMetrics holds quantum-related metrics for mesh nodes and cryptography.
type QuantumMetrics struct {
	Coherence          float64
	Phase              float64
	Amplitude          float64
	Harmonics          []float64
	EntropyScore       float64
	CoherenceThreshold float64
	KeyStrength        int
	EntropyQuality     float64
	QuantumResistance  float64
	KeyLength          int
	Signature          string
	Strength           float64
	PhaseShift         float64
	EntropyLevel       int
	Pattern            string
	MeshNodeID         string
	CoherenceHistory   []float64
	ValidationScore    float64
	NodeState          string
	Timestamp          int64
}

// EncryptionMetrics holds encryption-related metrics for quantum security.
type EncryptionMetrics struct {
	KeyStrength        int
	EntropyQuality     float64
	QuantumResistance  float64
	CoherenceThreshold float64
	EntropyScore       float64
	KeyLength          int
	Signature          string
	Harmonics          []float64
	Strength           float64
	PhaseShift         float64
	EntropyLevel       int
	Pattern            string
	MeshNodeID         string
	CoherenceHistory   []float64
	ValidationScore    float64
	NodeState          string
	Timestamp          int64
}

// GenerateSignature creates a new UUID signature for a mesh node or metric.
func GenerateSignature() string {
	return uuid.New().String()
}

// GlyphMemory tracks historical validation and improvement for each glyph
type GlyphMemory struct {
	History []bool
}

func (gm *GlyphMemory) Improving() bool {
	if len(gm.History) < 2 {
		return false
	}
	// If last two attempts were both valid, or last was valid and previous was invalid
	return gm.History[len(gm.History)-1] && (!gm.History[len(gm.History)-2] || gm.History[len(gm.History)-2])
}

func InitializeEncryptionMetricsWithGlyph(glyph LyraGlyph) EncryptionMetrics {
	return EncryptionMetrics{
		KeyStrength:        256,
		EntropyQuality:     0.99,
		QuantumResistance:  0.95,
		CoherenceThreshold: 0.9,
		EntropyScore:       0.98,
		KeyLength:          32,
		Signature:          GenerateSignature(),
		Harmonics:          GenerateDynamicHarmonics(glyph),
		Strength:           0.95,
		PhaseShift:         math.Pi / 4,
		EntropyLevel:       10,
	}
}

func InitializeHighSecurityEncryptionMetricsWithGlyph(glyph LyraGlyph) EncryptionMetrics {
	return EncryptionMetrics{
		KeyStrength:        256,
		EntropyQuality:     0.99,
		QuantumResistance:  0.95,
		CoherenceThreshold: 0.90,
		EntropyScore:       0.98,
		KeyLength:          4096,
		Signature:          GenerateSignature(),
		Harmonics:          GenerateDynamicHarmonics(glyph),
		Strength:           0.95,
		PhaseShift:         math.Pi / 4,
		EntropyLevel:       10,
	}
}

func InitializeQuantumMetricsWithGlyph(glyph LyraGlyph) QuantumMetrics {
	return QuantumMetrics{
		Coherence:          0.99,
		Phase:              math.Pi / 2,
		Amplitude:          1.0,
		Harmonics:          GenerateDynamicHarmonics(glyph),
		EntropyScore:       0.98,
		CoherenceThreshold: 0.90,
		KeyStrength:        256,
		EntropyQuality:     0.99,
		QuantumResistance:  0.95,
		KeyLength:          4096,
		Signature:          GenerateSignature(),
		Strength:           0.95,
		PhaseShift:         math.Pi / 4,
		EntropyLevel:       10,
		Pattern:            "default-pattern",
		MeshNodeID:         GenerateSignature(),
		CoherenceHistory:   []float64{0.99},
		ValidationScore:    1.0,
		NodeState:          "active",
		Timestamp:          glyph.Timestamp,
	}
}
