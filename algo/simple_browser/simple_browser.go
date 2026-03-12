package simple_browser

import arraystack "github.com/practice/pkg/stack/array"

type Browser struct {
	// store the page which need forward
	forward *arraystack.Stack

	// store the page which need backward
	backward *arraystack.Stack
}

func New() *Browser {
	return &Browser{
		forward:  arraystack.New(),
		backward: arraystack.New(),
	}
}

func (b *Browser) Forward(page string) string {
	// it means to open new page if page is not empty
	if 0 != len(page) {
		// if backward stack is not empty, flush backward
		if !b.backward.IsEmpty() {
			b.backward.Flush()
		}

		// insert page into forward stack directly
		b.forward.Push(page)
		return page
	}

	// if backward stack is empty, it means can not forward anymore
	if b.backward.IsEmpty() {
		return ""
	}

	bPage := b.backward.Pop()
	b.forward.Push(bPage)

	return bPage.(string)
}

func (b *Browser) Backward() string {
	// if forward stack is empty, it means can not backward anymore
	if b.forward.IsEmpty() {
		return ""
	}

	aPage := b.forward.Pop()
	b.backward.Push(aPage)

	return aPage.(string)
}
