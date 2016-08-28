# Return Path [![Build Status](https://travis-ci.org/wakawaka54/ReturnPath.svg?branch=master)](https://travis-ci.org/wakawaka54/ReturnPath)

####[Design Document](DESIGN.md)

## API Endpoints

### Sentences

#### GET `/api/sentences?limit=20&offset=0`
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Returns a paginated list of sentences in database.

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;**Reponse:**

**Headers**
`X-Total-Count` - Total number of sentences

```json
[
  {
    "id": 0,
    "sentence": "why is this great",
    "tags": [
      "why",
      "great"
    ]
  },
  {
    "id": 1,
    "sentence": "why is this great",
    "tags": [
      "why",
      "great"
    ]
  }
]
```

#### POST `/api/sentences`
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Creates a sentence object in the database and assigns cooresponding tags.

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;**Example Request:**

```json
{
  "sentence":"this is really amazing if you think about it"
}
```

#### DELETE `/api/sentences/{id}`
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Deletes a sentence object in the database

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;**Example Url:** `/api/sentences/26723478dede714`

#### GET `/api/sentences/statistics`
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Retrieves the top 15 tags and their frequency counts on the entire current dataset. 

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;**Example Response:**
```json
[
  {
    "tag": "great",
    "count": 30
  },
  {
    "tag": "why",
    "count": 30
  }
]
```
