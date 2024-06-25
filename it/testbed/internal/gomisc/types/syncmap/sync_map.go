package syncmap

import (
	"encoding/json"
	"sync"

	"github.com/NikolNikolaeva/project_weather/it/testbed/internal/gomisc/lang"
	arrays "github.com/NikolNikolaeva/project_weather/it/testbed/internal/gomisc/lang/array"
	"github.com/NikolNikolaeva/project_weather/it/testbed/internal/gomisc/lang/maps"
	"github.com/NikolNikolaeva/project_weather/it/testbed/internal/gomisc/types/entry"
	"github.com/NikolNikolaeva/project_weather/it/testbed/internal/gomisc/types/pair"
)

type Updater[V any] func(V, bool) (V, bool)

type SyncMap[K comparable, V any] interface {
	Keys() []K
	Values() []V
	OrElse(key K, def V) V
	Get(key K) (V, bool)
	Snapshot() map[K]V
	ContainsKey(key K) bool
	Entries() []entry.Entry[K, V]
	ForEach(consumer func(K, V)) SyncMap[K, V]
	IfPresent(key K, consumer func(V)) SyncMap[K, V]

	Delete(keys ...K) SyncMap[K, V]
	Put(key K, value V) SyncMap[K, V]
	Update(key K, updater Updater[V]) (V, bool)
	Clear(callback ...func(map[K]V)) SyncMap[K, V]
}

func New[K comparable, V any](initial ...map[K]V) SyncMap[K, V] {
	return &_SyncMap[K, V]{
		mutex: &sync.RWMutex{},
		inner: maps.Merge(initial...),
	}
}

type _SyncMap[K comparable, V any] struct {
	inner map[K]V
	mutex *sync.RWMutex
}

func (self *_SyncMap[K, V]) Keys() []K {
	return lang.DoWithLock(self.mutex.RLocker(), func() []K {
		return maps.Keys(self.inner)
	})
}

func (self *_SyncMap[K, V]) Values() []V {
	return lang.DoWithLock(self.mutex.RLocker(), func() []V {
		return maps.Values(self.inner)
	})
}

func (self *_SyncMap[K, V]) Snapshot() map[K]V {
	return lang.DoWithLock(self.mutex.RLocker(), func() map[K]V {
		return maps.Merge(self.inner) // merge creates a new map :))
	})
}

func (self *_SyncMap[K, V]) Get(key K) (V, bool) {
	return lang.DoWithLock(self.mutex.RLocker(), func() pair.Pair[V, bool] {
		return pair.NewPair(self.inner[key], maps.ContainsKey(self.inner, key))
	}).Unpack()
}

func (self *_SyncMap[K, V]) OrElse(key K, def V) V {
	return lang.DoWithLock(self.mutex.RLocker(), func() V {
		value, exists := self.inner[key]

		return lang.Ternary(exists, value, def)
	})

}

func (self *_SyncMap[K, V]) Entries() []entry.Entry[K, V] {
	return lang.DoWithLock(self.mutex.RLocker(), func() []entry.Entry[K, V] {
		return arrays.Map(maps.Keys(self.inner), func(_ int, key K) entry.Entry[K, V] {
			return entry.NewEntry[K, V](key, self.inner[key])
		})
	})
}

func (self *_SyncMap[K, V]) ContainsKey(key K) bool {
	return lang.DoWithLock(self.mutex.RLocker(), func() bool {
		return maps.ContainsKey(self.inner, key)
	})
}

func (self *_SyncMap[K, V]) ForEach(consumer func(K, V)) SyncMap[K, V] {
	return lang.DoWithLock(self.mutex.RLocker(), func() SyncMap[K, V] {
		for key, value := range self.inner {
			consumer(key, value)
		}

		return self
	})
}

func (self *_SyncMap[K, V]) IfPresent(key K, consumer func(V)) SyncMap[K, V] {
	return lang.DoWithLock(self.mutex.RLocker(), func() SyncMap[K, V] {
		if value, exists := self.inner[key]; exists {
			consumer(value)
		}

		return self
	})
}

func (self *_SyncMap[K, V]) MarshalJSON() ([]byte, error) {
	return lang.DoWithLock(self.mutex.RLocker(), func() pair.Pair[[]byte, error] {
		return pair.NewPair(json.Marshal(self.inner))
	}).Unpack()
}

func (self *_SyncMap[K, V]) Clear(callback ...func(map[K]V)) SyncMap[K, V] {
	return lang.DoWithLock(self.mutex, func() SyncMap[K, V] {
		defer append(callback, func(_ map[K]V) {})[0](self.inner)

		self.inner = map[K]V{}

		return self
	})
}

func (self *_SyncMap[K, V]) Delete(keys ...K) SyncMap[K, V] {
	return lang.DoWithLock(self.mutex, func() SyncMap[K, V] {
		for _, key := range keys {
			delete(self.inner, key)
		}

		return self
	})
}

func (self *_SyncMap[K, V]) Put(key K, value V) SyncMap[K, V] {
	return lang.DoWithLock(self.mutex, func() SyncMap[K, V] {
		self.inner[key] = value

		return self
	})
}

func (self *_SyncMap[K, V]) Update(key K, updater Updater[V]) (V, bool) {
	return lang.DoWithLock(self.mutex, func() pair.Pair[V, bool] {
		current, exists := self.inner[key]

		delete(self.inner, key)

		updated, ok := updater(current, exists)
		if ok {
			self.inner[key] = updated
		}

		return pair.NewPair(updated, ok)
	}).Unpack()
}

func (self *_SyncMap[K, V]) UnmarshalJSON(data []byte) error {
	return lang.DoWithLock(self.mutex, func() error {
		return json.Unmarshal(data, &self.inner)
	})
}
