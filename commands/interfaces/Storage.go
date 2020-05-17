package interfaces

import "time"

type TemporaryStorageInterface interface {
	Get(key interface{}) (interface{}, bool)
	Set(key interface{}, value interface{}, duration time.Duration)
	Delete(key interface{})
}
