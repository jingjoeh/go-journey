package orderprocessor

import (
	"context"
	"errors"

	"golang.org/x/sync/errgroup"
)

type Order struct {
	ID int
}

type Processor func(ctx context.Context, order Order) error

func ProcessOrders(
	ctx context.Context,
	orders []Order,
	workerCount int,
	process Processor,
) error {

	// ป้องกันการสร้าง worker pool ที่ไม่มี worker สำหรับรับงาน
	if workerCount <= 0 {
		return errors.New("workerCount must be greater than 0")
	}

	// ไม่มีงานให้ประมวลผล จึงจบได้ทันทีโดยไม่ต้องสร้าง goroutine
	if len(orders) == 0 {
		return nil
	}

	// jobs ใช้กระจาย Order ให้ worker แต่ละตัว โดยหนึ่ง Order จะถูกรับเพียงครั้งเดียว
	jobs := make(chan Order)

	// เมื่อ worker ใดคืน error, errgroup จะยกเลิก groupCtx เพื่อหยุดงานที่เหลือ
	g, groupCtx := errgroup.WithContext(ctx)

	// สร้าง worker ตามจำนวนที่กำหนด เพื่อจำกัดจำนวนงานที่ประมวลผลพร้อมกัน
	for range workerCount {
		g.Go(func() error {
			for {
				select {
				// หยุด worker เมื่อ parent context ถูกยกเลิกหรือ worker อื่นพบ error
				case <-groupCtx.Done():
					return groupCtx.Err()

				case order, ok := <-jobs:
					// jobs ถูกปิดและไม่มีงานเหลือแล้ว
					if !ok {
						return nil
					}

					// ส่ง error กลับให้ errgroup ซึ่งจะยกเลิก worker ตัวอื่นโดยอัตโนมัติ
					if err := process(groupCtx, order); err != nil {
						return err
					}
				}
			}
		})
	}

produceLoop:
	for _, order := range orders {
		select {
		// ส่งงานให้ worker เมื่อมี worker พร้อมรับ
		case jobs <- order:

		// หยุดป้อนงานทันทีเมื่อ context ถูกยกเลิก
		case <-groupCtx.Done():
			break produceLoop
		}
	}
	// ผู้ส่งเป็นผู้ปิด channel เพื่อแจ้ง worker ว่าไม่มีงานเพิ่มแล้ว
	close(jobs)

	// รอ worker ทุกตัวจบและคืน error แรกที่ errgroup ได้รับ
	return g.Wait()
}
