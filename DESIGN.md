<img src="https://returnpath.com/wp-content/uploads/2015/12/ReturnPath-Favicon-512.png" width=200 height=200 />

# Design & Implementation Doc

#### [API Documentation](README.md)

<img src="http://www.unixstickers.com/image/cache/data/stickers/golang/Go-aviator.sh-600x600.png" width=200 height=200 />
<img src="https://cdn.captora.com/media/docker.com/media/Icon-Cloud-Blue.png-1456879454393" width=200 height=200 />


### Previous Attempts

My inital attempt at this challenge was to use Docker, .Net, and MongoDb. I believe there was some trouble in getting that
iteration to build, it is possible that **docker** was tried to use directly instead of **docker-compose**. However, I was asked to retry this project
and attempt to use **Go** instead. I'm not one to back down from a challenge but having no prior development experience **at all** with
Go, I knew it was going to be a challenge.

## Development Timeline

####`Friday`

**8 AM** - Started learning Go through the the [Tour of Go](https://tour.golang.org/welcome/1) tutorials on the GoLang documentation site.

**5 PM** - Starting working on the Rest Api. Used some key resources such as [Making a Restful Api with Go](http://thenewstack.io/make-a-restful-json-api-go/)

**~2 AM** - Setup my .Net Core frontend to use the new Go Rest backend. Most endpoints were working except for `api/statistics`

####`Saturday`

**10 AM** - Setup **tests** and starting working on *api/statistics* and *prepopulate.go*

**1 PM** - Setup **rpfrontend** with the help of Go documentation to serve single page website

**2 PM** - :family: Had to spend some family time for the rest of the day!

**10 PM** - Started working on **rpfrontend** and got templating to inject ApiAddress into page

**12 AM** - Finished Middleware and read [Best Practices for a Pragmatic Rest API](http://www.vinaysahni.com/best-practices-for-a-pragmatic-restful-api)
and shed a tear that my API did not meet some of the criteria.

**2 AM** - Modified the `GET /api/sentence` endpoint with `limit, offset, and filters`

####`Sunday`

**10 AM** - Started working on deployment, this should be easy right?

**1 PM** - Realized that I can't use relative filepaths with Go. I finally got both Dockerfiles working. After quickly, setting up a
**docker-compose** file all that's left is to deploy to my new dedicated server! Should be easy right?

**3 PM** - Realized I am not a DevOps engineer, I can't use **docker-compose** easily on my server because it has Ubuntu 15.04.

**5 PM** - Finally got my Docker builds to work and DNS addresses all setup along with nginx routing. Wow.

**10 PM** - Added Draft README and DESIGN documents to GitHub repo.

####`Monday`

**4 PM** - Updated Documentation throughout project and tried to parallelize Statistics endpoint.

###Concept

The idea to create an application to do word counting arose from the accessibility of sentence based data from Wikipedia,
the simplicity of writing a small tool to extract this data from Wikipedia, and my desire to dig
through large datasets.

The prepopulated dataset is a collection of over 8000 sentences from the Wikipedia pages of the largest Chemical and Tech companies.

##Software Design&nbsp;&nbsp;![alt interesting](https://cdn0.iconfinder.com/data/icons/octicons/1024/server-16.png)

###Backend Framework

####Prepopulate
There is almost always a need to prepopulate data.

**`PrepopulateDataset()`**

Short function which reads a prepopulate data file with JSON sentence data and uploads it to the application datastore

####Config
This command file gives the application the ability to read a configuration from a `config.json` file.

**`Config struct`**

Structure which holds the configuration data

**`BuildConfig()`**

The difficulties with this function arose around the possibility of starting the Go binary from a folder location other than the project folder.
I added the ability to read an environment variable with the path to the Project Directory. The `Port` and `Hostname` can also be changed
from the terminal by modifying environment variables (`GO_HOSTNAME` and `GO_PORT`)

####Middleware

I am used to frameworks that allow you to attach middleware functions.

**`Middleware struct`**

Allows attaching `http.Handler` functions together to form a pipeline. This can be useful for adding headers such as `EnableCors`
and other obscure uses like handling `OPTIONS` requests when receiving AJAX requests from a client.

**`Middleware.AddService()`**

Attaches a service to the middleware "pipeline". They are called in a Last-In -> First-Out manner. Middleware have the ability
to short circuit the pipeline. They can decide not call the `next.ServeHTTP(w,r)` which will cause the request to end.

**`Middleware.AddHandler()`**

Attaches a final handler to the "pipeline", usually the final router goes here.

####Routes

This application uses the `gorilla/mux` router. The primary decision for this was ease of use. I liked how easy it is to setup
and it was highly recommended through several articles. Either way, once it is injected into the Middleware pipeline, it implements
`http.Handler`. The routes are setup in the `Routes.go` file which allows easy access to adding, removing, changing existing routes.

**`NewRouter()`**

Returns a new router with all the routes defined in the `routes` object in the `Routes.go` file.

####Handlers

All the http route handlers are implemented here. Usually it would preferable for buisness logic to not occur in the `Handlers`, however due to the
small nature of this application and my relative inexperience with Go, I took care of that logic in there.

These handlers are described pretty well in the API documentation.

####Sentence

Defines the primary **model** and **datastore** of the application.

**`Sentence struct`**

Sentence model used throughout application

**`SentenceCompare struct`**

Structure used for comparison purposes. Allows the defining of `nil` properties which is useful for identifying which fields are
being filtered by.

Additionally, the filtering logic for sentences happens here because I could not find generic methods to accomplish this.

####Deployment

Deployment is containerized using a Docker-Compose file or can be individually deployed using each subdirectories Dockerfile.

###Frontend

The frontend is a pretty simple project. It implements some of the same things as the backend, such as `config`, and `http pipeline`.

A static file server was setup to server CSS, JS, images, and the favicon ![PCMASTERRACE](rpfrontend/static/favicon.ico)

**`Index()`**

The method for serving the single html page. Uses `html/template` to process a template and parse the `IndexModel` onto the page. This
is how the frontend gets the `ApiAddress` information to process AJAX from the client.

####Site.js

This is where all the frontend magic happens for communicating with API with AJAX and parsing the response into something the user
can see. They are **not** beautifully written as I focused more on getting a solid backend.

##Overview

If you've made it this far then you've probably gotten a pretty good overview of the application framework! In general, there were a lot
of challenges along the way and I learned an entirely new langauge within about half a week and got a semi decent API running.
David mentioned `"we are looking for fast, agile engineers"`, I have demostrated that I'm a naturally smarter than average engineer and
I am always looking forward to learning something new and exciting! :thumbsup:

# Thank you :metal:
