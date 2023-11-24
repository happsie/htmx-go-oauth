# HTMX + Go + OAuth + MySQL

A login example with stated technologies. Stores JWT as cookie after login. Currently only working with twitch login

## Trying out the demo

1. Copy `config.example.yml` to `config.yml`
2. Create an app inside [Twitch Console](https://dev.twitch.tv/console)
3. Copy `client id` and `client secret` from Twitch Console and fill it inside `config.yml` under `oauth`
4. Run `docker-compose up`
5. Navigate to `http://gohtmx:8080` in your browser

## Run configuration (Fleet)

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