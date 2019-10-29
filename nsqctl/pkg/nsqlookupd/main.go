package nsqlookupd

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Client struct {
	Url string
	Cli *http.Client
}

type Producer struct {
	RemoteAddr string `json:"remote_address"`
	Hostname   string `json:"hostname"`
	BcastAddr  string `json:"broadcast_address"`
	TcpPort    int    `json:"tcp_port"`
	HttpPort   int    `json:"http_port"`
	Version    string `json:"version"`
}

type Producers struct {
	List []Producer `json:"producers"`
}

// GetProducersForTopic returns a slice of producers (ip:port).
// In the event an error is produced, an empty slice is returned along with the error.
func (c *Client) GetProducersForTopic(t string) ([]string, error) {
	req, err := http.NewRequest("GET", c.Url, nil)
	if err != nil {
		return []string{}, err
	}

	q := req.URL.Query()
	q.Add("topic", t)
	req.URL.RawQuery = q.Encode()

	resp, err := c.Cli.Do(req)
	if err != nil {
		return []string{}, err
	}
	defer resp.Body.Close()

	p := Producers{}
	err = json.NewDecoder(resp.Body).Decode(&p)
	if err != nil {
		return []string{}, err
	}

	hosts := make([]string, len(p.List))
	i := 0
	for _, v := range p.List {
		ip := v.RemoteAddr
		port := v.HttpPort
		hosts[i] = fmt.Sprintf("%s:%d", ip, port)
		i++
	}

	return hosts, nil
}
