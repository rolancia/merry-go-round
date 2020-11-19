# merry-go-round
The simplest synchronized, sized resource pool.

###example

---
```
pool := merry_go_round.NewPool(func() interface{} {
	return 1 // returns resource you need
}, 64)

rs := pool.Get() // get resource
pool.Put(rs) // release resource
```