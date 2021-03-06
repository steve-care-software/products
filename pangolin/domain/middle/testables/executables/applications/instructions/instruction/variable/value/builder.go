package value

import (
	"errors"

	"github.com/deepvalue-network/software/pangolin/domain/middle/testables/executables/applications/instructions/instruction/variable/value/computable"
)

type builder struct {
	isStackFrame bool
	comp         computable.Value
	variable     string
}

func createBuilder() Builder {
	out := builder{
		isStackFrame: false,
		comp:         nil,
		variable:     "",
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// IsStackFrame flags the builder as a stackframe
func (app *builder) IsStackFrame() Builder {
	app.isStackFrame = true
	return app
}

// WithComputable adds a computable to the builder
func (app *builder) WithComputable(computable computable.Value) Builder {
	app.comp = computable
	return app
}

// WithVariable adds a variable to the builder
func (app *builder) WithVariable(variable string) Builder {
	app.variable = variable
	return app
}

// Now builds a new Value instance
func (app *builder) Now() (Value, error) {
	if app.isStackFrame {
		return createValueWithStackFrame(), nil
	}

	if app.comp != nil {
		return createValueWithComputable(app.comp), nil
	}

	if app.variable != "" {
		return createValueWithVariable(app.variable), nil
	}

	return nil, errors.New("the Value is invalid")
}
