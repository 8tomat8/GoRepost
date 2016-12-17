# GoRepost

Application helps repost news into Social Networks **!!!Not stable!!!**

# Usage

### Download
    $ git clone https://github.com/8tomat8/GoRepost
    $ cd GoRepost/builds

or
    
    $ wget https://github.com/8tomat8/GoRepost/raw/master/builds/GoRepost-latest
### Run
    $ GoRepost-latest -port 8181 -host 127.0.0.123
    Creating router...
    Router created.
    Starting HTTP server at http://127.0.0.123:8181
Now you could open [http://127.0.0.123:8181](http://127.0.0.123:8181) to see application status.

### Make API call
**Example:**

    curl -X POST -H "Content-Type: application/json" -d '{
        "destinations": [{
                "social": "vk",
                "group_ids": [
                    "123456789",
                    "123456789"
                ]
        }],
        "message": "Text Text Text Text \n Text Text Text ",
        "attachments": [
            {
                "type": "photo",
                "link": "http://link.to/image/filename.jpg"
            },
            {
                "type": "photo",
                "link": "http://link.to/image/dynamic/"
            }
        ]
    }' "http://127.0.0.123:8181/tasks" 
