# choppa
chop, chop, chop urls!

--- 
Lightweight and blazing fast web app chopping long, boring URLs to custom branded vanity links.

Features:
- ⚡ **Excellent performance** - it leverages Go (Gin) and Redis to redirect users instantaneously
- 🎨 **Customise to your brand's content** - use your own domain name and fully customise link names
- 📱 **Target links by device type** - one branded short link can redirect to different URLs on per-platform basis (Windows, Linux, macOS, Android or iOS) 

## How to run

Requires Docker to be instaled. `git clone` the repository, then run the command `docker-compose build` in the root directory and `docker-compose up` after that.

## Configuration
Handled by environment variables specified in `.env` file. 

**Important: Don't forget to change the default values of `CHOPPA_PASSWORD` and `REDIS_PASSWORD`!**
```
# Password to access POST api endpoints
CHOPPA_PASSWORD=choppawontmissaurl # CHANGE

# Password for redis database
REDIS_PASSWORD=ilovedatabases # CHANGE

# debug / release
GIN_MODE=release  
```

By default, choppa runs on the port `8080`

## How to use

Choppa runs on your server, and it uses the domain name or IP associated with that server. 
A shortened url is called a *chop*. It has its path (ex. example.com/**choppa**) and URLs to redirect to (supports different links per OS).

Frontend is planned to be implemented in the future, but for now creating links is only accessible directly via HTTP requests to the endpoint:

```http request 
POST ip:8080/[path]
```

It requires the following parameters via form data
- `auth` - string: Password to access the API specified in `CHOPPA_PASSWORD` environment variable
- `url` - string : The URL to redirect to 

Also, there is an optional parameter:
- `platform` - string: Override the URL to redirect to for the specified platform
  (Windows, Linux, macOS, Android, iOS) - if left empty, determining the platform is omitted

# License
Shared under Apache License 2.0

# Contributions

Contributions are welcome, just open a pull request 😃
