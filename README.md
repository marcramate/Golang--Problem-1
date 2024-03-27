<!-- func checkSR(seq string) string {
	shots := 0
	rev := 0
	lastShot := 'R'

	for _, act := range seq {
		switch act {
		case 'S':
			shots++
			lastShot = 'S'
		case 'R':
			if shots == 0 {
				return "Bad boy"
			}
			rev++
			lastShot = 'R'
		}
	}

	if rev <= shots || lastShot != 'R' {
		return "Bad boy"
	}

	return "Good boy"
} -->


1.func checkRevenge(sequence string) string: กำหนดฟังก์ชั่นชื่อ checkRevenge ที่รับ string (sequence) เป็น input และ return string ที่บ่งบอกถึง "Good boy" หรือ "Bad boy"

2.shots := 0: inititalize ตัวแปร shots ไว้เก็บจำนวนการยิงทั้งหมดของเด็กๆ ในละแวก

3.revenge := 0: initialize ตัวแปร revenge ไว้เก็บจำนวนการแก้แค้นของ Boss Baby

4.lastShot := 'R': initialize ตัวแปร lastShot ด้วยตัวอักษร 'R'

5.for _, action := range sequence: วนซ้ำแต่ละตัวอักษร (action) ใน string sequence โดยใช้ range loop

6.switch action: ใช้ switch statement แยกแยะ action ต่างๆ

6.1. case 'S' เพิ่ม shots เปลี่ยน lastShot เป็น 'S'

6.2 case 'R'  ตรวจสอบ shots ว่าเท่ากับ 0 หรือไม่ (Boss Baby ยิงก่อน) ถ้า shots == 0 return "Bad boy" เพิ่ม revenge เปลี่ยน lastShot เป็น 'R'

7.if revenge <= shots || lastShot != 'R'

7.1.revenge <= shots: ตรวจสอบว่าจำนวนการแก้แค้นน้อยกว่าหรือเท่ากับจำนวนการยิง จำนวนการแก้แค้นน้อยกว่าจำนวนการยิง return "Bad boy" จำนวนการแก้แค้นเท่ากับจำนวนการยิง

7.2 lastShot != 'R': ตรวจสอบ lastShot ว่าไม่ใช่ 'R' return "Bad boy"

8. 