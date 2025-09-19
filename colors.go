package projetred

import (
	"fmt"
	"strings"
	"time"
)

const (
	Reset = "\033[0m"
	Bold  = "\033[1m"
	Dim   = "\033[2m"

	Black   = "\033[30m"
	Red     = "\033[31m"
	Green   = "\033[32m"
	Yellow  = "\033[33m"
	Blue    = "\033[34m"
	Magenta = "\033[35m"
	Cyan    = "\033[36m"
	White   = "\033[37m"

	BrightBlack   = "\033[90m"
	BrightRed     = "\033[91m"
	BrightGreen   = "\033[92m"
	BrightYellow  = "\033[93m"
	BrightBlue    = "\033[94m"
	BrightMagenta = "\033[95m"
	BrightCyan    = "\033[96m"
	BrightWhite   = "\033[97m"

	BgBlack   = "\033[40m"
	BgRed     = "\033[41m"
	BgGreen   = "\033[42m"
	BgYellow  = "\033[43m"
	BgBlue    = "\033[44m"
	BgMagenta = "\033[45m"
	BgCyan    = "\033[46m"
	BgWhite   = "\033[47m"
)

type ColorTheme struct {
	Primary    string
	Secondary  string
	Accent     string
	Text       string
	Background string
}

var (
	BartTheme = ColorTheme{
		Primary:    BrightYellow + Bold,
		Secondary:  Yellow,
		Accent:     BrightRed,
		Text:       BrightWhite,
		Background: BgBlue,
	}

	LisaTheme = ColorTheme{
		Primary:    BrightRed + Bold,
		Secondary:  Red,
		Accent:     BrightYellow,
		Text:       BrightWhite,
		Background: BgMagenta,
	}

	MaggieTheme = ColorTheme{
		Primary:    BrightBlue + Bold,
		Secondary:  Blue,
		Accent:     BrightCyan,
		Text:       BrightWhite,
		Background: BgBlue,
	}

	HomerTheme = ColorTheme{
		Primary:    BrightYellow + Bold,
		Secondary:  Yellow,
		Accent:     BrightGreen,
		Text:       BrightWhite,
		Background: BgYellow,
	}

	MargeTheme = ColorTheme{
		Primary:    BrightBlue + Bold,
		Secondary:  Blue,
		Accent:     BrightGreen,
		Text:       BrightWhite,
		Background: BgBlue,
	}

	NedTheme = ColorTheme{
		Primary:    BrightGreen + Bold,
		Secondary:  Green,
		Accent:     BrightYellow,
		Text:       BrightWhite,
		Background: BgGreen,
	}

	MoeTheme = ColorTheme{
		Primary:    BrightBlack + Bold,
		Secondary:  Black,
		Accent:     BrightRed,
		Text:       BrightWhite,
		Background: BgBlack,
	}

	BossTheme = ColorTheme{
		Primary:    BrightRed + Bold,
		Secondary:  Red,
		Accent:     BrightYellow,
		Text:       BrightWhite,
		Background: BgRed,
	}

	ShopTheme = ColorTheme{
		Primary:    BrightGreen + Bold,
		Secondary:  Green,
		Accent:     BrightYellow,
		Text:       BrightWhite,
		Background: BgGreen,
	}

	CombatTheme = ColorTheme{
		Primary:    BrightRed + Bold,
		Secondary:  Red,
		Accent:     BrightYellow,
		Text:       BrightWhite,
		Background: BgRed,
	}

	SystemTheme = ColorTheme{
		Primary:    BrightCyan + Bold,
		Secondary:  Cyan,
		Accent:     BrightWhite,
		Text:       White,
		Background: BgCyan,
	}
)

func Colorize(text string, color string) string {
	return color + text + Reset
}

func ColorizeThemed(text string, theme ColorTheme, style string) string {
	switch style {
	case "primary":
		return theme.Primary + text + Reset
	case "secondary":
		return theme.Secondary + text + Reset
	case "accent":
		return theme.Accent + text + Reset
	case "text":
		return theme.Text + text + Reset
	case "background":
		return theme.Background + " " + theme.Text + text + " " + Reset
	default:
		return theme.Primary + text + Reset
	}
}

func GetCharacterTheme(characterClass string) ColorTheme {
	switch strings.ToLower(characterClass) {
	case "bart":
		return BartTheme
	case "lisa":
		return LisaTheme
	case "maggie":
		return MaggieTheme
	default:
		return BartTheme
	}
}

func ColoredTypeWriter(text string, delay time.Duration, color string) {
	coloredText := Colorize(text, color)
	fmt.Print(coloredText)
	fmt.Println()
	time.Sleep(delay)
}

func ThemedTypeWriter(text string, delay time.Duration, theme ColorTheme, style string) {
	coloredText := ColorizeThemed(text, theme, style)
	fmt.Print(coloredText)
	fmt.Println()
	time.Sleep(delay)
}

func AnimatedText(text string, delay time.Duration, color string) {
	for _, char := range text {
		fmt.Print(Colorize(string(char), color))
		time.Sleep(delay / time.Duration(len(text)))
	}
	fmt.Println()
}

func FlashingText(text string, color1 string, color2 string, flashes int, delay time.Duration) {
	for i := 0; i < flashes; i++ {
		if i%2 == 0 {
			fmt.Print("\r" + Colorize(text, color1))
		} else {
			fmt.Print("\r" + Colorize(text, color2))
		}
		time.Sleep(delay)
	}
	fmt.Print("\r" + Colorize(text, color1))
	fmt.Println()
}

func GradientText(text string, startColor string, endColor string) {
	fmt.Print(Colorize(text, startColor))
	fmt.Println()
}

func BoxedText(text string, theme ColorTheme) {
	width := len(text) + 4
	topBorder := strings.Repeat("═", width-2)

	fmt.Println(ColorizeThemed("╔"+topBorder+"╗", theme, "accent"))
	fmt.Println(ColorizeThemed("║ "+text+" ║", theme, "primary"))
	fmt.Println(ColorizeThemed("╚"+topBorder+"╝", theme, "accent"))
}

func ProgressBar(current int, max int, width int, theme ColorTheme) string {
	if max <= 0 {
		return ColorizeThemed("[ERROR]", theme, "accent")
	}

	percentage := float64(current) / float64(max)
	filled := int(percentage * float64(width))

	bar := "["
	for i := 0; i < width; i++ {
		if i < filled {
			bar += ColorizeThemed("█", theme, "primary")
		} else {
			bar += ColorizeThemed("░", theme, "secondary")
		}
	}
	bar += "]"

	return bar + fmt.Sprintf(" %d/%d", current, max)
}

func HealthBar(current int, max int, width int) string {
	if max <= 0 {
		return Colorize("[ERROR]", BrightRed)
	}

	percentage := float64(current) / float64(max)
	filled := int(percentage * float64(width))

	var color string
	if percentage > 0.7 {
		color = BrightGreen
	} else if percentage > 0.3 {
		color = BrightYellow
	} else {
		color = BrightRed
	}

	bar := "["
	for i := 0; i < width; i++ {
		if i < filled {
			bar += Colorize("█", color)
		} else {
			bar += Colorize("░", BrightBlack)
		}
	}
	bar += "]"

	return bar + fmt.Sprintf(" %s%d%s/%d", color, current, Reset, max)
}

func DialogueBox(speaker string, text string, theme ColorTheme) {
	speakerColored := ColorizeThemed(speaker, theme, "primary")
	textColored := ColorizeThemed(text, theme, "text")

	fmt.Printf("%s : %s\n", speakerColored, textColored)
}

func MenuHeader(title string, theme ColorTheme) {
	width := len(title) + 4
	topBorder := strings.Repeat("═", width-2)

	fmt.Println()
	fmt.Println(ColorizeThemed("╔"+topBorder+"╗", theme, "accent"))
	fmt.Printf("%s║ %s%s%s ║%s\n",
		ColorizeThemed("", theme, "accent"),
		ColorizeThemed("", theme, "primary"),
		title,
		ColorizeThemed("", theme, "accent"),
		Reset)
	fmt.Println(ColorizeThemed("╚"+topBorder+"╝", theme, "accent"))
	fmt.Println()
}

func StatusEffect(effect string, isPositive bool) string {
	if isPositive {
		return Colorize("✨ "+effect, BrightGreen+Bold)
	} else {
		return Colorize("💀 "+effect, BrightRed+Bold)
	}
}

func DamageNumber(damage int, isCritical bool) string {
	if isCritical {
		return FlashingTextReturn(fmt.Sprintf("💥 %d CRITIQUE!", damage), BrightRed+Bold, BrightYellow+Bold)
	} else {
		return Colorize(fmt.Sprintf("💢 %d", damage), BrightRed)
	}
}

func FlashingTextReturn(text string, color1 string, color2 string) string {
	return Colorize(text, color1)
}

func HealingNumber(healing int) string {
	return Colorize(fmt.Sprintf("💚 +%d PV", healing), BrightGreen+Bold)
}

func CurrencyDisplay(amount int) string {
	return Colorize(fmt.Sprintf("💰 %d dollars", amount), BrightYellow+Bold)
}

func ItemDisplay(item string, rarity string) string {
	switch rarity {
	case "legendary":
		return Colorize("⭐ "+item, BrightYellow+Bold)
	case "rare":
		return Colorize("🔷 "+item, BrightBlue+Bold)
	case "common":
		return Colorize("⚪ "+item, BrightWhite)
	default:
		return Colorize("📦 "+item, White)
	}
}

func LevelUpEffect(newLevel int) {
	fmt.Println()
	FlashingText("🎉 NIVEAU SUPÉRIEUR! 🎉", BrightYellow+Bold, BrightGreen+Bold, 6, 200*time.Millisecond)
	fmt.Printf("%s\n", Colorize(fmt.Sprintf("✨ Nouveau niveau : %d ✨", newLevel), BrightCyan+Bold))
	fmt.Println()
}

func VictoryEffect() {
	fmt.Println()
	FlashingText("🏆 VICTOIRE! 🏆", BrightYellow+Bold, BrightGreen+Bold, 8, 150*time.Millisecond)
	fmt.Println()
}

func DefeatEffect() {
	fmt.Println()
	FlashingText("💀 DÉFAITE... 💀", BrightRed+Bold, Red, 5, 300*time.Millisecond)
	fmt.Println()
}
