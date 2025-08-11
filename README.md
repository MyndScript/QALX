[![SLSA Go releaser](https://github.com/MyndScript/QALX/actions/workflows/go-ossf-slsa3-publish.yml/badge.svg)](https://github.com/MyndScript/QALX/actions/workflows/go-ossf-slsa3-publish.yml)[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)[![Go Report Card](https://goreportcard.com/badge/github.com/MyndScript/QALX)](https://goreportcard.com/report/github.com/MyndScript/QALX)
[![Release](https://img.shields.io/github/v/release/MyndScript/QALX)](https://github.com/MyndScript/QALX/releases)

# QALX

Quantum-secure, emotionally tuned cryptography and mesh logic.

## Vision
Security as Empathy. Quantum-Resilient. Emotionally Tuned.

QALX is a living, ethically aware security organism. It combines quantum-resistant cryptography, emotional modulation (LYRA), and mesh network propagation to create proactive, multidimensional security.

## Features
- Quantum-secure key generation
- LYRA glyph modulation
- Mesh network propagation and validation
- Ethical revocation and stewardship

## How It Works
QALX is modular:
- `core/qalx.go`: Core types and metric initializers
- `core/lyra.go`: LYRA glyph logic and harmonics
- `core/mesh.go`: Mesh node/network logic and validation
- `core/entropy.go`: Entropy pool, key generation, and QRE security

## Quantum Resistance Entropy (QRE)
QRE is the heart of QALX’s quantum security. It combines classical and quantum metrics with emotional glyph modulation:

**QRE Formula:**

	 QRE = log₂(Eₛ × Pₚ × Φₕ × Ψₑ)

- **Eₛ**: Entropy spread across keyspace (Grover resistance)
- **Pₚ**: Period obfuscation factor (Shor resistance)
- **Φₕ**: Hybrid mimicry distortion (wave-emulated attack resistance)
- **Ψₑ**: Emotional glyph modulation (LYRA)

QRE is used in all key generation and mesh validation logic.

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
3. See `core/qalx.go` for main algorithm and usage examples.

## Getting Started
See `core/qalx.go` for the main algorithm. Explore `docs/manifesto.md` for the philosophical and ethical framing.
## Quickstart

1. Clone the repo:
	```sh
	git clone https://github.com/MyndScript/QALX.git
	cd QALX
	```
2. Build for your platform:
	```sh
	go build -o qalx-linux-amd64 ./main.go
	```
3. Run with CLI flags:
	```sh
	./qalx-linux-amd64 --emotion surprise --intensity 0.8 --ethics 0.95
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

## Build and Run

Build for your platform (example for Linux amd64):

```sh
go build -o qalx-linux-amd64 ./main.go
```

Run with CLI flags:

```sh
./qalx-linux-amd64 --emotion surprise --intensity 0.8 --ethics 0.95
```

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
