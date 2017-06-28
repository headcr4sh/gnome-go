package glib

/*
#cgo pkg-config: glib-2.0
#include <glib-2.0/glib.h>

// Macro wrappers
int __G_IS_DIR_SEPARATOR(gchar sep) { return G_IS_DIR_SEPARATOR(sep); }
*/
import "C"

// IsDirSeparator checks whether a character is a directory separator.
// It returns true for '/' on UNIX machines and for '\' or '/' under Windows.
func IsDirSeparator(sep rune) bool {
	return C.__G_IS_DIR_SEPARATOR((C.gchar)(sep)) == C.TRUE
}
