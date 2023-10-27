package module

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

func ReadArray(reader *bufio.Reader) ([]string, error) {
	line, err := reader.ReadString('\n')
	if err != nil {
		return nil, err
	}

	if line[0] != '*' {
		return nil, fmt.Errorf("expected '*', got %q", line[0])
	}

	count, err := strconv.Atoi(strings.TrimSpace(line[1:]))
	if err != nil {
		return nil, err
	}

	result := make([]string, count)
	for i := 0; i < count; i++ {
		result[i], err = readBulkString(reader)
		if err != nil {
			return nil, err
		}
	}

	return result, nil
}

func readBulkString(reader *bufio.Reader) (string, error) {
	line, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	if line[0] != '$' {
		return "", fmt.Errorf("expected '$', got %q", line[0])
	}

	length, err := strconv.Atoi(strings.TrimSpace(line[1:]))
	if err != nil {
		return "", err
	}

	data := make([]byte, length)
	_, err = io.ReadFull(reader, data)
	if err != nil {
		return "", err
	}

	_, err = reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	return string(data), nil
}

func WriteSimpleString(writer *bufio.Writer, s string) {
	_, _ = writer.WriteString("+" + s + "\r\n")
}

func WriteError(writer *bufio.Writer, s string) {
	_, _ = writer.WriteString("-" + s + "\r\n")
}

func WriteBulkString(writer *bufio.Writer, s string) {
	_, _ = writer.WriteString("$" + strconv.Itoa(len(s)) + "\r\n" + s + "\r\n")
}

func WriteNullBulkString(writer *bufio.Writer) {
	_, _ = writer.WriteString("$-1\r\n")
}
