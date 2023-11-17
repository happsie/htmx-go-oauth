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