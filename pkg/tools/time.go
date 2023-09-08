package tools

/*	License: GPLv3
	Authors:
		Mirko Brombin <mirko@fabricators.ltd>
		Lit Contributors <https://github.com/linux-immutability-tools/>
	Copyright: 2023
	Description:
		Containers Wrapper is a Go library that provides a convenient
		and unified interface for interacting with container engines
		such as Docker, Podman, and Containerd.
*/

import (
	"regexp"
	"strconv"
	"strings"
	"time"
)

// ParseTime parses the given time string, even a fancy one like
// "2 weeks ago" or "10 minutes ago" into a time.Time struct.
func ParseTime(timeStr string) (t time.Time, err error) {
	regex := regexp.MustCompile(`(\d+)\s+(second|minute|hour|day|week|month|year)s?\s+ago`)

	match := regex.FindStringSubmatch(timeStr)
	if match != nil {
		num, _ := strconv.Atoi(match[1])
		unit := strings.ToLower(match[2])

		var duration time.Duration
		switch unit {
		case "second":
			duration = time.Second * time.Duration(num)
		case "minute":
			duration = time.Minute * time.Duration(num)
		case "hour":
			duration = time.Hour * time.Duration(num)
		case "day":
			duration = 24 * time.Hour * time.Duration(num)
		case "week":
			duration = 7 * 24 * time.Hour * time.Duration(num)
		case "month":
			duration = 30 * 24 * time.Hour * time.Duration(num)
		case "year":
			duration = 365 * 24 * time.Hour * time.Duration(num)
		}

		t = time.Now().Add(-duration)
		return t, nil
	}

	t, err = time.Parse(time.RFC3339, timeStr)
	if err == nil {
		return t, nil
	}

	return t, err
}
