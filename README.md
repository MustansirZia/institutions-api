# musalleen-apis

> A set of commonly used APIs in Musalleen and it's other projects. 

https://apis.musalleen.com

These should essentially be stateless and quick to cold start. This ensures they can be scaled infinitely and used in serverless or faaS enviroments such as AWS Lambda. The APIs are RESTful and are exposed via a single lambda function. 

The APIs are written in [Golang (1.12)](https://golang.org/) and [fasthttp](https://github.com/valyala/fasthttp) as a [GO module](https://blog.golang.org/using-go-modules).

## APIs.

1) `/institution` 

To search for all Indian colleges and univeristies around the world.
```json
/institution?search=Kashmir&count=10

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

## Development.
* Prerequisites. 
    * Golang (1.12). https://golang.org
    * Up. https://up.docs.apex.sh

* `git clone git@github.com:qazimusab/musalleen-apis && cd musalleen-apis`.
* Run this once `go mod verify`.
* To start server locally. `go run main.go`. Server is now live on `localhost:5000`. First time around might take a while till all the third party libraries are download. After that the startup time is around ~1 second even taking into account all this data. ❤️

## Codebase
The data for the institution API is loaded into memory on each cold start via JSON files located in `data/json/*.json`. For now there are only three files which hold all the college and university data. Data can however come from any other source too as we will see shortly. 

Data sources for colleges or universites  actually come from something called providers or more specifically, implementations of an interface called `InstitutionProvider`. 
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
        //
        // Add your own `InstitutionProvider` instance here.
        //
	)
	if err := r.Load(); err != nil {
		log.Fatal(err)
	}
}
```

To faciliate additional JSON files a helper `jSONProvder` can be used. More on this inside `institutions/providers/json_provider.go`. All the three JSON files are loaded using this provider. 

## Deployment
* [Up](https://up.docs.apex.sh/) can be used to quickly deploy our API as an AWS Lambda function.
* Run `./scripts/deploy-staging.sh` to deploy staging version of our API to `https://staging-apis.musalleen.com` and `scripts/deploy.sh` to deploy the production version to `https://apis.musalleen.com`. Custom domain name  and API Gateway are already configured inside AWS. Proper AWS credentials for Musalleen are injected automatically.

## License.
Unlicensed.

Cheers!