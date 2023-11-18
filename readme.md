# HTMX + Go + OAuth + MySQL

A login example with stated technologies. Stores JWT as cookie after login. Currently only working with twitch login

### Run configuration (Fleet)

```json
{
    "configurations": [
        {
            "name": "run",
            "type": "go",
            "goExecPath": "C:/Program Files/Go/bin/go.exe",
            "buildParams": [
                "$PROJECT_DIR$/main.go"
            ],
            "runParams": [
                "config",
                "config.yml"
            ]
        }
    ]
}
```