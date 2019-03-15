package goserbench

import (
	"io"
	"time"
	"unsafe"
)

var (
	_ = unsafe.Sizeof(0)
	_ = io.ReadFull
	_ = time.Now()
)

type GencodeUser struct {
	Id       string
	Name     string
	Password string
	Age      int32
	BirthDay int64
	Spouse   bool
	Money    float64
}

func (d *GencodeUser) Size() (s uint64) {

	{
		l := uint64(len(d.Id))

		{

			t := l
			for t >= 0x80 {
				t >>= 7
				s++
			}
			s++

		}
		s += l
	}
	{
		l := uint64(len(d.Name))

		{

			t := l
			for t >= 0x80 {
				t >>= 7
				s++
			}
			s++

		}
		s += l
	}
	{
		l := uint64(len(d.Password))

		{

			t := l
			for t >= 0x80 {
				t >>= 7
				s++
			}
			s++

		}
		s += l
	}
	s += 21
	return
}
func (d *GencodeUser) Marshal(buf []byte) ([]byte, error) {
	size := d.Size()
	{
		if uint64(cap(buf)) >= size {
			buf = buf[:size]
		} else {
			buf = make([]byte, size)
		}
	}
	i := uint64(0)

	{
		l := uint64(len(d.Id))

		{

			t := uint64(l)

			for t >= 0x80 {
				buf[i+0] = byte(t) | 0x80
				t >>= 7
				i++
			}
			buf[i+0] = byte(t)
			i++

		}
		copy(buf[i+0:], d.Id)
		i += l
	}
	{
		l := uint64(len(d.Name))

		{

			t := uint64(l)

			for t >= 0x80 {
				buf[i+0] = byte(t) | 0x80
				t >>= 7
				i++
			}
			buf[i+0] = byte(t)
			i++

		}
		copy(buf[i+0:], d.Name)
		i += l
	}
	{
		l := uint64(len(d.Password))

		{

			t := uint64(l)

			for t >= 0x80 {
				buf[i+0] = byte(t) | 0x80
				t >>= 7
				i++
			}
			buf[i+0] = byte(t)
			i++

		}
		copy(buf[i+0:], d.Password)
		i += l
	}
	{

		buf[i+0+0] = byte(d.Age >> 0)

		buf[i+1+0] = byte(d.Age >> 8)

		buf[i+2+0] = byte(d.Age >> 16)

		buf[i+3+0] = byte(d.Age >> 24)

	}
	{

		buf[i+0+4] = byte(d.BirthDay >> 0)

		buf[i+1+4] = byte(d.BirthDay >> 8)

		buf[i+2+4] = byte(d.BirthDay >> 16)

		buf[i+3+4] = byte(d.BirthDay >> 24)

		buf[i+4+4] = byte(d.BirthDay >> 32)

		buf[i+5+4] = byte(d.BirthDay >> 40)

		buf[i+6+4] = byte(d.BirthDay >> 48)

		buf[i+7+4] = byte(d.BirthDay >> 56)

	}
	{
		if d.Spouse {
			buf[i+12] = 1
		} else {
			buf[i+12] = 0
		}
	}
	{

		v := *(*uint64)(unsafe.Pointer(&(d.Money)))

		buf[i+0+13] = byte(v >> 0)

		buf[i+1+13] = byte(v >> 8)

		buf[i+2+13] = byte(v >> 16)

		buf[i+3+13] = byte(v >> 24)

		buf[i+4+13] = byte(v >> 32)

		buf[i+5+13] = byte(v >> 40)

		buf[i+6+13] = byte(v >> 48)

		buf[i+7+13] = byte(v >> 56)

	}
	return buf[:i+21], nil
}

func (d *GencodeUser) Unmarshal(buf []byte) (uint64, error) {
	i := uint64(0)

	{
		l := uint64(0)

		{

			bs := uint8(7)
			t := uint64(buf[i+0] & 0x7F)
			for buf[i+0]&0x80 == 0x80 {
				i++
				t |= uint64(buf[i+0]&0x7F) << bs
				bs += 7
			}
			i++

			l = t

		}
		d.Id = string(buf[i+0 : i+0+l])
		i += l
	}
	{
		l := uint64(0)

		{

			bs := uint8(7)
			t := uint64(buf[i+0] & 0x7F)
			for buf[i+0]&0x80 == 0x80 {
				i++
				t |= uint64(buf[i+0]&0x7F) << bs
				bs += 7
			}
			i++

			l = t

		}
		d.Name = string(buf[i+0 : i+0+l])
		i += l
	}
	{
		l := uint64(0)

		{

			bs := uint8(7)
			t := uint64(buf[i+0] & 0x7F)
			for buf[i+0]&0x80 == 0x80 {
				i++
				t |= uint64(buf[i+0]&0x7F) << bs
				bs += 7
			}
			i++

			l = t

		}
		d.Password = string(buf[i+0 : i+0+l])
		i += l
	}
	{

		d.Age = 0 | (int32(buf[i+0+0]) << 0) | (int32(buf[i+1+0]) << 8) | (int32(buf[i+2+0]) << 16) | (int32(buf[i+3+0]) << 24)

	}
	{

		d.BirthDay = 0 | (int64(buf[i+0+4]) << 0) | (int64(buf[i+1+4]) << 8) | (int64(buf[i+2+4]) << 16) | (int64(buf[i+3+4]) << 24) | (int64(buf[i+4+4]) << 32) | (int64(buf[i+5+4]) << 40) | (int64(buf[i+6+4]) << 48) | (int64(buf[i+7+4]) << 56)

	}
	{
		d.Spouse = buf[i+12] == 1
	}
	{

		v := 0 | (uint64(buf[i+0+13]) << 0) | (uint64(buf[i+1+13]) << 8) | (uint64(buf[i+2+13]) << 16) | (uint64(buf[i+3+13]) << 24) | (uint64(buf[i+4+13]) << 32) | (uint64(buf[i+5+13]) << 40) | (uint64(buf[i+6+13]) << 48) | (uint64(buf[i+7+13]) << 56)
		d.Money = *(*float64)(unsafe.Pointer(&v))

	}
	return i + 21, nil
}
