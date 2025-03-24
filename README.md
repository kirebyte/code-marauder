# Code-Marauder 🏴‍☠️

---

**Code-Marauder** is a tool written in Go that crawls through a source code repository and generates a single, flattened text file containing all the code files in a clean, structured listing.

It's designed to help **AI systems** (especially large language models) read and understand entire projects more efficiently, without having to navigate dozens of files or complex directory structures.

---

## ✨ Features

- Converts full source code repos into a single `.txt` file.
- Designed for AI input, summarization, and programmatic inspection.
- Lightweight and fast.
- Highly configurable via environment variables.

---

## 🧰 Installation

Make sure you have Go 1.21+ installed, then run:

```bash
go install github.com/kirebyte/code-marauder@latest
```

## ⚙️ Usage

1. Set the required environment variables:

```bash
export CODE_LISTING_DIRS="./myproject"
export CODE_LISTING_FILE="./output/listing.txt"
export CODE_LISTING_ALLOWED_FILES=".go,.ts,.js,.py,.java"
```

CODE_LISTING_DIRS can be a comma-separated list of folders.

2. Run the tool:

```bash
code-marauder
```

You will get a full code listing at the path defined in CODE_LISTING_FILE.
Each file will be wrapped like this:

```text
==== File: ./path/to/file.go ====
<file content>
==== End of File ====
```

---

# License 🪪
This project is licensed under the MIT License.
See the LICENSE file for details.

---

# Contributing 🧪
There is currently no contribution guide yet. PRs and ideas are welcome!

---

# To-Do 💡

- [ ] Add CLI flags:
  - `--output` for custom output path
  - `--ext` to override allowed file extensions
  - `--ignore` to skip specific folders or files
- [ ] Support `.gitignore`-style exclusion
- [ ] Use `cobra` for structured CLI commands
- [ ] Add dry-run mode (`--dry-run`) to preview files being listed
- [ ] Support output in different formats:
  - JSON
  - Markdown
  - Embeddable HTML blocks
- [ ] Add optional config file (`.marauderrc.yaml`)
- [ ] Include a `--verbose` mode for logging each processed file
- [ ] Generate file index or table of contents
- [ ] Create `.env.example` template
- [ ] Write `Makefile` for build/install automation
- [ ] Create `Dockerfile` for containerized usage
- [ ] Add simple web viewer for the output file (static HTML)