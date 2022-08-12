package mock

type Client struct {
	Category *CategoryClient
}

func NewClient() *Client {
	client := &Client{}
	client.init()
	return client
}

func (c *Client) init() {
	c.Category = NewCategoryClient()
}

func (c *Client) Close() error {
	return nil
}
