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

 zis de posibile metode de a preveni atacuri etc.
