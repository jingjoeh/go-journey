package starter

import (
	"context"
	"errors"
	"net/http"
	"time"
)

// Server กำหนดความสามารถขั้นต่ำที่ Run ต้องใช้ ทำให้ทดสอบ lifecycle ของ server
// ได้โดยไม่ต้องเปิด HTTP server จริง
type Server interface {
	ListenAndServe() error
	Shutdown(context.Context) error
}

// Run เริ่ม server และรอจนกว่า server จะหยุดเอง หรือ ctx จะสั่งให้ปิดแบบ graceful
func Run(ctx context.Context, server Server, shutdownTimeout time.Duration) error {
	// ใช้ buffer ขนาดหนึ่งเพื่อให้ goroutine ส่งผลลัพธ์ได้แม้ Run กำลังคืนค่าจากอีก branch
	// จึงไม่ทิ้ง goroutine ที่ค้างอยู่กับการส่งเข้า channel
	errChan := make(chan error, 1)
	go func() {
		// ListenAndServe เป็น blocking call จึงต้องรันแยก เพื่อให้ Run ยังรับ cancellation ได้
		serverErr := server.ListenAndServe()
		errChan <- serverErr
	}()

	// รอ event แรกระหว่างคำสั่งยกเลิกกับการที่ server หยุดทำงาน
	select {
	case <-ctx.Done():
		// ctx เดิมถูกยกเลิกแล้ว จึงสร้าง context ใหม่สำหรับช่วง graceful shutdown
		// พร้อมกำหนด timeout เพื่อไม่ให้ shutdown ค้างโดยไม่มีกำหนด
		shutdownCtx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
		defer cancel()
		if err := server.Shutdown(shutdownCtx); err != nil {
			// ส่งต่อ shutdown error เพราะ server อาจปิด resource ไม่ครบภายในเวลาที่กำหนด
			return err
		}

		// หลัง Shutdown สำเร็จ รอให้ ListenAndServe สิ้นสุด เพื่อไม่ให้ Run คืนค่าก่อน server หยุดจริง
		err := <-errChan
		// http.Server คืน ErrServerClosed เมื่อถูกปิดตามปกติ จึงไม่ถือเป็น failure
		// errors.Is รองรับกรณีที่ error ถูก wrap มาก่อน
		if errors.Is(err, http.ErrServerClosed) {
			return nil
		}
		return err

	case err := <-errChan:
		// หาก server หยุดก่อน context ให้แยกการปิดตามปกติออกจาก startup/runtime error
		if errors.Is(err, http.ErrServerClosed) {
			return nil
		}
		return err
	}
}
