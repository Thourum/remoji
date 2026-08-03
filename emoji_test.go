package main

import "testing"

func TestStrip(t *testing.T) {
	cases := []struct{ in, want string }{
		{"❌ BEFORE: \"Managed social media\"", " BEFORE: \"Managed social media\""},
		{"✅ AFTER: grew 250%", " AFTER: grew 250%"},
		{"plain ascii stays", "plain ascii stays"},
		{"family 👨‍👩‍👧‍👦 gone", "family  gone"},
		{"flag 🇬🇧 gone", "flag  gone"},
		{"thumbs 👍🏽 gone", "thumbs  gone"},
		{"keycap 1️⃣ gone", "keycap 1 gone"},
		{"\\alpha × ✓ ™ é 日本 —", "\\alpha × ✓ ™ é 日本 —"},
		{"🚀", ""},
	}
	for _, c := range cases {
		if got := Strip(c.in); got != c.want {
			t.Errorf("Strip(%q) = %q, want %q", c.in, got, c.want)
		}
	}
}
