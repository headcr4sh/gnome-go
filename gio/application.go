package gio

/*
#cgo pkg-config: gio-2.0
#include <gio/gio.h>
*/
import "C"

type ApplicationFlag uint

const (
	ApplicationFlagNone                           ApplicationFlag = C.G_APPLICATION_FLAGS_NONE
	ApplicationFlagApplicationIsServer            ApplicationFlag = C.G_APPLICATION_IS_SERVICE
	ApplicationFlagApplicationIsLauncher          ApplicationFlag = C.G_APPLICATION_IS_LAUNCHER
	ApplicationFlagApplicationHandlesOpen         ApplicationFlag = C.G_APPLICATION_HANDLES_OPEN
	ApplicationFlagApplicationHandlersCommandLine ApplicationFlag = C.G_APPLICATION_HANDLES_COMMAND_LINE
	ApplicationFlagApplicationSendEnvironment     ApplicationFlag = C.G_APPLICATION_SEND_ENVIRONMENT
	ApplicationFlagApplicationNonUnique           ApplicationFlag = C.G_APPLICATION_NON_UNIQUE
	ApplicationFlagApplicationCanOverrideAppID    ApplicationFlag = C.G_APPLICATION_CAN_OVERRIDE_APP_ID
)
