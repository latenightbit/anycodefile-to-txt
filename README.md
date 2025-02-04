# Code File to Text Converter

A robust Go utility that converts code files to text format while preserving original structure and creating comprehensive combined output.

## Motivation

I created this program out of personal struggle. I was tired of manually copying and pasting my code files back and forth into AI assistants. The tedious process of copying each file individually and combining them into one big file was time-consuming and frustrating. Now, with this tool, those sufferings end! 😌

## Features

- 🚀 Converts **any file type** to `.txt` format
- 📂 Maintains original filenames with `.txt` extension
- 🧩 Creates combined output with smart separators
- 🛡️ Auto-creates required directories
- 📊 Provides detailed processing statistics
- ❌ Clear error handling and reporting

## Installation

1. **Install Go** (1.20+ required):  
   [golang.org/dl](https://golang.org/dl/)

2. **Clone repository**:
```
git clone https://github.com/latenightbit/anycodefile-to-txt.git
cd anycodefile-to-txt
```

## Usage

### Basic Setup
```
# Create project structure (auto-created on first run)
mkdir -p input

# Add your files to convert
cp ~/your-code/* input/
```

### Run Conversion
```
go run main.go
```

### Check Output
```
ls -l text_output/
```

## Example Output Structure
```
text_output/
├── main.py.txt
├── utils.js.txt
├── notes.md.txt
└── everything_combined.txt
```

## Combined File Format
```
========================================
FILE: main.py
========================================
def hello():
    print("Hello World!")

========================================
FILE: utils.js
========================================
function log(msg) {
    console.log(msg);
}
```

## Advanced Usage

### Build Executable
```
go build -o code2text
./code2text
```

### Custom Directories (Modify main.go)
```
const (
    inputDir  = "./custom_input"  // Change input directory
    outputDir = "./text_results"  // Change output directory
)
```

## Requirements

- Go 1.20+
- 500KB disk space
- Read/Write permissions in project directory

## License

MIT License  
Copyright (c) 2024 [Your Name]

---

**Pro Tip**: Add `alias code2text="go run /path/to/main.go"` to your `.bashrc`/`.zshrc` for quick access!
```

This README includes:
1. Clear installation instructions
2. Visual examples of file structures
3. Ready-to-use code blocks
4. Both basic and advanced usage
5. Licensing information
6. Professional formatting with emoji visual cues

Save this as `README.md` in your project root directory. Let me know if you need any adjustments!

---
Answer from Perplexity: pplx.ai/share