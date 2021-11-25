# Atlan Collect API

## Premise

There are various possible layouts for serving data through the API. Since the plug-in needs to have a
definitive control over the format and content of the data it needs, going for a GraphQL API seems
to be a viable option.

However, for the scope of this task, I'll go for setting up a REST API instead; for the sake
of simplicity and familiarity among developers.

## API Endpoints

**Note:** Every API endpoint requires the request to have `collect-auth-token` in its header, which is 
issued by Atlan.

This token would contain a payload having the userId,
name, email and all the fields required to identify
the individual.

### `/getForm`

```url
    /api/v1/getForm?form_id=<FORM_ID>&version=<VERSION>
```

- Parameters
    - `form_id` - The form ID of the form requested
    - `version` - The version of the form requested

- Request Body
    - No request body

- Response

```json
{
    "status": "Success",
    "message": "Data fetched successfully",
    "form_data": {
        "name": "Form Name",
        "description": "Description can be null",
        "createdAt": 1637758818834,
        "is_published": true,
        "user_id": 123123712,
        "version": "v1",
        "question_bank": [
            {
                "question_id": 12312334,
                "question_text": "How is the question?",
                "type": "Small Text",
                "meta_data": {...}
            },
            ...
        ]
    }
}
```

### `/getQuestion`

```url
    /api/v1/getQuestion?question_id=<QUESTION_ID>
```

- Parameters
    - `question_id` - The question ID of the question requested

- Request Body
    - No request body

- Response

```json
{
    "status": "Success",
    "message": "Data fetched successfully",
    "question_data": {
        "question_bank_id": 12312334,
        "form_id": 394875,
        "user_id": 2349872,
        "isPublished": true,
        "question_text": "How is the question?",
        "type": "Small Text",
        "meta_data": {...}
    }
}
```

### `/getResponse`

```url
    /api/v1/getResponse?form_id=<FORM_ID>&version=<VERSION>
```

- Parameters
    - `form_id` - The form ID of the form requested
    - `version` - The version of the form requested

- Request Body
    - No request body

- Response
```json
{
    "status": "Success",
    "message": "Data fetched successfully",
    "response_data": [
        {
            "response_bank_id": 123712836,
            "user_id": 1234234,
            "is_published": true,
            "responses": [
                {
                    "response_id": 1387987,
                    "question_id": 039485,
                    "type": "text",
                    "response": "This is a sample response",
                    "meta_data": {...}
                },
                ...
            ]
        },
        ...
    ]
}
```

### `/getResponseBank`

```url
    /api/v1/getResponse?response_bank_id=<RESPONSE_BANK_ID>
```

- Parameters
    - `response_bank_id` - The ID of the response bank requested

- Request Body
    - No request body

- Response
```json
{
    "status": "Success",
    "message": "Data fetched successfully",
    "response_bank_data": {
        "response_bank_id": 123712836,
        "user_id": 1234234,
        "is_published": true,
        "responses": [
            {
                "response_id": 1387987,
                "question_id": 039485,
                "type": "text",
                "response": "This is a sample response",
                "meta_data": {...}
            },
            ...
        ],
    }
}
```