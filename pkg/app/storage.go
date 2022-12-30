package app

import (
	"encoding/json"
	"sync"

	"github.com/maxence-charriere/go-app/v9/pkg/errors"
)

// BrowserStorage is the interface that describes a web browser storage.
type BrowserStorage interface {
	// Set sets the value to the given key. The value must be json convertible.
	Set(k string, v any) error

	// Get gets the item associated to the given key and store it in the given
	// value.
	// It returns an error if v is not a pointer.
	Get(k string, v any) error

	// Del deletes the item associated with the given key.
	Del(k string)

	// Len returns the number of items stored.
	Len() int

	// Key returns the key of the item associated to the given index.
	Key(i int) (string, error)

	// Clear deletes all items.
	Clear()
}

type MemoryStorage struct {
	mu   sync.RWMutex
	data map[string][]byte
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		data: make(map[string][]byte),
	}
}

func (s *MemoryStorage) Set(k string, v any) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	s.mu.Lock()
	s.data[k] = b
	s.mu.Unlock()
	return nil
}

func (s *MemoryStorage) Get(k string, v any) error {
	s.mu.RLock()
	d, ok := s.data[k]
	if !ok {
		s.mu.RUnlock()
		return nil
	}

	s.mu.RUnlock()
	return json.Unmarshal(d, v)
}

func (s *MemoryStorage) Del(k string) {
	s.mu.Lock()
	delete(s.data, k)
	s.mu.Unlock()
}

func (s *MemoryStorage) Clear() {
	s.mu.Lock()
	for k := range s.data {
		delete(s.data, k)
	}
	s.mu.Unlock()
}

func (s *MemoryStorage) Len() int {
	s.mu.RLock()
	l := len(s.data)
	s.mu.RUnlock()
	return l
}

func (s *MemoryStorage) Key(i int) (string, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	j := 0
	for k := range s.data {
		if i == j {
			return k, nil
		}
		j++
	}

	return "", errors.New("index out of range").
		Tag("index", i).
		Tag("len", s.Len())
}

type JsStorage struct {
	name  string
	mutex sync.RWMutex
}

func NewJSStorage(name string) *JsStorage {
	return &JsStorage{name: name}
}

func (s *JsStorage) Set(k string, v any) (err error) {
	defer func() {
		r := recover()
		if r != nil {
			err = errors.New("setting storage value failed").
				Tag("storage-type", s.name).
				Tag("key", k).
				Wrap(r.(error))
		}
	}()

	s.mutex.Lock()
	defer s.mutex.Unlock()

	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	Window().Get(s.name).Call("setItem", k, string(b))
	return nil
}

func (s *JsStorage) Get(k string, v any) error {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	item := Window().Get(s.name).Call("getItem", k)
	if !item.Truthy() {
		return nil
	}

	return json.Unmarshal([]byte(item.String()), v)
}

func (s *JsStorage) Del(k string) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	Window().Get(s.name).Call("removeItem", k)
}

func (s *JsStorage) Clear() {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	Window().Get(s.name).Call("clear")
}

func (s *JsStorage) Len() int {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	return s.len()
}

func (s *JsStorage) len() int {
	return Window().Get(s.name).Get("length").Int()
}

func (s *JsStorage) Key(i int) (string, error) {
	if l := s.len(); i < 0 || i >= l {
		return "", errors.New("index out of range").
			Tag("index", i).
			Tag("len", l)
	}

	return Window().Get(s.name).Call("key", i).String(), nil
}
