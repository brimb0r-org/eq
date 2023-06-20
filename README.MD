[![Build Status](https://your.jenkins.url/buildStatus/)](https://your.jenkins/url/buildStatus)
# Project Deps check
- Run script in the scripts directory to ensure you have all deps installed
    - ``` cd scripts .\deps.sh ```

## ENV Vars
```
export LOCALSTACK="true"
export LOCALSTACK_ENDPOINT="http://localhost:4566"
export AWS_REGION="us-east-1"
export CONFIG_PATH="../../../application/config_files/"
```
## Build and test on local and in ci pipeline with compose ( to build and push artifacts ) [ CI steps to be done in jenkins and push artifacts ]
-  prereq = docker volume create --name=mongodata ( only need to generate once )
-  docker-compose -f docker-compose.override.yaml -f docker-compose.yaml -p testing up --build -d
-  docker-compose -f docker-compose.application.yaml -p app up -d builder
-  docker-compose -f docker-compose.application.yaml -p app exec -T builder bash -c 'go build -buildvcs=false -gcflags="-l" ./application/cmd'
-  docker-compose -f docker-compose.application.yaml -p app push
-  docker-compose -f docker-compose.application.yaml -p app down