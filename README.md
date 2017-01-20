# GoRepost

Application helps repost news into Social Networks.

# Usage

### Download
    $ git clone https://github.com/8tomat8/GoRepost
    $ cd GoRepost/builds

or
    
    $ wget https://github.com/8tomat8/GoRepost/raw/master/builds/GoRepost-latest
### Run
    $ GoRepost-latest -port 8181 -host 127.0.0.222 -log_dir /var/log/gorepost
    Creating router...
    Router created.
    Starting HTTP server at http://127.0.0.222:8181
Now you could open [http://127.0.0.222:8181](http://127.0.0.222:8181) to see application status.

    HTTP status 422: Could not unmarshal json / List of destinations are empty!
    HTTP status 200: OK. The job is accepted for processing

### Make API call
**Example:**

    curl -X POST -H "Content-Type: application/json" -d '{
       "destinations": {
         "vk": [
           {
             "id": "1234567890",
             "access_key": "9ccd58639b306d26ccfc59d4919e036864a17a0fe98dcd6aa1b7d184eb776b4b5a1af6a289a83f025a7d4",
             "from_group": true  // true - post from groups name, false - post fron users name
           },
           {
             "id": "1234567890",
             "access_key": "9ccd58639b306d26ccfc59d4919e036864a17a0fe98dcd6aa1b7d184eb776b4b5a1af6a289a83f025a7d4",
             "from_group": false  // true - post from groups name, false - post fron users name
           }
         ]
       },
       "message": "Blah blah blah https://www.youtube.com/watch?v=dQw4w9WgXcQ Sign",
       "attachments": [
         {
           "type": "photo",
           "link": "https://cdn.spacetelescope.org/archives/images/wallpaper2/heic0910q.jpg"
         },
         {
           "type": "link",
           "link": "http://google.com"
         }
       ]
     }' "http://127.0.0.222:8181/tasks" 
