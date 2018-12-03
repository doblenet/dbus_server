package dbus_server

/*
 * Cfr. reference at
 * https://stackoverflow.com/questions/31702465/how-to-define-a-d-bus-activated-systemd-service
 */
import (
 	"strings"
)


type DbusServiceHelper struct {
	ObjName			string
	ServiceName 	string
	AllowedGroup	string
	Command			string
}


// /usr/share/dbus-1/system-services/<objName>.service
const dbusService_str = `
[D-BUS Service]
Name=%objName%
Exec=/bin/false
User=root
SystemdService=%serviceName%.service
`

// @ /etc/dbus-1/system.d/<objName>.conf
const dbusPolicy_str = `
<!DOCTYPE busconfig PUBLIC 
	"-//freedesktop//DTD D-BUS Bus Configuration 1.0//EN"
	"http://www.freedesktop.org/standards/dbus/1.0/busconfig.dtd">
	
<busconfig>
	<policy user="root">
		<allow own="%objName%"/>
		<allow send_destination="%objName%"/>
	</policy>
	<policy group="%group%">
		<allow send_destination="%objName%"/>
	</policy>
	<policy context="default">
		<deny send_destination="%objName%"/>
	</policy>
</busconfig>
`

// /usr/lib/systemd/system/<serviceName>.service
const systemdService_str = `
[Service]
Type=dbus
BusName=%objName%
ExecStart=%program%

[Install]
Alias=%serviceName%.service
`

func (o DbusServiceHelper) DbusService() string {
	return replaceVec(dbusService_str, map[string]string{
		"objName":		o.ObjName,
		"serviceName":	o.ServiceName,
	})
}

func (o DbusServiceHelper) SystemdService() string {
	return replaceVec(dbusService_str, map[string]string{
		"objName":		o.ObjName,
		"serviceName":	o.ServiceName,
		"program":		o.Command,
	})
}

func (o DbusServiceHelper) DbusPolicy() string {
	return replaceVec(dbusService_str, map[string]string{
		"objName":	o.ObjName,
		"group":	o.AllowedGroup,
	})
}

func replaceVec(s string, vals map[string]string) string {

	var t = s
	for k,v := range vals {
		t = strings.Replace(t,("%"+k+"%"),v,-1)
	}
	return t
}
