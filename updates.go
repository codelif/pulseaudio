package pulseaudio

func (c *Client) ensureSubscribed(mask uint32) error {
	c.subMu.Lock()
	defer c.subMu.Unlock()
	if c.subscribed && c.subMask == mask {
		return nil
	}
	if _, err := c.request(commandSubscribe, uint32Tag, uint32(mask)); err != nil {
		return err
	}
	c.subscribed = true
	c.subMask = mask
	return nil
}

const subscriptionMaskAll = 0x02ff

func (c *Client) Updates() (<-chan struct{}, error) {
	if err := c.ensureSubscribed(subscriptionMaskAll); err != nil {
		return nil, err
	}
	return c.updates, nil
}

func (c *Client) Events() (<-chan Event, error) {
	if err := c.ensureSubscribed(subscriptionMaskAll); err != nil {
		return nil, err
	}
	return c.events, nil
}
