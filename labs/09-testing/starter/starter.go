package starter

import (
	"context"
	"errors"
	"fmt"
	"strings"
)

var ErrInvalidEmail = errors.New("invalid email")

func NormalizeEmail(raw string) (string, error) {

	raw = strings.TrimSpace(raw)

	if raw == "" {
		return "", ErrInvalidEmail
	}

	parts := strings.Split(raw, "@")
	if len(parts) != 2 {
		return "", ErrInvalidEmail
	}
	if parts[0] == "" || parts[1] == "" {
		return "", ErrInvalidEmail
	}

	return fmt.Sprintf("%s@%s", parts[0], strings.ToLower(parts[1])), nil
}

/*
1. Process email ตามลำดับ

2. ก่อน process แต่ละ email
   → ถ้า ctx cancelled
   → return nil, ctx.Err()

3. NormalizeEmail แต่ละตัว

4. ถ้า NormalizeEmail error
   → return error โดยรักษา original cause
     เพื่อให้ errors.Is(err, ErrInvalidEmail) == true

5. ถ้าสำเร็จทั้งหมด
   → return normalized emails ตามลำดับเดิม

6. ห้ามสร้าง goroutine */

func ProcessEmails(
	ctx context.Context,
	rawEmails []string,
) ([]string, error) {

	results := make([]string, 0, len(rawEmails))
	for _, v := range rawEmails {
		if ctx.Err() != nil {
			return nil, ctx.Err()
		}

		nomalize, err := NormalizeEmail(v)
		if err != nil {
			return nil, fmt.Errorf("%s : %w", v, err)
		}
		results = append(results, nomalize)
	}

	return results, nil

}
