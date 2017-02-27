package server

import (
	"net"
	"errors"
	"time"
)

/*
    Custom TCP listener that could receive stop signal inspired by
    http://www.hydrogen18.com/blog/stop-listening-http-server-go.html
 */
type ServerListener struct {
	*net.TCPListener        // Backing TCP listener
	stop    chan int        // Indicator that listener should shutdown
	Open    bool            // Indicator that this listener opened or not
}

var StoppedError = errors.New("Server listener stopped")


func NewTCPListener(l net.Listener) (*ServerListener, error) {
	tcpl, ok := l.(*net.TCPListener)
	
	if !ok {
		return nil, errors.New("Cannot wrap listener")
	}
	
	newl := &ServerListener{
		TCPListener: tcpl,
		stop: make(chan int),
	}
	
	return newl, nil
}


/*
	Override accept method of backed listener
 */
func (sl *ServerListener) Accept() (net.Conn, error) {
	for {
		// Wait up to one second for a new connection
		sl.SetDeadline(time.Now().Add(time.Second))
		
		// Accept backed TCP listener
		newCon, err := sl.TCPListener.Accept()
		
		// Mark this listener opened
		sl.Open = true
		
		// Check for the channel being closed
		select {
		case <-sl.stop:
			sl.Open = false
			if err == nil {
				newCon.Close()
			}
			return nil, StoppedError
		default:
			// If the channel is still open, continue as normal
		}
		
		if err != nil {
			netErr, ok := err.(net.Error)
			
			// If this is a timeout, then continue to wait for new connections
			if ok && netErr.Timeout() && netErr.Temporary() {
				continue
			}
		}
		
		return newCon, err
	}
}


func (sl *ServerListener) Close() (error) {
	// Close this channel
	if sl.Open {
		close(sl.stop)
	}
	
	return nil
}
