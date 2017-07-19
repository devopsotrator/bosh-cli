package cmd

import (
	"github.com/cloudfoundry/bosh-cli/ui/table"
)

type ColumnOpt struct {
	table.Header
}

func (a *ColumnOpt) UnmarshalFlag(arg string) error {
	a.Header = table.NewHeader(arg)
	return nil
}
