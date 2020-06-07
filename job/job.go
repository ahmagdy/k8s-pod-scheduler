package job

// Job a representation of a running job
type Job struct {
	Name    string
	Version float32
	Kind    string
	Cron    string
	Specs   *Specs
}
type Specs struct {
	Image    string
	Args     string
	Commands []string
}
