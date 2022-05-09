package lib

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"
)

var rpcServer *rpc.Server

type HttpConn struct {
	in  io.Reader
	out io.Writer
}

func (c *HttpConn) Read(p []byte) (n int, err error)  { return c.in.Read(p) }
func (c *HttpConn) Write(d []byte) (n int, err error) { return c.out.Write(d) }
func (c *HttpConn) Close() error                      { return nil }
func startRPCHandler() {
	rpcServer = rpc.NewServer()

	l, err := net.Listen("tcp", ":"+RPCPort)
	if err != nil {
		fmt.Printf("Listener tcp err: %s", err)
		return
	}
	go func() {
		http.Serve(l, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			serverCodec := jsonrpc.NewServerCodec(&HttpConn{in: r.Body, out: w})
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,OPTIONS")
			w.Header().Set("Access-Control-Request-Method", "*")
			w.Header().Set("Access-Control-Allow-Headers", "Accept, Origin, XRequestedWith, Content-Type, LastModified")

			rpcServer.ServeRequest(serverCodec)

		}))
	}()

}
