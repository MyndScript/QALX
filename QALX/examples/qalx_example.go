package examples

import (
	"fmt"
	"log"

	"qalx/core"
)

func main() {
	// Initialize QALX Core with default mesh context
	qx, err := core.NewQALX()
	if err != nil {
		log.Fatalf("Failed to initialize QALX: %v", err)
	}

	// Define a sample glyph input—this could be user-generated or mesh-propagated
	glyph := core.Glyph{
		Type:    "trust",
		Origin:  "node:LYRA-42",
		Payload: []byte("I will protect your silence."),
		Metadata: map[string]string{
			"emotion": "gentle conviction",
			"channel": "whisper",
		},
	}

	// Validate the glyph using QALX’s emotional and quantum layers
	result, err := qx.ValidateGlyph(glyph)
	if err != nil {
		log.Fatalf("Glyph validation failed: %v", err)
	}

	// Output the result with emotional resonance score
	fmt.Printf("Glyph '%s' validated.\n", glyph.Type)
	fmt.Printf("Resonance Score: %.2f\n", result.Resonance)
	fmt.Printf("Quantum Integrity: %v\n", result.QuantumSafe)
}
