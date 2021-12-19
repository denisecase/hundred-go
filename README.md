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

## Running the App

To run the app, have go (golang) installed.

Copy the `.env.example` file to `.env` and provide your Auth0 credentials from your App (e.g. Hundred Go).

```bash
# .env

AUTH0_CLIENT_ID={CLIENT_ID}
AUTH0_DOMAIN={DOMAIN}
AUTH0_CLIENT_SECRET={CLIENT_SECRET}
AUTH0_CALLBACK_URL=https://hundred-go.herokuapp.com/callback
```

After editing .env, run `go mod vendor` to download Go dependencies.

Run `go run main.go` to start the app. If needed, click Allow Access.

Open browser to [http://localhost:3000/](http://localhost:3000/).

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

## Periodically Update Go & Dependencies

```PowerShell
go get -u
go mod tidy
```
In Go 1.17+, go.mod has a second require block for indirect dependencies.

## Resources

- [Auth0](https://auth0.com)
- [Auth0 Golang Web App](https://github.com/auth0-samples/auth0-golang-web-app)
- [Auth0 Go Quickstart](https://auth0.com/docs/quickstart/webapp/golang)

## License

This project is licensed under the MIT license. See the [LICENSE](LICENSE.txt) file for more info.
