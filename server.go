package dbus_server

import (
	"github.com/godbus/dbus"
	"github.com/pkg/errors"
)


// DbusServer models a DbusServer (object)
type DbusServer struct {
	ObjName	string
}




func (s *DbusServer) RegisterName(conn *dbus.Conn) error {
	
	reply, err := conn.RequestName(s.ObjName, dbus.NameFlagDoNotQueue)
	if nil!=err {
		return errors.Wrap(err,"DbusServer#Register")
	}
	
	if reply != dbus.RequestNameReplyPrimaryOwner {
		return errors.New("Could not claim ObjName: name already taken!")
	}
	return nil
}

func (s *DbusServer) Unregister(conn *dbus.Conn) error {
	
	r,err := conn.ReleaseName(s.ObjName)
	if err!=nil || r!=dbus.ReleaseNameReplyReleased {
		return errors.New("Could not release bus name")
	}
	return nil
}

func (s *DbusServer) BusName() string {
	return s.ObjName
}

func (s *DbusServer) String() string {
	return ("DbusServer["+s.ObjName+"]")
}
