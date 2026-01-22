package options

import (
	"fmt"
	"net"
	"strconv"
	"time"
)

type MySQLOptions struct {
	Addr	string `json:"add,omitempty" mapstructure:"addr"`
	Username	string	`json:"username,omitempty" mapstructure:"username"`
	Password	string	`json:"-", mapstructure:"password"`
	Database	string	`json:"database", mapstructure:"database"`
	MaxIdleConnections	int	`json:"max-idle-connections,omitempty", mapstructure:"max-idle-connections,omitempty"`
	MaxOpenConnections	int	`json:"max-open-connections,omitempty", mapstructure:"max-open-connections,omitempty"`
	MaxConnectionsLifeTime	time.Duration	`json:"max-connection-life-time,omitempty" mapstructure:"max-connection-life-time,omitempty"`
}

func NewMySQLOptions() *MySQLOptions {
	return &MySQLOptions{
		Addr: "127.0.0.1:3306",
		Username: "onex",
		Password: "onex(#)666#",
		Database: "onex",
		MaxIdleConnections: 100,
		MaxOpenConnections: 100,
		MaxConnectionsLifeTime: time.Duration(10) * time.Second,
	}
}


func (o *MySQLOptions) Validate() error {
	if o.Addr == "" {
		return fmt.Errorf("MySQL Server address cannot be empty")
	}

	host, portStr, err := net.SplitHostPort(o.Addr)
	if err != nil {
		return fmt.Errorf("Invalid MySQL address format '%s': %w", o.Addr, err)
	}

	port, err := strconv.Atoi(portStr)
	if err != nil || port < 1 || port > 65535 {
		return fmt.Errorf("Invalid MySQL port: %s", portStr)
	}

	if host == "" {
		return fmt.Errorf("MySQL hostname cannot be empty")
	}

	if o.Username == "" {
		return fmt.Errorf("MySQL username cannot be empty")
	}

	if o.Password == "" {
		return fmt.Errorf("MySQL password cannot be empty")
	}

	if o.Database == "" {
		return fmt.Errorf("MySQL database cannot be empty")
	}

	if o.MaxIdleConnections <= 0 {
		return fmt.Errorf("MySQL max idle connections must be greater than 0")
	}

	if o.MaxOpenConnections <= 0 {
		return fmt.Errorf("MySQL max open connections must be greater than 0")
	}

	if o.MaxIdleConnections > o.MaxOpenConnections {
		return  fmt.Errorf("MySQL max idle connections cannot be greater than max open connections")
	}

	if o.MaxConnectionsLifeTime <= 0 {
		return fmt.Errorf("MySQL max connections life time cannot be greater than 0")
	}

	return nil
}