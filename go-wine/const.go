package wine

import "errors"

var Win32 = arch{Name: "win32"}
var Win64 = arch{Name: "win64"}

var ErrFailedExecWine = errors.New("failed to execute wine command")
var ErrFailedConvToFloat = errors.New("failed to parse float")
var ErrFailedGetWineVer = errors.New("failed to get local wine version")
var ErrThisDirIsNotPrefix = errors.New("this directory is not wine prefix")
//var ErrFailedGetPath = errors.New("failed to get the path in wine")
//var ErrFailedGetSys32 = errors.New("failed to get the path of c:\\windows\\system32")
