/*
 * Copyright (c) 2000-2018, 达梦数据库有限公司.
 * All rights reserved.
 */
package dm

import (
	"container/list"
	"io"
)

type Dm_build_1549 struct {
	dm_build_1550 *list.List
	dm_build_1551 *dm_build_1603
	dm_build_1552 int
}

func Dm_build_1553() *Dm_build_1549 {
	return &Dm_build_1549{
		dm_build_1550: list.New(),
		dm_build_1552: 0,
	}
}

func (dm_build_1555 *Dm_build_1549) Dm_build_1554() int {
	return dm_build_1555.dm_build_1552
}

func (dm_build_1557 *Dm_build_1549) Dm_build_1556(dm_build_1558 *Dm_build_0, dm_build_1559 int) int {
	var dm_build_1560 = 0
	var dm_build_1561 = 0
	for dm_build_1560 < dm_build_1559 && dm_build_1557.dm_build_1551 != nil {
		dm_build_1561 = dm_build_1557.dm_build_1551.dm_build_1611(dm_build_1558, dm_build_1559-dm_build_1560)
		if dm_build_1557.dm_build_1551.dm_build_1606 == 0 {
			dm_build_1557.dm_build_1593()
		}
		dm_build_1560 += dm_build_1561
		dm_build_1557.dm_build_1552 -= dm_build_1561
	}
	return dm_build_1560
}

func (dm_build_1563 *Dm_build_1549) Dm_build_1562(dm_build_1564 []byte, dm_build_1565 int, dm_build_1566 int) int {
	var dm_build_1567 = 0
	var dm_build_1568 = 0
	for dm_build_1567 < dm_build_1566 && dm_build_1563.dm_build_1551 != nil {
		dm_build_1568 = dm_build_1563.dm_build_1551.dm_build_1615(dm_build_1564, dm_build_1565, dm_build_1566-dm_build_1567)
		if dm_build_1563.dm_build_1551.dm_build_1606 == 0 {
			dm_build_1563.dm_build_1593()
		}
		dm_build_1567 += dm_build_1568
		dm_build_1563.dm_build_1552 -= dm_build_1568
		dm_build_1565 += dm_build_1568
	}
	return dm_build_1567
}

func (dm_build_1570 *Dm_build_1549) Dm_build_1569(dm_build_1571 io.Writer, dm_build_1572 int) int {
	var dm_build_1573 = 0
	var dm_build_1574 = 0
	for dm_build_1573 < dm_build_1572 && dm_build_1570.dm_build_1551 != nil {
		dm_build_1574 = dm_build_1570.dm_build_1551.dm_build_1620(dm_build_1571, dm_build_1572-dm_build_1573)
		if dm_build_1570.dm_build_1551.dm_build_1606 == 0 {
			dm_build_1570.dm_build_1593()
		}
		dm_build_1573 += dm_build_1574
		dm_build_1570.dm_build_1552 -= dm_build_1574
	}
	return dm_build_1573
}

func (dm_build_1576 *Dm_build_1549) Dm_build_1575(dm_build_1577 []byte, dm_build_1578 int, dm_build_1579 int) {
	if dm_build_1579 == 0 {
		return
	}
	var dm_build_1580 = dm_build_1607(dm_build_1577, dm_build_1578, dm_build_1579)
	if dm_build_1576.dm_build_1551 == nil {
		dm_build_1576.dm_build_1551 = dm_build_1580
	} else {
		dm_build_1576.dm_build_1550.PushBack(dm_build_1580)
	}
	dm_build_1576.dm_build_1552 += dm_build_1579
}

func (dm_build_1582 *Dm_build_1549) dm_build_1581(dm_build_1583 int) byte {
	var dm_build_1584 = dm_build_1583
	var dm_build_1585 = dm_build_1582.dm_build_1551
	for dm_build_1584 > 0 && dm_build_1585 != nil {
		if dm_build_1585.dm_build_1606 == 0 {
			continue
		}
		if dm_build_1584 > dm_build_1585.dm_build_1606-1 {
			dm_build_1584 -= dm_build_1585.dm_build_1606
			dm_build_1585 = dm_build_1582.dm_build_1550.Front().Value.(*dm_build_1603)
		} else {
			break
		}
	}
	return dm_build_1585.dm_build_1624(dm_build_1584)
}
func (dm_build_1587 *Dm_build_1549) Dm_build_1586(dm_build_1588 *Dm_build_1549) {
	if dm_build_1588.dm_build_1552 == 0 {
		return
	}
	var dm_build_1589 = dm_build_1588.dm_build_1551
	for dm_build_1589 != nil {
		dm_build_1587.dm_build_1590(dm_build_1589)
		dm_build_1588.dm_build_1593()
		dm_build_1589 = dm_build_1588.dm_build_1551
	}
	dm_build_1588.dm_build_1552 = 0
}
func (dm_build_1591 *Dm_build_1549) dm_build_1590(dm_build_1592 *dm_build_1603) {
	if dm_build_1592.dm_build_1606 == 0 {
		return
	}
	if dm_build_1591.dm_build_1551 == nil {
		dm_build_1591.dm_build_1551 = dm_build_1592
	} else {
		dm_build_1591.dm_build_1550.PushBack(dm_build_1592)
	}
	dm_build_1591.dm_build_1552 += dm_build_1592.dm_build_1606
}

func (dm_build_1594 *Dm_build_1549) dm_build_1593() {
	var dm_build_1595 = dm_build_1594.dm_build_1550.Front()
	if dm_build_1595 == nil {
		dm_build_1594.dm_build_1551 = nil
	} else {
		dm_build_1594.dm_build_1551 = dm_build_1595.Value.(*dm_build_1603)
		dm_build_1594.dm_build_1550.Remove(dm_build_1595)
	}
}

func (dm_build_1597 *Dm_build_1549) Dm_build_1596() []byte {
	var dm_build_1598 = make([]byte, dm_build_1597.dm_build_1552)
	var dm_build_1599 = dm_build_1597.dm_build_1551
	var dm_build_1600 = 0
	var dm_build_1601 = len(dm_build_1598)
	var dm_build_1602 = 0
	for dm_build_1599 != nil {
		if dm_build_1599.dm_build_1606 > 0 {
			if dm_build_1601 > dm_build_1599.dm_build_1606 {
				dm_build_1602 = dm_build_1599.dm_build_1606
			} else {
				dm_build_1602 = dm_build_1601
			}
			copy(dm_build_1598[dm_build_1600:dm_build_1600+dm_build_1602], dm_build_1599.dm_build_1604[dm_build_1599.dm_build_1605:dm_build_1599.dm_build_1605+dm_build_1602])
			dm_build_1600 += dm_build_1602
			dm_build_1601 -= dm_build_1602
		}
		if dm_build_1597.dm_build_1550.Front() == nil {
			dm_build_1599 = nil
		} else {
			dm_build_1599 = dm_build_1597.dm_build_1550.Front().Value.(*dm_build_1603)
		}
	}
	return dm_build_1598
}

type dm_build_1603 struct {
	dm_build_1604 []byte
	dm_build_1605 int
	dm_build_1606 int
}

func dm_build_1607(dm_build_1608 []byte, dm_build_1609 int, dm_build_1610 int) *dm_build_1603 {
	return &dm_build_1603{
		dm_build_1608,
		dm_build_1609,
		dm_build_1610,
	}
}

func (dm_build_1612 *dm_build_1603) dm_build_1611(dm_build_1613 *Dm_build_0, dm_build_1614 int) int {
	if dm_build_1612.dm_build_1606 <= dm_build_1614 {
		dm_build_1614 = dm_build_1612.dm_build_1606
	}
	dm_build_1613.Dm_build_83(dm_build_1612.dm_build_1604[dm_build_1612.dm_build_1605 : dm_build_1612.dm_build_1605+dm_build_1614])
	dm_build_1612.dm_build_1605 += dm_build_1614
	dm_build_1612.dm_build_1606 -= dm_build_1614
	return dm_build_1614
}

func (dm_build_1616 *dm_build_1603) dm_build_1615(dm_build_1617 []byte, dm_build_1618 int, dm_build_1619 int) int {
	if dm_build_1616.dm_build_1606 <= dm_build_1619 {
		dm_build_1619 = dm_build_1616.dm_build_1606
	}
	copy(dm_build_1617[dm_build_1618:dm_build_1618+dm_build_1619], dm_build_1616.dm_build_1604[dm_build_1616.dm_build_1605:dm_build_1616.dm_build_1605+dm_build_1619])
	dm_build_1616.dm_build_1605 += dm_build_1619
	dm_build_1616.dm_build_1606 -= dm_build_1619
	return dm_build_1619
}

func (dm_build_1621 *dm_build_1603) dm_build_1620(dm_build_1622 io.Writer, dm_build_1623 int) int {
	if dm_build_1621.dm_build_1606 <= dm_build_1623 {
		dm_build_1623 = dm_build_1621.dm_build_1606
	}
	dm_build_1622.Write(dm_build_1621.dm_build_1604[dm_build_1621.dm_build_1605 : dm_build_1621.dm_build_1605+dm_build_1623])
	dm_build_1621.dm_build_1605 += dm_build_1623
	dm_build_1621.dm_build_1606 -= dm_build_1623
	return dm_build_1623
}
func (dm_build_1625 *dm_build_1603) dm_build_1624(dm_build_1626 int) byte {
	return dm_build_1625.dm_build_1604[dm_build_1625.dm_build_1605+dm_build_1626]
}
