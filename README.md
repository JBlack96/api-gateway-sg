## Getting Started

Pre-Req -> please ensure you are signed into your AWS account via the cli before running these commands -> you may also need go installed in order to build the code project.

In order to run this repository you will need to run the following scripts in order:

(If the files don't have execute permissions from the cli please run ``chmod +x ./scripts/SCRIPT_NAME_HERE)``

1. ``./scripts/create-s3-bucket.sh`` - used to create the s3 bucket so we can copy the lambda files up
2. ``./scripts/deploy.sh`` - used to create the api gateway + lambda for the Hello function

If these files execute correctly you should then be able to login to your AWS console (web interface) and search api gateway where you should find your created gateway api. 

### Need to add an api gateway stage so the DNS is available (Will only be testable via AWS currently)
If you get the DNS from here that AWS assigns, you will then be able to attempt a post a request with a json body ``{"name": "MYNAME"}`` where you get a simple hello message and a 200 status back.