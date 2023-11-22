package swagin

import "sync"

// SyncedLimitedInt is synchronized integer
type SyncedLimitedInt struct {
	mutex sync.RWMutex
	value int
	limit int
}

// NewLimitedInt
func NewLimitedInt(limit int) SyncedLimitedInt {
	return SyncedLimitedInt{limit: limit}
}

// Incr limited int or receive false
func (sli *SyncedLimitedInt) Incr() bool {
	sli.mutex.RLock()
	if sli.value >= sli.limit {
		sli.mutex.RUnlock()
		return false
	}
	sli.mutex.RUnlock()
	sli.mutex.Lock()
	defer sli.mutex.Unlock()
	if sli.value < sli.limit {
		sli.value++
		return true
	} else {
		return false
	}
}

// Decr limited int
func (sli *SyncedLimitedInt) Decr() {
	sli.mutex.Lock()
	defer sli.mutex.Unlock()
	sli.value--
}

// SetLimit
func (sli *SyncedLimitedInt) SetLimit(limit int) {
	sli.mutex.Lock()
	defer sli.mutex.Unlock()
	sli.limit = limit
}

// Get counter value
func (sli *SyncedLimitedInt) Get() int {
	sli.mutex.RLock()
	defer sli.mutex.RUnlock()
	return sli.value
}
