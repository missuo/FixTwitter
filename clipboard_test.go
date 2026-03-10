package main

import "testing"

func TestNormalizeUrls(t *testing.T) {
	monitor := &ClipboardMonitor{replaceService: "fxtwitter.com"}

	tests := []struct {
		name string
		in   string
		want string
	}{
		{
			name: "replaces x status url and drops query",
			in:   "https://x.com/example/status/12345?ref=share",
			want: "https://fxtwitter.com/example/status/12345",
		},
		{
			name: "strips instagram query on www host",
			in:   "https://www.instagram.com/yukfanha/?g=5",
			want: "https://www.instagram.com/yukfanha",
		},
		{
			name: "strips instagram query on subdomain",
			in:   "https://m.instagram.com/p/abc123/?utm_source=ig_web_copy_link",
			want: "https://m.instagram.com/p/abc123",
		},
		{
			name: "removes fragment when query exists",
			in:   "https://www.instagram.com/p/abc123?foo=bar#section",
			want: "https://www.instagram.com/p/abc123",
		},
		{
			name: "preserves trailing punctuation",
			in:   "(https://www.instagram.com/yukfanha/?g=5).",
			want: "(https://www.instagram.com/yukfanha).",
		},
		{
			name: "leaves non instagram urls alone",
			in:   "https://example.com/path?a=1",
			want: "https://example.com/path?a=1",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := monitor.normalizeUrls(tt.in); got != tt.want {
				t.Fatalf("normalizeUrls(%q) = %q, want %q", tt.in, got, tt.want)
			}
		})
	}
}
