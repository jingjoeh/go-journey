package starter

import (
	"context"
	"sync"
)

type call struct {
	// done จะถูกปิดเมื่อการโหลดเสร็จ เพื่อปลุก goroutine ทุกตัวที่รอ key เดียวกัน
	done  chan struct{}
	value string
	err   error
}

// Cache เก็บค่าที่โหลดสำเร็จ และรวม request ที่ขอ key เดียวกันในเวลาเดียวกัน
// ให้เหลือการเรียก load เพียงครั้งเดียว
type Cache struct {
	// mu ป้องกันการเข้าถึง values และ inflight พร้อมกันจากหลาย goroutine
	mu     sync.Mutex
	values map[string]string
	// inflight เก็บงานที่กำลังโหลด โดยใช้ key เป็นตัวระบุงาน
	inflight map[string]*call
}

// New สร้าง Cache พร้อม map ที่พร้อมใช้งาน เพื่อไม่ให้เกิด panic ตอนเขียนค่า
func New() *Cache { return &Cache{values: make(map[string]string), inflight: make(map[string]*call)} }

// Get คืนค่าจาก cache หากมีอยู่ มิฉะนั้นจะโหลดค่าผ่าน load
// request ที่ขอ key เดียวกันระหว่างโหลดจะรอและใช้ผลลัพธ์ร่วมกัน
func (c *Cache) Get(ctx context.Context, key string, load func(context.Context) (string, error)) (string, error) {
	// Lock เปิด critical section แรก เพื่อให้การตรวจ values และ inflight เป็นการตัดสินใจชุดเดียวกัน
	// ถ้าไม่ล็อก goroutine สองตัวอาจเห็นว่าไม่มีทั้งค่าและงาน แล้วเรียก load ซ้ำพร้อมกันได้
	// ฟังก์ชันนี้ไม่ใช้ defer Unlock เพราะต้องปลด lock ระหว่างรอและระหว่างเรียก load
	c.mu.Lock()
	// ตรวจ cache ภายใต้ lock เพราะ Go map ไม่ปลอดภัยสำหรับการใช้งานพร้อมกัน
	if value, ok := c.values[key]; ok {
		// เมื่อ copy value ออกจาก map แล้ว ไม่จำเป็นต้องถือ lock ระหว่าง return
		c.mu.Unlock()
		return value, nil
	}

	// ถ้ามี goroutine อื่นกำลังโหลด key นี้อยู่ ให้รอผลเดิมแทนการโหลดซ้ำ
	if current, ok := c.inflight[key]; ok {
		// ต้องปลด lock ก่อนรอ ไม่เช่นนั้นเจ้าของงานจะเข้ามาบันทึกผลและปิด done ไม่ได้
		c.mu.Unlock()
		select {
		case <-ctx.Done():
			// ผู้รอสามารถยกเลิกการรอของตัวเองได้ โดยไม่ยกเลิกงานที่กำลังโหลด
			return "", ctx.Err()

		case <-current.done:
			// การปิด done เป็นสัญญาณว่า value และ err พร้อมให้อ่านแล้ว
			return current.value, current.err
		}
	}

	// request แรกของ key นี้เป็นเจ้าของงานโหลด และประกาศงานไว้ให้ request อื่นมารอ
	work := &call{done: make(chan struct{})}
	c.inflight[key] = work
	// ปลด lock หลังลงทะเบียน inflight สำเร็จ จากนี้ request ของ key เดียวกันจะเห็น work นี้และรอ
	// ส่วน request ของ key อื่นสามารถเข้าถึง cache ได้ระหว่างที่ load กำลังทำงาน
	c.mu.Unlock()

	// ไม่ถือ lock ระหว่าง load เพราะอาจเป็นงานช้า และ key อื่นควรทำงานต่อได้
	work.value, work.err = load(ctx)

	// Lock อีกครั้งก่อนเปลี่ยน shared state เพื่อไม่ให้เขียน values/inflight แข่งกับ goroutine อื่น
	// critical section นี้ทำให้การบันทึกผลและการนำงานออกเกิดเป็นลำดับเดียวกัน
	c.mu.Lock()
	// เก็บเฉพาะผลสำเร็จ เพื่อให้ความผิดพลาดชั่วคราวสามารถลองใหม่ใน request ถัดไป
	if work.err == nil {
		c.values[key] = work.value
	}

	// นำงานออกก่อนปิด done เพื่อให้สถานะภายในสอดคล้องกันเมื่อผู้รอทำงานต่อ
	delete(c.inflight, key)
	// close ปลุกผู้รอทุกตัว และรับประกันว่าผู้รอจะเห็น value/err ที่เขียนก่อนหน้านี้
	close(work.done)

	// shared state ถูกอัปเดตครบแล้ว จึงปล่อยให้ goroutine อื่นตรวจหรือแก้ไข map ต่อได้
	c.mu.Unlock()
	return work.value, work.err
}
