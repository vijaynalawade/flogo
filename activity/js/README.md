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
          "name": "outputVars",
          "type": "string"
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
| inputVars   | Set of variables defined in JSON object that you want to input to your java script code |
| outputVars   | Comma separated list of variables that you want to output from your java script code |
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
    { "name": "inputVars", "value": "{\"number1\":2,\"number2\":30}" },
    { "name": "outputVars", "value": "sum" }
    { "name": "javascript", "value": "sum = number1 + number2;" };
  ]
}
```
