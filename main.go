package main

import (
	"github.com/godbus/dbus"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		log.Fatal("Not enough arguments.\n" +
			"Usage: dbusSignal <interface> <signal name> <path> [data string]\n" +
			"Example: dbusSignal de.toolboxbodensee.plane newSerial /serial TestString123")
	}

	iface := os.Args[1]
	signalName := os.Args[2]
	path := os.Args[3]
	//dataType := "string"
	data := concatStrings(os.Args[4:]...)

	conn, err := dbus.SessionBus()
	if err != nil {
		log.Fatal(err)
	}

	_, err = conn.RequestName(iface, dbus.NameFlagDoNotQueue)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.ReleaseName(iface)

	err = conn.Emit(dbus.ObjectPath(path), iface+"."+signalName, data)
	if err != nil {
		log.Fatal(err)
	}
}

func concatStrings(strings ...string) string {
	r := ""
	for _, s := range strings {
		r += s
	}
	return r
}
