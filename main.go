package main

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

//go:embed glyphs.json
var glyphsData []byte

type Glyph struct {
	Value string `json:"value"`
	Name  string `json:"name"`
	Group string `json:"group"`
}

func main() {
	// Load glyphs from embedded data
	var glyphs []Glyph
	if err := json.Unmarshal(glyphsData, &glyphs); err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing embedded JSON: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Loaded %d glyphs\n", len(glyphs))

	// Prepare fzf input
	var lines []string
	for _, glyph := range glyphs {
		line := fmt.Sprintf("%s %s [%s]", glyph.Value, glyph.Name, glyph.Group)
		lines = append(lines, line)
	}

	// Run fzf
	selected, err := runFzf(lines)
	if err != nil {
		if err.Error() == "exit status 130" {
			// User cancelled (Ctrl+C)
			os.Exit(0)
		}
		fmt.Fprintf(os.Stderr, "fzf error: %v\n", err)
		os.Exit(1)
	}

	if selected == "" {
		os.Exit(0)
	}

	// Extract glyph value (first character/rune)
	glyphValue := strings.Fields(selected)[0]

	// Copy to clipboard
	if err := copyToClipboard(glyphValue); err != nil {
		fmt.Fprintf(os.Stderr, "Error copying to clipboard: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Copied '%s' to clipboard!\n", glyphValue)
}

func copyToClipboard(text string) error {
	var cmd *exec.Cmd

	// Try different clipboard commands
	if _, err := exec.LookPath("wl-copy"); err == nil {
		// Wayland
		cmd = exec.Command("wl-copy")
	} else if _, err := exec.LookPath("xclip"); err == nil {
		// X11 with xclip
		cmd = exec.Command("xclip", "-selection", "clipboard")
	} else if _, err := exec.LookPath("xsel"); err == nil {
		// X11 with xsel
		cmd = exec.Command("xsel", "--clipboard", "--input")
	} else {
		return fmt.Errorf("no clipboard utility found (install xclip, xsel, or wl-clipboard)")
	}

	stdin, err := cmd.StdinPipe()
	if err != nil {
		return err
	}

	go func() {
		defer stdin.Close()
		stdin.Write([]byte(text))
	}()

	return cmd.Run()
}

func runFzf(lines []string) (string, error) {
	cmd := exec.Command("fzf",
		"--ansi",
		"--preview-window=up:5:wrap",
		"--preview=printf '\\n  %s %s %s %s\\n\\nName: %s\\nGroup: %s\\n' {1} {1} {1} {1} {2} {3}",
		"--height=80%",
		"--border",
		"--prompt=Search glyphs: ",
		"--header=Press Enter to copy glyph to clipboard • Use Ctrl+C to exit",
	)

	stdin, err := cmd.StdinPipe()
	if err != nil {
		return "", err
	}

	go func() {
		defer stdin.Close()
		for _, line := range lines {
			fmt.Fprintln(stdin, line)
		}
	}()

	output, err := cmd.Output()
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(output)), nil
}
