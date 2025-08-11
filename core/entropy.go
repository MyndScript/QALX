package entropy

import (
	"crypto/rand"
	"crypto/sha512"
	"encoding/base64"
	"math"
)

const EntropyPoolSize = 1024

var entropyPool []byte = initializeEntropyPool()

func initializeEntropyPool() []byte {
	pool := make([]byte, EntropyPoolSize)
	_, err := rand.Read(pool)
	if err != nil {
		panic("Failed to initialize entropy pool")
	}
	return pool
}

// Normalization helper for QRE
func normalize(x float64, min float64, max float64) float64 {
	return math.Max(0.1, math.Min(1.0, (x-min)/(max-min)))
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

func QALXGenerateSecureKey(metrics QuantumMetrics, glyph LyraGlyph, meshScore float64) ([]byte, error) {
	if metrics.Coherence < MinCoherence {
		return nil, ErrInsufficientCoherence
	}
	quantumEntropy := harvestQuantumEntropy(metrics)
	mixedEntropy := mixEntropy(quantumEntropy, glyph, meshScore)
	return amplifyQuantumKey(mixedEntropy, metrics, glyph), nil
}

var ErrInsufficientCoherence = &QALXError{"Insufficient quantum coherence for secure key generation"}

type QALXError struct {
	msg string
}

func (e *QALXError) Error() string {
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

func mixEntropy(quantumEntropy []byte, glyph LyraGlyph, meshScore float64) []byte {
	h := sha512.New()
	h.Write(entropyPool)
	h.Write(quantumEntropy)
	bias := byte((glyph.Intensity * glyph.EthicsScore * meshScore) * 128)
	for i := range quantumEntropy {
		quantumEntropy[i] ^= bias
	}
	h.Write(quantumEntropy)
	entropyPool = h.Sum(nil)
	return entropyPool
}

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

// QRE-based validation: composite and QRE threshold
func QALXValidateQuantumSecurity(metrics QuantumMetrics, resonance float64, glyph LyraGlyph, meshScore float64) bool {
	threshold := 0.5
	compositeThreshold := 0.5
	volatilityBonus := 0.0
	driftComp := 1.0
	if glyph.Emotion == "trust" {
		if glyph.Intensity < 0.90 && metrics.Coherence < 0.90 {
			threshold = 0.35
			compositeThreshold = 0.40
			glyph.EthicsScore += 0.02 * meshScore
			driftComp = 1.0 + 0.01*math.Abs(float64(glyph.Timestamp%1000))/1000
			volatilityBonus = 0.02 * math.Abs(math.Sin(float64(glyph.Timestamp%360)*math.Pi/180))
		} else {
			threshold = 0.4
			compositeThreshold = 0.45
			glyph.EthicsScore += 0.01 * meshScore
		}
	}
	if resonance < threshold {
		return false
	}
	composite := metrics.Coherence*resonance*glyph.EthicsScore*meshScore*driftComp + volatilityBonus
	qre := ComputeQRE(metrics, glyph, meshScore, resonance)
	return composite > compositeThreshold && qre > 2.0
}

func QALXSign(data []byte) string {
	return base64.StdEncoding.EncodeToString(data)
}
