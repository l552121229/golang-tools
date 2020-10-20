package cache

import (
	"encoding/binary"
	"errors"
	"sync"
)

const (
	headerEntrySize = 4
	defaultValue    = 1024 // For this example we use 1024 like default value.
)

// Shard Shard存储
type Shard struct {
	items        map[uint64]uint32
	lock         sync.RWMutex
	array        []byte
	tail         int
	headerBuffer []byte
	holeChannel  chan *hole
	dataChannel  chan *data
}

type hole struct {
	startKey int
	endKey   int
}

type data struct {
	key      uint64
	startKey uint32
	endKey   uint32
}

func initNewShard() *Shard {
	return &Shard{
		items:        make(map[uint64]uint32, defaultValue),
		array:        make([]byte, defaultValue),
		tail:         1,
		headerBuffer: make([]byte, headerEntrySize),
	}
}

func (s *Shard) set(hashedKey uint64, entry []byte) {
	// w := wrapEntry(entry)
	s.lock.Lock()
	index := s.push(entry)
	s.items[hashedKey] = uint32(index)
	s.lock.Unlock()
}

func (s *Shard) push(data []byte) int {
	index := s.tail
	s.save(data, len(data))
	return index
}

func (s *Shard) save(data []byte, len int) {
	// Put in the first 4 bytes the size of the value
	binary.LittleEndian.PutUint32(s.headerBuffer, uint32(len))
	s.copy(s.headerBuffer, headerEntrySize)
	s.copy(data, len)
}

func (s *Shard) copy(data []byte, len int) {
	// Using the tail to keep the order to write in the array
	s.tail += copy(s.array[s.tail:], data[:len])
}

func (s *Shard) get(key string, hashedKey uint64) ([]byte, error) {
	s.lock.RLock()
	itemIndex := int(s.items[hashedKey])
	if itemIndex == 0 {
		s.lock.RUnlock()
		return nil, errors.New("This key not found")
	}

	// Read the first 4 bytes after the index, remember these 4 bytes have the size of the value, so
	// you can use this to get the size and get the value in the array using index+blockSize to know until what point
	// you need to read
	blockSize := int(binary.LittleEndian.Uint32(s.array[itemIndex : itemIndex+headerEntrySize]))
	entry := s.array[itemIndex+headerEntrySize : itemIndex+headerEntrySize+blockSize]
	s.lock.RUnlock()
	return entry, nil
}

// del 删除缓存
func (s *Shard) del(hashedKey uint64) bool {
	s.lock.Lock()
	if s.items[hashedKey] == 0 {
		s.lock.Unlock()
		return true
	}
	s.items[hashedKey] = 0
	s.lock.Unlock()
	return true
}

// getDataNoHole 获取没有空穴的数据
func (s *Shard) getDataNoHole() {
	var endKey uint32
	for key, itemIndex := range s.items {
		if itemIndex == 0 {
			continue
		}
		endKey = itemIndex + headerEntrySize + binary.LittleEndian.Uint32(s.array[itemIndex:itemIndex+headerEntrySize])

		s.dataChannel <- &data{key, itemIndex, endKey}
	}
}

// CheckHoleProportion 获取空穴所占比例
func (s *Shard) CheckHoleProportion() float64 {
	var nextKey, holeLenght uint32
	for _, itemIndex := range s.items {
		if itemIndex == 0 {
			continue
		}
		if nextKey < itemIndex {
			holeLenght += itemIndex - nextKey
		}
		nextKey = itemIndex + headerEntrySize + binary.LittleEndian.Uint32(s.array[itemIndex:itemIndex+headerEntrySize])
	}
	allLenght := len(s.array)
	return float64(holeLenght) / float64(allLenght)
}

// CleanHole 清除空穴
func (s *Shard) CleanHole() bool {
	// 写锁, 整理空穴期间禁止写入操作
	s.lock.Lock()
	defer s.lock.Unlock()

	go s.getDataNoHole()
	temp := make([]byte, defaultValue)
	items := make(map[uint64]uint32, defaultValue)
	for data := range s.dataChannel {
		temp = append(temp, s.array[data.startKey:data.endKey]...)
		items[data.key] = data.startKey
	}

	s.lock.RLock() // 读锁, 转移数据期间禁止读取操作
	defer s.lock.RUnlock()

	s.items = items
	copy(s.array, temp)

	return true
}

func readEntry(data []byte) []byte {
	dst := make([]byte, len(data))
	copy(dst, data)

	return dst
}

func wrapEntry(entry []byte) []byte {
	// You can put more information like a timestamp if you want.
	blob := make([]byte, len(entry))
	copy(blob, entry)
	return blob
}
