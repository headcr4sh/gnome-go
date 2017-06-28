package gtk

/*
#cgo pkg-config: gtk+-3.0
#include <gtk/gtk.h>
*/
import "C"

type Widget struct {
	CWidget *C.GtkWidget
}

func (w *Widget) Destroy() {
	C.gtk_widget_destroy(w.CWidget)
}
