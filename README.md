# mu-ref-spa

A reference implementation for a Single Page App that uses a CloudFront distribution to serve static content from an S3 bucket, and dynamic content from a Go application running on ECS.


## The application

The application uses the Monte Carlo method to approximate the value of Pi. Inside a square of size 2, a circle with radius 1 is inscribed. Random points are selected in the square. Some will fall inside the circle, and some will not. The ratio of points inside the circle to the total number of points can be used to estimate the value of Pi.

See here for more information: https://en.wikipedia.org/wiki/Pi#Monte_Carlo_methods

The application has a single HTML file that is stored in an S3 bucket, and served by CloudFront. There is JavaScript in this file that makes REST calls. The Go program that implements this REST API has a single endpoint: /api/pi. It takes an optional query string parameter called "count" which is the number of random points to use when estimating Pi. The UI has a slider input control. Every time you change the slider value, another call is made to /api/pi with a new value for "count". The API returns a JSON payload that includes a list of random points in the square, along with an estimate for Pi. The UI renders the random values using D3, as well as the estimated value for Pi.


## To Deploy

```
mu pipeline up
```

## mu-cloudfront extension

This reference app uses the mu-cloudfront extension to set up the S3 Bucket and the CloudFront distribution.

