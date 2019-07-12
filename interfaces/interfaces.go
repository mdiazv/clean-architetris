package interfaces

import (
	"github.com/mdiazv/clean-architetris/domain"
)

type OutputAdapter interface {
	Render(w domain.World)
}
