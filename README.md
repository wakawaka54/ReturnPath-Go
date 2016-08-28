# Return Path with Go [![Build Status](https://travis-ci.org/wakawaka54/ReturnPath-Go.svg?branch=master)](https://travis-ci.org/wakawaka54/ReturnPath-Go)

![Image of Yaktocat](rpfrontend/static/images/go.png)

####[Design Document](DESIGN.md)

## Building this Repo

**ReturnPath-Go** is made up of two Go packages, *rpfrontend* and *rpapi*. In order to launch both of these packages, I have facilated a **docker-compose** file which should make it very easy to launch both of these.

1. Make sure you have `docker-compose` and `docker` installed on your machine
2. Clone the repo using `git clone https://github.com/wakawaka54/ReturnPath-Go.git`
3. Enter into the repo `cd ReturnPath-Go`
4. Run `docker-compose build`
  * This should build both images, depending on how you have `docker-compose` configured, you may need to run this as sudo
5. Run `docker-compose up`
  * This should start up both images
6. The frontend should be mapped to `http://localhost:1400`
7. The backend should be mapped to `http://localhost:1337`

## API Endpoints

### Sentences

#### GET `/api/sentences`
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Returns a paginated list of sentences in database.

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Supported Query String Parameters:

* `limit` - limits response to certain number of sentences (default is 20)
* `offset` - offsets the results by a certain number of sentences (default is 0)
* `id` - get sentence by id
* `sentence` - get sentences that contains sentence
* `tags` - get sentences that contain tags

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;**Reponse:**

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;**Headers**

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;`X-Total-Count` - Total number of sentences

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
