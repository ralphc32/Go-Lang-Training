wait-for "${DATABASE_HOST}:${DATABASE_PORT}" -- "$@"

# tracks every changes in the main.go file
CompileDaemon --build="go build -o main main.go" --command=./main