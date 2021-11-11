# Lift Tracker API

## Installation

You need [Go](https://golang.org/doc/install) installed and set your `Go` workplace first.
For editor, I recommend [goland](https://www.jetbrains.com/go/download/) or [vscode](https://code.visualstudio.com/download).

```bash
# Create a .env file
cp .env.template .env
# Run this project
go run main.go
```

In Vscode, simply create a `launch.json` file in a `.vscode` folder, paste the content below, and run debugger.

```json
{
  "version": "0.2.0",
  "configurations": [
    {
      "type": "go",
      "request": "launch",
      "name": "Launch Program",
      "skipFiles": [
        "<node_internals>/**"
      ],
      "program": "${workspaceFolder}/main.go",
      "envFile": "${workspaceFolder}/.env"
    }
  ]
}
```