Used go Container Action Template

### How to use

include into your action

```   
    name: Process repository content
    uses: polupanovaanna/parse_repo_docker_action@master
        with:
          headHash: ${{ github.event.pull_request.base.sha }}
          diff: ${{ $DIFF }}
          host: ${{ $HOST_URL }}
          command: ${{ $COMMAND }}
          targets: ${{ $TARGETS }}
          
```

Where

`diff`: the difference between branches and master HEAD, got by the git diff command
`host`: the url of the server
`command`: the command to run tests/build the target project
`targets`: affected by the changes, Bazel targets, listed divided by spaces