## Remote Assignment V3

How To Build & Run:

To build the server, first ensure that the $GOPATH is set to the path where 
you downloaded this .zip from GitHub. Then, from the top-level directory run "go
install twitch". Lastly, run "go run main.go" to start the server.

Request/Response Flow:

To obtain the information specified in the remote assignment in JSON format, 
navigate to the route "localhost:8080/users/USERNAME", where USERNAME is the name 
of the Twitch user whose data you wish to access. The HTTP response will be in
JSON format with appropriate headers added for readability.

An Example: [user=esl_csgo]

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

