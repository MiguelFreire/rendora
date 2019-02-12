package rendora

type Pool struct {
	pool chan *headlessClient
}

func newPool(size int) *Pool {
	return &Pool{
		pool: make(chan *headlessClient, size),
	}
}

func (p *Pool) getHeadlessClient() *headlessClient {
	var h *headlessClient
	select {
	case h = <-p.pool:
	}

	return h
}

func (p *Pool) putHeadlessClient(h *headlessClient) {
	select {
	case p.pool <- h:
	}
}
