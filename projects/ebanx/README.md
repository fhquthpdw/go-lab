# Dependencies
- Mysql  
- Redis

# Init and setup dependencies
- #### Mysql:   
  import data from `docs/init/mysql.init.sql` into mysql db

# Configuration
Devops API using [viper](https://github.com/spf13/viper) to do configuration

# golang ci lint
Please do `golangci-lint run ./...` [⇱](https://golangci-lint.run/) and resolve all problems before you commit your codes, or you will fail in CI process.

# Swagger
- Localhost: http://localhost:8090/swagger/index.html
- La-test: https://devops-api.internal.iherbtest.io/swagger/index.html

# Host
- la-test: devops-api.internal.iherbtest.io
- production: 

# Setup dependencies
- Redis
  > $ docker run --rm -d --name redis -p 6379:6379 redis:6.2.4

# Test

# Deploy with Skaffold CD
- Install [skaffold](https://github.com/GoogleContainerTools/skaffold)
- Deploy `using current kube context`.
- Run `cd ./pipeline && skaffold run`
  - by default, `devops-api` deployed to `platform` namespace.
  - to deploy to `la-test`, use:
    > $ skaffold run --kube-context la-test
