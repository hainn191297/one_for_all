package smtp

import (
	"crypto/tls"
	"errors"
	"fmt"

	"github.com/zeromicro/go-zero/core/logx"
)

type SmtpClient struct {
	Client    *Client
	Host      string
	Port      string
	UserName  string
	a         Auth
	KeepAlive bool
}

func Connect(host string, port string, a Auth, keepAlive bool) (*SmtpClient, error) {
	var err error
	var c *Client

	c, err = smtpConnect(host+":"+port, a)
	if err != nil {
		return nil, err
	}

	return &SmtpClient{
		Client:    c,
		a:         a,
		KeepAlive: keepAlive,
	}, nil
}

func (sm *SmtpClient) ResetConn() error {
	return sm.Client.Reset()
}

func (sm *SmtpClient) CloseConn() error {
	err := sm.Client.Close()
	if err != nil {
		logx.Error(err)
	}
	return sm.Client.Close()
}

func (sm *SmtpClient) SendQuit() error {
	err := sm.Client.Quit()
	if err != nil {
		logx.Error(err)
	}
	return sm.Client.Quit()
}

func (sm *SmtpClient) SendNood() error {
	return sm.Client.Noop()
}

func (sm *SmtpClient) SendMail(from string, recipients []string, msg string) error {
	if len(recipients) == 0 {
		return errors.New("no recipients")
	}

	if from == "" {
		return errors.New("no sender")
	}

	err := sm.Client.Noop()
	if err != nil {
		logx.Error(err)
		sm.SendQuit()
		sm.CloseConn()
		sm.Client, err = smtpConnect(sm.Host+":"+sm.Port, sm.a)
		if err != nil {
			logx.Error(err)
		}
	}
	err = send(from, recipients, msg, sm)
	if err != nil {
		return errors.New("send err" + err.Error())
	}
	return nil
}

func send(from string, to []string, msg string, c *SmtpClient) error {
	if c != nil {
		if c.Client != nil {
			err := sendProcess(from, to, msg, c)
			if err != nil {
				logx.Error(err)
				checkKeepAlive(c)
			}

		}
	}
	return nil
}

func smtpConnect(addr string, a Auth) (*Client, error) {
	c, err := Dial(addr)
	if err != nil {
		return nil, err
	}
	if err = c.hello(); err != nil {
		c.Close()
		return nil, err
	}

	if ok, _ := c.Extension("STARTTLS"); ok {
		config := &tls.Config{
			ServerName:         c.serverName,
			InsecureSkipVerify: true,
		}
		if testHookStartTLS != nil {
			testHookStartTLS(config)
		}
		if err = c.StartTLS(config); err != nil {
			return nil, err
		}
	}
	if a != nil && c.ext != nil {
		if _, ok := c.ext["AUTH"]; !ok {
			return nil, errors.New("smtp: server doesn't support AUTH")
		}
		if err = c.Auth(a); err != nil {
			c.Close()
			return nil, err
		}
	}

	return c, nil
}

func sendProcess(from string, to []string, msg string, c *SmtpClient) error {

	if err := validateLine(from); err != nil {
		return err
	}

	for _, recp := range to {
		if err := validateLine(recp); err != nil {
			return err
		}
	}

	if err := c.Client.Mail(from); err != nil {
		return err
	}

	for _, address := range to {
		if err := c.Client.Rcpt(address); err != nil {
			return err
		}
	}

	w, err := c.Client.Data()
	if err != nil {
		return err
	}

	_, err = fmt.Fprint(w, msg)
	if err != nil {
		return err
	}

	err = w.Close()
	if err != nil {
		return err
	}

	return nil
}

func checkKeepAlive(c *SmtpClient) {
	if c.KeepAlive {
		c.ResetConn()
	} else {
		c.SendQuit()
		c.CloseConn()
	}
}
