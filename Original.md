// QHA256.go - Quantum-secure cryptography and mesh logic
//
// Usage Example:
//
//	func main() {
//	    metrics := InitializeQuantumMetrics()
//	    key, err := QHA256GenerateSecureKey(metrics)
//	    if err != nil {
//	        panic(err)
//	    }
//	    node := GenerateMeshNode(metrics)
//	    err = ValidateMeshNode(node)
//	    if err != nil {
//	        panic(err)
//	    }
//	    pattern := GeneratePattern(metrics)
//	    err = ValidatePattern(pattern, []string{})
//	    if err != nil {
//	        panic(err)
//	    }
//	    signature := QHA256Sign(key)
//	    println("Quantum Key Signature:", signature)
//	}
//
// QuantumMetrics holds all quantum and classical metrics for key generation and mesh logic.
// EncryptionMetrics holds classical encryption metrics for compatibility.
// QuantumMeshNode represents a node in the quantum mesh network.
// GenerateMeshNode creates a mesh node from metrics.
// GeneratePattern creates a pattern string from metrics.
// UpdateCoherenceHistory updates a mesh node's coherence history.
// ValidateMeshNode checks mesh node integrity and quantum security.
// ValidatePattern checks pattern uniqueness and entropy.
// QHA256GenerateSecureKey generates a quantum-secure key from metrics.
// QHA256Sign encodes data as a quantum signature.
// QHA256Error is used for error handling.
package coherra

import (
	"crypto/rand"
	"crypto/sha512"
	"encoding/base64"
	"fmt"
	"math"

	"github.com/google/uuid"
)

const (
	EntropyPoolSize = 1024
	MinCoherence    = 0.85
)

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
	ResonanceHistory   []float64 // for smoothing
}

// Generates dynamic harmonics for a glyph (stub)

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
}

// Normalization helper for QRE
func normalize(x float64, min float64, max float64) float64 {
	return math.Max(0.1, math.Min(1.0, (x-min)/(max-min)))
}

// Golden ratio constant for harmonics
var Phi float64 = (1 + math.Sqrt(5)) / 2

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

// Computes Quantum Resistance Entropy (QRE) for future-proof quantum validation (log-sum model)
func ComputeQRE(metrics QuantumMetrics, glyph LyraGlyph, meshScore float64, resonance float64) float64 {
	entropySpread := normalize(metrics.EntropyScore*metrics.EntropyQuality, 0.1, 1.0)
	periodObfuscation := normalize(math.Abs(math.Sin(metrics.PhaseShift*Phi)), 0.1, 1.0)
	hybridDistortion := normalize(math.Abs(math.Sin(metrics.Phase*Phi)), 0.1, 1.0)
	glyphModulation := normalize(glyph.Intensity*glyph.EthicsScore*meshScore, 0.1, 1.0)
	weightedResonance := normalize(1.0/(1.0+resonance), 0.1, 1.0)

	sum := entropySpread + periodObfuscation + hybridDistortion + glyphModulation + weightedResonance
	return math.Log2(sum + 1)
}

// Generates harmonics for quantum metrics
// Deprecated: use GenerateDynamicHarmonics
func GenerateHarmonics() []float64 {
	return []float64{1, Phi, math.Pow(Phi, 2), math.Pow(Phi, 3)}
}

// Generates a unique signature for metrics or keys
func GenerateSignature() string {
	return uuid.New().String()
}

// Initializes EncryptionMetrics with default values
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

// Initializes EncryptionMetrics for high-security use
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

// Helper to initialize QuantumMetrics with all fields
func InitializeQuantumMetricsWithGlyph(glyph LyraGlyph) QuantumMetrics {
	return QuantumMetrics{
		Coherence:         0.99,
		Phase:             math.Pi / 2,
		Amplitude:         1.0,
		Harmonics:         GenerateDynamicHarmonics(glyph),
		EntropyScore:      0.98,
		KeyStrength:       256,
		EntropyQuality:    0.99,
		QuantumResistance: 0.95,
		KeyLength:         4096,
		Signature:         GenerateSignature(),
		Strength:          0.95,
		PhaseShift:        math.Pi / 4,
		EntropyLevel:      10,
		Pattern:           "default-pattern",
		MeshNodeID:        GenerateSignature(),
		CoherenceHistory:  []float64{0.99},
		ValidationScore:   1.0,
		NodeState:         "active",
		Timestamp:         glyph.Timestamp,
		ResonanceHistory:  []float64{},
	}
}

var entropyPool []byte = initializeEntropyPool()

func initializeEntropyPool() []byte {
	pool := make([]byte, EntropyPoolSize)
	_, err := rand.Read(pool)
	if err != nil {
		panic("Failed to initialize entropy pool")
	}
	return pool
}

// Updated to use Adaptive Entropy Amplification
func QHA256GenerateSecureKey(metrics QuantumMetrics, glyph LyraGlyph, meshScore float64) ([]byte, error) {
	if metrics.Coherence < MinCoherence {
		return nil, ErrInsufficientCoherence
	}
	quantumEntropy := harvestQuantumEntropy(metrics)
	mixedEntropy := mixEntropy(quantumEntropy, glyph, meshScore)
	return amplifyQuantumKey(mixedEntropy, metrics, glyph), nil
}

var ErrInsufficientCoherence = &QHA256Error{"Insufficient quantum coherence for secure key generation"}

type QHA256Error struct {
	msg string
}

func (e *QHA256Error) Error() string {
	return e.msg
}

func harvestQuantumEntropy(metrics QuantumMetrics) []byte {
	entropyBuffer := make([]byte, 32)
	quantumFactors := append([]float64{metrics.Coherence, metrics.Phase, metrics.Amplitude}, metrics.Harmonics...)
	for i := 0; i < len(entropyBuffer); i++ {
		quantum := math.Sin(quantumFactors[i%len(quantumFactors)] * math.Pi)
		entropyBuffer[i] = byte((quantum + 1) * 128)
	}
	return entropyBuffer
}

// Adaptive Entropy Amplification: mix entropy with glyph and mesh feedback
func mixEntropy(quantumEntropy []byte, glyph LyraGlyph, meshScore float64) []byte {
	h := sha512.New()
	h.Write(entropyPool)
	h.Write(quantumEntropy)
	// Emotional bias
	bias := byte((glyph.Intensity * glyph.EthicsScore * meshScore) * 128)
	for i := range quantumEntropy {
		quantumEntropy[i] ^= bias
	}
	h.Write(quantumEntropy)
	entropyPool = h.Sum(nil)
	return entropyPool
}

// Emotionally-Evolved Key Generation: modulate key with glyph and harmonics
func amplifyQuantumKey(entropy []byte, metrics QuantumMetrics, glyph LyraGlyph) []byte {
	keyBuffer := make([]byte, 64)
	glyphPhase := math.Sin(glyph.Intensity * glyph.EthicsScore * math.Pi)
	for i := 0; i < len(keyBuffer); i++ {
		harmonicFactor := metrics.Harmonics[i%len(metrics.Harmonics)]
		quantumPhase := math.Sin(metrics.Phase * harmonicFactor * glyphPhase)
		keyBuffer[i] = entropy[i%len(entropy)] ^ byte((quantumPhase+1)*128)
	}
	return keyBuffer
}

// Validation Layer Upgrade: multidimensional composite score
// QRE-based validation: composite and QRE threshold
func QHA256ValidateQuantumSecurity(metrics QuantumMetrics, resonance float64, glyph LyraGlyph, meshScore float64) bool {
	// Adaptive φ resonance threshold by glyph type
	threshold := 0.5
	compositeThreshold := 0.5
	volatilityBonus := 0.0
	driftComp := 1.0
	if glyph.Emotion == "trust" {
		// Edge state calibration
		if glyph.Intensity < 0.90 && metrics.Coherence < 0.90 {
			threshold = 0.35                                                                      // allow more dissonance for trust in Edge
			compositeThreshold = 0.40                                                             // lower composite threshold for trust in Edge
			glyph.EthicsScore += 0.02 * meshScore                                                 // amplify mesh feedback
			driftComp = 1.0 + 0.01*math.Abs(float64(glyph.Timestamp%1000))/1000                   // compensate for timestamp drift
			volatilityBonus = 0.02 * math.Abs(math.Sin(float64(glyph.Timestamp%360)*math.Pi/180)) // pattern volatility bonus
		} else {
			threshold = 0.4
			compositeThreshold = 0.45
			glyph.EthicsScore += 0.01 * meshScore
		}
	}
	if resonance < threshold {
		return false // Glyph is too dissonant
	}
	composite := metrics.Coherence*resonance*glyph.EthicsScore*meshScore*driftComp + volatilityBonus
	qre := ComputeQRE(metrics, glyph, meshScore, resonance)
	// Require both composite and QRE to pass
	return composite > compositeThreshold && qre > 2.0
}

func QHA256Sign(data []byte) string {
	return base64.StdEncoding.EncodeToString(data)
}

// Mesh node and pattern generation logic

type QuantumMeshNode struct {
	ID               string
	Metrics          QuantumMetrics
	Pattern          string
	CoherenceHistory []float64
	State            string
	Timestamp        int64
}

// Generates a new mesh node from QuantumMetrics
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

// Generates a pattern string based on metrics
func GeneratePattern(metrics QuantumMetrics) string {
	return base64.StdEncoding.EncodeToString([]byte{
		byte(metrics.Coherence * 100),
		byte(metrics.Phase * 100),
		byte(metrics.Amplitude * 100),
	})
}

// Updates mesh node coherence history
func UpdateCoherenceHistory(node *QuantumMeshNode, newCoherence float64) {
	node.CoherenceHistory = append(node.CoherenceHistory, newCoherence)
	node.Metrics.Coherence = newCoherence
}

// Temporal entropy: Evolve keys and patterns using timestamps and phase shifts

// Evolves QuantumMetrics over time
func EvolveQuantumMetrics(metrics *QuantumMetrics, newTimestamp int64) {
	phaseShift := float64(newTimestamp%360) * (math.Pi / 180)
	metrics.Phase += phaseShift
	metrics.Timestamp = newTimestamp
	metrics.EntropyScore *= 1 + (phaseShift / math.Pi * 0.01)
	metrics.Pattern = metrics.Pattern + ":t" + fmt.Sprintf("%d", newTimestamp)
}

// Advanced validation and error handling for mesh and pattern

// Validates mesh node integrity and quantum security
func ValidateMeshNode(node QuantumMeshNode, resonance float64, glyph LyraGlyph) error {
	if node.Metrics.Coherence < MinCoherence {
		return &QHA256Error{"Mesh node coherence below threshold"}
	}
	if node.State != "active" {
		return &QHA256Error{"Mesh node is not active"}
	}
	if !QHA256ValidateQuantumSecurity(node.Metrics, resonance, glyph, 1.0) {
		return &QHA256Error{"Quantum security validation failed"}
	}
	return nil
}

// Validates pattern uniqueness and entropy
func ValidatePattern(pattern string, history []string) error {
	for _, p := range history {
		if p == pattern {
			return &QHA256Error{"Pattern is not unique in mesh history"}
		}
	}
	if len(pattern) < 8 {
		return &QHA256Error{"Pattern entropy too low"}
	}
	return nil
}

// Peer-to-peer propagation logic for QuantumMeshNode

type MeshNetwork struct {
	Nodes map[string]QuantumMeshNode
}

// Generates dynamic harmonics for a glyph (stub)

// Initializes a new mesh network
func NewMeshNetwork() *MeshNetwork {
	return &MeshNetwork{Nodes: make(map[string]QuantumMeshNode)}
}

// Adds a node to the mesh network
func (net *MeshNetwork) AddNode(node QuantumMeshNode) {
	net.Nodes[node.ID] = node
}

// Shares metrics between two nodes and validates trust
func (net *MeshNetwork) PropagateMetrics(fromID, toID string) error {
	fromNode, ok1 := net.Nodes[fromID]
	toNode, ok2 := net.Nodes[toID]
	if !ok1 || !ok2 {
		return &QHA256Error{"Node not found in mesh network"}
	}
	// Share metrics (simple merge for demo)
	toNode.Metrics.CoherenceHistory = append(toNode.Metrics.CoherenceHistory, fromNode.Metrics.Coherence)
	toNode.Metrics.ValidationScore = (toNode.Metrics.ValidationScore + fromNode.Metrics.ValidationScore) / 2
	net.Nodes[toID] = toNode
	// Validate trust
	// For demo, use resonance=1.0 and a default glyph
	defaultGlyph := LyraGlyph{Emotion: "trust", Intensity: 1.0, EthicsScore: 1.0, Timestamp: toNode.Timestamp}
	return ValidateMeshNode(toNode, 1.0, defaultGlyph)
}

// Validates all nodes in the mesh network
func (net *MeshNetwork) ValidateAllNodes() map[string]error {
	results := make(map[string]error)
	for id, node := range net.Nodes {
		// For demo, use resonance=1.0 and a default glyph
		defaultGlyph := LyraGlyph{Emotion: "trust", Intensity: 1.0, EthicsScore: 1.0, Timestamp: node.Timestamp}
		results[id] = ValidateMeshNode(node, 1.0, defaultGlyph)
	}
	return results
}

// Revocation protocol: Mark nodes inactive/compromised and propagate revocation

// Marks a mesh node as inactive or compromised
func (net *MeshNetwork) RevokeNode(nodeID string, reason string) {
	node, ok := net.Nodes[nodeID]
	if ok {
		node.State = "revoked"
		node.Pattern = node.Pattern + ":revoked"
		net.Nodes[nodeID] = node
	}
}

// Propagates revocation to all connected nodes
func (net *MeshNetwork) PropagateRevocation(nodeID string) {
	for id, node := range net.Nodes {
		if id != nodeID {
			node.State = "revoked"
			node.Pattern = node.Pattern + ":revoked"
			net.Nodes[id] = node
		}
	}
}

// LYRA integration: Modulate QuantumMetrics using emotional glyphs

type LyraGlyph struct {
	Emotion     string
	Intensity   float64
	EthicsScore float64
	Timestamp   int64
}

// Generates dynamic harmonics for a glyph (stub)
func GenerateDynamicHarmonics(glyph LyraGlyph) []float64 {
	// Real harmonics logic: modulate harmonics using glyph intensity, ethics, emotion, and timestamp
	base := glyph.Intensity * glyph.EthicsScore
	phi := Phi
	t := float64(glyph.Timestamp%1000) / 1000.0
	var emotionFactor float64
	switch glyph.Emotion {
	case "trust":
		emotionFactor = 1.25
	case "hope":
		emotionFactor = 1.1
	case "resolve":
		emotionFactor = 1.15
	default:
		emotionFactor = 1.0
	}
	h1 := base * phi * emotionFactor
	h2 := math.Sin(base*phi+t) * emotionFactor
	h3 := math.Pow(base+phi+t, 2) * 0.5 * emotionFactor
	return []float64{h1, h2, h3}
}

// Modulates QuantumMetrics using a LyraGlyph
func ModulateQuantumMetricsWithLyra(metrics *QuantumMetrics, glyph LyraGlyph) {
	// Example modulation logic
	metrics.EntropyQuality *= glyph.Intensity
	metrics.QuantumResistance += glyph.EthicsScore * 0.01
	metrics.Coherence += glyph.Intensity * 0.01
	metrics.ValidationScore += glyph.EthicsScore * 0.05
	metrics.Timestamp = glyph.Timestamp
	metrics.Pattern = metrics.Pattern + ":" + glyph.Emotion
}
