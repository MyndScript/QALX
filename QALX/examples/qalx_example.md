// qalx_example.go - Example usage of QALX core logic
package main

import (
	"coherra"
	"fmt"
)

func main() {
	glyph := coherra.LyraGlyph{Emotion: "trust", Intensity: 1.0, EthicsScore: 1.0, Timestamp: 1234567890}
	metrics := coherra.InitializeQuantumMetricsWithGlyph(glyph)
	key, err := coherra.QALXGenerateSecureKey(metrics, glyph, 1.0)
	if err != nil {
		panic(err)
	}
	node := coherra.GenerateMeshNode(metrics)
	err = coherra.ValidateMeshNode(node, 1.0, glyph, 1.0)
	if err != nil {
		panic(err)
	}
	qre := coherra.ComputeQRE(metrics, glyph, 1.0, 1.0)
	fmt.Println("Quantum Key Signature:", coherra.QALXSign(key))
	fmt.Printf("QRE: %.2f\n", qre)
}
