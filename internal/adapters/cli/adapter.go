package cli

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"github.com/Jsamson33/tide/internal/domain"
)

const (
	ColorReset  = "\033[0m"
	ColorBold   = "\033[1m"
	ColorBlue   = "\033[34m"
	ColorCyan   = "\033[36m"
	ColorYellow = "\033[33m"
	ColorGray   = "\033[90m"
)

type App struct {
	service *domain.TideService
}

func NewApp(service *domain.TideService) *App {
	return &App{service: service}
}

func (a *App) Run(args []string) {
	isJSON := false
	for _, arg := range args {
		if arg == "--json" {
			isJSON = true
		}
	}

	cmd := ""
	if len(args) > 1 && !strings.HasPrefix(args[1], "-") {
		cmd = args[1]
	}

	var res domain.Result
	var err error

	switch cmd {
	case "ask":
		res, err = a.service.LuckyDraw()
	default:
		res, err = a.service.DailyDraw()
	}

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	if isJSON {
		a.displayJSON(res)
	} else {
		a.displayASCII(res)
	}
}

func (a *App) displayJSON(res domain.Result) {
	output, err := json.MarshalIndent(res, "", "  ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error formatting JSON: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(string(output))
}

func (a *App) displayASCII(res domain.Result) {
	fmt.Printf("\n%s--- %s ---%s\n\n", ColorBold+ColorYellow, strings.ToUpper(res.Data.Title), ColorReset)

	// Display lines from top (5) to bottom (0)
	for i := 5; i >= 0; i-- {
		color := ColorCyan
		if i < 3 {
			color = ColorBlue
		}

		if res.Hexagram.Lines[i] == domain.Solid {
			fmt.Printf("  %s━━━━━━━━━━━━%s  ", color, ColorReset)
		} else {
			fmt.Printf("  %s━━━━    ━━━━%s  ", color, ColorReset)
		}

		if i == 4 {
			fmt.Printf("  %s %s %s", ColorCyan, res.Upper.Title, ColorReset)
		} else if i == 1 {
			fmt.Printf("  %s %s %s", ColorBlue, res.Lower.Title, ColorReset)
		}
		fmt.Println()
	}

	fmt.Printf("\n%s%s%s\n", ColorGray, "──────────────────────────────────────────────────", ColorReset)

	// Display content skipping the first line (title)
	contentLines := strings.Split(res.Data.Content, "\n")
	if len(contentLines) > 1 {
		body := strings.Join(contentLines[1:], "\n")
		fmt.Println(strings.TrimSpace(body))
	}
	fmt.Println()
}
