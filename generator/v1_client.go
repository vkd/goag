package generator

type Client struct {
	Handlers []Handler

	IsDecodeJSONFunc bool
}

func NewClient(handlers []Handler) Client {
	c := Client{
		Handlers: handlers,
	}
	return c
}

func (c Client) Execute() (string, error) { return templates.ExecuteTemplate("Client", c) }
