package telegram

import (
	"EfimBot/e"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"path"
	"strconv"
)

type Client struct {
	host     string
	basePath string
	client   http.Client
}

const (
	getUpdatesMethod  = "getUpdates"
	sendMessageMethod = "sendMessage"
)

func New(host string, token string) *Client {
	return &Client{
		host:     host,
		basePath: newBasePath(token),
		client:   http.Client{},
	}
}

func (c *Client) Updates(offset int, limit int) (update []Update, err error) {

	q := url.Values{}
	q.Add("offset", strconv.Itoa(offset))
	q.Add("limit", strconv.Itoa(limit))

	data, err := c.doRequest(getUpdatesMethod, q)
	if err != nil {
		return nil, e.Wrap("can't get updates", err)
	}

	var resp UpdatesResponse

	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, e.Wrap("can't unmarshal updates", err)
	}

	return resp.Result, nil

}

func (c *Client) SendMessage(chatID int, text string) (err error) {
	defer func() { _ = e.WrapIfErr("error occured in send message method", err) }()

	q := url.Values{}
	q.Add("chat_id", strconv.Itoa(chatID))
	q.Add("text", text)

	_, err = c.doRequest(sendMessageMethod, q)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) doRequest(method string, query url.Values) (data []byte, err error) {
	u := url.URL{
		Scheme: "https",
		Host:   c.host,
		Path:   path.Join(c.basePath, method),
	}

	req, err := http.NewRequest(http.MethodGet, u.String(), nil)

	if err != nil {
		return nil, err
	}

	req.URL.RawQuery = query.Encode()

	resp, err := c.client.Do(req)

	if err != nil {
		return nil, err
	}

	defer func() { _ = resp.Body.Close() }()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	return body, nil

}

func newBasePath(token string) string {
	return "bot" + token
}
