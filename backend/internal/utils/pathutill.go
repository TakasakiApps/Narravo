package utils

import (
	"fmt"
	"github.com/TakasakiApps/Narravo/backend/internal/global"
	"github.com/TakasakiApps/Narravo/backend/internal/types"
	"github.com/gookit/goutil/fsutil"
	"github.com/gookit/goutil/sysutil"
	"github.com/ohanakogo/exceptiongo"
	"os"
	"path/filepath"
)

// EnsureDir creates a directory named dir and any necessary intermediate directories.
// It takes in a string dir which is the initial directory to create, and subs which are
// additional subdirectories to create within the initial directory.
// If the target path already exists as a file, an exception will be thrown.
// If there is an error creating the directory, an exception will also be thrown.
// The function returns the full path of the created directory.
func EnsureDir(dir string, subs ...string) string {
	dir = filepath.Join(dir, filepath.Join(subs...))
	// Check if path exists
	if fsutil.PathExists(dir) {
		// Check if path is a file
		if fsutil.FileExists(dir) {
			// Throw exception if path is a file
			exceptiongo.QuickThrowMsg[types.FileSystemException](fmt.Sprintf(`unable to create directory "%s": target path is a file`, dir))
		}
	} else {
		// Create directory if it doesn't exist
		err := fsutil.MkDirs(os.ModeDir, dir)
		exceptiongo.QuickThrow[types.FileSystemException](err)
	}
	return dir
}

// GetAppHomeDir returns the home directory for the current user and appends the global.AppName
// with a period in front to create a new subdirectory within the home directory.
// The function returns the full path of the new subdirectory.
func GetAppHomeDir() string {
	homeDir := sysutil.HomeDir()
	// Append global.AppName with a period to create a subdirectory within the home directory
	appHomeDir := filepath.Join(homeDir, "."+global.AppName)
	return appHomeDir
}
