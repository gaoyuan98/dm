/*
 * Copyright (c) 2000-2018, 达梦数据库有限公司.
 * All rights reserved.
 */
package dm

import (
	"container/list"
	"io"
)

type Dm_build_1613 struct {
	dm_build_1614 *list.List
	dm_build_1615 *dm_build_1667
	dm_build_1616 int
}

func Dm_build_1617() *Dm_build_1613 {
	return &Dm_build_1613{
		dm_build_1614: list.New(),
		dm_build_1616: 0,
	}
}

func (dm_build_1619 *Dm_build_1613) Dm_build_1618() int {
	return dm_build_1619.dm_build_1616
}

func (dm_build_1621 *Dm_build_1613) Dm_build_1620(dm_build_1622 *Dm_build_0, dm_build_1623 int) int {
	var dm_build_1624 = 0
	var dm_build_1625 = 0
	for dm_build_1624 < dm_build_1623 && dm_build_1621.dm_build_1615 != nil {
		dm_build_1625 = dm_build_1621.dm_build_1615.dm_build_1675(dm_build_1622, dm_build_1623-dm_build_1624)
		if dm_build_1621.dm_build_1615.dm_build_1670 == 0 {
			dm_build_1621.dm_build_1657()
		}
		dm_build_1624 += dm_build_1625
		dm_build_1621.dm_build_1616 -= dm_build_1625
	}
	return dm_build_1624
}

func (dm_build_1627 *Dm_build_1613) Dm_build_1626(dm_build_1628 []byte, dm_build_1629 int, dm_build_1630 int) int {
	var dm_build_1631 = 0
	var dm_build_1632 = 0
	for dm_build_1631 < dm_build_1630 && dm_build_1627.dm_build_1615 != nil {
		dm_build_1632 = dm_build_1627.dm_build_1615.dm_build_1679(dm_build_1628, dm_build_1629, dm_build_1630-dm_build_1631)
		if dm_build_1627.dm_build_1615.dm_build_1670 == 0 {
			dm_build_1627.dm_build_1657()
		}
		dm_build_1631 += dm_build_1632
		dm_build_1627.dm_build_1616 -= dm_build_1632
		dm_build_1629 += dm_build_1632
	}
	return dm_build_1631
}

func (dm_build_1634 *Dm_build_1613) Dm_build_1633(dm_build_1635 io.Writer, dm_build_1636 int) int {
	var dm_build_1637 = 0
	var dm_build_1638 = 0
	for dm_build_1637 < dm_build_1636 && dm_build_1634.dm_build_1615 != nil {
		dm_build_1638 = dm_build_1634.dm_build_1615.dm_build_1684(dm_build_1635, dm_build_1636-dm_build_1637)
		if dm_build_1634.dm_build_1615.dm_build_1670 == 0 {
			dm_build_1634.dm_build_1657()
		}
		dm_build_1637 += dm_build_1638
		dm_build_1634.dm_build_1616 -= dm_build_1638
	}
	return dm_build_1637
}

func (dm_build_1640 *Dm_build_1613) Dm_build_1639(dm_build_1641 []byte, dm_build_1642 int, dm_build_1643 int) {
	if dm_build_1643 == 0 {
		return
	}
	var dm_build_1644 = dm_build_1671(dm_build_1641, dm_build_1642, dm_build_1643)
	if dm_build_1640.dm_build_1615 == nil {
		dm_build_1640.dm_build_1615 = dm_build_1644
	} else {
		dm_build_1640.dm_build_1614.PushBack(dm_build_1644)
	}
	dm_build_1640.dm_build_1616 += dm_build_1643
}

func (dm_build_1646 *Dm_build_1613) dm_build_1645(dm_build_1647 int) byte {
	var dm_build_1648 = dm_build_1647
	var dm_build_1649 = dm_build_1646.dm_build_1615
	for dm_build_1648 > 0 && dm_build_1649 != nil {
		if dm_build_1649.dm_build_1670 == 0 {
			continue
		}
		if dm_build_1648 > dm_build_1649.dm_build_1670-1 {
			dm_build_1648 -= dm_build_1649.dm_build_1670
			dm_build_1649 = dm_build_1646.dm_build_1614.Front().Value.(*dm_build_1667)
		} else {
			break
		}
	}
	return dm_build_1649.dm_build_1688(dm_build_1648)
}
func (dm_build_1651 *Dm_build_1613) Dm_build_1650(dm_build_1652 *Dm_build_1613) {
	if dm_build_1652.dm_build_1616 == 0 {
		return
	}
	var dm_build_1653 = dm_build_1652.dm_build_1615
	for dm_build_1653 != nil {
		dm_build_1651.dm_build_1654(dm_build_1653)
		dm_build_1652.dm_build_1657()
		dm_build_1653 = dm_build_1652.dm_build_1615
	}
	dm_build_1652.dm_build_1616 = 0
}
func (dm_build_1655 *Dm_build_1613) dm_build_1654(dm_build_1656 *dm_build_1667) {
	if dm_build_1656.dm_build_1670 == 0 {
		return
	}
	if dm_build_1655.dm_build_1615 == nil {
		dm_build_1655.dm_build_1615 = dm_build_1656
	} else {
		dm_build_1655.dm_build_1614.PushBack(dm_build_1656)
	}
	dm_build_1655.dm_build_1616 += dm_build_1656.dm_build_1670
}

func (dm_build_1658 *Dm_build_1613) dm_build_1657() {
	var dm_build_1659 = dm_build_1658.dm_build_1614.Front()
	if dm_build_1659 == nil {
		dm_build_1658.dm_build_1615 = nil
	} else {
		dm_build_1658.dm_build_1615 = dm_build_1659.Value.(*dm_build_1667)
		dm_build_1658.dm_build_1614.Remove(dm_build_1659)
	}
}

func (dm_build_1661 *Dm_build_1613) Dm_build_1660() []byte {
	var dm_build_1662 = make([]byte, dm_build_1661.dm_build_1616)
	var dm_build_1663 = dm_build_1661.dm_build_1615
	var dm_build_1664 = 0
	var dm_build_1665 = len(dm_build_1662)
	var dm_build_1666 = 0
	for dm_build_1663 != nil {
		if dm_build_1663.dm_build_1670 > 0 {
			if dm_build_1665 > dm_build_1663.dm_build_1670 {
				dm_build_1666 = dm_build_1663.dm_build_1670
			} else {
				dm_build_1666 = dm_build_1665
			}
			copy(dm_build_1662[dm_build_1664:dm_build_1664+dm_build_1666], dm_build_1663.dm_build_1668[dm_build_1663.dm_build_1669:dm_build_1663.dm_build_1669+dm_build_1666])
			dm_build_1664 += dm_build_1666
			dm_build_1665 -= dm_build_1666
		}
		if dm_build_1661.dm_build_1614.Front() == nil {
			dm_build_1663 = nil
		} else {
			dm_build_1663 = dm_build_1661.dm_build_1614.Front().Value.(*dm_build_1667)
		}
	}
	return dm_build_1662
}

type dm_build_1667 struct {
	dm_build_1668 []byte
	dm_build_1669 int
	dm_build_1670 int
}

func dm_build_1671(dm_build_1672 []byte, dm_build_1673 int, dm_build_1674 int) *dm_build_1667 {
	return &dm_build_1667{
		dm_build_1672,
		dm_build_1673,
		dm_build_1674,
	}
}

func (dm_build_1676 *dm_build_1667) dm_build_1675(dm_build_1677 *Dm_build_0, dm_build_1678 int) int {
	if dm_build_1676.dm_build_1670 <= dm_build_1678 {
		dm_build_1678 = dm_build_1676.dm_build_1670
	}
	dm_build_1677.Dm_build_83(dm_build_1676.dm_build_1668[dm_build_1676.dm_build_1669 : dm_build_1676.dm_build_1669+dm_build_1678])
	dm_build_1676.dm_build_1669 += dm_build_1678
	dm_build_1676.dm_build_1670 -= dm_build_1678
	return dm_build_1678
}

func (dm_build_1680 *dm_build_1667) dm_build_1679(dm_build_1681 []byte, dm_build_1682 int, dm_build_1683 int) int {
	if dm_build_1680.dm_build_1670 <= dm_build_1683 {
		dm_build_1683 = dm_build_1680.dm_build_1670
	}
	copy(dm_build_1681[dm_build_1682:dm_build_1682+dm_build_1683], dm_build_1680.dm_build_1668[dm_build_1680.dm_build_1669:dm_build_1680.dm_build_1669+dm_build_1683])
	dm_build_1680.dm_build_1669 += dm_build_1683
	dm_build_1680.dm_build_1670 -= dm_build_1683
	return dm_build_1683
}

func (dm_build_1685 *dm_build_1667) dm_build_1684(dm_build_1686 io.Writer, dm_build_1687 int) int {
	if dm_build_1685.dm_build_1670 <= dm_build_1687 {
		dm_build_1687 = dm_build_1685.dm_build_1670
	}
	dm_build_1686.Write(dm_build_1685.dm_build_1668[dm_build_1685.dm_build_1669 : dm_build_1685.dm_build_1669+dm_build_1687])
	dm_build_1685.dm_build_1669 += dm_build_1687
	dm_build_1685.dm_build_1670 -= dm_build_1687
	return dm_build_1687
}
func (dm_build_1689 *dm_build_1667) dm_build_1688(dm_build_1690 int) byte {
	return dm_build_1689.dm_build_1668[dm_build_1689.dm_build_1669+dm_build_1690]
}
