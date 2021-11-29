package decimal

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/Rhymond/go-money"
)

func NewFromString(s string) (*money.Money, error) {
	re, err := regexp.Compile(`^(-?)(\d+)(,(\d{2}))?$`)
	if err != nil {
		return nil, err
	}
	s = strings.TrimSpace(s)
	if !re.MatchString(s) {
		return nil, fmt.Errorf("failed to parse %q", s)
	}
	m := re.FindStringSubmatch(s)
	amt := m[1] + m[2] + m[4]

	amtInt, err := strconv.ParseInt(amt, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("failed parsing %q: %w", amt, err)
	}
	return money.New(amtInt, "SEK"), nil
}
