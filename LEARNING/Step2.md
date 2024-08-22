# Installing Helper packages:

1. We are going to use `Chi` for our server:

```sh
go get github.com/go-chi/chi
```

2. We are going to install `cors` package for handling CORS:

```sh
go get github.com/go-chi/cors
```

3. In order to write a simple server, add this in `main()` func.

```go
    router := chi.NewRouter()

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + PORT,
	}

	fmt.Printf("Server is running on port: http://localhost:%v", PORT)
	err := srv.ListenAndServe()
	if err != nil {
		panic(err)
	}
```

4. Add `cors` setting for development:

```go
	// Cores setup
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))
```
