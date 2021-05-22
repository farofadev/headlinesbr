package payloads

type Payload struct {
	Meta   map[string]interface{} `json:"meta,omitempty"`
	Extras map[string]interface{} `json:"extras,omitempty"`
	Data   interface{}            `json:"data"`
}

type PayloadOption func(payload *Payload) *Payload

func NewPayload(options ...PayloadOption) *Payload {
	payload := &Payload{
		Meta:   make(map[string]interface{}),
		Extras: make(map[string]interface{}),
		Data:   nil,
	}

	for _, opt := range options {
		payload = opt(payload)
	}

	return payload
}

func (p *Payload) SetData(data interface{}) *Payload {
	p.Data = data

	return p
}

func (p *Payload) SetMeta(meta map[string]interface{}) *Payload {
	p.Meta = meta

	return p
}

func (p *Payload) AddMeta(key string, value interface{}) *Payload {
	p.Meta[key] = value

	return p
}

func (p *Payload) SetExtras(extras map[string]interface{}) *Payload {
	p.Extras = extras

	return p
}

func (p *Payload) AddExtra(key string, value interface{}) *Payload {
	p.Extras[key] = value

	return p
}

func WithMeta(meta map[string]interface{}) func(payload *Payload) *Payload {
	return func(payload *Payload) *Payload {
		return payload.SetMeta(meta)
	}
}

func AddMeta(key string, value interface{}) func(payload *Payload) *Payload {
	return func(payload *Payload) *Payload {
		return payload.AddMeta(key, value)
	}
}

func AddMetaWhen(condition bool, key string, value func() interface{}) func(payload *Payload) *Payload {
	if !condition {
		return func(p *Payload) *Payload { return p }
	}

	return AddMeta(key, value())
}

func WithData(data interface{}) func(payload *Payload) *Payload {
	return func(payload *Payload) *Payload {
		return payload.SetData(data)
	}
}

func WithExtras(extras map[string]interface{}) func(payload *Payload) *Payload {
	return func(payload *Payload) *Payload {
		return payload.SetExtras(extras)
	}
}

func AddExtra(key string, value interface{}) func(payload *Payload) *Payload {
	return func(payload *Payload) *Payload {
		return payload.AddExtra(key, value)
	}
}

func AddExtraWhen(condition bool, key string, value func() interface{}) func(payload *Payload) *Payload {
	if !condition {
		return func(p *Payload) *Payload { return p }
	}

	return AddExtra(key, value())
}
