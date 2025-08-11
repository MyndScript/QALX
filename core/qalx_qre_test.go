package coherra

import (
	"math"
	"strings"
	"testing"
)

func BenchmarkMeshNodeStateCheck(b *testing.B) {
	glyph := LyraGlyph{Emotion: "trust", Intensity: 1.0, EthicsScore: 1.0, Timestamp: 1234567890}
	metrics := InitializeQuantumMetricsWithGlyph(glyph)
	node := GenerateMeshNode(metrics)
	for i := 0; i < b.N; i++ {
		_ = node.State == "active"
	}
}

func BenchmarkMeshNodeCoherenceCheck(b *testing.B) {
	glyph := LyraGlyph{Emotion: "trust", Intensity: 1.0, EthicsScore: 1.0, Timestamp: 1234567890}
	metrics := InitializeQuantumMetricsWithGlyph(glyph)
	node := GenerateMeshNode(metrics)
	for i := 0; i < b.N; i++ {
		_ = node.Metrics.Coherence > 0.85
	}
}

func BenchmarkPatternNormalization(b *testing.B) {
	pattern := "AbC12345"
	for i := 0; i < b.N; i++ {
		_ = strings.ToLower(pattern)
	}
}

func BenchmarkQREComponents(b *testing.B) {
	glyph := LyraGlyph{Emotion: "trust", Intensity: 1.0, EthicsScore: 1.0, Timestamp: 1234567890}
	metrics := InitializeQuantumMetricsWithGlyph(glyph)
	for i := 0; i < b.N; i++ {
		_ = metrics.EntropyScore * metrics.EntropyQuality
		_ = math.Abs(math.Sin(metrics.PhaseShift * Phi))
		_ = math.Abs(math.Sin(metrics.Phase * Phi))
		_ = glyph.Intensity * glyph.EthicsScore
	}
}

func TestRevokedNodeValidation(t *testing.T) {
	glyph := LyraGlyph{Emotion: "trust", Intensity: 1.0, EthicsScore: 1.0, Timestamp: 1234567890}
	metrics := InitializeQuantumMetricsWithGlyph(glyph)
	node := GenerateMeshNode(metrics)
	node.State = "revoked"
	err := ValidateMeshNode(node, 1.0, glyph, 1.0)
	if err == nil {
		t.Error("Expected error for revoked node state")
	}
}

func TestMalformedGlyphValidation(t *testing.T) {
	glyph := LyraGlyph{Emotion: "", Intensity: -1.0, EthicsScore: -0.5, Timestamp: 0}
	metrics := InitializeQuantumMetricsWithGlyph(glyph)
	node := GenerateMeshNode(metrics)
	err := ValidateMeshNode(node, 0.0, glyph, 0.0)
	if err == nil {
		t.Error("Expected error for malformed glyph and zero scores")
	}
}

func TestPatternCollisionHandling(t *testing.T) {
	history := []string{"abc12345", "def67890"}
	err := ValidatePattern("abc12345", history)
	if err == nil {
		t.Error("Expected error for pattern collision")
	}
	err = ValidatePattern("unique9876", history)
	if err != nil {
		t.Errorf("Unexpected error for unique pattern: %v", err)
	}
}

func TestQuantumHybridAttackSim(t *testing.T) {
	// Simulate a hybrid quantum-wave attack scenario
	glyph := LyraGlyph{Emotion: "uncertainty", Intensity: 0.8, EthicsScore: 0.7, Timestamp: 9999999999}
	metrics := InitializeQuantumMetricsWithGlyph(glyph)
	// Simulated attack vector: rapid phase shifts, low amplitude, high entropy
	metrics.Phase += math.Pi
	metrics.Amplitude = 0.1
	metrics.EntropyScore = 1.0
	someThreshold := 2.0 // Future-proof threshold for QRE
	qre := ComputeQRE(metrics, glyph, 0.5, 2.0)
	if qre < someThreshold {
		t.Errorf("QRE failed under hybrid quantum attack: got %f, want >= %f", qre, someThreshold)
	}
}

// qalx_qre_test.go - Basic QRE and mesh validation tests for QALX

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

func TestValidatePattern_TableDriven(t *testing.T) {
	cases := []struct {
		pattern string
		history []string
		wantErr bool
	}{
		{"abc12345", []string{"xyz12345"}, false},
		{"abc12345", []string{"abc12345"}, true},
		{"short", []string{}, true},
		{"ABC12345", []string{"abc12345"}, true}, // case-insensitive collision
	}
	for _, tc := range cases {
		err := ValidatePattern(tc.pattern, tc.history)
		if (err != nil) != tc.wantErr {
			t.Errorf("pattern: %s, wantErr: %v, got: %v", tc.pattern, tc.wantErr, err)
		}
	}
}

func TestMeshNodeRevocation(t *testing.T) {
	glyph := LyraGlyph{Emotion: "trust", Intensity: 1.0, EthicsScore: 1.0, Timestamp: 1234567890}
	metrics := InitializeQuantumMetricsWithGlyph(glyph)
	node := GenerateMeshNode(metrics)
	net := NewMeshNetwork()
	net.AddNode(node)
	net.RevokeNode(node.ID, "test reason")
	revoked := net.Nodes[node.ID]
	if revoked.State != "revoked" {
		t.Errorf("Node state not revoked: got %s", revoked.State)
	}
	if !strings.HasSuffix(revoked.Pattern, ":revoked") {
		t.Errorf("Node pattern not marked revoked: got %s", revoked.Pattern)
	}
	// Should not double append
	net.RevokeNode(node.ID, "test reason")
	if strings.Count(revoked.Pattern, ":revoked") > 1 {
		t.Errorf("Node pattern has multiple revoked tags: got %s", revoked.Pattern)
	}
}

func TestValidatePattern_EdgeCases(t *testing.T) {
	// Zero/negative values
	err := ValidatePattern("", []string{})
	if err == nil {
		t.Error("Expected error for empty pattern")
	}
	err = ValidatePattern("short", []string{})
	if err == nil {
		t.Error("Expected error for short pattern")
	}
	err = ValidatePattern("validpattern", []string{"validpattern"})
	if err == nil {
		t.Error("Expected error for duplicate pattern")
	}
}

func TestMeshNodeValidation_EdgeCases(t *testing.T) {
	glyph := LyraGlyph{Emotion: "trust", Intensity: 0.0, EthicsScore: 0.0, Timestamp: 1234567890}
	metrics := InitializeQuantumMetricsWithGlyph(glyph)
	metrics.Coherence = -1.0
	node := GenerateMeshNode(metrics)
	err := ValidateMeshNode(node, 0.0, glyph, 0.0)
	if err == nil {
		t.Error("Expected error for negative coherence and zero scores")
	}
	node.State = "revoked"
	err = ValidateMeshNode(node, 1.0, glyph, 1.0)
	if err == nil {
		t.Error("Expected error for revoked node state")
	}
}

func BenchmarkComputeQRE(b *testing.B) {
	glyph := LyraGlyph{Emotion: "trust", Intensity: 1.0, EthicsScore: 1.0, Timestamp: 1234567890}
	metrics := InitializeQuantumMetricsWithGlyph(glyph)
	for i := 0; i < b.N; i++ {
		_ = ComputeQRE(metrics, glyph, 1.0, 1.0)
	}
}

func BenchmarkMeshNodeValidation(b *testing.B) {
	glyph := LyraGlyph{Emotion: "trust", Intensity: 1.0, EthicsScore: 1.0, Timestamp: 1234567890}
	metrics := InitializeQuantumMetricsWithGlyph(glyph)
	node := GenerateMeshNode(metrics)
	for i := 0; i < b.N; i++ {
		_ = ValidateMeshNode(node, 1.0, glyph, 1.0)
	}
}
