# Learning GO: Building an API with GIN, Swagger GoORM and SQLITE

This API was built by watching the tutorial [Desvendando Go: Aprenda a Construir APIS Poderosas com Gin, Swagger, GoORM e SQLite | Gopportunities](https://youtu.be/wyEYpX5U4Vg?si=7YSoAM7oKCYFjULu)

```bash
go mod init github.com/vickttor/gooportunities
```

- Compiled Language: Go is a compiled language, that means we compile our source code and get an executable file.
- When creating a new project in Go, don't use `src` as the name of the main files directory, since the project at all is the source code. Go projects doesn't follow the convention of other technologies such as NodeJS.
- When we create a new directory we consider it as a sub package.
- It's a convention to give the same name of the file into a folder. Example: `router/router.go`
- When we declare a function with Captial letter, it becomes exportable
