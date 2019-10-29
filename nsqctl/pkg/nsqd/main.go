package nsqd

import (
	"fmt"
	"net/http"
)

type Client struct {
	Url string
	Cli *http.Client
}

// PauseTopic Pauses a topic.
func (c *Client) PauseTopic(t string) error {
	req, err := http.NewRequest("POST", c.Url, nil)
	if err != nil {
		return err
	}

	q := req.URL.Query()
	q.Add("topic", t)
	req.URL.RawQuery = q.Encode()

	resp, err := c.Cli.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return fmt.Errorf("pause of topic: %s on host: %s http code: %s", t, c.Url, resp.StatusCode)
	}

	return nil
}
