// qalx_qre_test.go - Basic QRE and mesh validation tests for QALX
package coherra

import (
	"testing"
)

func TestComputeQRE(t *testing.T) {
	glyph := LyraGlyph{Emotion: "trust", Intensity: 1.0, EthicsScore: 1.0, Timestamp: 1234567890}
	metrics := InitializeQuantumMetricsWithGlyph(glyph)
	qre := ComputeQRE(metrics, glyph, 1.0, 1.0)
	if qre < 2.0 {
		t.Errorf("QRE too low: got %f, want >= 2.0", qre)
	}
}

func TestMeshNodeValidation(t *testing.T) {
	glyph := LyraGlyph{Emotion: "trust", Intensity: 1.0, EthicsScore: 1.0, Timestamp: 1234567890}
	metrics := InitializeQuantumMetricsWithGlyph(glyph)
	node := GenerateMeshNode(metrics)
	err := ValidateMeshNode(node, 1.0, glyph, 1.0)
	if err != nil {
		t.Errorf("Mesh node validation failed: %v", err)
	}
}
