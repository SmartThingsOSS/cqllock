cqllock
=======

Command-line tools to create and delete distributed locks using Cassandra and lightweight transactions. These tools are useful for managing distributed jobs on multiple hosts.

Usage
-----

Set up a config file at `~/.cqllockrc` in the following YAML format:

```
seeds:
- cassandra-seed1.domain.com
- optional-cassandra-seed2.domain.com
- etc
keyspace: keyspace_containing_lock_table
table: lock_table_name
username: optional-username-if-using-auth
password: optional-password-if-using-auth
certpath: /path/to/optional/client/cert/if/using/SSL.pem
keypath: /path/to/optional/client/key/if/using/SSL.key
```

### Possible scenarios

##### Ensure a process only runs once ever, regardless of how many times or on what host it's attempted:
  ```
  cqllock ProcessLock && ./myProcess.sh
  ```

  The first call to cqllock that successfully acquires the lock will hold it forever. Further calls to cqllock will return non-zero, so ./myProcess.sh won't run again.
  This command can be run on multiple hosts.

##### Multiple hosts run a job, but only allow one at a time:
  ```
  cqllock -t 1d -r 5m jobLock && ./myJob.sh ; cqlunlock jobLock
  ```

  The first host to acquire the lock will hold it until ./myJob.sh finishes. Other hosts will attempt to acquire the lock every five minutes and will wait up to one day
  trying to reaquire the lock.

There are likely other useful scenarios.
