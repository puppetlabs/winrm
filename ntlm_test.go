package winrm

import (
	"net/http"

	. "gopkg.in/check.v1"
)

func (s *WinRMSuite) TestHttpNTLMRequest(c *C) {
	ts, host, port, err := StartTestServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/soap+xml")
		w.Write([]byte(response))
	}))
	c.Assert(err, IsNil)
	defer ts.Close()
	endpoint := NewEndpoint(host, port, false, false, nil, nil, nil, 0,0)

	params := DefaultParameters
	params.TransportDecorator = func() Transporter { return &ClientNTLM{} }
	client, err := NewClientWithParameters(endpoint, "test", "test", params)

	c.Assert(err, IsNil)
	shell, err := client.CreateShell()
	c.Assert(err, IsNil)
	c.Assert(shell.id, Equals, "67A74734-DD32-4F10-89DE-49A060483810")
}
