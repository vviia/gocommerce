package entities

type CommonResult struct {
	ResCode    int    `json:"-"`
	ResMessage string `json:"-"`
}

func (c *CommonResult) SetCode(code int, err error) {
	c.ResCode = code
	if err != nil {
		c.ResMessage = err.Error()
	}

	if code >= 500 {
		// todo: sendo to slack, sentry, etc.
		panic(err)
	}
}
