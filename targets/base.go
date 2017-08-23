package targets

type BaseTarget struct {
	IsLog      bool
	Name       string
	TargetType string
	Layout     string
	Encode     string
}

func (t *BaseTarget) GetName() string {
	return t.Name
}

func (t *BaseTarget) GetLayout() string {
	return t.Layout
}
