## API

### /data/:key/:part - GET
* key: access ID
* part: integer index of data to be accessed
Returns blob of data.

### /data/:key/meta - GET
* key: access ID
Returns JSON of metadata, including data names.

### /data - POST
Form data within body:
* data: 1 or more blobs to be uploaded
Returns JSON with access ID (Key).
