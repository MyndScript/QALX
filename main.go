package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	coherra "github.com/MyndScript/QALX/core"
)

func main() {
	// Example CLI flags
	emotion := flag.String("emotion", "joy", "LyraGlyph emotion")
	intensity := flag.Float64("intensity", 1.0, "LyraGlyph intensity")
	ethics := flag.Float64("ethics", 0.9, "LyraGlyph ethics score")
	flag.Parse()

	// Create a LyraGlyph from CLI args
	glyph := coherra.LyraGlyph{
		Emotion:     *emotion,
		Intensity:   *intensity,
		EthicsScore: *ethics,
		Timestamp:   time.Now().Unix(),
	}

	// Generate metrics and print
	metrics := coherra.InitializeQuantumMetricsWithGlyph(glyph)
	fmt.Printf("QuantumMetrics: %+v\n", metrics)
	os.Exit(0)
}
