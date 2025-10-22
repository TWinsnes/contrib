package contrib

import (
	"context"
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/go-git/go-git/v6"
	"github.com/go-git/go-git/v6/plumbing/object"
)

type ContribOptions struct {
	// Tag to open the contrib block
	openTag string
	// Tag to close the contrib block
	closeTag string

	// Path to the file to add the contrib block to
	Path string

	// Header to add to the top of the contrib block
	Header string
}

func NewContribOptions() *ContribOptions {
	return &ContribOptions{
		openTag:  "<!---Contrib Block Start-->",
		closeTag: "<!---Contrib Block End-->",
		Path:     "README.md",
		Header:   "## Contributors\n\nWithout these amazing people, this project wouldn't exist.",
	}
}

func (c *ContribOptions) Run(ctx context.Context) error {

	fmt.Println("Scanning git logs for contributors...")
	authors, err := c.getAuthors(ctx)

	if err != nil {
		return err
	}

	fmt.Println("Found ", len(authors), " contributors")

	sortedAuthors := c.sorthByCommits(authors)

	fmt.Printf("Writing contributor list to %s\n", c.Path)
	err = c.writeAuthors(sortedAuthors)
	if err != nil {
		return err
	}
	fmt.Println("Done")
	return nil
}

func (c *ContribOptions) writeAuthors(authors []Author) error {
	var content string

	data, err := os.ReadFile(c.Path)
	if err != nil && !os.IsNotExist(err) {
		return err
	}

	if err == nil {
		content = string(data)
	}

	var builder strings.Builder
	startIdx := strings.Index(content, c.openTag)
	endIdx := strings.Index(content, c.closeTag)

	if startIdx == -1 || endIdx == -1 {
		builder.WriteString(content)
		if len(content) > 0 && !strings.HasSuffix(content, "\n") {
			builder.WriteString("\n")
		}
	} else {
		builder.WriteString(content[:startIdx])
	}

	builder.WriteString(c.openTag + "\n")
	builder.WriteString(c.Header + "\n\n")
	for _, author := range authors {
		builder.WriteString(fmt.Sprintf("- %s\n", author.Name))
	}

	if startIdx == -1 && endIdx == -1 {
		builder.WriteString(c.closeTag + "\n")
	} else {
		builder.WriteString(content[endIdx:])
	}

	return os.WriteFile(c.Path, []byte(builder.String()), 0644)
}

func (c *ContribOptions) sorthByCommits(authors map[string]Author) []Author {
	result := make([]Author, 0, len(authors))
	for _, author := range authors {
		result = append(result, author)
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i].Commits > result[j].Commits
	})

	return result
}

func (c *ContribOptions) getAuthors(ctx context.Context) (map[string]Author, error) {
	r, err := git.PlainOpen(".")
	if err != nil {
		return map[string]Author{}, err
	}

	ref, err := r.Head()
	if err != nil {
		return map[string]Author{}, err
	}

	opt := &git.LogOptions{
		From: ref.Hash(),
	}

	commits, err := r.Log(opt)

	if err != nil {
		return map[string]Author{}, err
	}

	authors := map[string]Author{}

	err = commits.ForEach(func(c *object.Commit) error {
		if ctx.Err() != nil {
			return ctx.Err()
		}

		val, ok := authors[c.Author.Name]
		if !ok {
			val = Author{
				Name:    c.Author.Name,
				Email:   c.Author.Email,
				Commits: 0,
			}
		}

		val.Commits++
		authors[c.Author.Name] = val

		return nil
	})

	if err != nil {
		return map[string]Author{}, err
	}
	return authors, nil
}
