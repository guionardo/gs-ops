# gs-ops

Self hosted OPS application

## Architecture

### Local repository

```
└── .gsops
    ├── host.yaml
    └── project.yaml
```

.gsops/project.yaml

```yaml
name: test-project
description: "Description of the project"
repository-url: https://github.com/guionardo/project
deploy-branch: main
```

.gsops/host.yaml

```yaml
host: https://ops.guiosoft.info
api-key: XXXXXXXXXXX
```

```mermaid
sequenceDiagram

box Local development environment
participant con as Console
participant cli as GS-OPS CLI
participant fs as Project configuration file
participant git as Local git
end

box External
participant gh as GitHub
end

box Application host
participant backend as Backend
participant frontend as Frontend
participant docker as Container Manager
end

note left of con: CLI operations

note over con: Local setup
con -->> +cli: 
cli ->> fs : Parse arguments <br>and save files
cli -->> -con: Show project data

note over con: Set credentials
con -->> +cli: Receive host API-KEY, github API-TOKEN
cli -->> backend: Send credentials for project
backend -->> gh: Validate API-TOKEN
backend -->> cli: Return status
cli -->> -con: Show status 

note over con: Publish setup
con -->> +cli: 
cli ->> backend: Send project data
backend ->> gh: Setup webhook
gh -->> backend: Return status
backend --) frontend : Notify<br>webhook availibility
backend --) frontend: Notify<br>project data
backend -->> cli: Return status
cli -->> -con: Shows project publish<br>information

note over con: Apply deploy
con -->> +cli: 
cli ->> backend: Apply deploy,<br> or drop container
backend -->> frontend: Notify<br>start apply data
alt deploy
    backend ->> gh : pull files from repository
    backend -->> frontend: Notify<br>git pull
    backend ->> docker : build new image and create tag
    backend -->> frontend: Notify<br>build image
else drop
    backend ->> docker : stop container
    backend -->> frontend: Notify<br>stop container
end

docker -->> backend: Return status
backend -->> frontend: Notify<br>apply status
backend -->> cli: Return status
cli -->> -con: Shows apply status

con -->> +cli : Get log from running application
cli ->> backend : 
backend ->> docker : Get log
docker --> backend : Return log
backend --> cli : Return log
cli -->> -con: Shows log

```


```mermaid
C4Context



      title GS-OPS SAAS OPS Tool
      Enterprise_Boundary(external, "External"){
          System_Ext(gh, "GitHub")
        
      }

System_Boundary(local,"Local development environment"){
System(cli,"GS-OPS CLI")
System(git,"Local git")
}

System_Boundary(host,"Application host"){
System(backend,"GS-OPS Backend")
System(frontend,"GS-OPS Frontend")
System(docker,"Container manager")
}

Rel(cli,backend,"Deploy rules setup","HTTP")
Rel(git,gh,"commit")
Rel(gh,backend,"Notifies","webhook")
Rel(backend,docker,"Run deploy")
    

```

