package format

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"bytes"
	"unicode"
	"github.com/fatih/structs"
)

func getRowMaps(rows []interface{}) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, row := range rows {
		result = append(result, structs.Map(row))
	}
	return result
}

func extractAttrs(item map[string]interface{}, attrs []string) ([]string, error) {
	result := make([]string, len(attrs))
	for i, attr := range attrs {
		value, ok := item[attr]
		if !ok {
			return nil, fmt.Errorf("Item has no attribute: %s", attr)
		}
		switch value.(type) {
		case map[string]interface{}:
			txt, err := json.Marshal(value)
			if err != nil {
				return nil, err
			}
			result[i] = fmt.Sprintf("%+v", string(txt))
		default:
			result[i] = fmt.Sprintf("%v", value)
		}
	}
	return result, nil
}

func extractSliceAttrs(items []map[string]interface{}, attrs []string) ([][]string, error) {
	result := make([][]string, len(items))
	for i, item := range items {
		values, err := extractAttrs(item, attrs)
		if err != nil {
			return nil, err
		}
		result[i] = values
	}
	return result, nil
}

func extractColumn(rows [][]string, column int) []string {
	values := make([]string, len(rows))
	for i, row := range rows {
		values[i] = row[column]
	}
	return values
}

func columnWidths(rows [][]string, columnLabels []string, includeCols bool, maxCellLength int) []int {
	if len(rows) == 0 {
		return nil
	}
	columnCount := len(columnLabels)

	widths := make([]int, columnCount)

	if includeCols {
		for i, label := range columnLabels {
			widths[i] = len(label)
		}
	}

	for _, row := range rows {
		for colIndex, colValue := range row {
			valLen := len(colValue)
			if maxCellLength > 0 && valLen > maxCellLength {
				valLen = maxCellLength
			}
			if valLen > widths[colIndex] {
				widths[colIndex] = valLen
			}
		}
	}
	return widths
}

func columnFormats(widths []int) []string {
	formats := make([]string, len(widths))
	for i, width := range widths {
		formats[i] = fmt.Sprintf("%%-%ds", width)
	}
	return formats
}

func sum(items []int) int {
	result := 0
	for _, item := range items {
		result += item
	}
	return result
}

// TableOpts are options used when rendering a table
type TableOpts struct {
	Rows       []interface{}
	Columns    []string
	Separator  string
	ShowHeader bool
	MaxCellWidth int
}

// Table builds a text table from the given data items and chosen columns.
// It returns a list of rows that can be printed.
func Table(opts TableOpts) ([]string, error) {

	if len(opts.Rows) == 0 {
		return nil, errors.New("No rows to display")
	}
	if len(opts.Columns) == 0 {
		return nil, errors.New("No columns to display")
	}

	columnLabels := make([]string, len(opts.Columns))
	for i, name := range opts.Columns {
		columnLabels[i] = strings.ToUpper(toSnakeCase(name))
	}

	rowMaps := getRowMaps(opts.Rows)

	tableData, err := extractSliceAttrs(rowMaps, opts.Columns)
	if err != nil {
		return nil, err
	}

	separator := " | "
	if opts.Separator != "" {
		separator = opts.Separator
	}

	separatorLen := len(separator)
	columnWidths := columnWidths(tableData, columnLabels, opts.ShowHeader, opts.MaxCellWidth)
	columnFormats := columnFormats(columnWidths)
	tableWidth := sum(columnWidths) + separatorLen*(len(opts.Columns)-1)

	rowCount := len(tableData)
	rowOffset := 0
	if opts.ShowHeader {
		rowCount += 3
		rowOffset = 3
	}
	rows := make([]string, rowCount)

	if opts.ShowHeader {
		headers := make([]string, len(opts.Columns))
		for i, label := range columnLabels {
			headers[i] = fmt.Sprintf(columnFormats[i], label)
		}
		rows[0] = strings.Repeat("=", tableWidth)
		rows[1] = strings.Join(headers, separator)
		rows[2] = strings.Repeat("=", tableWidth)
	}

	for i, row := range tableData {
		rowItems := make([]string, len(row))
		columnWidthProgress := 0
		for h, item := range row {
			if(opts.MaxCellWidth > 0) {
				itemLength := len(item)
				columnWidth := columnWidths[h]
				if itemLength > columnWidth {
					item = cellWrap(item, uint(columnWidth), columnWidthProgress + len(separator))
				}
				columnWidthProgress += columnWidth
			}
			rowItems[h] = fmt.Sprintf(columnFormats[h], item)
		}
		rows[i+rowOffset] = strings.Join(rowItems, separator)
	}

	return rows, nil
}

// NormalizeStrings normalizes a slice of strings to all uppercase
func NormalizeStrings(input []string) []string {
	output := make([]string, len(input))
	for i, s := range input {
		output[i] = strings.ToUpper(s)
	}
	return output
}

const nbsp = 0xA0

// Based off https://github.com/mitchellh/go-wordwrap
func cellWrap(s string, lim uint, indentLength int) string {
	// Initialize a buffer with a slightly larger size to account for breaks
	init := make([]byte, 0, len(s))
	buf := bytes.NewBuffer(init)

	var current uint
	var wordBuf, spaceBuf bytes.Buffer
	var wordBufLen, spaceBufLen uint

	for _, char := range s {
		if char == '\n' {
			if wordBuf.Len() == 0 {
				if current+spaceBufLen > lim {
					current = 0
				} else {
					current += spaceBufLen
					spaceBuf.WriteTo(buf)
				}
				spaceBuf.Reset()
				spaceBufLen = 0
			} else {
				current += spaceBufLen + wordBufLen
				spaceBuf.WriteTo(buf)
				spaceBuf.Reset()
				spaceBufLen = 0
				wordBuf.WriteTo(buf)
				wordBuf.Reset()
				wordBufLen = 0
			}
			buf.WriteRune(char)
			current = 0
		} else if unicode.IsSpace(char) && char != nbsp {
			if spaceBuf.Len() == 0 || wordBuf.Len() > 0 {
				current += spaceBufLen + wordBufLen
				spaceBuf.WriteTo(buf)
				spaceBuf.Reset()
				spaceBufLen = 0
				wordBuf.WriteTo(buf)
				wordBuf.Reset()
				wordBufLen = 0
			}

			spaceBuf.WriteRune(char)
			spaceBufLen++
		} else {
			wordBuf.WriteRune(char)
			wordBufLen++

			if current+wordBufLen+spaceBufLen > lim && wordBufLen < lim {
				buf.WriteString("\n" + strings.Repeat(" ", indentLength))
				current = 0
				spaceBuf.Reset()
				spaceBufLen = 0
			}
		}
	}

	if wordBuf.Len() == 0 {
		if current+spaceBufLen <= lim {
			spaceBuf.WriteTo(buf)
		}
	} else {
		spaceBuf.WriteTo(buf)
		wordBuf.WriteTo(buf)
	}

	return buf.String()
}
