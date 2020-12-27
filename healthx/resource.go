package health

import "context"

const (
	UP   = "UP"
	DOWN = "DOWN"
)

type Resource struct {
	Name    string
	Checker func(ctx context.Context) error
}

func (r *Resource) Report(ctx context.Context) Report {
	var (
		status = UP
		errMsg = ""
	)

	if err := r.Checker(ctx); err != nil {
		status = DOWN
		errMsg = err.Error()
	}

	return Report{
		Name:   r.Name,
		Status: status,
		Error:  errMsg,
	}
}
