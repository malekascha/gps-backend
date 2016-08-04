#Using the API

This API currently has two paths.

/api/addmessage

This route takes a POST request, and will add a message to the database. 
Location data is encoded using the GeoJSON format.
The following is an example request body:

```
{
  "geojson": {
    "type": "Point"
    "coordinates": [25, 25]
  },
  "message": "test",
  "owner": "swarco"
}

```

/api/getmessagesbyradius

This route takes a GET request, and will return all messages within a certain radius from a set of coordinates.
There are 3 query parameters: latitude, longitude, and radius.
For example, to get all messages within 5000 meters of 25,25, simply make a GET request to ${API_URL}/api/getmessage?latitude=25&longitude=25&radius=5000.

#Notes

API url is highwaymessage.com.