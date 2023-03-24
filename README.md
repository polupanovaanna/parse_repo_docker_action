Used go Container Action Template

### how to use

include into your action

```   
    uses: polupanovaanna/parse_repo_docker_action@master
        with:
          headHash: ${{ github.event.pull_request.base.sha }}
          diff: ${{ $DIFF }}
          host: ${{ $HOST_URL }}
          command: ${{ $COMMAND }}
          
```