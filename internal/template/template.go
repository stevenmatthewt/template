package template

type Templater struct {
	source      string
	destination string
	data        interface{}
}

// New returns a new templater, ready to output a new
// project at the specified destination
func New(source string, dest string, data interface{}) Templater {
	return Templater{
		source:      source,
		destination: dest,
		data:        data,
	}
}

// File is the action that the templater performs on files
func (t Templater) File(path string) error {
	// fmt.Printf("%s -- FILE\n", path)
	return t.templateFile(path)
}

// Folder is the action that the templater performs on folders
func (t Templater) Folder(path string) (ignore bool, err error) {
	// fmt.Printf("%s -- FOLDER\n", path)
	err = t.templateFolder(path)
	return false, err
}
