# goCounter

A page view counter web service written in Go. Stores the views as a save file so that the app can be quit and reloaded without losing all of your page views. works with static sites by injecting the page vies as externally loaded JavaScript.

## API

### - Get all page view statistics:
##### `Get` /
Returns a complete list of all pages and their current view count

### - Increment and return a single page's view count:
##### `Get` /[page]/count.js
Increments and returns [page]'s current view count. The response is in the form
of externally loaded JavaScript code, e.g:

```
document.getElementById('viewCount').innerHTML = '3 Page Views';
```

This means that you can use the goCounter service externally of a static website
without hitting CORS issues by importing the view count as externally loaded
JavaScript. For instance with the snippet of code below, the view count would
be loaded and displayed in the paragraph html block:

```
<p id="viewCount"></p>
<script src="http://[address of goCounter service]/index/count.js"></script>
```

### - Set a page's view count:
##### `PUT` /[page]?count=[new view count]
Sets [page]'s view count to [new view count]. Returns 200 on success.


## Run with Docker
```
$ docker build -t gocounter
$ docker run -p [host port here]:3000 gocounter
```

## Run natively (with Go installed)
```
$ go build
$ ./goCounter
```
This will run on port :3000
