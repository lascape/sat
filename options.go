package sat

type Option func(*Options)

type Options struct {
	Path string `json:"path"`
}

func SetPath(path string) Option {
	return func(args *Options) {
		args.Path = path
	}
}
