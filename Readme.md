## What is a tcp server?
 - It is a simple process that runs in a machine that 'listens' to a port.
 - Any machine who wants to talk to server has to connect over the port and establish the connection.
## Steps
    1. Start listening on the port
        - when process starts, pick a port and start listening to it
    2. Wait for a client to connect
        - Invoke the `Accept` system call and wait for a client to connect(blocking call)
        - This is a blocking call and server would not proceed until some client connects
    3. Read the request and send response
        once the connection is established
        - Invoke the `read` system call to read the request(blocking call)
        - Invoke the `write` system call to send the response(blocking call)
        - Close the connection
    4. Do this over and over again
        - Put these entire thing in an infinite loop
            - Wait for client to connect
            - Read the request
            - Send the response
            - Close the connection
    5. Handle multiple request concurrently
        - Till step-4, server processes requests sequentially.
        - To handle multiple requests concurrently
            - When client connects, fork a thread to process the request and respond
            - Let main thread come back to `accept` system call as soon as possible

## Drawback
    The drawback is when a large number of client connects to the server, it creates thread for each requests and causes therad overloading
## Improvements
    1. Limiting the number of threads - We can not just create thread for each requests. There should be some limit on the number of threads.
    2. Add thread pool to save on thread creation time - We can have a thread pool with specfic number of threads. For each requests, we can pick the thread from thread pool, process the request and put the thread back to thread pool
    3. Connection timeout - We can not let client to connect the server for infinite amount of time without making a request. Therefore, we add a timeout with each thread and kill the connection if time has passed.
    4. TCP backlog queue configuration - We can configure how many connections we want for the server. 