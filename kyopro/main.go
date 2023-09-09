package main

import (
	"bufio"
	"errors"
	"io"
	"strconv"
)

func main() {
	// cio := NewCustomIO(os.Stdin)
	// n := cio.MustScanInt()
}

//
// CustomIO is i/o for kyopro
type CustomIO struct {
	scanner *bufio.Scanner
}

func NewCustomIO(r io.Reader) *CustomIO {
	sc := bufio.NewScanner(r)
	sc.Split(bufio.ScanWords)

	return &CustomIO{
		scanner: sc,
	}
}

func (c *CustomIO) scan() (string, error) {
	if !c.scanner.Scan() {
		if err := c.scanner.Err(); err != nil {
			return "", err
		}
		return "", errors.New("EOF error")
	}
	return c.scanner.Text(), nil
}

func (c *CustomIO) ScanStr() (string, error) {
	return c.scan()
}

func (c *CustomIO) MustScanStr() string {
	s, err := c.ScanStr()
	if err != nil {
		panic(err)
	}
	return s
}

func (c *CustomIO) ScanStrN(n int) ([]string, error) {
	out := make([]string, 0, n)
	for range make([]struct{}, n) {
		s, err := c.ScanStr()
		if err != nil {
			return nil, err
		}
		out = append(out, s)
	}
	return out, nil
}

func (c *CustomIO) MustScanStrN(n int) []string {
	out := make([]string, 0, n)
	for range make([]struct{}, n) {
		out = append(out, c.MustScanStr())
	}
	return out
}

func (c *CustomIO) ScanInt() (int, error) {
	s, err := c.scan()
	if err != nil {
		return 0, err
	}
	return strconv.Atoi(s)
}

func (c *CustomIO) MustScanInt() int {
	i, err := c.ScanInt()
	if err != nil {
		panic(err)
	}
	return i
}

func (c *CustomIO) ScanIntN(n int) ([]int, error) {
	out := make([]int, 0, n)
	for range make([]struct{}, n) {
		i, err := c.ScanInt()
		if err != nil {
			return nil, err
		}
		out = append(out, i)
	}
	return out, nil
}

func (c *CustomIO) MustScanIntN(n int) []int {
	out := make([]int, 0, n)
	for range make([]struct{}, n) {
		out = append(out, c.MustScanInt())
	}
	return out
}
