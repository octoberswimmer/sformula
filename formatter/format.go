package formatter

import (
	"bytes"
	"fmt"
	"io"
	"regexp"
	"strconv"
	"strings"

	"github.com/antlr4-go/antlr/v4"
	"github.com/octoberswimmer/sformula/parser"
)

type errorListener struct {
	*antlr.DefaultErrorListener
	logger io.Writer
}

func (e *errorListener) SyntaxError(_ antlr.Recognizer, _ interface{}, line, column int, msg string, _ antlr.RecognitionException) {
	_, _ = fmt.Fprintln(e.logger, "line "+strconv.Itoa(line)+":"+strconv.Itoa(column)+" "+msg)
}

func Format(f string) (string, error) {
	input := antlr.NewInputStream(f)
	lexer := parser.NewFormulaLexer(input)
	lexer.RemoveErrorListeners()

	var buf bytes.Buffer
	errors := &errorListener{logger: &buf}
	lexer.AddErrorListener(errors)

	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	p := parser.NewFormulaParser(stream)
	p.RemoveErrorListeners()
	p.AddErrorListener(errors)
	// p.AddErrorListener(antlr.NewDiagnosticErrorListener(false))

	v := NewFormatVisitor(stream)
	out, ok := v.visitRule(p.CompilationUnit()).(string)
	if !ok {
		return f, fmt.Errorf("Unexpected result parsing formula")
	}
	if buf.Len() != 0 {
		return f, fmt.Errorf("syntax error: %s", string(buf.Bytes()))
	}
	out = removeExtraCommentIndentation(out)
	return out, nil
}

func removeIndentationFromComment(comment string) string {
	// Find the position of the initial \uFFFA and the final \uFFFB
	startIndex := strings.Index(comment, "\uFFFA")
	endIndex := strings.LastIndex(comment, "\uFFFB")
	if startIndex == -1 || endIndex == -1 || endIndex <= startIndex {
		// \uFFFA or \uFFFB not found, or the indices are invalid, return the original comment
		return comment
	}

	// Determine the indentation level from the first line
	firstLine := comment[:startIndex]
	leadingTabs := strings.Count(firstLine, "\t")

	// Extract the content between \uFFFA and \uFFFB
	commentBody := comment[startIndex+len("\uFFFA") : endIndex]

	// Create a regex to match the leading tabs in subsequent lines
	tabPattern := fmt.Sprintf(`\n\t{%d}`, leadingTabs)
	re := regexp.MustCompile(tabPattern)

	// Replace the matched pattern with a newline, effectively removing the leading tabs
	modifiedComment := re.ReplaceAllString(commentBody, "\n")

	// Return the modified comment, reattaching any text outside the comment block if necessary
	return firstLine + modifiedComment
}

// Comments are annotated in FormatVisitor.visitRule.  We preserve whitespace
// within multi-line comments by removing the indentation added within the
// comment.
func removeExtraCommentIndentation(input string) string {
	commentPattern := regexp.MustCompile(`(?s)\t*` + "\uFFFA" + `.*?` + "\uFFFB")
	return commentPattern.ReplaceAllStringFunc(input, removeIndentationFromComment)
}
