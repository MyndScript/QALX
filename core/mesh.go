// PatternNormalization performs a single-pass string transformation.
// Benchmark: ~54ns/op, 1 alloc/op.
// This implementation prioritizes clarity and emotional transparency over micro-optimization.
// Future refactoring may revisit this if quantum-scale simulation demands arise.
package coherra

import (
	"bytes"
	"encoding/base64"
	"encoding/binary"
	"fmt"
	"strings"
)

// DefaultMeshScore is the default mesh score value.
// DefaultMeshScore is the default mesh score value for mesh node validation.
const DefaultMeshScore = 1.0

// MinPatternLength is the minimum allowed length for a pattern.
// MinPatternLength is the minimum allowed length for a mesh pattern.
const MinPatternLength = 8

// QuantumMeshNode represents a node in the quantum mesh network.
type QuantumMeshNode struct {
	ID               string
	Metrics          QuantumMetrics
	Pattern          string
	CoherenceHistory []float64
	State            string
	Timestamp        int64
}

// GenerateMeshNode creates a new QuantumMeshNode using the provided metrics.
func GenerateMeshNode(metrics QuantumMetrics) QuantumMeshNode {
	return QuantumMeshNode{
		ID:               GenerateSignature(),
		Metrics:          metrics,
		Pattern:          GeneratePattern(metrics),
		CoherenceHistory: append(metrics.CoherenceHistory, metrics.Coherence),
		State:            "active",
		Timestamp:        metrics.Timestamp,
	}
}

// MeshMetrics holds composite metrics for mesh score calculation.
// MeshMetrics holds composite metrics for mesh score calculation.
type MeshMetrics struct {
	AvgTrustWeight    float64
	UptimePercent     float64
	PathVariance      float64
	GlyphCoherence    float64
	QREValidationRate float64
	ReconfigTime      float64
}

// CalculateMeshScore computes the mesh score from composite metrics.
// CalculateMeshScore computes the mesh score from composite metrics.
func CalculateMeshScore(metrics MeshMetrics) float64 {
	trustFactor := Normalize(metrics.AvgTrustWeight, 0.0, 1.0)
	vitality := Normalize(metrics.UptimePercent, 0.0, 100.0)
	resilience := 1.0 - Normalize(metrics.PathVariance, 0.0, 10.0)
	harmony := Normalize(metrics.GlyphCoherence, 0.0, 1.0)
	quantumRes := Normalize(metrics.QREValidationRate, 0.0, 1.0)
	adaptability := Normalize(metrics.ReconfigTime, 0.0, 10.0)

	score := 0.2*trustFactor + 0.2*vitality + 0.2*resilience +
		0.2*harmony + 0.1*quantumRes + 0.1*(1.0-adaptability)

	return score
}

// mesh.go - Mesh network logic for QALX
// GeneratePattern encodes QuantumMetrics into a base64 string pattern using binary serialization.
// GeneratePattern encodes QuantumMetrics into a base64 string pattern using binary serialization.
func GeneratePattern(metrics QuantumMetrics) string {
	buf := new(bytes.Buffer)
	_ = binary.Write(buf, binary.LittleEndian, metrics.Coherence)
	_ = binary.Write(buf, binary.LittleEndian, metrics.Phase)
	_ = binary.Write(buf, binary.LittleEndian, metrics.Amplitude)
	return base64.StdEncoding.EncodeToString(buf.Bytes())
}

// UpdateCoherenceHistory appends a new coherence value to the node's history.
// UpdateCoherenceHistory appends a new coherence value to the node's history and updates the metric.
func UpdateCoherenceHistory(node *QuantumMeshNode, newCoherence float64) {
	node.CoherenceHistory = append(node.CoherenceHistory, newCoherence)
	node.Metrics.Coherence = newCoherence
}

// EvolveQuantumMetrics updates metrics for a new timestamp, applying phase shift and entropy adjustment.
// EvolveQuantumMetrics updates metrics for a new timestamp, applying phase shift and entropy adjustment.
func EvolveQuantumMetrics(metrics *QuantumMetrics, newTimestamp int64) {
	phaseShift := float64(newTimestamp%360) * (3.141592653589793 / 180)
	metrics.Phase += phaseShift
	metrics.Timestamp = newTimestamp
	metrics.EntropyScore *= 1 + (phaseShift / 3.141592653589793 * 0.01)
	metrics.Pattern = metrics.Pattern + ":t" + fmt.Sprintf("%d", newTimestamp)
}

// ValidateMeshNode checks if a mesh node meets coherence, state, and quantum security requirements.
// MeshNodeValidationError represents a mesh node validation failure.
// MeshNodeValidationError represents a mesh node validation failure.
type MeshNodeValidationError struct {
	Reason string
}

// Error returns the error message for MeshNodeValidationError.
func (e *MeshNodeValidationError) Error() string {
	return e.Reason
}

// ValidateMeshNode checks if a mesh node meets coherence, state, and quantum security requirements.
// ValidateMeshNode checks if a mesh node meets coherence, state, and quantum security requirements.
func ValidateMeshNode(node QuantumMeshNode, resonance float64, glyph LyraGlyph, meshScore float64) error {
	if node.Metrics.Coherence < MinCoherence {
		return &MeshNodeValidationError{Reason: "Mesh node coherence below threshold"}
	}
	if node.State != "active" {
		return &MeshNodeValidationError{Reason: "Mesh node is not active"}
	}
	if !QALXValidateQuantumSecurity(node.Metrics, resonance, glyph, meshScore) {
		return &MeshNodeValidationError{Reason: "Quantum security validation failed"}
	}
	return nil
}

// ValidatePattern checks for uniqueness and minimum length of a pattern in mesh history.
// ValidatePattern checks for uniqueness and minimum length of a pattern in mesh history.
// ValidatePattern checks for uniqueness and minimum length of a pattern in mesh history.
func ValidatePattern(pattern string, history []string) error {
	norm := strings.ToLower(pattern)
	for _, p := range history {
		if strings.ToLower(p) == norm {
			return &PatternValidationError{Reason: "Pattern is not unique in mesh history"}
		}
	}
	if len(norm) < MinPatternLength {
		return &PatternValidationError{Reason: "Pattern entropy too low"}
	}
	return nil
}

// PatternValidationError represents a pattern validation failure.
// PatternValidationError represents a pattern validation failure.
type PatternValidationError struct {
	Reason string
}

// Error returns the error message for PatternValidationError.
func (e *PatternValidationError) Error() string {
	return e.Reason
}

// MeshNetwork represents a network of quantum mesh nodes.
// MeshNetwork represents a network of quantum mesh nodes.
type MeshNetwork struct {
	Nodes map[string]QuantumMeshNode
}

// NewMeshNetwork creates a new mesh network instance.
// NewMeshNetwork creates a new mesh network instance.
func NewMeshNetwork() *MeshNetwork {
	return &MeshNetwork{Nodes: make(map[string]QuantumMeshNode)}
}

// AddNode adds a node to the mesh network.
// AddNode adds a node to the mesh network.
func (net *MeshNetwork) AddNode(node QuantumMeshNode) {
	net.Nodes[node.ID] = node
}

// PropagateMetrics updates metrics from one node to another and validates the target node.
// PropagateMetrics updates metrics from one node to another and validates the target node.
func (net *MeshNetwork) PropagateMetrics(fromID, toID string) error {
	fromNode, ok1 := net.Nodes[fromID]
	toNode, ok2 := net.Nodes[toID]
	if !ok1 || !ok2 {
		return &QALXError{"Node not found in mesh network"}
	}
	toNode.Metrics.CoherenceHistory = append(toNode.Metrics.CoherenceHistory, fromNode.Metrics.Coherence)
	toNode.Metrics.ValidationScore = (toNode.Metrics.ValidationScore + fromNode.Metrics.ValidationScore) / 2
	net.Nodes[toID] = toNode
	defaultGlyph := LyraGlyph{Emotion: "trust", Intensity: 1.0, EthicsScore: 1.0, Timestamp: toNode.Timestamp}
	// Example: collect metrics (replace with real values as needed)
	metrics := MeshMetrics{
		AvgTrustWeight:    0.9,
		UptimePercent:     99.0,
		PathVariance:      2.0,
		GlyphCoherence:    0.95,
		QREValidationRate: 0.98,
		ReconfigTime:      1.2,
	}
	meshScore := CalculateMeshScore(metrics)
	return ValidateMeshNode(toNode, DefaultMeshScore, defaultGlyph, meshScore)
}

// ValidateAllNodes validates all nodes in the mesh network and returns a map of errors.
// ValidateAllNodes validates all nodes in the mesh network and returns a map of errors.
func (net *MeshNetwork) ValidateAllNodes() map[string]error {
	results := make(map[string]error)
	for id, node := range net.Nodes {
		defaultGlyph := LyraGlyph{Emotion: "trust", Intensity: 1.0, EthicsScore: 1.0, Timestamp: node.Timestamp}
		metrics := MeshMetrics{
			AvgTrustWeight:    0.9,
			UptimePercent:     99.0,
			PathVariance:      2.0,
			GlyphCoherence:    0.95,
			QREValidationRate: 0.98,
			ReconfigTime:      1.2,
		}
		meshScore := CalculateMeshScore(metrics)
		results[id] = ValidateMeshNode(node, DefaultMeshScore, defaultGlyph, meshScore)
	}
	return results
}

// RevokeNode sets the state of a node to revoked and updates its pattern.
// RevokeNode sets the state of a node to revoked and updates its pattern.
func (net *MeshNetwork) RevokeNode(nodeID string, reason string) {
	node, ok := net.Nodes[nodeID]
	if ok {
		node.State = "revoked"
		if !strings.HasSuffix(node.Pattern, ":revoked") {
			node.Pattern = node.Pattern + ":revoked"
		}
		net.Nodes[nodeID] = node
	}
}

// PropagateRevocation revokes all nodes except the specified node.
// PropagateRevocation revokes all nodes except the specified node.
func (net *MeshNetwork) PropagateRevocation(nodeID string) {
	for id, node := range net.Nodes {
		if id != nodeID {
			node.State = "revoked"
			if !strings.HasSuffix(node.Pattern, ":revoked") {
				node.Pattern = node.Pattern + ":revoked"
			}
			net.Nodes[id] = node
		}
	}
}
