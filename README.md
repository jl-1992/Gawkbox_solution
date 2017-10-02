## Twitch Server

How To Build & Run (from downloaded files):

Step 1: Ensure GOPATH is set to wherever this downloaded directory is<br />
Step 2: Ensure PATH has included $GOPATH/bin<br />
Step 3: Run "go install main" from top-level directory<br />
Step 4: Run "main"

How To Build & Run (as a Docker container from Docker Hub):

Step 1: Run "docker pull jl1992/gawkbox"<br />
Step 2: Run "docker run --rm -p 8080:8080 jl1992/gawkbox"

How To Build & Run (using the Dockerfile):

Step 1: Run "docker build -t your_filename ."<br />
Step 2: Run "docker run --rm -p 8080:8080 your_filename"

The application will be running on localhost:8080

Request/Response Flow:

To obtain the information specified in the remote assignment in JSON format, 
navigate to the route "localhost:8080/users/USERNAME", where USERNAME is the name 
of the Twitch user whose data you wish to access. The HTTP response will be in
JSON format with appropriate headers added for readability.

An Example: [USERNAME=esl_csgo]

    Request=localhost:8080/users/esl_csgo

    Response=

    User info:
    {
    "bio": "http://eslgaming.com",
    "created_at": "2012-06-11T13:36:21Z",
    "display_name": "ESL_CSGO"
    }

    Channel info:
    {
    "views": 254371862,
    "followers": 2188903,
    "game": "Counter-Strike: Global Offensive",
    "language": "en"
    }

    Stream info:
    {
    "stream": true <Indicates whether or not a channel is currently streaming>
    }
