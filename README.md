# Return Path with Go [![Build Status](https://travis-ci.org/wakawaka54/ReturnPath-Go.svg?branch=master)](https://travis-ci.org/wakawaka54/ReturnPath-Go)

![Image of Yaktocat](rpfrontend/static/images/go.png)

####[Design Document](DESIGN.md)

[Live Demo](http://dev.waka.run)

[Live API Running](http://api.waka.run/api/sentences)

## Building this Repo

**`ReturnPath-Go`** is made up of two Go packages, **`rpfrontend`** and **`rpapi`**. In order to launch both of these packages, I have facilated a **docker-compose** file which should make it very easy to launch both of these.

1. Make sure you have `docker-compose` and `docker` installed on your machine (don't want to use docker? See below!)
2. Clone the repo using `git clone https://github.com/wakawaka54/ReturnPath-Go.git`
3. Enter into the repo `cd ReturnPath-Go`
4. Run `docker-compose build`
  * This should build both images, depending on how you have `docker-compose` configured, you may need to run this as sudo
5. Run `docker-compose up`
  * This should start up both images
6. The frontend should be mapped to `http://localhost:1400`
7. The backend should be mapped to `http://localhost:1337`

### What if I don't want to use Docker?

Well, you don't have to use Docker! :ok_woman:

1. You will need one of the latest versions of Go. I used `go1.7` for this particular project.
2. Find the **`src/`** folder that goes along with your installation of `Go`. If you don't know what I am talking about, [Go](https://golang.org/doc/install) follow the guide. No pun intended.
3. You can `git clone https://github.com/wakawaka54/ReturnPath-Go.git` directly into the **/src** folder
4. `cd ReturnPath-Go.git` then `cd rpapi` to get into the **`rpapi`** directory.
5. Run `go get` to fetch all the dependencies and use `go install` to generate the binary into the **/bin** folder of your **Go** installation
6. You can now `rpapi` to run the binary if you have your **$PATH** variable modified to include the **/bin** folder
  * If you don't run `rpapi` from inside the **rpapi** directory, you will get some configuration errors. You can fix these by setting the **rpapi** directory path to the `GO_HOME` environment variable
7. You can now `cd ..` and `cd rpfrontend` and use `go get` to fetch all the dependencies and `go install` to generate the binary
8. Again, if you have the **$PATH** variable modified correctly, you should be able to run `rpfrontend` to begin the frontend.
 * If you don't see any data showing up, this is probably because the port mappings are changed on your machine. Set the `ApiAddress` in `config.json` in the **rpfrontend** directory to the correct **rpapi** address on your machine.


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

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;**Status Codes**

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;**`200`** - OK - Everything worked

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;**`500`** - Internal Server Error - Couldn't parse a JSON response for some reason.

#### POST `/api/sentences`
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Creates a sentence object in the database and assigns cooresponding tags.

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;**Example Request:**

```json
{
  "sentence":"this is really amazing if you think about it"
}
```

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;**Status Codes**

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;**`201`** - Created - New sentence was created

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;**`409`** - Conflict - Issue parsing request JSON data

#### DELETE `/api/sentences/{id}`
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Deletes a sentence object in the database

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;**Example Url:** `/api/sentences/8520`

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;**Status Codes**

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;**`202`** - Accepted - Sentence was deleted

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;**`404`** - Not Found - Sentence with ID was not found

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

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;**Status Codes**

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;**`200`** - OK - Everything worked, statistics was sent back

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;**`500`** - Internal Server Error - There was an issue parsing JSON data back to you
