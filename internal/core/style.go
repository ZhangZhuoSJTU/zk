package core

import "fmt"

// Style is a key representing a single styling rule.
type Style string

// Semantic styling rules.
var (
	StyleTitle      = Style("title")     // Title of a note.
	StylePath       = Style("path")      // Path to notebook file.
	StyleTerm       = Style("term")      // Searched for term in a note.
	StyleEmphasis   = Style("emphasis")  // Element to emphasize, e.g. [y]es.
	StyleUnderstate = Style("understate") // Element to understate, e.g. note content in fzf.
)

// Text decoration styles.
var (
	StyleBold          = Style("bold")
	StyleItalic        = Style("italic")
	StyleFaint         = Style("faint")
	StyleUnderline     = Style("underline")
	StyleStrikethrough = Style("strikethrough")
	StyleBlink         = Style("blink")
	StyleReverse       = Style("reverse")
	StyleHidden        = Style("hidden")
)

// baseColors lists the color names used to generate foreground, background,
// and bright variant styles.
var baseColors = []string{
	"black", "red", "green", "yellow",
	"blue", "magenta", "cyan", "white",
}

// Color styles — generated from baseColors at init time.
var (
	StyleBlack, StyleRed, StyleGreen, StyleYellow     Style
	StyleBlue, StyleMagenta, StyleCyan, StyleWhite    Style

	StyleBlackBg, StyleRedBg, StyleGreenBg, StyleYellowBg     Style
	StyleBlueBg, StyleMagentaBg, StyleCyanBg, StyleWhiteBg    Style

	StyleBrightBlack, StyleBrightRed, StyleBrightGreen, StyleBrightYellow     Style
	StyleBrightBlue, StyleBrightMagenta, StyleBrightCyan, StyleBrightWhite    Style

	StyleBrightBlackBg, StyleBrightRedBg, StyleBrightGreenBg, StyleBrightYellowBg     Style
	StyleBrightBlueBg, StyleBrightMagentaBg, StyleBrightCyanBg, StyleBrightWhiteBg    Style
)

func init() {
	colorVars := []*Style{
		&StyleBlack, &StyleRed, &StyleGreen, &StyleYellow,
		&StyleBlue, &StyleMagenta, &StyleCyan, &StyleWhite,
	}
	bgVars := []*Style{
		&StyleBlackBg, &StyleRedBg, &StyleGreenBg, &StyleYellowBg,
		&StyleBlueBg, &StyleMagentaBg, &StyleCyanBg, &StyleWhiteBg,
	}
	brightVars := []*Style{
		&StyleBrightBlack, &StyleBrightRed, &StyleBrightGreen, &StyleBrightYellow,
		&StyleBrightBlue, &StyleBrightMagenta, &StyleBrightCyan, &StyleBrightWhite,
	}
	brightBgVars := []*Style{
		&StyleBrightBlackBg, &StyleBrightRedBg, &StyleBrightGreenBg, &StyleBrightYellowBg,
		&StyleBrightBlueBg, &StyleBrightMagentaBg, &StyleBrightCyanBg, &StyleBrightWhiteBg,
	}

	for i, name := range baseColors {
		*colorVars[i] = Style(name)
		*bgVars[i] = Style(name + "-bg")
		*brightVars[i] = Style("bright-" + name)
		*brightBgVars[i] = Style("bright-" + name + "-bg")
	}
}

// Styler stylizes text according to predefined styling rules.
//
// A rule key can be either semantic, e.g. "title" or explicit, e.g. "red".
type Styler interface {
	// Style formats the given text according to the provided styling rules.
	Style(text string, rules ...Style) (string, error)
	// Style formats the given text according to the provided styling rules,
	// panicking if the rules are unknown.
	MustStyle(text string, rules ...Style) string
}

// ProxyStyler is a styler delegating to an underlying styler.
// Can be used to change the active styler during runtime.
type ProxyStyler struct {
	Styler Styler
}

func NewProxyStyler(styler Styler) *ProxyStyler {
	return &ProxyStyler{styler}
}

// Style implements Styler.
func (s ProxyStyler) Style(text string, rule ...Style) (string, error) {
	return s.Styler.Style(text, rule...)
}

// MustStyle implements Styler.
func (s ProxyStyler) MustStyle(text string, rule ...Style) string {
	return s.Styler.MustStyle(text, rule...)
}

// NullStyler is a Styler with no styling rules.
var NullStyler = nullStyler{}

type nullStyler struct{}

// Style implements Styler.
func (s nullStyler) Style(text string, rule ...Style) (string, error) {
	return text, nil
}

// MustStyle implements Styler.
func (s nullStyler) MustStyle(text string, rule ...Style) string {
	return text
}

// TagStyler is a Styler which outputs XML tags.
var TagStyler = tagStyler{}

type tagStyler struct{}

// Style implements Styler.
func (s tagStyler) Style(text string, rules ...Style) (string, error) {
	return s.MustStyle(text, rules...), nil
}

// MustStyle implements Styler.
func (s tagStyler) MustStyle(text string, rules ...Style) string {
	for _, rule := range rules {
		text = fmt.Sprintf("<%s>%s</%s>", rule, text, rule)
	}
	return text
}
