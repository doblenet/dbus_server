package dbus_server

import (
	"github.com/godbus/dbus"
)

type DbusService struct {
	ObjPath			string
	Interfaces		[]string
	Implementation	interface{}
}




func (s *DbusService) Export(conn *dbus.Conn) (err error) {
	
	p := dbus.ObjectPath(s.ObjPath)
	for _,in := range s.Interfaces {
		err = conn.Export(s.Implementation, p, in)
		if nil!=err {
			return
		}
	}
	return nil
}

func (s *DbusService) Unexport(conn *dbus.Conn) error {
	
	p := dbus.ObjectPath(s.ObjPath)
	for _,in := range s.Interfaces {
		conn.Export(nil, p, in)
	}
	return nil
}
