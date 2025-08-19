# FixTwitter

A macOS clipboard monitoring tool that automatically converts X.com (Twitter) links to alternative embedding services for better previews and functionality.

## What it does

FixTwitter runs as a background service that monitors your clipboard for X.com links. When it detects a Twitter/X.com status URL, it automatically replaces the domain with an alternative service that provides better embeds, previews, and functionality.

### Supported URL transformation:
- `https://x.com/username/status/123456789` → `https://fxtwitter.com/username/status/123456789`
- `https://x.com/username/status/123456789` → `https://no.sb/username/status/123456789`

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

The service will start monitoring your clipboard and automatically convert any X.com links you copy.

### Custom service
You can specify a custom replacement service:

```bash
fixtwitter -service your-custom-domain.com
```

### Stopping the service
Press `Ctrl+C` to stop the monitoring service.

## How it works

1. The application monitors your macOS clipboard for changes every 500ms
2. When clipboard content changes, it checks for X.com status URLs using regex pattern matching
3. If found, it replaces `x.com` with the configured service domain
4. The modified URL is automatically placed back into your clipboard
5. You can then paste the converted link with better embed support

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
