# institutions-api

> A RESTful API to query all colleges in India, universities in India and all international universities around the world using their name, a prefix of their name or any part of their name. This is not a wrapper over any third party service or dependency but instead houses all the data within itself and can act as a standalone microservice.

<br />

## Use case.
It can act as a backend service for an autocomplete input that searches for all colleges in India, universities in India and accompanied by all international universities.

<br />

## API.

1) `/institutions` 

To search for all Indian colleges and universities around the world. As you can see the keyword
can either be a prefix to a result or be somewhere inside the result itself. This more or less works like a text index.
```json
GET /institution?search=Kashmir&count=10

[
  "University of Kashmir, Srinagar",
  "KASHMIR WOMENS COLLEGE OF EDUCATION, BARAMULLAH",
  "KASHMIR VALLEY COLLEGE OF EDUCATION, BUDGAM",
  "KASHMIR TIBBIA COLLEGE",
  "KASHMIR PARADISE COLLEGE OF EDUCATION, BARAMULLAH",
  "KASHMIR LAW COLLEGE",
  "KASHMIR COLLEGE OF EDUCATION, BARAMULLAH",
  "Azad Jammu and Kashmir University"
]
```

<br />

## Codebase
The API is written in [Golang (1.12)](https://golang.org/) as a [GO module](https://blog.golang.org/using-go-modules).

The data for the institutions API is lazily loaded once into memory the first time an API call is made via JSON files located in `data/json/*.json`. This in-memory data is then queried for matches. For now there are only three files which hold all the college and university data. Data can however come from any other source too as we will see shortly.

Data sources for colleges or universites actually come from something called providers or more specifically, implementations of an interface called `InstitutionProvider`. 
The default set of providers are located inside `institutions/providers`.

This makes it fairly easy to add a new colleges and universities by adding another data provider which implements the `InstitutionProvider` interface. We can then pass an instance of this newly made provider to the`institutions.NewInstitutionRepository` call inside `main.go`.
Here's the `loadRepository` function inside `main.go`.

```go
func loadRepository() {
	r = institutions.NewInstitutionRepository(
          // Exisiting providers
          providers.NewIndianCollegesProvider(),
          providers.NewIndianUniversitiesProvider(),
          providers.NewWorldUniversitiesProvider(),
          // Add your own `InstitutionProvider` instance here.
	)
	if err := r.Load(); err != nil {
		log.Fatal(err)
	}
}
```

To faciliate additional JSON files a helper `jSONProvder` can be used. More on this inside `institutions/providers/json_provider.go`. All the three JSON files are loaded using this provider. 

### How the data is queried? 
Data that's loaded inside the memory is stored in a data structure called `Trie`. Here's a Wikipedia [article](https://en.wikipedia.org/wiki/Trie) that describes Tries. A Trie stores data as key-value pairs in the form of interconnected nodes and leaves whose position is based on a key they are associated with. Each node has descendants which share a common string prefix. That prefix string is associated with that particular node and thus using a particular prefix key can be mapped to multiple values, in a way. This way a number of institutions can be queried using a common prefix. 
This implementation can however be also used for a string that's located inside the name of the institution such as `Kashmir` which is inside `University of Kashmir`. Before storing an institution we split its name into its individual words and then store the same institution of over and over again for each of the words inside its name which act as keys. This way the institition can be queried by not only its prefix but also using any part of its name. 

<br />

## Development.
* Prerequisites. 
    * Golang (1.12). https://golang.org
* `git clone git@github.com:mustansirzia/institutions-api && cd institutions-api`.
* Run this once `go mod verify`.
* To start server locally. `go run main.go`. An HTTP Server is now live on `localhost:5000`.
* To install `go install`. Then run `institutions-api`.

<br />

## Deployment (Bonus)
* Serverless.
    * Prerequisites.
        * Up. https://up.docs.apex.sh.
* [Up](https://up.docs.apex.sh/) can be used to quickly deploy this API as an AWS Lambda function.
* While inside the project root, run `up` to deploy a version of our API to AWS Lambda. Proper AWS credentials for this to be configured for this though.

<br />
* Docker.
    * Prerequisites.
        * Docker. https://docker.com.
* [Docker](https://up.docs.apex.sh/) is already configured for this repository. Tu run it inside a docker container use these commands.
* `docker build . -t institutions-api`. This will build the docker image and tag it as `institutions-api`. This does not need the Go runtime to be installed locally. Everything is first complied inside a golang base image container and then the executable is copied inside an alpine container which eventually runs.
* `docker run -p 81:5000 institutions-api`. Point browser to `http://localhost:81/institutions` to get the ball rolling. 

<br />

## License.
MIT.
