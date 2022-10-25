
# 1. 概念

原子操作 atomic 包
加锁操作涉及到内核态的上下文切换，比较耗时，代价高，
针对基本数据类型我们还可以使用原子操作来保证并发的安全，
因为原子操作是go语言提供的方法，我们在用户态就可以完成，因此性能比加锁操作更好
go语言的原子操作由内置的库 sync/atomic 完成

# 2. atomic包

| 方法 | 解释 | example |
| :-----| ----: |:-------:|
| func LoadInt32(addr int32) (val int32)
func LoadInt64(addr `int64) (val int64)<br>func LoadUint32(addruint32) (val uint32)<br>func LoadUint64(addruint64) (val uint64)<br>func LoadUintptr(addruintptr) (val uintptr)<br>func LoadPointer(addrunsafe.Pointer`) (val unsafe.Pointer) | 读取操作 |   单元格   |
| 单元格 | 单元格 |   单元格   |