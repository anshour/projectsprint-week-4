package formatTime

import "time"

func FormatToISO8601WithNano(datetime string) (string, error) {
	t, err := time.Parse(time.RFC3339Nano, datetime)
	if err != nil {
		return "", err
	}
	return t.Format(time.RFC3339Nano), nil
}
