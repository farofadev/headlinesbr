package payloads

type Payload struct {
	Meta map[string]interface{} `json:"meta,omitempty"`
	Data interface{}            `json:"data"`
}

type PayloadOption func(payload *Payload) *Payload

func NewPayload(options ...PayloadOption) *Payload {
	payload := &Payload{
		Meta: make(map[string]interface{}),
		Data: nil,
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

func WithData(data interface{}) func(payload *Payload) *Payload {
	return func(payload *Payload) *Payload {
		return payload.SetData(data)
	}
}

func WithMeta(meta map[string]interface{}) func(payload *Payload) *Payload {
	return func(payload *Payload) *Payload {
		return payload.SetMeta(meta)
	}
}

func WithAddMeta(key string, value interface{}) func(payload *Payload) *Payload {
	return func(payload *Payload) *Payload {
		return payload.AddMeta(key, value)
	}
}
