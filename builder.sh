

# Windows
env GOOS=windows GOARCH=amd64 go build -o ./runner/windows/glrd_windows_arm64 .

# Linux
env GOOS=linux GOARCH=amd64 go build -o ./runner/linux/glrd_linux_arm64 .

# MacOS
env GOOS=darwin GOARCH=amd64 go build -o ./runner/macos/glrd_macos_arm64 .