package upgradeinspectorv4

import (
	"github.com/chef/automate/components/automate-deployment/pkg/cli"
)

type PlannedDownTimeInspection struct {
	writer *cli.Writer
}

func (pd *PlannedDownTimeInspection) ShowInfo(index *int) error {
	pd.writer.Printf("%d. You have scheduled downtime for the duration of the upgrade\n", *index)
	*index++
	return nil
}

func NewPlannedDownTimeInspection(w *cli.Writer) *PlannedDownTimeInspection {
	return &PlannedDownTimeInspection{
		writer: w,
	}
}
