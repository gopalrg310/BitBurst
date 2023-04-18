# BitBurst

**This task is to write a small MVP ledger service that manages user balances**

Expose endpoints

/users/{uid}/add → takes a positive amount in USD to be added to the users account
/users/{uid}/balance → that shows how much money the user currently has 
/users/{uid}/history?page= → exposes the full transaction history of the user with optional parameter to send pagenumber each page have 10.


How could the service be developed to handle thousands of concurrent users with hundreds of transactions each?
Asynchronous processing.

Optimize database interactions

Use caching such as Redis or Memcached.

Scalable architecture.

Implement rate limiting.

Monitor and optimize performance.

Use fault-tolerant techniques.

Plan for scalability.

What has to be paid attention to if we assume that the service is going to run in multiple instances for high availability?

Shared data consistency such as distributed transactions or event-driven architectures.

Load balancing among all instances.

Session management.

Distributed caching like Redis or Memcached that support distributed caching.

Cross-instance communication.

Health monitoring and failover.

Scalable infrastructure with scalability and resilience.

Testing and deployment.

Disaster recovery.

How does the the add endpoint have to be designed, if the caller cannot guarantee that it will call exactely once for the same money transfer?
Request Deduplication.

Idempotency Tokens.

Database Transactions.

Idempotent API Design.

Proper Error Handling.