# quic-test
# QUIC / HTTP/3 Test Artifacts

Collection of the tools and scripts I used to test QUIC. 

* **api**: Simple REST API written in GO, It generates a random response. I wanted the response to be roughly 100KB. You can adjust the size of the response by changing the loop count and text length in main.go
* **curl**: script to update curl on your environment to 7.67.0, to support http3. And a file to format the output of the curl resposne to display info
* **nginx**: nginx conf file, You need nginx 1.16.1 for this, you'll need to clone the nginx repo and build it. a guide can be found here [https://blog.cloudflare.com/experiment-with-http-3-using-nginx-and-quiche/]
* **load test**: script to generate a load test using curl and append the output to a file
