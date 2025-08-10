// mesh.go - Mesh network logic for QALX
package coherra

import (
	"encoding/base64"
	"fmt"
)

type QuantumMeshNode struct {
	ID               string
	Metrics          QuantumMetrics
	Pattern          string
	CoherenceHistory []float64
	State            string
	Timestamp        int64
}

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

func GeneratePattern(metrics QuantumMetrics) string {
	return base64.StdEncoding.EncodeToString([]byte{
		byte(metrics.Coherence * 100),
		byte(metrics.Phase * 100),
		byte(metrics.Amplitude * 100),
	})
}

func UpdateCoherenceHistory(node *QuantumMeshNode, newCoherence float64) {
	node.CoherenceHistory = append(node.CoherenceHistory, newCoherence)
	node.Metrics.Coherence = newCoherence
}

func EvolveQuantumMetrics(metrics *QuantumMetrics, newTimestamp int64) {
	phaseShift := float64(newTimestamp%360) * (3.141592653589793 / 180)
	metrics.Phase += phaseShift
	metrics.Timestamp = newTimestamp
	metrics.EntropyScore *= 1 + (phaseShift / 3.141592653589793 * 0.01)
	metrics.Pattern = metrics.Pattern + ":t" + fmt.Sprintf("%d", newTimestamp)
}

func ValidateMeshNode(node QuantumMeshNode, resonance float64, glyph LyraGlyph, meshScore float64) error {
	if node.Metrics.Coherence < MinCoherence {
		return &QALXError{"Mesh node coherence below threshold"}
	}
	if node.State != "active" {
		return &QALXError{"Mesh node is not active"}
	}
	if !QALXValidateQuantumSecurity(node.Metrics, resonance, glyph, meshScore) {
		return &QALXError{"Quantum security validation failed"}
	}
	return nil
}

func ValidatePattern(pattern string, history []string) error {
	for _, p := range history {
		if p == pattern {
			return &QALXError{"Pattern is not unique in mesh history"}
		}
	}
	if len(pattern) < 8 {
		return &QALXError{"Pattern entropy too low"}
	}
	return nil
}

type MeshNetwork struct {
	Nodes map[string]QuantumMeshNode
}

func NewMeshNetwork() *MeshNetwork {
	return &MeshNetwork{Nodes: make(map[string]QuantumMeshNode)}
}

func (net *MeshNetwork) AddNode(node QuantumMeshNode) {
	net.Nodes[node.ID] = node
}

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
	meshScore := 1.0 // TODO: Replace with real mesh score logic if available
	return ValidateMeshNode(toNode, 1.0, defaultGlyph, meshScore)
}

func (net *MeshNetwork) ValidateAllNodes() map[string]error {
	results := make(map[string]error)
	for id, node := range net.Nodes {
		defaultGlyph := LyraGlyph{Emotion: "trust", Intensity: 1.0, EthicsScore: 1.0, Timestamp: node.Timestamp}
		meshScore := 1.0 // TODO: Replace with real mesh score logic if available
		results[id] = ValidateMeshNode(node, 1.0, defaultGlyph, meshScore)
	}
	return results
}

func (net *MeshNetwork) RevokeNode(nodeID string, reason string) {
	node, ok := net.Nodes[nodeID]
	if ok {
		node.State = "revoked"
		node.Pattern = node.Pattern + ":revoked"
		net.Nodes[nodeID] = node
	}
}

func (net *MeshNetwork) PropagateRevocation(nodeID string) {
	for id, node := range net.Nodes {
		if id != nodeID {
			node.State = "revoked"
			node.Pattern = node.Pattern + ":revoked"
			net.Nodes[id] = node
		}
	}
}
