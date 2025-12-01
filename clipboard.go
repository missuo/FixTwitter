package main

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework Cocoa
#import <Cocoa/Cocoa.h>

// Get clipboard text content as C string
const char* getClipboardText() {
    NSPasteboard *pasteboard = [NSPasteboard generalPasteboard];
    NSString *string = [pasteboard stringForType:NSPasteboardTypeString];
    if (string == nil) {
        return NULL;
    }
    return [string UTF8String];
}

// Set clipboard text content from C string
void setClipboardText(const char* text) {
    NSPasteboard *pasteboard = [NSPasteboard generalPasteboard];
    [pasteboard clearContents];
    [pasteboard setString:[NSString stringWithUTF8String:text] forType:NSPasteboardTypeString];
}

// Get clipboard change counter to detect changes
long getChangeCount() {
    NSPasteboard *pasteboard = [NSPasteboard generalPasteboard];
    return [pasteboard changeCount];
}
*/
import "C"
import (
	"regexp"
	"time"
	"unsafe"
)

// ClipboardMonitor monitors clipboard changes and processes Twitter/X.com URLs
type ClipboardMonitor struct {
	lastChangeCount int64
	lastContent     string
	replaceService  string // The service to replace x.com with (e.g., "fxtwitter.com")
}

// NewClipboardMonitor creates a new clipboard monitor with the specified replacement service
func NewClipboardMonitor(replaceService string) *ClipboardMonitor {
	return &ClipboardMonitor{
		lastChangeCount: int64(C.getChangeCount()),
		replaceService:  replaceService,
	}
}

// GetClipboardText retrieves current clipboard text content
func (cm *ClipboardMonitor) GetClipboardText() string {
	cStr := C.getClipboardText()
	if cStr == nil {
		return ""
	}
	return C.GoString(cStr)
}

// SetClipboardText sets clipboard content to the specified text
func (cm *ClipboardMonitor) SetClipboardText(text string) {
	cStr := C.CString(text)
	defer C.free(unsafe.Pointer(cStr))
	C.setClipboardText(cStr)
}

// HasChanged checks if clipboard content has changed since last check
func (cm *ClipboardMonitor) HasChanged() bool {
	currentChangeCount := int64(C.getChangeCount())
	if currentChangeCount != cm.lastChangeCount {
		cm.lastChangeCount = currentChangeCount
		return true
	}
	return false
}

// ProcessClipboard checks clipboard changes and processes Twitter/X.com URLs
func (cm *ClipboardMonitor) ProcessClipboard() {
	if !cm.HasChanged() {
		return
	}

	content := cm.GetClipboardText()
	if content == "" || content == cm.lastContent {
		return
	}

	cm.lastContent = content
	newContent := cm.replaceXComUrls(content)
	
	if newContent != content {
		cm.SetClipboardText(newContent)
		cm.lastContent = newContent
	}
}

// replaceXComUrls replaces x.com status URLs with the configured replacement service
func (cm *ClipboardMonitor) replaceXComUrls(text string) string {
	// Match x.com status links with optional query parameters
	re := regexp.MustCompile(`https://x\.com/([^/]+/status/\d+)(?:\?[^\s]*)?`)
	return re.ReplaceAllString(text, "https://"+cm.replaceService+"/$1")
}

// Start begins monitoring clipboard changes in an infinite loop
func (cm *ClipboardMonitor) Start() {
	for {
		cm.ProcessClipboard()
		time.Sleep(500 * time.Millisecond) // Check every 500ms
	}
}