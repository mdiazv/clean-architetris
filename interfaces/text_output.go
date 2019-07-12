package interfaces

import (
	"fmt"
	"github.com/mdiazv/clean-architetris/domain"
)

type textOutputAdapter struct{}

func MakeTextOutputAdapter() OutputAdapter {
	return &textOutputAdapter{}
}

func (a *textOutputAdapter) Render(w domain.World) {
	taken := w.Taken()
	for _, p := range taken {
		fmt.Printf("%+v\n", p)
	}
}
