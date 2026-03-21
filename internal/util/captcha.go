package util

import (
	"math/rand"
	"strconv"
	"sync"
	"time"
)

var CaptchaStoreHandler = NewCaptchaStore(10240, 5*time.Minute)

type CaptchaStore struct {
	store map[string]string
	ttl   time.Duration
	mutex sync.RWMutex
}

func NewCaptchaStore(size int, ttl time.Duration) *CaptchaStore {
	return &CaptchaStore{
		store: make(map[string]string, size),
		ttl:   ttl,
	}
}

func (s *CaptchaStore) Set(id, value string) {
	s.mutex.Lock()
	s.store[id] = value
	s.mutex.Unlock()
	time.AfterFunc(s.ttl, func() {
		s.mutex.Lock()
		delete(s.store, id)
		s.mutex.Unlock()
	})
}

func (s *CaptchaStore) Get(id string, clear bool) string {
	s.mutex.RLock()
	value, ok := s.store[id]
	if !ok {
		s.mutex.RUnlock()
		return ""
	}
	if clear {
		s.mutex.RUnlock()
		s.mutex.Lock()
		delete(s.store, id)
		s.mutex.Unlock()
		return value
	}
	s.mutex.RUnlock()
	return value
}

func (s *CaptchaStore) Verify(id, value string, clear bool) bool {
	stored := s.Get(id, clear)
	return stored == value
}

// 使用全局随机源
var randomSource = rand.NewSource(time.Now().UnixNano())
var random = rand.New(randomSource)

func GenerateCaptcha() (string, string) {
	digits := "0123456789"
	result := make([]byte, 4)
	for i := range result {
		result[i] = digits[random.Intn(len(digits))]
	}
	code := string(result)
	id := strconv.FormatInt(time.Now().UnixNano()+int64(random.Int63()), 36)
	return id, code
}

func VerifyCaptcha(id, input string) bool {
	return CaptchaStoreHandler.Verify(id, input, true)
}
