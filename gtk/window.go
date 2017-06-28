package gtk

// #cgo pkg-config: gtk+-3.0
// #include <gtk/gtk.h>
// GtkWindow* __GTK_WINDOW(gpointer obj) { return GTK_WINDOW(obj); }
import "C"

type WindowType int

const (
	Toplevel WindowType = C.GTK_WINDOW_TOPLEVEL
	Popup    WindowType = C.GTK_WINDOW_POPUP
)

func NewWindow(t WindowType) *Window {
	cWidget := C.gtk_window_new(C.GtkWindowType(t))
	cWindow := C.__GTK_WINDOW((C.gpointer)(cWidget))
	w := &Window{
		cWindow: cWindow,
	}
	w.CWidget = cWidget
	return w
}

type Window struct {
	Widget
	cWindow *C.GtkWindow
}

func (w *Window) SetTitle(title string) {
	cTitle := (*C.gchar)(C.CString(title))
	C.gtk_window_set_title(w.cWindow, cTitle)
}

func (w *Window) GetTitle() string {
	cTitle := (*C.char)(C.gtk_window_get_title(w.cWindow))
	return C.GoString(cTitle)
}
