package constraints

type Validation interface {
	IsBlank(param string) bool
}

type validation struct {
}

func DefaultValidationStruct() Validation {
	return &validation{}
}

func (v *validation) IsBlank(param string) bool {

	if param == "" {
		return true
	}

	return false
}
