# MemHold

Uility for consuming and holding a certain amount of memory (for testing).


## Usage

```
memhold 1000
```
Will cause `memhold` to hold onto 1,000 MiB of memory until interrupted.

Can be used in kubernetes to fabricate / simulate an `OOMKilled` event:

```
kubectl apply -f deployment.yaml
```