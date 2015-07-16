package winprinter

func NewJob(file string, title string) (job *Job) {
	job = new(Job)
	job.file = file
	job.title = title
	return
}

type Job struct {
	file string
	title string
}

func (job *Job) Print(printer *Printer) (jobId uint32, err error) {
	jobId, err = printer.PrintPostScriptFile(job.file, job.title)
	return
}