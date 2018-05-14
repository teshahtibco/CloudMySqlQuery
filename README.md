---
title: CloudMySqlQuery Activity
---

# CloudMySqlQuery Activity
This activity allows you to Connect to Google cloud MySQL instance and perform Query/Update operations on the database

## Installation
### Flogo CLI
```bash
flogo install github.com/teshahtibco/CloudMySqlQuery
```

## Schema
Inputs and Outputs:

```json
{   
  "inputs":[
    {
      "name": "hostname",
      "type": "string"
    },
    {
      "name": "port",
      "type": "string"
    },
    {
      "name": "username",
      "type": "string"
    },
    {
      "name": "password",
      "type": "string"
    },
    {
      "name": "instance",
      "type": "string"
    },
	{
      "name": "method",
      "type": "string",
	  "allowed": [
        "QUERY",
        "OTHER"
      ],
	   "value": "QUERY"
    },
	{
      "name": "query",
      "type": "string"
    }
  ],
  "outputs": [
    {
      "name": "result",
      "type": "any"
    }
  ]
 }
```
## Settings
| Setting        | Required | Description |
|:---------------|:---------|:------------|
| hostname      | True     | The Hostname/IP Address of Google Cloud MySql instance |         
| port       | True     | The Port number of Google Cloud MySql instance
| username      | True     | The username to connect to the database
| password         | True		| The username to connect to the database
| instance          | True     | The Database instance to be connected |
| method      | True     | The Operation to be performed. QUERY/OTHER |   
| query      | True     | Database Query to be executed |      

```