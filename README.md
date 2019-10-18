# ci-wake


## Usage 

# ci-wake.sh
You will need to change the url and team name in the `check_pipelines` function. 
then simply run 

```
./ci-wake.sh
```

*Warning*  if you currently have "green" CI the monitor will switch off immediately 

# server

This tool allows you to force the CI machine to wake up. To run;

```
export PORT=8090
cd ci-wake
./server
```

To force wake up the CI machine. 
```
curl http://{address of CI machine}:${PORT}/stop
```

## TO DO
- Paramaterise the concourse URL and Team
- Force to take a port number for `server` or have a sensible default
- Have the `server` app return the address it is listening on, on start up. 
