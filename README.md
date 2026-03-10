# FixTwitter

A macOS clipboard monitoring tool that automatically normalizes X.com (Twitter) and Instagram links for better previews and cleaner sharing.

## What it does

FixTwitter runs as a background service that monitors your clipboard for supported URLs and normalizes them automatically.

### Supported URL transformations
- `https://x.com/username/status/123456789?ref=share` → `https://fxtwitter.com/username/status/123456789`
- `https://x.com/username/status/123456789?ref=share` → `https://no.sb/username/status/123456789`
- `https://www.instagram.com/yukfanha/?g=5` → `https://www.instagram.com/yukfanha`
- `https://m.instagram.com/p/abc123/?utm_source=ig_web_copy_link#section` → `https://m.instagram.com/p/abc123#section`

### Instagram support

When an Instagram URL is copied, FixTwitter removes the entire suffix from `?` onward
for `instagram.com` links, including subdomains such as `www.instagram.com` and
`m.instagram.com`. This also trims the trailing `/` before the query so
`/path/?q=...` becomes `/path`, and `#fragment` parts are removed in the same
operation.

## Installation

Install via Homebrew:

```bash
brew tap owo-network/brew
brew install fixtwitter
```

The `fixtwitter` package uses the fxtwitter.com service.

Or install the no.sb version:

```bash
brew install fixtwitter-nosb
```

The `fixtwitter-nosb` package uses the no.sb FxEmbed service.

## Usage

### Basic usage
Simply run the command to start monitoring your clipboard:

```bash
fixtwitter
```

The service will start monitoring your clipboard and automatically normalize supported links you copy.

### Custom service
You can specify a custom replacement service:

```bash
fixtwitter -service your-custom-domain.com
```

### Stopping the service
Press `Ctrl+C` to stop the monitoring service.

## How it works

1. The application monitors your macOS clipboard for changes every 500ms
2. When clipboard content changes, it checks the text for supported URLs
3. X.com status links are rewritten to the configured replacement service
4. Instagram links have their query parameters removed while keeping the rest of the URL intact
5. The modified text is automatically placed back into your clipboard

## Build from source

Requirements:
- Go 1.21 or later
- macOS (uses Cocoa framework)

```bash
git clone https://github.com/missuo/FixTwitter
cd FixTwitter
go build -o fixtwitter .
```

## License

Copyright © 2025 by Vincent Yang, All Rights Reserved.
