# BitBurst

**This task is to write a small MVP ledger service that manages user balances**

To access this code 

try go get 

Expose endpoints run `go get github.com/gopalrg310/bitburst` in terminal.

**In terminal 1:**

if docker not available install the docker setup in machine
run below command
`docker-compose up`

**In terminal 2:**
`cd postdeploymenttesting`

`go test -v`

Able to see below result

```
MacBook-Pro:postdeploymenttesting gravip214$ go test -v
=== RUN   Test_Postiveflow
Test passed:  {"Message":"Transaction of $500.000000 added for user gopal"}  Status:  200
Test passed:  {"Message":"User gopal has a balance of $500.000000"}  Status:  200
Test passed:  [{"UserID":"gopal","Amount":500,"Timestamp":"2023-04-18T17:03:00.309113Z"}]  Status:  200
--- PASS: Test_Postiveflow (0.02s)
=== RUN   Test_Negativeflow1
Test passed:    Status:  405
Test passed:    Status:  405
Test passed:    Status:  405
--- PASS: Test_Negativeflow1 (0.00s)
=== RUN   Test_Negativeflow2
Test passed:  {"Message":"Amount must be positive"}  Status:  500
--- PASS: Test_Negativeflow2 (0.00s)
=== RUN   Test_Negativeflow3
Test passed:  {"Message":"User unknown has a balance of $0.000000"}  Status:  200
--- PASS: Test_Negativeflow3 (0.00s)
=== RUN   Test_Negativeflow4
Test passed:  {}  Status:  200
--- PASS: Test_Negativeflow4 (0.00s)
PASS
ok      github.com/gopalrg310/bitburst/postdeploymenttesting    0.700s
```

/users/{uid}/add → takes a positive amount in USD to be added to the users account
/users/{uid}/balance → that shows how much money the user currently has 
/users/{uid}/history?page= → exposes the full transaction history of the user with optional parameter to send pagenumber each page have 10.


How could the service be developed to handle thousands of concurrent users with hundreds of transactions each?
Asynchronous processing?

* Optimize database interactions

* Use caching such as Redis or Memcached.

* Scalable architecture.

* Implement rate limiting.

* Monitor and optimize performance.

* Use fault-tolerant techniques.

* Plan for scalability.

What has to be paid attention to if we assume that the service is going to run in multiple instances for high availability?

* Shared data consistency such as distributed transactions or event-driven architectures.

* Load balancing among all instances.

* Session management.

* Distributed caching like Redis or Memcached that support distributed caching.

* Cross-instance communication.

* Health monitoring and failover.

* Scalable infrastructure with scalability and resilience.

* Testing and deployment.

* Disaster recovery.

How does the the add endpoint have to be designed, if the caller cannot guarantee that it will call exactely once for the same money transfer?
Request Deduplication?

* Idempotency Tokens.

* Database Transactions.

* Idempotent API Design.

* Proper Error Handling.
