package gtk

/*
#cgo pkg-config: gtk+-3.0
#include <gtk/gtk.h>
*/
import "C"
import "github.com/headcr4sh/gnome/gio"

func IsValidApplicationID(applicationID string) bool {
	cApplicationID := (*C.gchar)(C.CString(applicationID))
	return C.g_application_id_is_valid(cApplicationID) == 0
}

func NewApplication(applicationID string, flags ...gio.ApplicationFlag) *Application {
	cFlags := (C.GApplicationFlags)(gio.ApplicationFlagNone)
	for _, flag := range flags {
		cFlags |= (C.GApplicationFlags)(flag)
	}
	cApplicationID := (*C.gchar)(C.CString(applicationID))
	return &Application{
		cApp: C.gtk_application_new(cApplicationID, cFlags),
	}
}

type Application struct {
	cApp *C.GtkApplication
}

func (app *Application) NewApplicationWindow() *Widget {
	cWidget := C.gtk_application_window_new(app.cApp)
	return &Widget{
		CWidget: cWidget,
	}
}
