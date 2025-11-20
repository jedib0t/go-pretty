# Analysis: 256-Color Support for color.go

## Current Implementation

The current `color.go` implementation supports:
- **16 basic colors**: 8 standard + 8 hi-intensity (foreground and background)
- **10 text attributes**: Reset, Bold, Faint, Italic, Underline, Blink, Reverse, Concealed, CrossedOut
- **Color type**: `Color int` - simple integer type
- **Escape sequence format**: `\x1b[<code>m` where `<code>` is the Color value directly

Current color ranges:
- Attributes: 0-9
- Foreground: 30-37 (standard), 90-97 (hi-intensity)
- Background: 40-47 (standard), 100-107 (hi-intensity)

## 256-Color ANSI Specification

256-color mode uses different escape sequences:
- **Foreground 256-color**: `\x1b[38;5;<n>m` where `n` is 0-255
- **Background 256-color**: `\x1b[48;5;<n>m` where `n` is 0-255

The 256-color palette consists of:
- 0-15: Standard 16 colors (matches current 8 basic + 8 hi-intensity)
- 16-231: 216 colors arranged in a 6x6x6 RGB cube
- 232-255: 24 grayscale colors

## Required Changes

### 1. Color Type and Constants

**Option A: Extend existing Color type (Recommended)**
- Keep `Color int` as-is (backward compatible)
- Define new constants for 256-color indices
- Use a reserved range or sentinel values to identify 256-color mode

**Option B: Add new types**
- Create `Color256` type (breaks existing interfaces)
- Not recommended as it would require changes throughout codebase

**Recommended approach:**
```go
// Reserve a high range for 256-color indices
// Use negative values or values > 1000 to distinguish 256-color mode
const (
    // 256-color foreground: use 1000-1255 range
    Fg256Start Color = 1000
    Fg256 Color = 1000  // Base for 256-color foreground
    // Helper functions: Fg256Color(n int) Color { return Fg256Start + Color(n) }
    
    // 256-color background: use 2000-2255 range  
    Bg256Start Color = 2000
    Bg256 Color = 2000  // Base for 256-color background
    // Helper functions: Bg256Color(n int) Color { return Bg256Start + Color(n) }
)
```

### 2. EscapeSeq() Method Changes

**Current implementation:**
```go
func (c Color) EscapeSeq() string {
    return EscapeStart + strconv.Itoa(int(c)) + EscapeStop
}
```

**Required changes:**
- Detect if Color is in 256-color range
- Generate appropriate escape sequence format
- Handle mixing with regular colors

**New implementation:**
```go
func (c Color) EscapeSeq() string {
    // Check if it's a 256-color foreground (1000-1255)
    if c >= Fg256Start && c < Fg256Start+256 {
        colorIndex := int(c - Fg256Start)
        return fmt.Sprintf("%s38;5;%d%s", EscapeStart, colorIndex, EscapeStop)
    }
    // Check if it's a 256-color background (2000-2255)
    if c >= Bg256Start && c < Bg256Start+256 {
        colorIndex := int(c - Bg256Start)
        return fmt.Sprintf("%s48;5;%d%s", EscapeStart, colorIndex, EscapeStop)
    }
    // Regular color (existing behavior)
    return EscapeStart + strconv.Itoa(int(c)) + EscapeStop
}
```

### 3. Colors.EscapeSeq() Method Changes

**Current implementation:**
- Joins all color codes with semicolons: `\x1b[30;40m`

**Required changes:**
- Handle mixed regular and 256-color codes
- 256-color codes need their full sequence (e.g., `38;5;123`) not just the index
- Need to properly combine sequences

**New implementation:**
```go
func (c Colors) EscapeSeq() string {
    if len(c) == 0 {
        return ""
    }
    
    colorsKey := fmt.Sprintf("%#v", c)
    escapeSeq, ok := colorsSeqMap.Load(colorsKey)
    if !ok || escapeSeq == "" {
        var codes []string
        for _, color := range c {
            codes = append(codes, c.colorToCode(color))
        }
        escapeSeq = EscapeStart + strings.Join(codes, ";") + EscapeStop
        colorsSeqMap.Store(colorsKey, escapeSeq)
    }
    return escapeSeq.(string)
}

func (c Colors) colorToCode(color Color) string {
    if color >= Fg256Start && color < Fg256Start+256 {
        colorIndex := int(color - Fg256Start)
        return fmt.Sprintf("38;5;%d", colorIndex)
    }
    if color >= Bg256Start && c < Bg256Start+256 {
        colorIndex := int(color - Bg256Start)
        return fmt.Sprintf("48;5;%d", colorIndex)
    }
    return strconv.Itoa(int(color))
}
```

### 4. CSS Classes Support

**Current implementation:**
- `colorCSSClassMap` maps Color values to CSS class names
- Returns empty string for unmapped colors

**Required changes:**
- Add CSS class mappings for 256 colors OR
- Generate CSS with RGB values for 256-color indices
- Option: Use CSS custom properties or inline styles

**Options:**
1. **Map to CSS classes**: Create 256 entries in `colorCSSClassMap`
   - Pros: Consistent with current approach
   - Cons: Large map, need to define 256 class names

2. **Generate RGB-based CSS**: Convert 256-color index to RGB
   - Pros: More flexible, no large map needed
   - Cons: Different approach from current implementation

3. **Hybrid**: Use CSS classes for standard colors, RGB for 256 colors
   - Pros: Best of both worlds
   - Cons: Inconsistent return format

**Recommended:**
```go
func (c Color) CSSClasses() string {
    // Check for 256-color and convert to RGB-based class
    if c >= Fg256Start && c < Fg256Start+256 {
        r, g, b := color256ToRGB(int(c - Fg256Start))
        return fmt.Sprintf("fg-256-%d-%d-%d", r, g, b)
    }
    if c >= Bg256Start && c < Bg256Start+256 {
        r, g, b := color256ToRGB(int(c - Bg256Start))
        return fmt.Sprintf("bg-256-%d-%d-%d", r, g, b)
    }
    // Existing behavior
    if class, ok := colorCSSClassMap[c]; ok {
        return class
    }
    return ""
}

func color256ToRGB(index int) (r, g, b int) {
    if index < 16 {
        // Standard 16 colors - map to predefined RGB
        return standardColorRGB[index]
    } else if index < 232 {
        // 216-color RGB cube
        index -= 16
        r = (index / 36) * 51
        g = ((index / 6) % 6) * 51
        b = (index % 6) * 51
    } else {
        // 24 grayscale colors
        gray := 8 + (index-232)*10
        r, g, b = gray, gray, gray
    }
    return
}
```

### 5. Helper Functions

Add convenience functions for creating 256-color values:

```go
// Fg256Color returns a foreground 256-color Color value
func Fg256Color(index int) Color {
    if index < 0 || index > 255 {
        return Reset // or panic
    }
    return Fg256Start + Color(index)
}

// Bg256Color returns a background 256-color Color value
func Bg256Color(index int) Color {
    if index < 0 || index > 255 {
        return Reset // or panic
    }
    return Bg256Start + Color(index)
}

// Fg256RGB returns a foreground 256-color from RGB values (0-5 each)
func Fg256RGB(r, g, b int) Color {
    if r < 0 || r > 5 || g < 0 || g > 5 || b < 0 || b > 5 {
        return Reset
    }
    index := 16 + (r*36 + g*6 + b)
    return Fg256Color(index)
}

// Bg256RGB returns a background 256-color from RGB values (0-5 each)
func Bg256RGB(r, g, b int) Color {
    if r < 0 || r > 5 || g < 0 || g > 5 || b < 0 || b > 5 {
        return Reset
    }
    index := 16 + (r*36 + g*6 + b)
    return Bg256Color(index)
}
```

### 6. Testing Requirements

- Test 256-color escape sequence generation
- Test mixing 256-colors with regular colors
- Test CSS class generation for 256 colors
- Test edge cases (invalid indices, boundary values)
- Test backward compatibility (existing colors still work)

### 7. Documentation Updates

- Update README.md to document 256-color support
- Add examples showing 256-color usage
- Document the color index ranges and RGB cube structure

## Backward Compatibility

✅ **Maintained**: All existing code will continue to work unchanged
- Existing Color constants remain valid
- Existing EscapeSeq() behavior preserved for standard colors
- No breaking changes to public interfaces

## Implementation Complexity

**Low to Medium**:
- Core logic is straightforward
- Main complexity is in CSS mapping strategy
- Testing required for edge cases
- Need to ensure escape sequence parser handles 256-color codes correctly

## Potential Issues

1. **Escape Sequence Parser**: The `EscSeqParser` in `escape_seq_parser.go` currently parses sequences by splitting on semicolons and storing individual codes. For 256-color sequences like `38;5;123`, it would store codes `38`, `5`, and `123` separately. This is **not a blocker** for basic 256-color support, but if you want the parser to recognize 256-color sequences as semantic units (e.g., for proper reset handling), it would need enhancement. For now, the parser will work but won't treat `38;5;123` as a single color code.

2. **Color Validation**: Need to validate 256-color indices are in valid range (0-255) in helper functions

3. **CSS Support**: HTML rendering may need additional CSS definitions for 256-color classes. The CSS class generation strategy needs to be decided (see section 4 above).

4. **Terminal Support**: Not all terminals support 256 colors - may need detection logic (though current `areANSICodesSupported()` might handle this). The escape sequences will be generated correctly regardless, and unsupported terminals will simply ignore them.

## Summary

The changes required are **minimal and backward-compatible**:
1. Extend `EscapeSeq()` to detect and format 256-color codes
2. Update `Colors.EscapeSeq()` to handle mixed color types
3. Add helper functions for creating 256-color values
4. Extend CSS class support (strategy TBD)
5. Add tests and documentation

The existing `Color int` type can be reused without breaking changes by using reserved value ranges to distinguish 256-color mode from standard colors.

## Quick Reference: Implementation Checklist

- [ ] Define reserved value ranges for 256-color indices (e.g., 1000-1255 for foreground, 2000-2255 for background)
- [ ] Update `Color.EscapeSeq()` to detect 256-color range and generate `38;5;n` or `48;5;n` format
- [ ] Update `Colors.EscapeSeq()` to handle mixed regular and 256-color codes
- [ ] Add helper functions: `Fg256Color()`, `Bg256Color()`, `Fg256RGB()`, `Bg256RGB()`
- [ ] Implement `color256ToRGB()` conversion function
- [ ] Update `Color.CSSClasses()` to handle 256-color indices (choose strategy: map, RGB-based, or hybrid)
- [ ] Update `Colors.CSSClasses()` to handle 256-color indices
- [ ] Add comprehensive tests for 256-color functionality
- [ ] Update documentation (README.md, examples)
- [ ] Verify backward compatibility with existing code

