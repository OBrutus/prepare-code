go build .

# alias backup clone
go build -o pc ./main.go

if [[ $? -ne 0 ]]; then
    echo "Build failed!"
    echo "Maybe go is not installed?"
    exit 1
fi