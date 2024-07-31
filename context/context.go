package context

import (
	"context"
	"fmt"
	"net/http"
)

// type Store interface {
// 	Fetch() string
// 	Cancel()
// }

type Store interface {
	Fetch(ctx context.Context) (string, error)
}

// func Server(store Store) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		// store.Cancel()
// 		// fmt.Fprint(w, store.Fetch())
// 		ctx := r.Context()

// 		data := make(chan string, 1)

// 		go func() {
// 			data <- store.Fetch()
// 		}()

// 		select {
// 		case d := <-data:
// 			fmt.Fprint(w, d)
// 		case <-ctx.Done():
// 			store.Cancel()
// 		}

// 	}
// }

func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := store.Fetch(r.Context())

		if err != nil {
			return //todo: log error
		}
		fmt.Fprint(w, data)
	}
}
