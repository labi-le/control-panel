package pm

type PackageManager interface {
	Install(name string) error
	Uninstall(name string) error
	Update() error
	Custom(command string, args ...string) error
}
