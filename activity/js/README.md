# activity-js
This activity executes simple embedded java script in your Flogo application.


## Installation

```bash
flogo install github.com/vijaynalawade/flogo/activity/js
```

## Schema
Inputs and Outputs:

```json
"inputs":[
    {
      "name": "inputVars",
      "type": "object"
    },
    {
      "name": "javascript",
      "type": "string",
      "required": true
    }
  ],
  "outputs": [
    {
      "name": "output",
      "type": "any"
    }
  ]
```
## Settings
| Setting   | Description    |
|:----------|:---------------|
| inputVars   | Set of dynamic variables that you want to pass to your java script code |
| javascript  | Javascript code|



## Configuration Examples
### Simple
Configure a task to log a message 'test message':

```json
{
  "id": 3,
  "type": 1,
  "ref": "github.com/vijaynalawade/flogo/activity/js",
  "name": "JavaScript Activity",
  "attributes": [
    { "name": "inputVars", "value": "{\"number1\":2,\"number2\":30}" },
    { "name": "javascript", "value": "number1 + number2;" }
  ]
}
```
