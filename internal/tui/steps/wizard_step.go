package steps

type WizardStep interface {
	Filled() bool
}
