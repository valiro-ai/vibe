package sep

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"gopkg.in/yaml.v3"
)

// Status constants for SEP lifecycle
const (
	StatusDraft     = "DRAFT"
	StatusAccepted  = "ACCEPTED"
	StatusBlocked   = "BLOCKED"
	StatusCancelled = "CANCELLED"
	StatusDone      = "DONE"
)

// ValidStatuses lists all valid SEP statuses
var ValidStatuses = []string{
	StatusDraft,
	StatusAccepted,
	StatusBlocked,
	StatusCancelled,
	StatusDone,
}

// Frontmatter represents the YAML frontmatter of a SEP
type Frontmatter struct {
	Title     string   `yaml:"title"`
	Status    string   `yaml:"status"`
	Created   string   `yaml:"created"`
	DependsOn []string `yaml:"depends_on"`
	Areas     []string `yaml:"areas,omitempty"`
	Assigned  string   `yaml:"assigned,omitempty"`
}

// SEP represents a Software Enhancement Proposal
type SEP struct {
	Number         string   // e.g., "0001"
	Title          string   // e.g., "User Authentication"
	Status         string   // DRAFT, ACCEPTED, BLOCKED, CANCELLED, DONE
	Created        string   // YYYY-MM-DD
	DependsOn      []string // e.g., ["0001", "0002"]
	Areas          []string // e.g., ["auth/*", "api/routes/login.go"]
	Assigned       string   // e.g., "@alice" - pilot assigned to implement
	WhatAndWhy     string   // Content of What & Why section
	DoneWhen       []string // Acceptance criteria
	DoneWhenStatus []bool   // Checked status of each criterion
	FilePath       string   // Full path to file
}

// Parse reads a SEP file and extracts its content
func Parse(filePath string) (*SEP, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	sep := &SEP{FilePath: filePath}

	// Extract number from filename
	base := filepath.Base(filePath)
	numMatch := regexp.MustCompile(`^(\d{4})-`).FindStringSubmatch(base)
	if len(numMatch) == 2 {
		sep.Number = numMatch[1]
	}

	scanner := bufio.NewScanner(file)
	var inFrontmatter bool
	var frontmatterLines strings.Builder
	var currentSection string
	var whatAndWhy strings.Builder

	for scanner.Scan() {
		line := scanner.Text()

		// Handle YAML frontmatter
		if line == "---" {
			if !inFrontmatter {
				inFrontmatter = true
				continue
			} else {
				// End of frontmatter, parse it
				inFrontmatter = false
				var fm Frontmatter
				if err := yaml.Unmarshal([]byte(frontmatterLines.String()), &fm); err == nil {
					sep.Title = fm.Title
					sep.Status = fm.Status
					sep.Created = fm.Created
					sep.DependsOn = fm.DependsOn
					sep.Areas = fm.Areas
					sep.Assigned = fm.Assigned
				}
				continue
			}
		}

		if inFrontmatter {
			frontmatterLines.WriteString(line)
			frontmatterLines.WriteString("\n")
			continue
		}

		// Track sections
		if after, found := strings.CutPrefix(line, "## "); found {
			currentSection = after
			continue
		}

		// Parse section content
		switch currentSection {
		case "What & Why":
			if line != "" && !strings.HasPrefix(line, "[") {
				whatAndWhy.WriteString(line)
				whatAndWhy.WriteString("\n")
			}
		case "Done When":
			if strings.HasPrefix(line, "- [") {
				checked := strings.HasPrefix(line, "- [x]") || strings.HasPrefix(line, "- [X]")
				criterion := strings.TrimSpace(line[5:])
				if criterion != "" && !strings.HasPrefix(criterion, "[") {
					sep.DoneWhen = append(sep.DoneWhen, criterion)
					sep.DoneWhenStatus = append(sep.DoneWhenStatus, checked)
				}
			}
		}
	}

	sep.WhatAndWhy = strings.TrimSpace(whatAndWhy.String())

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return sep, nil
}

// List finds and parses all SEPs in the given directory
func List(dir string) ([]*SEP, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	re := regexp.MustCompile(`^\d{4}-.*\.md$`)
	var seps []*SEP

	for _, entry := range entries {
		if entry.IsDir() || !re.MatchString(entry.Name()) {
			continue
		}

		sep, err := Parse(filepath.Join(dir, entry.Name()))
		if err != nil {
			continue // Skip files that can't be parsed
		}
		seps = append(seps, sep)
	}

	return seps, nil
}

// GroupByStatus groups SEPs by their status
func GroupByStatus(seps []*SEP) map[string][]*SEP {
	groups := make(map[string][]*SEP)
	for _, sep := range seps {
		groups[sep.Status] = append(groups[sep.Status], sep)
	}
	return groups
}

// FindByNumber finds a SEP by its number in the given directory
func FindByNumber(dir, number string) (*SEP, error) {
	// Pad number if needed
	if len(number) < 4 {
		number = fmt.Sprintf("%04s", number)
	}

	seps, err := List(dir)
	if err != nil {
		return nil, err
	}

	for _, sep := range seps {
		if sep.Number == number {
			return sep, nil
		}
	}

	return nil, fmt.Errorf("SEP not found: %s", number)
}

// NextNumber determines the next available SEP number
func NextNumber(dir string) (string, error) {
	seps, err := List(dir)
	if err != nil && !os.IsNotExist(err) {
		return "", err
	}

	maxNum := 0
	for _, sep := range seps {
		var num int
		fmt.Sscanf(sep.Number, "%d", &num)
		if num > maxNum {
			maxNum = num
		}
	}

	return fmt.Sprintf("%04d", maxNum+1), nil
}

// ID returns the full SEP ID (e.g., "SEP-0001")
func (s *SEP) ID() string {
	return fmt.Sprintf("SEP-%s", s.Number)
}

// Conflict represents an overlap between two SEPs
type Conflict struct {
	SEP1         *SEP
	SEP2         *SEP
	OverlapAreas []string
}

// FindConflicts detects SEPs with overlapping areas
func FindConflicts(seps []*SEP) []Conflict {
	var conflicts []Conflict

	for i := 0; i < len(seps); i++ {
		for j := i + 1; j < len(seps); j++ {
			// Skip if either SEP has no areas defined
			if len(seps[i].Areas) == 0 || len(seps[j].Areas) == 0 {
				continue
			}

			// Skip if either is DONE or CANCELLED
			if seps[i].Status == StatusDone || seps[i].Status == StatusCancelled {
				continue
			}
			if seps[j].Status == StatusDone || seps[j].Status == StatusCancelled {
				continue
			}

			overlaps := findOverlappingAreas(seps[i].Areas, seps[j].Areas)
			if len(overlaps) > 0 {
				conflicts = append(conflicts, Conflict{
					SEP1:         seps[i],
					SEP2:         seps[j],
					OverlapAreas: overlaps,
				})
			}
		}
	}

	return conflicts
}

// findOverlappingAreas checks if two area lists have overlaps
func findOverlappingAreas(areas1, areas2 []string) []string {
	var overlaps []string

	for _, a1 := range areas1 {
		for _, a2 := range areas2 {
			if areasOverlap(a1, a2) {
				// Add the more specific one, or both if equal
				if len(a1) >= len(a2) {
					overlaps = append(overlaps, a1)
				} else {
					overlaps = append(overlaps, a2)
				}
			}
		}
	}

	return overlaps
}

// areasOverlap checks if two area patterns overlap
func areasOverlap(a1, a2 string) bool {
	// Exact match
	if a1 == a2 {
		return true
	}

	// Check if one is a prefix of the other (directory containment)
	// e.g., "internal/user/*" overlaps with "internal/user/model.go"
	a1Base := strings.TrimSuffix(a1, "/*")
	a2Base := strings.TrimSuffix(a2, "/*")

	if strings.HasPrefix(a1, a2Base) || strings.HasPrefix(a2, a1Base) {
		return true
	}

	return false
}

// UpdateStatus updates the status field in a SEP file
func (s *SEP) UpdateStatus(newStatus string) error {
	content, err := os.ReadFile(s.FilePath)
	if err != nil {
		return err
	}

	// Extract frontmatter
	parts := strings.SplitN(string(content), "---", 3)
	if len(parts) < 3 {
		return fmt.Errorf("invalid frontmatter format")
	}

	// Parse frontmatter
	var fm Frontmatter
	if err := yaml.Unmarshal([]byte(parts[1]), &fm); err != nil {
		return fmt.Errorf("failed to parse frontmatter: %w", err)
	}

	// Update status
	fm.Status = newStatus

	// Marshal back to YAML
	newFrontmatter, err := yaml.Marshal(&fm)
	if err != nil {
		return fmt.Errorf("failed to marshal frontmatter: %w", err)
	}

	// Rebuild file content
	newContent := "---\n" + string(newFrontmatter) + "---" + parts[2]

	// Write back
	if err := os.WriteFile(s.FilePath, []byte(newContent), 0644); err != nil {
		return err
	}

	s.Status = newStatus
	return nil
}

// Assign sets the assigned pilot for a SEP
func (s *SEP) Assign(pilot string) error {
	content, err := os.ReadFile(s.FilePath)
	if err != nil {
		return err
	}

	// Extract frontmatter
	parts := strings.SplitN(string(content), "---", 3)
	if len(parts) < 3 {
		return fmt.Errorf("invalid frontmatter format")
	}

	// Parse frontmatter
	var fm Frontmatter
	if err := yaml.Unmarshal([]byte(parts[1]), &fm); err != nil {
		return fmt.Errorf("failed to parse frontmatter: %w", err)
	}

	// Update assigned
	fm.Assigned = pilot

	// Marshal back to YAML
	newFrontmatter, err := yaml.Marshal(&fm)
	if err != nil {
		return fmt.Errorf("failed to marshal frontmatter: %w", err)
	}

	// Rebuild file content
	newContent := "---\n" + string(newFrontmatter) + "---" + parts[2]

	// Write back
	if err := os.WriteFile(s.FilePath, []byte(newContent), 0644); err != nil {
		return err
	}

	s.Assigned = pilot
	return nil
}
