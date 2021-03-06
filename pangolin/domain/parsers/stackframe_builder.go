package parsers

import "errors"

type stackFrameBuilder struct {
	isPush bool
	isPop  bool
	index  Index
	skip   Skip
}

func createStackFrameBuilder() StackFrameBuilder {
	out := stackFrameBuilder{
		isPush: false,
		isPop:  false,
		index:  nil,
		skip:   nil,
	}

	return &out
}

// Create initializes the builder
func (app *stackFrameBuilder) Create() StackFrameBuilder {
	return createStackFrameBuilder()
}

// IsPush flags the builder as push
func (app *stackFrameBuilder) IsPush() StackFrameBuilder {
	app.isPush = true
	return app
}

// IsPop flags the builder as pop
func (app *stackFrameBuilder) IsPop() StackFrameBuilder {
	app.isPop = true
	return app
}

// WithIndex adds an index to the builder
func (app *stackFrameBuilder) WithIndex(index Index) StackFrameBuilder {
	app.index = index
	return app
}

// WithSkip adds a skip to the builder
func (app *stackFrameBuilder) WithSkip(skip Skip) StackFrameBuilder {
	app.skip = skip
	return app
}

// Now builds a new StackFrame instance
func (app *stackFrameBuilder) Now() (StackFrame, error) {
	if app.isPush {
		return createStackFrameWithPush(), nil
	}

	if app.isPop {
		return createStackFrameWithPop(), nil
	}

	if app.index != nil {
		return createStackFrameWithIndex(app.index), nil
	}

	if app.skip != nil {
		return createStackFrameWithSkip(app.skip), nil
	}

	return nil, errors.New("the StackFrame is invalid")
}
