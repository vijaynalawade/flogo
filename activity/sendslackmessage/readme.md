![gofmt status](https://img.shields.io/badge/gofmt-compliant-green.svg?style=flat-square) ![golint status](https://img.shields.io/badge/golint-compliant-green.svg?style=flat-square) ![automated test coverage](https://img.shields.io/badge/test%20coverage-1%20testcase-orange.svg?style=flat-square)

# Send Slack Message
This activity sends a message to a Slack channel.

## Installation

```bash
flogo add activity github.com/vijaynalawade/flogo/activity/sendslackmessage
```

## Schema
Inputs and Outputs:

```json
{
"inputs":[
    {
      "name": "Webhook",
      "type": "string",
      "required": true
    },
    {
      "name": "Channel",
      "type": "string",
      "required": false
    },
    {
      "name": "Message",
      "type": "string",
      "required": true
    },
    {
      "name": "Username",
      "type": "string",
      "required": false
    },
    {
      "name": "Iconemoji",
      "type": "string",
      "required": false
    }
  ],
  "outputs": [
    {
      "name": "result",
      "type": "string"
    }
  ]
}
```

## Settings
| Setting     | Description    |
|:------------|:---------------|
| Webhook     | The Webhook URL to connect to Slack |         
| Channel     | The channel to send the message to (e.g. `#channel`)   |
| Message     | The message to send |
| Username    | The username to use |
| Iconemoji   | The icon emoji to send |

## Configuration Examples
The below configuration would connect to a Slack and send a message saying `Hello World`
```json
      {
        "id": 2,
        "name": "Send a message to a Slack channel",
        "type": 1,
        "activityType": "sendSlackMessage",
        "attributes": [
          {
            "name": "Webhook",
            "value": "https://hooks.slack.com/services/xxxxxx/xxxxxx",
            "type": "string"
          },
          {
            "name": "Channel",
            "value": "#channel",
            "type": "string"
          },
          {
            "name": "Message",
            "value": "Hello World",
            "type": "string"
          },
          {
            "name": "Username",
            "value": "user",
            "type": "string"
          },
          {
            "name": "Iconemoji",
            "value": ":ghost:",
            "type": "string"
          }
        ]
      }
```

## Contributors
[Vijay Nalawade](https://github.com/vijaynalawade)
