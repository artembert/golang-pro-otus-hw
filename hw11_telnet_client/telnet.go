package main

import (
	"errors"
	"io"
	"net"
	"time"
)

var (
	ErrConnection              = errors.New("connection error")
	ErrNoEstablishedConnection = errors.New("no established connection")
	ErrCloseConnections        = errors.New("failed to close connection")
	ErrSend                    = errors.New("failed to send message")
	ErrReceive                 = errors.New("failed to receive message")
)

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
		return errors.Join(err, ErrConnection)
	}

	tc.connection = connection

	return nil
}

func (tc *telnetClient) Close() error {
	if tc.connection == nil {
		return ErrNoEstablishedConnection
	}

	err := tc.connection.Close()
	if err != nil {
		return errors.Join(err, ErrCloseConnections)
	}

	return nil
}

func (tc *telnetClient) Send() error {
	if tc.connection == nil {
		return ErrNoEstablishedConnection
	}

	_, err := io.Copy(tc.connection, tc.in)
	if err != nil {
		return errors.Join(err, ErrSend)
	}

	return nil
}

func (tc *telnetClient) Receive() error {
	_, err := io.Copy(tc.out, tc.connection)
	if err != nil {
		return errors.Join(err, ErrReceive)
	}

	return nil
}
