package interfaces

import (
	"github.com/mdiazv/clean-architetris/domain"

	"fmt"
	"net/http"
)

type httpOutputAdapter struct {
	w http.ResponseWriter
	r *http.Request
}

func MakeHttpOutputAdapter(w http.ResponseWriter, r *http.Request) OutputAdapter {
	return &httpOutputAdapter{w: w, r: r}
}

func (a *httpOutputAdapter) Render(w domain.World) {
	taken := w.Taken()
	for _, p := range taken {
		fmt.Fprintf(a.w, "%+v\n", p)
	}
}
