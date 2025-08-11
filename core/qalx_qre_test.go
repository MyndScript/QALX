// qalx_qre_test.go - Basic QRE and mesh validation tests for QALX
package coherra

import (
	"strings"
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
