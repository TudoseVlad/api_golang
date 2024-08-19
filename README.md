# api_golang
 Hello, this is a Golang API.
 
 It receives data from the Body when post is called where is sent a text message that is split into words and stored in memory.
 
 When someone wants to interogate the data, it should be called like this localhost.../cuvinte?words=word1_word2_.... . This is how you can find out how many times word1, word2 etc. appear in the stored data.
 
 The data sent after an GET is in a JSON format where it looks like this {"word1" : "how many times it appears,  "word2": how many times it appears ...}
 
 The data is saved locally in a file located in src/info where the data is dumped at the end of every life cycle or after 20 POSTs.
 
 Also every POST/GET need to have a basicAuthentication credentials which will be checked before every execution.
 
 The credentials are stored in src/info/credentials.json



 HOW TO RUN
 
 **go build**
 
 **./golang_api**

 HOW TO RUN UNIT TESTS
 
**go test -v**

**FUTURE IMPROVEMENTS**
- Creating a variable that decides if we use the HMTL or the gRPC protocol and implement the infrastructure necessary to handle the gRPC communication
- Creating a thread manager that remembers the sequence in which commands are given, in order to parallelize the tasks, for we can have mupltiple POSTs where which we can process in at the same time splitting the data and then sending it to the main thread where we will put all of it in the memory
- Enabling multiple ways of authentication, or only the basic one
- Creating protections against attacks:
  * restricting which users cand POST data and creating multiple tiers of user privileges
  * remembering the amount of requests one user has made in last seconds, limiting the amount of commands one user can do
  * Sanitizing the input to ensure that the format is respected, eliminating in the process any commands that may be given to interfere
  * Using services such as CloudFlare that after deployment would ensure DDOS protection and providing us with a web application firewall
