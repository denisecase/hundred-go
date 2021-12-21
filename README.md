# hundred-go

This project demonstrates a web app written in Go that uses free Auth0 for authentication and authorization.

## Web Apps - Plan for Auth Early

Most web apps require registration/login/logout - 
authenticating that a user is who they say - 
and then authorizing specific functions based on their credentials. 
The UI and communictions required to implement this are critical and can take a significant amount of time and expertise. 
Early research into how this can be managed should be considered early in the project. 
Using existing solutions can help developers focus on the unique business aspects of their project sooner. 

For this educational app, I wanted something free, 
easy to understand and implement, 
and able to be implemented in a variety of languages and platforms. 
To support student experimentation, 
all selected solutions must provide no-credit card access. 

Auth0 looks like a great fit. 
Permanent free plan - up to 7000 users with 2 social options and no credit card. 
It offers options for a wide variety of server-side platforms including Java, .NET, Node.js, and Go. 

To try Auth0, go to [Auth0](https://auth0.com/) and Sign Up.
You can create Apps and APIs. 
For this, I'll create "Hundred Go" Regular Wed App (views are generated server-side rather than a SPA). 
"Token Endpoint Authentication Method" = Basic sends client secret via the header.
"Token Endpoint Authentication Method" = Post sends client secret via the payload. 
We'll use Basic. 

## App Focus

This app will hold 100 things - could be habits, reminders, developer profiles, events, locations, or other. 
It's a customizable narrow focus that allows developers to develop skills across the entire vertical slice through a complex web application without getting bogged down with unnecessary repetition. 
Students can add images, sound, videos, geolocation, or related entities, and can focus on their preferred skills including front-end design, APIs, data, testing, devops, security, client communication, or management and planning. 

## Hosting on Heroku

Heroku offers a free, no-credit card, easy to use way to deploy directly from GitHub. 

For this app, we'll create a "Regular Web App"/ Go and use the following endpoints (see web/app for implementation):

- <https://hundred-go.herokuapp.com/callback>
- <https://hundred-go.herokuapp.com/home>
- <https://hundred-go.herokuapp.com/login>
- <https://hundred-go.herokuapp.com/logout>
- <https://hundred-go.herokuapp.com/user>

Allowed Logout URLs:

- <https://hundred-go.herokuapp.com/>

Allowed web origins:

- <https://hundred-go.herokuapp.com>

Allowed Origins (CORS) - match the URLS in home.html:

- <https://code.jquery.com>
- <https://maxcdn.bootstrapcdn.com>

## Add GitHub Social Connection

Most students will have GitHub accounts. 
It makes sense to allow users to login to our app using their GitHub Account. 
To do so, follow the instructions at [Auth0 - GitHub Social Connection](https://marketplace.auth0.com/integrations/github-social-connection). In GitHub:

- App Name: Hundred Go
- Homepage URL: <https://hundred-go.herokuapp.com/>
- Description: Example web app written in Go
- Callback: <https://hundred-go.herokuapp.com/login/callback>
- Note use of login/callback, not just callback. 

## Run the App Locally (Quick Test)

To do a quick test, run the app locally. 

- Install Go (aka Golang) on your machine.
- Copy the `.env.example` file to `.env` and provide your Auth0 credentials from your App (e.g. Hundred Go).
- Run `go mod vendor` to create vendor directory with copies of all packages needed
- Run `go run main.go` to start the app. If needed, click Allow Access.

Open browser to [http://localhost:3000/](http://localhost:3000/).

The command `go mod download` is similar to `go mod vendor`, but it  downloads named modules into the module cache. We want local copies in the vendor directory. 

## Make/Makefile - Simplify Common Commands

See the Makefile in the repo. 
Using make allows us to write complicated commands once and then call them by executing a Make target. 
When creating new make targets use a TAB - not spaces. 

Open PowerShell as Admin, install Make, add it to your path, and exit PS.  

- `winget install GnuWin32.Make`
- Win Start Key / Edit the System Environment Variables / Environment Variables / Sytem Variables / Path / Edit / New & add an entry for `C:\Program Files (x86)\GnuWin32\bin` / OK / OK to exit.
- `refreshenv`

Open a new PowerShell as Admin and run the echo target from the Makefile (in this root project folder):
- `make help`

## Initial New Project Creation

Create repo in GitHub cloud with no files. On local machine, open Git Bash in project folder (e.g., hundred-go):

```Bash
git init
git remote add origin https://github.com/denisecase/hundred-go.git
git branch -M main
git add .
git commit -m "initial commit"
git push -u origin main
```

Then, run `heroku login` to set up your app in Heroku. Settings / Config Vars / Reveal / add each variable from your .env file. 

## Create Favicon

Go to <https://favicon.io/favicon-generator/>. 
Create favicon (e.g., Text = 100, circle, Titan One, Regular 400 Normal, 60pt).
Extract to web/static/image/favicon_io folder.

## Installing Go on Windows

- Go installs to $GOROOT, which is C:\Program Files\Go. See bin, docs, and more.  
- User $GOPATH is C:\Users\username\go. See the bin folder as you install things like the linter. 

## Go Linting

We'll use golangci-lint. Don't fetch it with go get, instead, run the following in Git Bash.
This will install the binary to $GOPATH/bin/golangci-lint. Use Edit System Environment variables to explicitly set $GOPATH and $GOPATH\bin to your path. 

```
curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.43.0

golangci-lint --version
```

Once installed, it can be invoked from PowerShell. DOS/Windows ends text lines with uses carriage return and line feed ("\r\n") while Linux-based systems use just line feed ("\n"). When auto-formatting code edited on different operating systems, you'll want to avoid switching back and forth. To ensure line endings for Windows builds are properly formatted, add .gitattributes file with: `*.go text eol=lf`.

## Periodically Update Go & Dependencies

```PowerShell
go get -u
go mod tidy
```
In Go 1.17+, go.mod has a second require block for indirect dependencies.

## Optional: Dockerize the App

We can run the app locally using Go. 
We can deploy to Heroku and run using Procfile and the Go build kit. 
We can also prepare our app to run in a Docker container 

To Dockerize the app:

- Create a Dockerfile
- On Windows, in VS Code, change CRLF to LF (see line endings discussion elsewhere).
- Use a multistage build to keep the final image smaller. 
- See the Dockerfile and references for more information about this file.
- Use the Dockerfile to build the docker image and then run it using either exec.ps1 (on Windows) or exec.sh (on Bash).

```PowerShell
.\exec.ps1
```

Monitor local Docker by opening Docker Desktop and watching in the GUI environment. See:

- Containers/Apps  / click one / Run
- Images / click one / Inspect / Run
- Volumes
- Dev Environments 

Or use the Docker CLI. Use `docker --ps` to list images (should include hundred-go), run an image (e.g., `docker run hundred-go`), and view the running 

```PowerShell
docker version
docker --help
docker --config
docker image ls 
docker image inspect hundred-go
docker run hundred-go
docker ps
docker ps -all
```

For debugging (if an image doesn't start):

```PowerShell
docker create hundred-go  (returns the <container ID>)
docker cp <container ID>:./ ./tmp
docker rm <container ID>
```

If the image builds, but gives an error on trying to `docker run` that says exec user process caused: no such file or directory. In VS Code, change line endings at the bottom of the page from CRFL to LF. 

For reference, on Windows, docker keeps information at:
- C:\Users\<username>\.docker (client config info)

## Resources

- [Auth0](https://auth0.com)
- [Auth0 Golang Web App](https://github.com/auth0-samples/auth0-golang-web-app)
- [Auth0 Go Quickstart](https://auth0.com/docs/quickstart/webapp/golang)
- [Install Make on Windows](https://www.technewstoday.com/install-and-use-make-in-windows/)
- [GNU Make Manual](https://www.gnu.org/software/make/manual/)
- [Makefiles with Go](https://golangdocs.com/makefiles-golang)
- [Make clean fails in Windows](https://stackoverflow.com/questions/49384301/make-clean-failed-in-windows)
- [Go Linting](https://golangci-lint.run/)
- [Dockerize Go App](https://docs.docker.com/language/golang/build-images/)
- [Docker build command](https://docs.docker.com/engine/reference/commandline/build/)
- [Docker run command](https://docs.docker.com/engine/reference/commandline/run/)

## License

This project is licensed under the MIT license. See the [LICENSE](LICENSE.txt) file for more info.
