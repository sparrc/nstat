package nstat

import (
	"bytes"
	"io/ioutil"
	"os"
	"strconv"
)

var zeroByte = []byte("0")

const (
	NET_NETSTAT = "/net/netstat"
	NET_SNMP    = "/net/snmp"
	NET_SNMP6   = "/net/snmp6"
)

func proc(path string) string {
	if root := os.Getenv("PROC_ROOT"); root != "" {
		return root + path
	}
	return "/proc" + path
}

// Counters is an object
type Counters struct {
	DumpZeros bool

	counters map[string]int64
}

func (c *Counters) Get() map[string]int64 {
	if c.counters == nil {
		c.counters = make(map[string]int64)
	}
	c.readProcNetstat()
	c.readProcNetSnmp()
	c.readProcNetSnmp6()
	return c.counters
}

// Read /proc/net/netstat and apply counters to given map
func (c *Counters) readProcNetstat() {
	fd, err := os.Open(proc(NET_NETSTAT))
	if err != nil {
		return
	}
	defer fd.Close()
	if netstatFile, err := ioutil.ReadAll(fd); err == nil {
		c.parseUglyFile(netstatFile)
	}
}

// Read /proc/net/snmp and apply counters to given map
func (c *Counters) readProcNetSnmp() {
	fd, err := os.Open(proc(NET_SNMP))
	if err != nil {
		return
	}
	defer fd.Close()
	if snmpFile, err := ioutil.ReadAll(fd); err == nil {
		c.parseUglyFile(snmpFile)
	}
}

// Read /proc/net/snmp6 and apply counters to given map
func (c *Counters) readProcNetSnmp6() {
	fd, err := os.Open(proc(NET_SNMP6))
	if err != nil {
		return
	}
	defer fd.Close()
	if snmp6File, err := ioutil.ReadAll(fd); err == nil {
		c.parseNiceFile(snmp6File)
	}
}

// parse an ugly counter file (netstat & snmp)
func (c *Counters) parseUglyFile(file []byte) {
	lines := bytes.Split(file, []byte("\n"))
	var value int64
	var err error
	for i := 0; i < len(lines); i = i + 2 {
		if len(lines[i]) == 0 {
			continue
		}
		headers := bytes.Fields(lines[i])
		prefix := bytes.TrimSuffix(headers[0], []byte(":"))
		metrics := bytes.Fields(lines[i+1])
		for j := 1; j < len(headers); j++ {
			// counter is zero
			if bytes.Equal(metrics[j], zeroByte) {
				if !c.DumpZeros {
					continue
				} else {
					c.counters[string(append(prefix, headers[j]...))] = 0
					continue
				}
			}
			// the counter is not zero, so parse it.
			value, err = strconv.ParseInt(string(metrics[j]), 10, 64)
			if err == nil {
				c.counters[string(append(prefix, headers[j]...))] = value
			}
		}
	}
}

// parse a nice counter file (snmp6)
func (c *Counters) parseNiceFile(file []byte) {
	fields := bytes.Fields(file)
	var value int64
	var err error
	for i := 0; i < len(fields); i = i + 2 {
		// counter is zero
		if bytes.Equal(fields[i+1], zeroByte) {
			if !c.DumpZeros {
				continue
			} else {
				c.counters[string(fields[i])] = 0
				continue
			}
		}
		// the counter is not zero, so parse it.
		value, err = strconv.ParseInt(string(fields[i+1]), 10, 64)
		if err == nil {
			c.counters[string(fields[i])] = value
		}
	}
}
