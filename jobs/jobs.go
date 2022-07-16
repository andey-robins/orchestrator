package jobs

import "errors"

type JobHandler func()
type ReportStatus func() string

const (
	SETUP = iota
	LOADING
	RUNNING
	REPORTING
	STOPPED
)

type Job struct {
	jobConfigFname string
	jobResultFname string
	status         ReportStatus
	state          int
	handler        JobHandler
	callback       string
}

func NewJob(config, result string) *Job {
	return &Job{
		jobConfigFname: config,
		jobResultFname: result,
		state:          SETUP,
	}
}

func (j *Job) RegisterHandler(h JobHandler) error {
	if j.state != LOADING {
		return errors.New("job in an unexpected state for call to RegisterHandler()")
	}
	j.handler = h
	return nil
}

func (j *Job) RegisterReporter(r ReportStatus) error {
	if j.state != LOADING {
		return errors.New("job in an unexpected state for call to RegisterReporter()")
	}
	j.status = r
	return nil
}

func (j *Job) LoadJob() error {
	if j.state != SETUP {
		return errors.New("job in an unexpected state for call to LoadJob()")
	}

	j.nextState()
	return nil
}

func (j *Job) StartJob() error {
	if j.state != LOADING {
		return errors.New("job in an unexpected state for call to StartJob()")
	}
	go j.handler()
	j.nextState()
	return nil
}

func (j *Job) StopJob() error {
	if j.state != RUNNING {
		return errors.New("job in an unexpected state for call to StopJob()")
	}
	j.nextState()
	return nil
}

func (j *Job) Callback() error {
	if j.state != REPORTING {
		return errors.New("job in an unexpected state for call to Callback()")
	}
	j.status()
	j.nextState()
	return nil
}

func (j *Job) nextState() {
	j.state++
}
