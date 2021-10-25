package tests

type ContextParams struct {
	Keys   []string
	Values []string
}

func NewCtxParams() *ContextParams {
	return &ContextParams{}
}

func (c *ContextParams) Add(key, value string) *ContextParams {
	c.Keys = append(c.Keys, key)
	c.Values = append(c.Values, value)
	return c
}
