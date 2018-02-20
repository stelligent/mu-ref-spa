# mu-ref-spa

A reference implementation for a Single Page App that uses a CloudFront distribution to serve static content from an S3 bucket, and dynamic content from a Go application running on ECS.


## The application

The application uses the Monte Carlo method to approximate the value of Pi. Inside a square of size 2, a circle with radius 1 is inscribed. Random points are selected in the square. Some will fall inside the circle, and some will not. The ratio of points inside the circle to the total number of points can be used to estimate the value of Pi.

See here for more information: https://en.wikipedia.org/wiki/Pi#Monte_Carlo_methods

The application has a single HTML file that is stored in an S3 bucket, and served by CloudFront. There is JavaScript in this file that makes REST calls. The Go program that implements this REST API has a single endpoint: /api/pi. It takes an optional query string parameter called "count" which is the number of random points to use when estimating Pi. The UI has a slider input control. Every time you change the slider value, another call is made to /api/pi with a new value for "count". The API returns a JSON payload that includes a list of random points in the square, along with an estimate for Pi. The UI renders the random values using D3, as well as the estimated value for Pi.


## Prerequisites

* Configure your AWS credentials. A simple way to do this to install the AWS command line tool, and then run `aws configure`. Instructions for installing the CLI can be found here: https://docs.aws.amazon.com/cli/latest/userguide/installing.html

* Install mu, using these instructions: https://github.com/stelligent/mu/wiki/Installation

## Fork this repo

Fork this repo. Then edit the mu.yml file, and change the repo name to the be the name of your new repo. A CodePipeline will be created that is started when changes are committed to this repo, so you want it to have your own repo name, not the stelligent one.

```
service:
  name: spa
  ...
  pipeline:
    source:
      repo: [Put your repo name here]
  ...
```

Commit this change and push to GitHub.

```
git add mu.yml
git commit -m'Use new repo'
git push origin master
```

## Deploy to acceptance environment

You will need a personal access token from GitHub that CodePipeline will use to access your repo.
If you don't alredy have one of these, you can  go here to set one up: https://github.com/settings/tokens

```
mu pipeline up -t [GitHub Personal Token]
```

You can monitor the progress of the pipeline by using this command:

```
mu svc show
```

## Find the domain name for your CloudFront distribution

Once the application has been deployed to the acceptance environment you need to find the CloudFront distribution that was created for your environment. You can go the AWS Console and select the CloudFront service.
You can also use the AWS CLI to get your CloudFront distribution domain name. The following command will list all of your CloudFront domain names:

```
aws cloudfront list-distributions --query 'DistributionList.Items[*].DomainName'
```

Point your browser at the domain name returned to see the running application. If you already have other CloudFront distributions provisioned, it will probably be simpler to use the AWS console to find the correct domain name. The CloudFormation console will have a stack named 'mu-loadbalancer-acceptance' (and eventually mu-loadbalancer-production). If you check the Resources tab for those stacks you can find the Physical ID for CloudFrontDist. Find that ID on the CloudFront console, and click it to show more detail including the domain name.

## Deploy to production

The app will be automatically deployed to the acceptance envinroment. A manual approval is needed before the application will be deployed to the production environment. Go to the CodePipeline console, select the mu-spa pipline, then click the Review button. From there you can approve or reject the deployment to production.

## Extension: mu-cloudfront

This reference app uses the mu-cloudfront extension to set up the S3 Bucket and the CloudFront distribution. The mu.yml file has an extensions attribute which contains a single extension in this case. The parameters attribute in the mu.yml file also specifies a value for "SourcePath". This parameter is used by the extension, and is a relative path for the directory that contains static files that should be served by CloudFront.

## Next steps

This example uses static files that already exist in the public directory of the project. But you can also use the buildspec.yml file to specify commands that generate the static files. For example, if you are writing your web client using React, you might be using webpack to transpile your React components and bundle them into a single JavaScript file. In that case, put the webpack command in buildspec.yml in the build section. And you will need to update the SourcePath parameter in mu.yml to refer to the path where webpack bundles the file you want CloudFront to serve.

