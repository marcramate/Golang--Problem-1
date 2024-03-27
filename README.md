# Golang Problem 1: Boss Baby's Revenge
### Description:
* English Version :
  * Boss Baby, known for his cool and clever ways, deals with teasing from the neighborhood kids who shoot water guns at his house. In response, Boss Baby seeks revenge by shooting at least one shot back, but only if the kids have already shot at him first. As Boss Baby's assistant
* Thai Version :
  *  บอสเบบี้มีความฉลาด แต่ได้มีเด็กในบริเวณบ้านมายิงปืนฉีดน้ำใส่บ้าน บอสเบบี้จึงทำการแก้แค้น โดย ยิงคืน 1 ครั้ง แต่จะยิงเมื่อเด็กยิงมาก่อน และมีเงื่อนไขการเป็น Goodboy คือบอสเบบี้ไม่ได้เริ่มยิงก่อนและทุกนัดที่โดนเด็กยิงได้แก้แค้นแล้ว แต่ถ้าไม่เข้าเงื่อนไขเหล่านี้ จะเป็น Badboy
## Use Language
 * Golang
## EX.
* Englist Version : 
   * Create Function  checkSR with 'shot' represents a shot and 'rev'  represents a revenge, It counting number shot and rev. If Boss Baby already shot at him fist will become 'Badboy' or Boss Baby doesn't get revenge in every shot. But if Boss Baby doesn't take him first and all shots have been revenged, it will be'Goodboy'
 * Thai Version :
   * สร้างฟังก์ชั่น checkSR โดยมี ตัวแปร shots แทน การยิงของเด็ก และ rev แทนการแก้แค้น และทำการนับจำนวนของการยิงและการแก้แค้น โดยถ้าบอสเบบี้เริ่มก่อนจะเป็น Badboyทันที หรือถ้า บอสเบบี้แก้แค้นไม่ครบทุกนัดก็จะเช่นเดียวกัน แต่ถ้าบอสเบบี้ไม่ได้เริ่มก่อนและแก้แค้นครบทุกนัดตามที่เด็กยิงมา จะเป็น Goodboy
## Result
```
  Input    |  Output  
 ------    | -------
SRSSRRR    | Good boy
RSSRR      | Bad boy
SSSRRRRS   | Bad boy
SRRSSR     | Bad boy
SSRSRRR    | Good boy
SRSRSSRSRRR| Good boy
```
>SRSSRRR :  Good boy
S Count : 3
R Count : 4

>RSSRR :  Bad boy
R fist : 0

>SSSRRRRS :  Bad boy
S Count : 4
R Count : 4

>SRRSSR :  Bad boy
S Count : 3
R Count : 3

>SSRSRRR : Good boy
S Count : 3
R Count : 4

>SRSRSSRSRRR : Good boy
S Count : 5
R Count : 6
