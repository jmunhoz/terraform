package rgw

type Config struct{}

// Client() returns a new client.
func (c *Config) Client() (interface{}, error) {
	return nil, nil
}
