package glib

/*
#cgo pkg-config: glib-2.0
#include <glib-2.0/glib.h>
*/
import "C"

var MajorVersion uint
var MinorVersion uint
var MicroVersion uint

func init() {
	MajorVersion = (uint)(C.glib_major_version)
	MinorVersion = (uint)(C.glib_minor_version)
	MicroVersion = (uint)(C.glib_micro_version)
}

// CheckVersion checks that the GLib library in use is compatible with the given version.
func CheckVersion(requiredMajor, requiredMinor, requiredMicro uint) (ok bool, msg string) {
	cmsg := C.glib_check_version(C.guint(requiredMajor), C.guint(requiredMinor), C.guint(requiredMicro))
	if ok = cmsg == nil; !ok {
		msg = C.GoString((*C.char)(cmsg))
	}
	return
}
