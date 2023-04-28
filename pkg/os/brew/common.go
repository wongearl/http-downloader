package brew

import (
	"fmt"

	fakeruntime "github.com/linuxsuren/go-fake-runtime"
)

const (
	// Tool is the tool name of brew
	Tool = "brew"
)

// CommonInstaller is the installer of a common brew
type CommonInstaller struct {
	Name   string
	Execer fakeruntime.Execer
}

// Available check if support current platform
func (d *CommonInstaller) Available() (ok bool) {
	if d.Execer.OS() == fakeruntime.OSDarwin || d.Execer.OS() == fakeruntime.OSLinux {
		_, err := d.Execer.LookPath(Tool)
		ok = err == nil
	}
	return
}

// Install installs the target package
func (d *CommonInstaller) Install() (err error) {
	if err = d.Execer.RunCommand(Tool, "install", d.Name); err != nil {
		return
	}
	return
}

// Uninstall uninstalls the Conntrack
func (d *CommonInstaller) Uninstall() (err error) {
	err = d.Execer.RunCommand(Tool, "remove", d.Name)
	return
}

// WaitForStart waits for the service be started
func (d *CommonInstaller) WaitForStart() (ok bool, err error) {
	ok = true
	return
}

// Start starts the Conntrack service
func (d *CommonInstaller) Start() error {
	fmt.Println("not supported yet")
	return nil
}

// Stop stops the Conntrack service
func (d *CommonInstaller) Stop() error {
	fmt.Println("not supported yet")
	return nil
}
