#WEBSITE INFORMATION API

This is a Go API server to get websites information with their relevance score and view count.

## Run DockerFile and Create container image

###build container image
```
docker build -t test/server .
```
###run container image
```
docker run -d -p 5000:5000 test/server:latest
```

###use curl to Test the API server
```
curl "http://localhost:5000/info?sortKey=views&limit=15"
```

#Usage

Server will start on Port :5000
```
curl "http://localhost:5000/info?sortKey=relevanceScore&limit=10"
```

##Mandatory Query Parameters :
###Sort Param
sortKey string 

###limit Param
limit int

