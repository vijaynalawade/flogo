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
      "name": "jsInput",
      "type": "any"
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
| jsInput   | Input to your java script code. It can be primitive or JSON object type. |
| javascript  | Javascript code|



## Configuration Examples
### Simple
Configure a task that adds two numbers in java script code:

```json
{
  "id": 3,
  "type": 1,
  "ref": "github.com/vijaynalawade/flogo/activity/js",
  "name": "JavaScript Activity",
  "attributes": [
    { "name": "jsInput", "value": "{\"number1\":2,\"number2\":3}" },
    { "name": "javascript", "value": "jsInput.number1 + jsInput.number2;" };
  ]
}
```
