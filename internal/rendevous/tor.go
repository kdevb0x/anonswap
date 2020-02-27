package rendevous

import (
	"context"
	"net"

	"github.com/cretz/bine/tor"
	lt "github.com/ipsn/go-libtor"
)

// newHiddenService creates a new tor hidden service and returns the created
// onion address, a context.CancelFunc, and a nil error on success.
func (s XferServer) newHiddenService(listener net.Listener) ([]byte, func(), error) {
	t, err := tor.Start(context.Background(), &tor.StartConf{ProcessCreator: lt.Creator, UseEmbeddedControlConn: true})
	if err != nil {

		return nil, nil, err
	}
	ctx, cancelfn := context.WithCancel(context.Background())
	service, err := t.Listen(ctx, &tor.ListenConf{Version3: true})
	if err != nil {
		return nil, nil, err
	}
	// we have to return the onion address as a []byte to be able to pass
	// 'nil' on failure instead of "".
	return []byte(service.String()), cancelfn, nil
}
