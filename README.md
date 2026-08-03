<p align="center">
  <picture>
    <source media="(prefers-color-scheme: dark)" srcset="https://shieldcn.dev/header/graph.svg?title=remoji&subtitle=Remove+emoji+from+text+files&logo=go&align=center&mode=dark" />
    <img alt="remoji" src="https://shieldcn.dev/header/graph.svg?title=remoji&subtitle=Remove+emoji+from+text+files&logo=go&align=center&mode=light" />
  </picture>
</p>

<p align="center">
  <a href="https://go.dev"><img alt="Go 1.26" src="https://shieldcn.dev/badge/go-1.26-00ADD8?variant=secondary&logo=go" /></a>
  <a href="https://github.com/Thourum/remoji/stargazers"><img alt="GitHub stars" src="https://shieldcn.dev/github/stars/Thourum/remoji?variant=secondary" /></a>
  <a href="https://github.com/Thourum/remoji/commits/main"><img alt="last commit" src="https://shieldcn.dev/github/last-commit/Thourum/remoji?variant=secondary" /></a>
</p>

AI have a tendence to overuse emoji — this simple Go CLI removes them.

## Install

```sh
git clone https://github.com/Thourum/remoji && cd remoji
go build -o remoji .
sudo mv remoji /usr/local/bin/     # or: mv remoji ~/.local/bin/
```

Now `remoji` works from any directory. If `~/.local/bin` isn't on your `PATH`,
add `export PATH="$HOME/.local/bin:$PATH"` to your shell profile.

Or in one step, straight into `$(go env GOPATH)/bin`:

```sh
go install github.com/Thourum/remoji@latest
```

## Usage

```sh
remoji input.txt          # cleaned text to stdout
remoji -i a.md b.tex      # rewrite the files in place
remoji < input.txt        # read stdin
```

Works with any UTF-8 text file (`.txt`, `.md`, `.tex`, `.org`, source code, ...).
Files with NUL bytes or invalid UTF-8 are rejected rather than mangled — the
extension is never used to decide.

## What it removes

Emoji plus the invisible glue that binds them: ZWJ sequences (`👨‍👩‍👧‍👦`), regional
indicator flags (`🇬🇧`), skin-tone modifiers (`👍🏽`), keycaps (`1️⃣` → `1`), and
variation selectors.

Text-presentation symbols survive, so prose and math stay intact:
`✓ ✗ ★ ♩ ← → ™ © × —`, accents, CJK.

## Example

```
❌ BEFORE: "Managed social media accounts"
✅ AFTER: "Grew Instagram following by 250% (5K to 17.5K)"
```

becomes

```
 BEFORE: "Managed social media accounts"
 AFTER: "Grew Instagram following by 250% (5K to 17.5K)"
```

Whitespace is untouched — only the emoji code points are dropped.
