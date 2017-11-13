#### TYK Bringup

# What is Tyk
Tyk acts as a Proxy or ("Gateway") to your API Services. 

# What is Tyk-plugin-demo-golang:
Tyk-plugin-demo-gloang is the middleware where Schema Validation or any Validation checks can be done.
The project implements a simple middleware for header injection (MyPreHook), using a **Pre** hook (see [Tyk custom middleware hooks](https://tyk.io/docs/tyk-api-gateway-v1-9/javascript-plugins/middleware-scripting/)). An authentication hook is also provided (MyAuthCheck), see [hooks.go](hooks.go).

# Tyk over NGINX
You can specicy Tyk servers as upstream in the nginx.conf
https://tyk.io/docs/tyk-api-gateway-v1-9/configuration/working-with-nginx/


```
worker_processes 1;

error_log /var/log/nginx/error.log;
pid /var/run/nginx.pid;

events {
    worker_connections 1024;
    use epoll;
}

http {
    # Enumerate all the Tyk servers here
    upstream tyk {
        server 127.0.0.1:5000, 127.0.0.1:5001, 127.0.0.1:5002;
    }

```

# Tyk Benchmarks
Steady pace between ~20ms to ~65ms latency all the way up to ~3,000 requests per second. (while Analytics is being recorded in redis in parallel)
https://tyk.io/features/tyk-benchmarks/



# Flow
User --> REST API --> tyk[API Gateway] --> tyk-plugin-demo-golang (Middleware)-->Applications.

git clone https://github.com/TykTechnologies/tyk.git
Build tyk with "coprocess grpc" to support GRPC plugin
root@elastic-stack:~/go/src/github.com/TykTechnologies/tyk# go build -tags "coprocess grpc"

Modify tyk.conf to include the coprocess_options. In this case, tyk-plugin-demo-golang Plugin will be listening on Port 27018. 

```
"coprocess_options": {
        "enable_coprocess": true,
        "coprocess_grpc_server": "tcp://localhost:27018",
        "python_path_prefix": ""
    },
```


git clone https://github.com/TykTechnologies/tyk-plugin-demo-golang.git
Run - "tyk-cli bundle build" (Get tyk-cli "go get github.com/TykTechnologies/tyk-cli"
      This command generates bundles.zip that contains the manifest (manifest.json)
      Copy the middle_ware block from manifest.json to test_api.

```
"custom_middleware_bundle": "test-bundle",
    "enable_coprocess_auth" : true,
    "active": true,
    "custom_middleware" : {
      "auth_check" : {
         "name" : "MyAuthCheck",
         "require_session" : false,
         "path" : ""
      },
      "post" : null,
      "pre" : [
         {
            "name" : "MyPreHook",
            "require_session" : false,
            "path" : ""
         }
      ],
      "id_extractor" : {
         "extract_with" : "",
         "extractor_config" : null,
         "extract_from" : ""
      },
      "post_key_auth" : null,
      "response" : null,
      "driver" : "grpc"
   },

```


# Start tyk

```
root@elastic-stack:~/go/src/github.com/TykTechnologies/tyk# go build -tags "coprocess grpc"
root@elastic-stack:~/go/src/github.com/TykTechnologies/tyk# ./tyk 
[Nov  4 06:00:08]  INFO Connection dropped, reconnecting...
[Nov  4 06:00:08]  INFO main: PIDFile location set to: /var/run/tyk-gateway.pid
[Nov  4 06:00:08]  INFO main: Initialising Tyk REST API Endpoints
2017/11/04 06:00:08 grpc: addrConn.resetTransport failed to create client transport: connection error: desc = "transport: dial tcp [::1]:27018: getsockopt: connection refused"; Reconnecting to { <nil>}
[Nov  4 06:00:08]  INFO host-check-mgr: Starting Poller
[Nov  4 06:00:08]  INFO main: --> Standard listener (http)
[Nov  4 06:00:08]  INFO main: Setting up Server
[Nov  4 06:00:08]  INFO main: Initialising distributed rate limiter
[Nov  4 06:00:08]  INFO Loading API Specification from apps/1.json
[Nov  4 06:00:08]  INFO Starting gateway rate limiter notifications...
[Nov  4 06:00:08]  INFO Loading API Specification from apps/app_sample.json
[Nov  4 06:00:08]  INFO Loading API Specification from apps/coprocess_app_sample.json
[Nov  4 06:00:08]  INFO Loading API Specification from apps/coprocess_app_sample_protected.json
[Nov  4 06:00:08]  INFO Loading API Specification from apps/coprocess_grpc_app_sample.json
[Nov  4 06:00:08]  INFO Loading API Specification from apps/coprocess_grpc_app_sample_protected.json
[Nov  4 06:00:08]  INFO Loading API Specification from apps/coprocess_lua_app_sample.json
[Nov  4 06:00:08]  INFO Loading API Specification from apps/quickstart.json
[Nov  4 06:00:08]  INFO main: Detected 8 APIs
[Nov  4 06:00:08]  INFO main: Loading API configurations.
[Nov  4 06:00:08]  INFO main: Tracking hostname api_name=Tyk Test API (Coprocess auth) domain=(no host)
[Nov  4 06:00:08]  INFO main: Tracking hostname api_name=Tyk Test API (Coprocess + Lua) domain=(no host)
[Nov  4 06:00:08]  INFO main: Tracking hostname api_name=Tyk Test API (Coprocess) domain=(no host)
[Nov  4 06:00:08]  INFO main: Tracking hostname api_name=Tyk Test API (Coprocess + gRPC) domain=(no host)
[Nov  4 06:00:08]  INFO main: Tracking hostname api_name=Tyk Test API domain=(no host)
[Nov  4 06:00:08]  INFO main: Tracking hostname api_name=Tyk Test API domain=(no host)
[Nov  4 06:00:08]  INFO main: Tracking hostname api_name=Test API domain=(no host)
[Nov  4 06:00:08]  INFO main: Tracking hostname api_name=Tyk Test API (Coprocess + gRPC) domain=(no host)
[Nov  4 06:00:08]  INFO main: Loading API api_name=Tyk Test API (Coprocess + gRPC)
[Nov  4 06:00:08]  INFO main: Loading API api_name=Tyk Test API (Coprocess + gRPC)
[Nov  4 06:00:08] ERROR main: No bundle base URL set, skipping bundle: <nil> api_id=7 org_id=default path=- server_name=http://httpbin.org user_id=- user_ip=-
[Nov  4 06:00:08]  INFO main: Loading API api_name=Tyk Test API (Coprocess auth)
[Nov  4 06:00:08]  INFO main: Loading API api_name=Tyk Test API
[Nov  4 06:00:08]  INFO main: Loading API api_name=Tyk Test API
[Nov  4 06:00:08]  INFO main: Loading API api_name=Test API
[Nov  4 06:00:08]  INFO main: Loading API api_name=Tyk Test API (Coprocess)
[Nov  4 06:00:08]  INFO main: Loading API api_name=Tyk Test API (Coprocess + Lua)
[Nov  4 06:00:08]  INFO main: Checking security policy: Token api_name=Tyk Test API (Coprocess + gRPC)
[Nov  4 06:00:08] ERROR coprocess: CP Driver not supported: python
[Nov  4 06:00:08]  INFO main: Checking security policy: Token api_name=Tyk Test API (Coprocess auth)
[Nov  4 06:00:08]  INFO main: Checking security policy: Token api_name=Tyk Test API
[Nov  4 06:00:08]  INFO main: Checking security policy: Open api_name=Tyk Test API
[Nov  4 06:00:08] ERROR main: No bundle base URL set, skipping bundle: <nil> api_id=1 org_id=53ac07777cbb8c2d53000002 path=- server_name=http://httpbin.org/ user_id=- user_ip=-
[Nov  4 06:00:08]  INFO main: Checking security policy: Open api_name=Tyk Test API (Coprocess + gRPC)
[Nov  4 06:00:08]  INFO main: Checking security policy: Open api_name=Tyk Test API (Coprocess + Lua)
[Nov  4 06:00:08] ERROR coprocess: CP Driver not supported: lua
[Nov  4 06:00:08]  INFO main: Checking security policy: Open api_name=Tyk Test API (Coprocess)
[Nov  4 06:00:08] ERROR coprocess: CP Driver not supported: python
[Nov  4 06:00:08] ERROR coprocess: CP Driver not supported: python
[Nov  4 06:00:08]  INFO main: Processed and listening on: /coprocess-auth-tyk-api-test/{rest:.*}
[Nov  4 06:00:08]  INFO main: Processed and listening on: /coprocess-lua-tyk-api-test/{rest:.*}
[Nov  4 06:00:08]  INFO main: Processed and listening on: /coprocess-tyk-api-test/{rest:.*}
[Nov  4 06:00:08]  INFO main: Processed and listening on: /grpc-protected/{rest:.*}
[Nov  4 06:00:08]  INFO main: Processed and listening on: /tyk-api-test/{rest:.*}
[Nov  4 06:00:08]  INFO main: Processed and listening on: /quickstart/{rest:.*}
[Nov  4 06:00:08]  INFO main: Processed and listening on: /test-api/{rest:.*}
[Nov  4 06:00:08]  INFO main: Processed and listening on: /grpc/{rest:.*}
[Nov  4 06:00:08]  INFO host-check-mgr: Loading uptime tests...
[Nov  4 06:00:08]  INFO main: Initialised API Definitions
[Nov  4 06:00:08]  INFO main: Loading policies
[Nov  4 06:00:08]  INFO main: Gateway started (v2.3.99)
[Nov  4 06:00:08]  INFO main: --> Listening on address: (open interface)
[Nov  4 06:00:08]  INFO main: --> Listening on port: 8080
[Nov  4 06:00:08]  INFO main: --> PID: 16802
2017/11/04 06:00:09 grpc: addrConn.resetTransport failed to create client transport: connection error: desc = "transport: dial tcp [::1]:27018: getsockopt: connection refused"; Reconnecting to { <nil>}
2017/11/04 06:00:10 grpc: addrConn.resetTransport failed to create client transport: connection error: desc = "transport: dial tcp [::1]:27018: getsockopt: connection refused"; Reconnecting to { <nil>}
2017/11/04 06:00:12 grpc: addrConn.resetTransport failed to create client transport: connection error: desc = "transport: dial tcp [::1]:27018: getsockopt: connection refused"; Reconnecting to { <nil>}
2017/11/04 06:00:17 grpc: addrConn.resetTransport failed to create client transport: connection error: desc = "transport: dial tcp [::1]:27018: getsockopt: connection refused"; Reconnecting to { <nil>}
2017/11/04 06:00:24 grpc: addrConn.resetTransport failed to create client transport: connection error: desc = "transport: dial tcp [::1]:27018: getsockopt: connection refused"; Reconnecting to { <nil>}


----

[Nov  4 06:00:54]  WARN API Definition with this ID already exists, deleting file...
[Nov  4 06:01:06]  INFO api: Reload URL Structure - Scheduled
[Nov  4 06:01:06]  INFO Reload queued
[Nov  4 06:01:06]  INFO Initiating reload
[Nov  4 06:01:06]  INFO main: Loading policies
[Nov  4 06:01:06]  INFO Loading API Specification from apps/1.json
[Nov  4 06:01:06]  INFO Loading API Specification from apps/app_sample.json
[Nov  4 06:01:06]  INFO Loading API Specification from apps/coprocess_app_sample.json
[Nov  4 06:01:06]  INFO Loading API Specification from apps/coprocess_app_sample_protected.json
[Nov  4 06:01:06]  INFO Loading API Specification from apps/coprocess_grpc_app_sample.json
[Nov  4 06:01:06]  INFO Loading API Specification from apps/coprocess_grpc_app_sample_protected.json
[Nov  4 06:01:06]  INFO Loading API Specification from apps/coprocess_lua_app_sample.json
[Nov  4 06:01:06]  INFO Loading API Specification from apps/quickstart.json
[Nov  4 06:01:06]  INFO main: Detected 8 APIs
[Nov  4 06:01:06]  INFO main: Preparing new router
[Nov  4 06:01:06]  INFO main: Initialising Tyk REST API Endpoints
[Nov  4 06:01:06]  INFO main: Loading API configurations.
[Nov  4 06:01:06]  INFO main: Tracking hostname api_name=Tyk Test API (Coprocess auth) domain=(no host)
[Nov  4 06:01:06]  INFO main: Tracking hostname api_name=Tyk Test API (Coprocess + Lua) domain=(no host)
[Nov  4 06:01:06]  INFO main: Tracking hostname api_name=Tyk Test API (Coprocess) domain=(no host)
[Nov  4 06:01:06]  INFO main: Tracking hostname api_name=Tyk Test API (Coprocess + gRPC) domain=(no host)
[Nov  4 06:01:06]  INFO main: Tracking hostname api_name=Tyk Test API domain=(no host)
[Nov  4 06:01:06]  INFO main: Tracking hostname api_name=Tyk Test API domain=(no host)
[Nov  4 06:01:06]  INFO main: Tracking hostname api_name=Test API domain=(no host)
[Nov  4 06:01:06]  INFO main: Tracking hostname api_name=Tyk Test API (Coprocess + gRPC) domain=(no host)
[Nov  4 06:01:06]  INFO main: Loading API api_name=Tyk Test API (Coprocess + gRPC)
[Nov  4 06:01:06]  INFO main: Checking security policy: Token api_name=Tyk Test API (Coprocess + gRPC)
[Nov  4 06:01:06]  INFO main: Loading API api_name=Tyk Test API (Coprocess auth)
[Nov  4 06:01:06] ERROR coprocess: CP Driver not supported: python
[Nov  4 06:01:06]  INFO main: Checking security policy: Token api_name=Tyk Test API (Coprocess auth)
[Nov  4 06:01:06]  INFO main: Loading API api_name=Tyk Test API (Coprocess + Lua)
[Nov  4 06:01:06]  INFO main: Checking security policy: Open api_name=Tyk Test API (Coprocess + Lua)
[Nov  4 06:01:06] ERROR coprocess: CP Driver not supported: lua
[Nov  4 06:01:06]  INFO main: Loading API api_name=Tyk Test API (Coprocess)
[Nov  4 06:01:06]  INFO main: Checking security policy: Open api_name=Tyk Test API (Coprocess)
[Nov  4 06:01:06] ERROR coprocess: CP Driver not supported: python
[Nov  4 06:01:06] ERROR coprocess: CP Driver not supported: python
[Nov  4 06:01:06]  INFO main: Loading API api_name=Tyk Test API
[Nov  4 06:01:06]  INFO main: Checking security policy: Open api_name=Tyk Test API
[Nov  4 06:01:06]  INFO main: Loading API api_name=Tyk Test API
[Nov  4 06:01:06]  INFO main: Loading API api_name=Tyk Test API (Coprocess + gRPC)
[Nov  4 06:01:06] ERROR main: No bundle base URL set, skipping bundle: <nil> api_id=7 org_id=default path=- server_name=http://httpbin.org user_id=- user_ip=-
[Nov  4 06:01:06]  INFO main: Checking security policy: Open api_name=Tyk Test API (Coprocess + gRPC)
[Nov  4 06:01:06]  INFO main: Checking security policy: Token api_name=Tyk Test API
[Nov  4 06:01:06]  INFO main: Loading API api_name=Test API
[Nov  4 06:01:06] ERROR main: No bundle base URL set, skipping bundle: <nil> api_id=1 org_id=53ac07777cbb8c2d53000002 path=- server_name=http://httpbin.org/ user_id=- user_ip=-
[Nov  4 06:01:06]  INFO main: Processed and listening on: /coprocess-auth-tyk-api-test/{rest:.*}
[Nov  4 06:01:06]  INFO main: Processed and listening on: /coprocess-lua-tyk-api-test/{rest:.*}
[Nov  4 06:01:06]  INFO main: Processed and listening on: /coprocess-tyk-api-test/{rest:.*}
[Nov  4 06:01:06]  INFO main: Processed and listening on: /grpc-protected/{rest:.*}
[Nov  4 06:01:07]  INFO main: Processed and listening on: /tyk-api-test/{rest:.*}
[Nov  4 06:01:07]  INFO main: Processed and listening on: /quickstart/{rest:.*}
[Nov  4 06:01:07]  INFO main: Processed and listening on: /test-api/{rest:.*}
[Nov  4 06:01:07]  INFO main: Processed and listening on: /grpc/{rest:.*}
[Nov  4 06:01:07]  INFO host-check-mgr: Loading uptime tests...
[Nov  4 06:01:07]  INFO main: Initialised API Definitions
[Nov  4 06:01:07]  INFO main: API reload complete
[Nov  4 06:01:07]  INFO Initiating coprocess reload
[Nov  4 06:01:07]  INFO coprocess: Reloading middlewares
[Nov  4 06:01:10]  WARN Incorrect key expiry setting detected, correcting

```

# Start the plugin

```
root@elastic-stack:~/go/src/github.com/TykTechnologies/tyk-plugin-demo-golang# ./tyk-plugin-demo-golang 
2017/11/04 06:00:31 Listening...
2017/11/04 06:01:10 Receiving object: hook_type:Pre hook_name:"MyPreHook" request:<headers:<key:"Accept" value:"*/*" > headers:<key:"Authorization" value:"d29e8f389a6cf39a72bc7156c5e710885e4b5b89" > headers:<key:"User-Agent" value:"curl/7.52.1" > url:"/test-api/get" return_overrides:<response_code:-1 > > spec:<key:"APIID" value:"1" > spec:<key:"OrgID" value:"53ac07777cbb8c2d53000002" > 
Receiving object: hook_type:Pre hook_name:"MyPreHook" request:<headers:<key:"Accept" value:"*/*" > headers:<key:"Authorization" value:"d29e8f389a6cf39a72bc7156c5e710885e4b5b89" > headers:<key:"User-Agent" value:"curl/7.52.1" > url:"/test-api/get" return_overrides:<response_code:-1 > > spec:<key:"APIID" value:"1" > spec:<key:"OrgID" value:"53ac07777cbb8c2d53000002" > 
2017/11/04 06:01:10 MyPreHook is called!
2017/11/04 06:01:10 Receiving object: hook_type:CustomKeyCheck hook_name:"MyAuthCheck" request:<headers:<key:"Accept" value:"*/*" > headers:<key:"Authorization" value:"d29e8f389a6cf39a72bc7156c5e710885e4b5b89" > headers:<key:"Myheader" value:"Myvalue" > headers:<key:"User-Agent" value:"curl/7.52.1" > url:"/test-api/get" return_overrides:<response_code:-1 > > spec:<key:"APIID" value:"1" > spec:<key:"OrgID" value:"53ac07777cbb8c2d53000002" > 
Receiving object: hook_type:CustomKeyCheck hook_name:"MyAuthCheck" request:<headers:<key:"Accept" value:"*/*" > headers:<key:"Authorization" value:"d29e8f389a6cf39a72bc7156c5e710885e4b5b89" > headers:<key:"Myheader" value:"Myvalue" > headers:<key:"User-Agent" value:"curl/7.52.1" > url:"/test-api/get" return_overrides:<response_code:-1 > > spec:<key:"APIID" value:"1" > spec:<key:"OrgID" value:"53ac07777cbb8c2d53000002" > 
2017/11/04 06:01:10 MyAuthCheck is called!
2017/11/04 06:01:10 Successful authentication on MyAuthCheck
```




# POST
```

1. Create test api
./test_api.sh

2. Hot Reload tyk for the config to take effect
curl -H "x-tyk-authorization: 352d20ee67be67f6340b4c0605b044b7" -s -H "Content-Type: application/json" http://localhost:8080/tyk/reload/ | python -mjson.tool

3. GET api to see if the above creation succeeded. Pick the Authorization token from hooks.go
curl -H "Authorization: d29e8f389a6cf39a72bc7156c5e710885e4b5b89" http://localhost:8080/test-api/get

4. Delete the API
curl -H "x-tyk-authorization: 352d20ee67be67f6340b4c0605b044b7"  -s  -H "Content-Type: application/json"  -X DELETE   http://localhost:8080/tyk/apis/1 | python -mjson.tool

5. Reload for the Delete to take effect
curl -H "x-tyk-authorization: 352d20ee67be67f6340b4c0605b044b7" -s -H "Content-Type: application/json" http://localhost:8080/tyk/reload/ | python -mjson.tool

6. GET api to see if the above delete succeeded
curl -H "Authorization: d29e8f389a6cf39a72bc7156c5e710885e4b5b89" http://localhost:8080/test-api/get

```

# Reads
https://tyk.io/docs/customise-tyk/plugins/rich-plugins/python/tutorial-add-demo-plugin-api/
root@elastic-stack:~/go/src/github.com/TykTechnologies/tyk# 

