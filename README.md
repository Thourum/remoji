# Remoji - Remove Emoji from text files

AI have a tendence to overuse emoji this simple go cli tool removes them.

## Install

```sh
go build -o remoji .
```

## Usage

```sh
remoji input.txt          # cleaned text to stdout
remoji -i a.md b.tex      # rewrite the files in place
remoji < input.txt        # read stdin
```

Works with any UTF-8 text file (`.txt`, `.md`, `.tex`, `.org`, source code, ...).
Files containing NUL bytes or invalid UTF-8 are rejected rather than mangled;
the extension is not used to decide.

## What it removes

Emoji and the invisible glue that binds them: ZWJ sequences (`👨‍👩‍👧‍👦`), regional
indicator flags (`🇬🇧`), skin-tone modifiers (`👍🏽`), keycaps (`1️⃣` becomes `1`),
and variation selectors.

Text-presentation symbols are kept, so prose and math stay intact:
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

Whitespace is left untouched — only the emoji code points are dropped.
