# Mootex

> A light-weight and central mutex source for distributed systems.

## Building

### The Binary

```bash
make
```

### The Image

```bash
docker build -t mootex .
```

## Execution

### Binary

```bash
./mootex
```

### Container

```bash
docker run -d --rm -p 3000 --name mootex mootex
```

## Usage

With any of the execution methods, you should now be able to use your central mutex in a few different ways.

### Locking

#### Query Parameter

```bash
curl http://localhost:3000/lock?key=MUTEX-NAME
```

or

```bash
curl http://localhost:3000/lockParameter?key=MUTEX-NAME
```

#### Plaintext Payload

```bash
curl -X POST -d 'MUTEX-NAME' http://localhost:3000/lockPlaintextBody
```

#### JSON Payload

```bash
curl -X POST -d '{"key":"MUTEX-NAME"}' http://localhost:3000/lockJSONBody
```

### Unlocking

#### Query Parameter

```bash
curl http://localhost:3000/lock?key=MUTEX-NAME
```

or

```bash
curl http://localhost:3000/lockParameter?key=MUTEX-NAME
```

#### Plaintext Payload

```bash
curl -X POST -d 'MUTEX-NAME' http://localhost:3000/lockPlaintextBody
```

#### JSON Payload

```bash
curl -X POST -d '{"key":"MUTEX-NAME"}' http://localhost:3000/lockJSONBody
```

### Response

When the lock has been activated, the http body will match `locked`.

When the lock has been deactivated, the http body will match `unlocked`.

Until activation or deactivation occurs, it will wait - so ensure your timeouts are set accordingly for your workload.
