package functional_opts

import "os"

type FileOptions struct {
	uid         int
	gid         int
	flags       int
	contents    string
	permissions os.FileMode
}

type Option func(*FileOptions)

func UID(userID int) Option {
	return func(args *FileOptions) {
		args.uid = userID
	}
}

func GID(groupID int) Option {
	return func(args *FileOptions) {
		args.gid = groupID
	}
}

func Contents(c string) Option {
	return func(args *FileOptions) {
		args.contents = c
	}
}

func Permissions(perms os.FileMode) Option {
	return func(args *FileOptions) {
		args.permissions = perms
	}
}

func NewFile(filepath string, setters ...Option) (*os.File, error) {
	// Default FileOptions
	args := &FileOptions{
		uid:         os.Getuid(),
		gid:         os.Getgid(),
		contents:    "",
		permissions: 0666,
		flags:       os.O_CREATE | os.O_EXCL | os.O_WRONLY,
	}

	for _, setter := range setters {
		setter(args)
	}

	os.Remove(filepath)
	f, err := os.OpenFile(filepath, args.flags, args.permissions)
	if err != nil {
		return f, err
	} else {
		defer f.Close()
	}

	if _, err := f.WriteString(args.contents); err != nil {
		return f, err
	}

	return f, nil//f.Chown(args.uid, args.gid)
}