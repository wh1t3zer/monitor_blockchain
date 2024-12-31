package push

import (
	"fmt"
	"net/http"
	"pushbot/config"
)

func PushBark(msg string) error {
	url := config.HandleYaml().Push.Bark.Hook
	pushUrl := url + msg
	resp, err := http.Get(pushUrl)
	if err != nil {
		return fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	return nil
}
