package cmd_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"errors"
	. "github.com/cloudfoundry/bosh-cli/cmd"
	boshdir "github.com/cloudfoundry/bosh-cli/director"
	fakedir "github.com/cloudfoundry/bosh-cli/director/directorfakes"
	fakeui "github.com/cloudfoundry/bosh-cli/ui/fakes"
	boshtbl "github.com/cloudfoundry/bosh-cli/ui/table"
)

var _ = Describe("VarsCmd", func() {
	var (
		ui         *fakeui.FakeUI
		deployment *fakedir.FakeDeployment
		command    VarsCmd
	)

	BeforeEach(func() {
		ui = &fakeui.FakeUI{}
		deployment = &fakedir.FakeDeployment{}
		command = NewVarsCmd(ui, deployment)
	})

	Describe("Run", func() {
		act := func() error { return command.Run() }

		It("lists placeholder variables (ID and Name)", func() {
			varsResults := []boshdir.VarsResult{
				{ID: "1", Name: "foo-1"},
				{ID: "2", Name: "foo-2"},
				{ID: "3", Name: "foo-3"},
				{ID: "4", Name: "foo-4"},
			}
			deployment.VarsReturns(varsResults, nil)

			err := act()
			Expect(err).ToNot(HaveOccurred())

			Expect(ui.Table).To(Equal(boshtbl.Table{

				Header: []string{"ID", "Name"},

				SortBy: []boshtbl.ColumnSort{
					{Column: 0, Asc: true},
					{Column: 1},
				},

				Rows: [][]boshtbl.Value{
					{
						boshtbl.NewValueString("1"),
						boshtbl.NewValueString("foo-1"),
					},
					{
						boshtbl.NewValueString("2"),
						boshtbl.NewValueString("foo-2"),
					},
					{
						boshtbl.NewValueString("3"),
						boshtbl.NewValueString("foo-3"),
					},
					{
						boshtbl.NewValueString("4"),
						boshtbl.NewValueString("foo-4"),
					},
				},
			}))
		})

		It("errors if getting placeholder variables fails", func() {
			deployment.VarsReturns(nil, errors.New("fake-err"))

			err := act()
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("fake-err"))
		})
	})
})
