package main

import (
	"errors"
	"io"
	"net"
	"time"
)

var ErrNoEstablishedConnection = errors.New("no established connection")

type TelnetClient interface {
	Connect() error
	io.Closer
	Send() error
	Receive() error
}

type telnetClient struct {
	address    string
	timeout    time.Duration
	connection net.Conn
	in         io.ReadCloser
	out        io.Writer
}

func NewTelnetClient(address string, timeout time.Duration, in io.ReadCloser, out io.Writer) TelnetClient {
	return &telnetClient{
		address:    address,
		timeout:    timeout,
		in:         in,
		out:        out,
		connection: nil,
	}
}

func (tc *telnetClient) Connect() error {
	connection, err := net.DialTimeout("tcp", tc.address, tc.timeout)
	if err != nil {
		return err
	}
	tc.connection = connection
	return nil
}

func (tc *telnetClient) Close() error {
	if tc.connection == nil {
		return ErrNoEstablishedConnection
	}
	err := tc.connection.Close()
	return err
}

func (tc *telnetClient) Send() error {
	if tc.connection == nil {
		return ErrNoEstablishedConnection
	}
	_, err := io.Copy(tc.connection, tc.in)
	return err
}

func (tc *telnetClient) Receive() error {
	_, err := io.Copy(tc.out, tc.connection)
	return err
}
