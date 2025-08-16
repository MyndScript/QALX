# QALX: Quantum-Emotional Security Paradigm

[![SLSA Go releaser](https://github.com/MyndScript/QALX/actions/workflows/go-ossf-slsa3-publish.yml/badge.svg)](https://github.com/MyndScript/QALX/actions/workflows/go-ossf-slsa3-publish.yml) [![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE) [![Go Report Card](https://goreportcard.com/badge/github.com/MyndScript/QALX)](https://goreportcard.com/report/github.com/MyndScript/QALX)
[![Release](https://img.shields.io/github/v/release/MyndScript/QALX)](https://github.com/MyndScript/QALX/releases)

## Vision
Security as Empathy. Quantum-Resilient. Emotionally Tuned.

QALX is a living, ethically aware security organism. It combines quantum-resistant cryptography, emotional modulation (LYRA), and mesh network propagation to create proactive, multidimensional security.

## Features
- Quantum-secure key generation
- LYRA glyph modulation
- Mesh network propagation and validation
- Ethical revocation and stewardship
- Emotional/quantum primitives for any platform
- Cross-language API (Go, Python, Node.js, etc.)

## How It Works
QALX is modular:
- `core/qalx.go`: Core types and metric initializers
- `core/lyra.go`: LYRA glyph logic and harmonics
- `core/mesh.go`: Mesh node/network logic and validation
- `core/entropy.go`: Entropy pool, key generation, and QRE security
- `qglyph/glyph.go`: Emotional, quantum, and privacy-first primitives

## API Documentation & Examples

### Emotional & Quantum Primitives (Go)
```go
import "github.com/MyndScript/QALX/qglyph"

// Create an emotional glyph
glyph := qglyph.NewEmotionalGlyph("trust", 0.95, map[string]interface{}{"context": "login"})

// Validate glyph
validator := qglyph.NewGlyphValidator()
result := validator.Validate(glyph)
if !result.Allowed {
    panic(result.Reason)
}

// Quantum hash and sign
hash := qglyph.QuantumHash(glyph)
signature := qglyph.QuantumSign(glyph, "my-secret")
```

### Mesh Logic Example
```go
import "github.com/MyndScript/QALX/core/mesh"
metrics := QuantumMetrics{Coherence: 0.95, Phase: 0.5, Amplitude: 0.8}
node := mesh.GenerateMeshNode(metrics)
net := mesh.NewMeshNetwork()
net.AddNode(node)
err := mesh.ValidateMeshNode(node, 1.0, LyraGlyph{Emotion: "trust", Intensity: 1.0, EthicsScore: 1.0, Timestamp: node.Timestamp}, 1.0)
if err != nil {
    // handle error
}
```

### Cross-Platform Usage
#### Python
```python
import subprocess
glyph = {"emotion": "trust", "intensity": 0.95, "meta": {"context": "login"}}
subprocess.run(["./qalx-linux-amd64", "--emotion", glyph["emotion"], "--intensity", str(glyph["intensity"]), "--ethics", "0.95"])
```
#### Node.js
```js
const { execFile } = require('child_process');
execFile('./qalx-linux-amd64', ['--emotion', 'trust', '--intensity', '0.95', '--ethics', '0.95'], (err, stdout, stderr) => {
    console.log(stdout);
});
```

## Quantum Resistance Entropy (QRE)
QRE is the heart of QALX’s quantum security. It combines classical and quantum metrics with emotional glyph modulation.

**QRE Formula:**

    QRE = log₂(Eₛ × Pₚ × Φₕ × Ψₑ)

- **Eₛ**: Entropy spread across keyspace (Grover resistance)
- **Pₚ**: Period obfuscation factor (Shor resistance)
- **Φₕ**: Hybrid mimicry distortion (wave-emulated attack resistance)
- **Ψₑ**: Emotional glyph modulation (LYRA)

QRE is used in all key generation and mesh validation logic.

## Emotional & Ethical Modulation
QALX uses emotional states (via LYRA glyphs) and ethical scores to modulate cryptographic parameters and mesh propagation. Emotional intensity and ethics scores directly influence entropy, validation, and node behavior.

**Impact:**
- High emotional intensity increases entropy quality and mesh coherence.
- Ethics score modulates quantum resistance and validation scores.
- Emotional states propagate through the mesh, affecting node trust and revocation logic.

**Adversarial Models:**
- QALX anticipates adversaries that may attempt to mimic or manipulate emotional states. Future tests simulate these scenarios to ensure resilience.

**Stewardship & Audit:**
- Contributors are stewards of emotional and ethical integrity.
- All changes to LYRA logic or mesh propagation require rationale and review.
- Ongoing audit and algorithm evolution are encouraged—see CONTRIBUTING.md for stewardship roles and upgrade guidelines.

## Quickstart
1. Initialize Go module:
    ```sh
    go mod init qalx
    go mod tidy
    ```
2. Install dependencies:
    ```sh
    go get github.com/google/uuid
    ```
3. See `core/qalx.go` and `qglyph/glyph.go` for main algorithm and usage examples.

## Getting Started
See `core/qalx.go` and `qglyph/glyph.go` for the main algorithm. Explore `docs/manifesto.md` for the philosophical and ethical framing.

## Build and Run
Build for your platform (example for Linux amd64):
```sh
go build -o qalx-linux-amd64 ./main.go
```
Run with CLI flags:
```sh
./qalx-linux-amd64 --emotion trust --intensity 0.95 --ethics 0.95
```

## Install
Download prebuilt binaries from the [Releases](https://github.com/MyndScript/QALX/releases) page, or build from source as above.

## CLI Flags
| Flag      | Type    | Description                | Default   |
|-----------|---------|---------------------------|-----------|
| --emotion | string  | LyraGlyph emotion         | joy       |
| --intensity | float | LyraGlyph intensity       | 1.0       |
| --ethics  | float   | LyraGlyph ethics score    | 0.9       |

## Output
The CLI prints a `QuantumMetrics` struct with all calculated fields.

## QALX CLI Usage

## Use from Python
```python
import subprocess
subprocess.run(["./qalx-linux-amd64", "--emotion", "surprise", "--intensity", "0.8", "--ethics", "0.95"])
```

## Use from Node.js
```js
const { execFile } = require('child_process');
execFile('./qalx-linux-amd64', ['--emotion', 'surprise', '--intensity', '0.8', '--ethics', '0.95'], (err, stdout, stderr) => {
	console.log(stdout);
});
```

## Cross-Language Integration
QALX is designed for any platform, any language, and any security need. See `qglyph/glyph.go` for emotional/quantum primitives and API docs.

## License
MIT © Synchronicity pal
