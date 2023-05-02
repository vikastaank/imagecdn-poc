# imagecdn-poc
Sample image cdn API which will accept image identifier key in query string params and will set image into response stream on webpage. Will be using local disk cache for images.

### [GET] "v1/images?url={encoded-img-identifier-key}"
1. Validate and check in-memory cache if image data existing for that key/filename.
2. If in-memory cache not having required data, check into local disk cache if image is residing by that name, and update the in-memory cache.
3. Image file raw data will be set into response stream to show on the webpage.

* Sample images in `resources/images` dir.
* Sample API urls : 
```
http://localhost:8080/v1/images?url=dGVzdGluZy5qcGVn
http://localhost:8080/v1/images?url=c2FtcGxlX21lZGlhX2ZpbGUucG5n
```
