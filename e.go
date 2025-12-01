/*
 * Copyright (c) 2000-2018, 达梦数据库有限公司.
 * All rights reserved.
 */
package dm

import (
	"bytes"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/ianaindex"
	"golang.org/x/text/transform"
	"io"
	"io/ioutil"
	"math"
)

type dm_build_1354 struct{}

var Dm_build_1355 = &dm_build_1354{}

func (Dm_build_1357 *dm_build_1354) Dm_build_1356(dm_build_1358 []byte, dm_build_1359 int, dm_build_1360 byte) int {
	dm_build_1358[dm_build_1359] = dm_build_1360
	return 1
}

func (Dm_build_1362 *dm_build_1354) Dm_build_1361(dm_build_1363 []byte, dm_build_1364 int, dm_build_1365 int8) int {
	dm_build_1363[dm_build_1364] = byte(dm_build_1365)
	return 1
}

func (Dm_build_1367 *dm_build_1354) Dm_build_1366(dm_build_1368 []byte, dm_build_1369 int, dm_build_1370 int16) int {
	dm_build_1368[dm_build_1369] = byte(dm_build_1370)
	dm_build_1369++
	dm_build_1368[dm_build_1369] = byte(dm_build_1370 >> 8)
	return 2
}

func (Dm_build_1372 *dm_build_1354) Dm_build_1371(dm_build_1373 []byte, dm_build_1374 int, dm_build_1375 int32) int {
	dm_build_1373[dm_build_1374] = byte(dm_build_1375)
	dm_build_1374++
	dm_build_1373[dm_build_1374] = byte(dm_build_1375 >> 8)
	dm_build_1374++
	dm_build_1373[dm_build_1374] = byte(dm_build_1375 >> 16)
	dm_build_1374++
	dm_build_1373[dm_build_1374] = byte(dm_build_1375 >> 24)
	dm_build_1374++
	return 4
}

func (Dm_build_1377 *dm_build_1354) Dm_build_1376(dm_build_1378 []byte, dm_build_1379 int, dm_build_1380 int64) int {
	dm_build_1378[dm_build_1379] = byte(dm_build_1380)
	dm_build_1379++
	dm_build_1378[dm_build_1379] = byte(dm_build_1380 >> 8)
	dm_build_1379++
	dm_build_1378[dm_build_1379] = byte(dm_build_1380 >> 16)
	dm_build_1379++
	dm_build_1378[dm_build_1379] = byte(dm_build_1380 >> 24)
	dm_build_1379++
	dm_build_1378[dm_build_1379] = byte(dm_build_1380 >> 32)
	dm_build_1379++
	dm_build_1378[dm_build_1379] = byte(dm_build_1380 >> 40)
	dm_build_1379++
	dm_build_1378[dm_build_1379] = byte(dm_build_1380 >> 48)
	dm_build_1379++
	dm_build_1378[dm_build_1379] = byte(dm_build_1380 >> 56)
	return 8
}

func (Dm_build_1382 *dm_build_1354) Dm_build_1381(dm_build_1383 []byte, dm_build_1384 int, dm_build_1385 float32) int {
	return Dm_build_1382.Dm_build_1401(dm_build_1383, dm_build_1384, math.Float32bits(dm_build_1385))
}

func (Dm_build_1387 *dm_build_1354) Dm_build_1386(dm_build_1388 []byte, dm_build_1389 int, dm_build_1390 float64) int {
	return Dm_build_1387.Dm_build_1406(dm_build_1388, dm_build_1389, math.Float64bits(dm_build_1390))
}

func (Dm_build_1392 *dm_build_1354) Dm_build_1391(dm_build_1393 []byte, dm_build_1394 int, dm_build_1395 uint8) int {
	dm_build_1393[dm_build_1394] = byte(dm_build_1395)
	return 1
}

func (Dm_build_1397 *dm_build_1354) Dm_build_1396(dm_build_1398 []byte, dm_build_1399 int, dm_build_1400 uint16) int {
	dm_build_1398[dm_build_1399] = byte(dm_build_1400)
	dm_build_1399++
	dm_build_1398[dm_build_1399] = byte(dm_build_1400 >> 8)
	return 2
}

func (Dm_build_1402 *dm_build_1354) Dm_build_1401(dm_build_1403 []byte, dm_build_1404 int, dm_build_1405 uint32) int {
	dm_build_1403[dm_build_1404] = byte(dm_build_1405)
	dm_build_1404++
	dm_build_1403[dm_build_1404] = byte(dm_build_1405 >> 8)
	dm_build_1404++
	dm_build_1403[dm_build_1404] = byte(dm_build_1405 >> 16)
	dm_build_1404++
	dm_build_1403[dm_build_1404] = byte(dm_build_1405 >> 24)
	return 3
}

func (Dm_build_1407 *dm_build_1354) Dm_build_1406(dm_build_1408 []byte, dm_build_1409 int, dm_build_1410 uint64) int {
	dm_build_1408[dm_build_1409] = byte(dm_build_1410)
	dm_build_1409++
	dm_build_1408[dm_build_1409] = byte(dm_build_1410 >> 8)
	dm_build_1409++
	dm_build_1408[dm_build_1409] = byte(dm_build_1410 >> 16)
	dm_build_1409++
	dm_build_1408[dm_build_1409] = byte(dm_build_1410 >> 24)
	dm_build_1409++
	dm_build_1408[dm_build_1409] = byte(dm_build_1410 >> 32)
	dm_build_1409++
	dm_build_1408[dm_build_1409] = byte(dm_build_1410 >> 40)
	dm_build_1409++
	dm_build_1408[dm_build_1409] = byte(dm_build_1410 >> 48)
	dm_build_1409++
	dm_build_1408[dm_build_1409] = byte(dm_build_1410 >> 56)
	return 3
}

func (Dm_build_1412 *dm_build_1354) Dm_build_1411(dm_build_1413 []byte, dm_build_1414 int, dm_build_1415 []byte, dm_build_1416 int, dm_build_1417 int) int {
	copy(dm_build_1413[dm_build_1414:dm_build_1414+dm_build_1417], dm_build_1415[dm_build_1416:dm_build_1416+dm_build_1417])
	return dm_build_1417
}

func (Dm_build_1419 *dm_build_1354) Dm_build_1418(dm_build_1420 []byte, dm_build_1421 int, dm_build_1422 []byte, dm_build_1423 int, dm_build_1424 int) int {
	dm_build_1421 += Dm_build_1419.Dm_build_1401(dm_build_1420, dm_build_1421, uint32(dm_build_1424))
	return 4 + Dm_build_1419.Dm_build_1411(dm_build_1420, dm_build_1421, dm_build_1422, dm_build_1423, dm_build_1424)
}

func (Dm_build_1426 *dm_build_1354) Dm_build_1425(dm_build_1427 []byte, dm_build_1428 int, dm_build_1429 []byte, dm_build_1430 int, dm_build_1431 int) int {
	dm_build_1428 += Dm_build_1426.Dm_build_1396(dm_build_1427, dm_build_1428, uint16(dm_build_1431))
	return 2 + Dm_build_1426.Dm_build_1411(dm_build_1427, dm_build_1428, dm_build_1429, dm_build_1430, dm_build_1431)
}

func (Dm_build_1433 *dm_build_1354) Dm_build_1432(dm_build_1434 []byte, dm_build_1435 int, dm_build_1436 string, dm_build_1437 string, dm_build_1438 *DmConnection) int {
	dm_build_1439 := Dm_build_1433.Dm_build_1571(dm_build_1436, dm_build_1437, dm_build_1438)
	dm_build_1435 += Dm_build_1433.Dm_build_1401(dm_build_1434, dm_build_1435, uint32(len(dm_build_1439)))
	return 4 + Dm_build_1433.Dm_build_1411(dm_build_1434, dm_build_1435, dm_build_1439, 0, len(dm_build_1439))
}

func (Dm_build_1441 *dm_build_1354) Dm_build_1440(dm_build_1442 []byte, dm_build_1443 int, dm_build_1444 string, dm_build_1445 string, dm_build_1446 *DmConnection) int {
	dm_build_1447 := Dm_build_1441.Dm_build_1571(dm_build_1444, dm_build_1445, dm_build_1446)

	dm_build_1443 += Dm_build_1441.Dm_build_1396(dm_build_1442, dm_build_1443, uint16(len(dm_build_1447)))
	return 2 + Dm_build_1441.Dm_build_1411(dm_build_1442, dm_build_1443, dm_build_1447, 0, len(dm_build_1447))
}

func (Dm_build_1449 *dm_build_1354) Dm_build_1448(dm_build_1450 []byte, dm_build_1451 int) byte {
	return dm_build_1450[dm_build_1451]
}

func (Dm_build_1453 *dm_build_1354) Dm_build_1452(dm_build_1454 []byte, dm_build_1455 int) int16 {
	var dm_build_1456 int16
	dm_build_1456 = int16(dm_build_1454[dm_build_1455] & 0xff)
	dm_build_1455++
	dm_build_1456 |= int16(dm_build_1454[dm_build_1455]&0xff) << 8
	return dm_build_1456
}

func (Dm_build_1458 *dm_build_1354) Dm_build_1457(dm_build_1459 []byte, dm_build_1460 int) int32 {
	var dm_build_1461 int32
	dm_build_1461 = int32(dm_build_1459[dm_build_1460] & 0xff)
	dm_build_1460++
	dm_build_1461 |= int32(dm_build_1459[dm_build_1460]&0xff) << 8
	dm_build_1460++
	dm_build_1461 |= int32(dm_build_1459[dm_build_1460]&0xff) << 16
	dm_build_1460++
	dm_build_1461 |= int32(dm_build_1459[dm_build_1460]&0xff) << 24
	return dm_build_1461
}

func (Dm_build_1463 *dm_build_1354) Dm_build_1462(dm_build_1464 []byte, dm_build_1465 int) int64 {
	var dm_build_1466 int64
	dm_build_1466 = int64(dm_build_1464[dm_build_1465] & 0xff)
	dm_build_1465++
	dm_build_1466 |= int64(dm_build_1464[dm_build_1465]&0xff) << 8
	dm_build_1465++
	dm_build_1466 |= int64(dm_build_1464[dm_build_1465]&0xff) << 16
	dm_build_1465++
	dm_build_1466 |= int64(dm_build_1464[dm_build_1465]&0xff) << 24
	dm_build_1465++
	dm_build_1466 |= int64(dm_build_1464[dm_build_1465]&0xff) << 32
	dm_build_1465++
	dm_build_1466 |= int64(dm_build_1464[dm_build_1465]&0xff) << 40
	dm_build_1465++
	dm_build_1466 |= int64(dm_build_1464[dm_build_1465]&0xff) << 48
	dm_build_1465++
	dm_build_1466 |= int64(dm_build_1464[dm_build_1465]&0xff) << 56
	return dm_build_1466
}

func (Dm_build_1468 *dm_build_1354) Dm_build_1467(dm_build_1469 []byte, dm_build_1470 int) float32 {
	return math.Float32frombits(Dm_build_1468.Dm_build_1484(dm_build_1469, dm_build_1470))
}

func (Dm_build_1472 *dm_build_1354) Dm_build_1471(dm_build_1473 []byte, dm_build_1474 int) float64 {
	return math.Float64frombits(Dm_build_1472.Dm_build_1489(dm_build_1473, dm_build_1474))
}

func (Dm_build_1476 *dm_build_1354) Dm_build_1475(dm_build_1477 []byte, dm_build_1478 int) uint8 {
	return uint8(dm_build_1477[dm_build_1478] & 0xff)
}

func (Dm_build_1480 *dm_build_1354) Dm_build_1479(dm_build_1481 []byte, dm_build_1482 int) uint16 {
	var dm_build_1483 uint16
	dm_build_1483 = uint16(dm_build_1481[dm_build_1482] & 0xff)
	dm_build_1482++
	dm_build_1483 |= uint16(dm_build_1481[dm_build_1482]&0xff) << 8
	return dm_build_1483
}

func (Dm_build_1485 *dm_build_1354) Dm_build_1484(dm_build_1486 []byte, dm_build_1487 int) uint32 {
	var dm_build_1488 uint32
	dm_build_1488 = uint32(dm_build_1486[dm_build_1487] & 0xff)
	dm_build_1487++
	dm_build_1488 |= uint32(dm_build_1486[dm_build_1487]&0xff) << 8
	dm_build_1487++
	dm_build_1488 |= uint32(dm_build_1486[dm_build_1487]&0xff) << 16
	dm_build_1487++
	dm_build_1488 |= uint32(dm_build_1486[dm_build_1487]&0xff) << 24
	return dm_build_1488
}

func (Dm_build_1490 *dm_build_1354) Dm_build_1489(dm_build_1491 []byte, dm_build_1492 int) uint64 {
	var dm_build_1493 uint64
	dm_build_1493 = uint64(dm_build_1491[dm_build_1492] & 0xff)
	dm_build_1492++
	dm_build_1493 |= uint64(dm_build_1491[dm_build_1492]&0xff) << 8
	dm_build_1492++
	dm_build_1493 |= uint64(dm_build_1491[dm_build_1492]&0xff) << 16
	dm_build_1492++
	dm_build_1493 |= uint64(dm_build_1491[dm_build_1492]&0xff) << 24
	dm_build_1492++
	dm_build_1493 |= uint64(dm_build_1491[dm_build_1492]&0xff) << 32
	dm_build_1492++
	dm_build_1493 |= uint64(dm_build_1491[dm_build_1492]&0xff) << 40
	dm_build_1492++
	dm_build_1493 |= uint64(dm_build_1491[dm_build_1492]&0xff) << 48
	dm_build_1492++
	dm_build_1493 |= uint64(dm_build_1491[dm_build_1492]&0xff) << 56
	return dm_build_1493
}

func (Dm_build_1495 *dm_build_1354) Dm_build_1494(dm_build_1496 []byte, dm_build_1497 int) []byte {
	dm_build_1498 := Dm_build_1495.Dm_build_1484(dm_build_1496, dm_build_1497)

	dm_build_1499 := make([]byte, dm_build_1498)
	copy(dm_build_1499[:int(dm_build_1498)], dm_build_1496[dm_build_1497+4:dm_build_1497+4+int(dm_build_1498)])
	return dm_build_1499
}

func (Dm_build_1501 *dm_build_1354) Dm_build_1500(dm_build_1502 []byte, dm_build_1503 int) []byte {
	dm_build_1504 := Dm_build_1501.Dm_build_1479(dm_build_1502, dm_build_1503)

	dm_build_1505 := make([]byte, dm_build_1504)
	copy(dm_build_1505[:int(dm_build_1504)], dm_build_1502[dm_build_1503+2:dm_build_1503+2+int(dm_build_1504)])
	return dm_build_1505
}

func (Dm_build_1507 *dm_build_1354) Dm_build_1506(dm_build_1508 []byte, dm_build_1509 int, dm_build_1510 int) []byte {

	dm_build_1511 := make([]byte, dm_build_1510)
	copy(dm_build_1511[:dm_build_1510], dm_build_1508[dm_build_1509:dm_build_1509+dm_build_1510])
	return dm_build_1511
}

func (Dm_build_1513 *dm_build_1354) Dm_build_1512(dm_build_1514 []byte, dm_build_1515 int, dm_build_1516 int, dm_build_1517 string, dm_build_1518 *DmConnection) string {
	return Dm_build_1513.Dm_build_1607(dm_build_1514[dm_build_1515:dm_build_1515+dm_build_1516], dm_build_1517, dm_build_1518)
}

func (Dm_build_1520 *dm_build_1354) Dm_build_1519(dm_build_1521 []byte, dm_build_1522 int, dm_build_1523 string, dm_build_1524 *DmConnection) string {
	dm_build_1525 := Dm_build_1520.Dm_build_1484(dm_build_1521, dm_build_1522)
	dm_build_1522 += 4
	return Dm_build_1520.Dm_build_1512(dm_build_1521, dm_build_1522, int(dm_build_1525), dm_build_1523, dm_build_1524)
}

func (Dm_build_1527 *dm_build_1354) Dm_build_1526(dm_build_1528 []byte, dm_build_1529 int, dm_build_1530 string, dm_build_1531 *DmConnection) string {
	dm_build_1532 := Dm_build_1527.Dm_build_1479(dm_build_1528, dm_build_1529)
	dm_build_1529 += 2
	return Dm_build_1527.Dm_build_1512(dm_build_1528, dm_build_1529, int(dm_build_1532), dm_build_1530, dm_build_1531)
}

func (Dm_build_1534 *dm_build_1354) Dm_build_1533(dm_build_1535 byte) []byte {
	return []byte{dm_build_1535}
}

func (Dm_build_1537 *dm_build_1354) Dm_build_1536(dm_build_1538 int8) []byte {
	return []byte{byte(dm_build_1538)}
}

func (Dm_build_1540 *dm_build_1354) Dm_build_1539(dm_build_1541 int16) []byte {
	return []byte{byte(dm_build_1541), byte(dm_build_1541 >> 8)}
}

func (Dm_build_1543 *dm_build_1354) Dm_build_1542(dm_build_1544 int32) []byte {
	return []byte{byte(dm_build_1544), byte(dm_build_1544 >> 8), byte(dm_build_1544 >> 16), byte(dm_build_1544 >> 24)}
}

func (Dm_build_1546 *dm_build_1354) Dm_build_1545(dm_build_1547 int64) []byte {
	return []byte{byte(dm_build_1547), byte(dm_build_1547 >> 8), byte(dm_build_1547 >> 16), byte(dm_build_1547 >> 24), byte(dm_build_1547 >> 32),
		byte(dm_build_1547 >> 40), byte(dm_build_1547 >> 48), byte(dm_build_1547 >> 56)}
}

func (Dm_build_1549 *dm_build_1354) Dm_build_1548(dm_build_1550 float32) []byte {
	return Dm_build_1549.Dm_build_1560(math.Float32bits(dm_build_1550))
}

func (Dm_build_1552 *dm_build_1354) Dm_build_1551(dm_build_1553 float64) []byte {
	return Dm_build_1552.Dm_build_1563(math.Float64bits(dm_build_1553))
}

func (Dm_build_1555 *dm_build_1354) Dm_build_1554(dm_build_1556 uint8) []byte {
	return []byte{byte(dm_build_1556)}
}

func (Dm_build_1558 *dm_build_1354) Dm_build_1557(dm_build_1559 uint16) []byte {
	return []byte{byte(dm_build_1559), byte(dm_build_1559 >> 8)}
}

func (Dm_build_1561 *dm_build_1354) Dm_build_1560(dm_build_1562 uint32) []byte {
	return []byte{byte(dm_build_1562), byte(dm_build_1562 >> 8), byte(dm_build_1562 >> 16), byte(dm_build_1562 >> 24)}
}

func (Dm_build_1564 *dm_build_1354) Dm_build_1563(dm_build_1565 uint64) []byte {
	return []byte{byte(dm_build_1565), byte(dm_build_1565 >> 8), byte(dm_build_1565 >> 16), byte(dm_build_1565 >> 24), byte(dm_build_1565 >> 32), byte(dm_build_1565 >> 40), byte(dm_build_1565 >> 48), byte(dm_build_1565 >> 56)}
}

func (Dm_build_1567 *dm_build_1354) Dm_build_1566(dm_build_1568 []byte, dm_build_1569 string, dm_build_1570 *DmConnection) []byte {
	if dm_build_1569 == "UTF-8" {
		return dm_build_1568
	}

	if dm_build_1570 == nil {
		if e := dm_build_1612(dm_build_1569); e != nil {
			tmp, err := ioutil.ReadAll(
				transform.NewReader(bytes.NewReader(dm_build_1568), e.NewEncoder()),
			)
			if err != nil {
				panic("UTF8 To Charset error!")
			}

			return tmp
		}

		panic("Unsupported Charset!")
	}

	if dm_build_1570.encodeBuffer == nil {
		dm_build_1570.encodeBuffer = bytes.NewBuffer(nil)
		dm_build_1570.encode = dm_build_1612(dm_build_1570.getServerEncoding())
		dm_build_1570.transformReaderDst = make([]byte, 4096)
		dm_build_1570.transformReaderSrc = make([]byte, 4096)
	}

	if e := dm_build_1570.encode; e != nil {

		dm_build_1570.encodeBuffer.Reset()

		n, err := dm_build_1570.encodeBuffer.ReadFrom(
			Dm_build_1626(bytes.NewReader(dm_build_1568), e.NewEncoder(), dm_build_1570.transformReaderDst, dm_build_1570.transformReaderSrc),
		)
		if err != nil {
			panic("UTF8 To Charset error!")
		}
		var tmp = make([]byte, n)
		if _, err = dm_build_1570.encodeBuffer.Read(tmp); err != nil {
			panic("UTF8 To Charset error!")
		}
		return tmp
	}

	panic("Unsupported Charset!")
}

func (Dm_build_1572 *dm_build_1354) Dm_build_1571(dm_build_1573 string, dm_build_1574 string, dm_build_1575 *DmConnection) []byte {
	return Dm_build_1572.Dm_build_1566([]byte(dm_build_1573), dm_build_1574, dm_build_1575)
}

func (Dm_build_1577 *dm_build_1354) Dm_build_1576(dm_build_1578 []byte) byte {
	return Dm_build_1577.Dm_build_1448(dm_build_1578, 0)
}

func (Dm_build_1580 *dm_build_1354) Dm_build_1579(dm_build_1581 []byte) int16 {
	return Dm_build_1580.Dm_build_1452(dm_build_1581, 0)
}

func (Dm_build_1583 *dm_build_1354) Dm_build_1582(dm_build_1584 []byte) int32 {
	return Dm_build_1583.Dm_build_1457(dm_build_1584, 0)
}

func (Dm_build_1586 *dm_build_1354) Dm_build_1585(dm_build_1587 []byte) int64 {
	return Dm_build_1586.Dm_build_1462(dm_build_1587, 0)
}

func (Dm_build_1589 *dm_build_1354) Dm_build_1588(dm_build_1590 []byte) float32 {
	return Dm_build_1589.Dm_build_1467(dm_build_1590, 0)
}

func (Dm_build_1592 *dm_build_1354) Dm_build_1591(dm_build_1593 []byte) float64 {
	return Dm_build_1592.Dm_build_1471(dm_build_1593, 0)
}

func (Dm_build_1595 *dm_build_1354) Dm_build_1594(dm_build_1596 []byte) uint8 {
	return Dm_build_1595.Dm_build_1475(dm_build_1596, 0)
}

func (Dm_build_1598 *dm_build_1354) Dm_build_1597(dm_build_1599 []byte) uint16 {
	return Dm_build_1598.Dm_build_1479(dm_build_1599, 0)
}

func (Dm_build_1601 *dm_build_1354) Dm_build_1600(dm_build_1602 []byte) uint32 {
	return Dm_build_1601.Dm_build_1484(dm_build_1602, 0)
}

func (Dm_build_1604 *dm_build_1354) Dm_build_1603(dm_build_1605 []byte, dm_build_1606 string) []byte {
	if dm_build_1606 == "UTF-8" {
		return dm_build_1605
	}

	if e := dm_build_1612(dm_build_1606); e != nil {

		tmp, err := ioutil.ReadAll(
			transform.NewReader(bytes.NewReader(dm_build_1605), e.NewDecoder()),
		)
		if err != nil {

			panic("Charset To UTF8 error!")
		}

		return tmp
	}

	panic("Unsupported Charset!")

}

func (Dm_build_1608 *dm_build_1354) Dm_build_1607(dm_build_1609 []byte, dm_build_1610 string, dm_build_1611 *DmConnection) string {
	return string(Dm_build_1608.Dm_build_1603(dm_build_1609, dm_build_1610))
}

func dm_build_1612(dm_build_1613 string) encoding.Encoding {
	if e, err := ianaindex.MIB.Encoding(dm_build_1613); err == nil && e != nil {
		return e
	}
	return nil
}

type Dm_build_1614 struct {
	dm_build_1615 io.Reader
	dm_build_1616 transform.Transformer
	dm_build_1617 error

	dm_build_1618                []byte
	dm_build_1619, dm_build_1620 int

	dm_build_1621                []byte
	dm_build_1622, dm_build_1623 int

	dm_build_1624 bool
}

const dm_build_1625 = 4096

func Dm_build_1626(dm_build_1627 io.Reader, dm_build_1628 transform.Transformer, dm_build_1629 []byte, dm_build_1630 []byte) *Dm_build_1614 {
	dm_build_1628.Reset()
	return &Dm_build_1614{
		dm_build_1615: dm_build_1627,
		dm_build_1616: dm_build_1628,
		dm_build_1618: dm_build_1629,
		dm_build_1621: dm_build_1630,
	}
}

func (dm_build_1632 *Dm_build_1614) Read(dm_build_1633 []byte) (int, error) {
	dm_build_1634, dm_build_1635 := 0, error(nil)
	for {

		if dm_build_1632.dm_build_1619 != dm_build_1632.dm_build_1620 {
			dm_build_1634 = copy(dm_build_1633, dm_build_1632.dm_build_1618[dm_build_1632.dm_build_1619:dm_build_1632.dm_build_1620])
			dm_build_1632.dm_build_1619 += dm_build_1634
			if dm_build_1632.dm_build_1619 == dm_build_1632.dm_build_1620 && dm_build_1632.dm_build_1624 {
				return dm_build_1634, dm_build_1632.dm_build_1617
			}
			return dm_build_1634, nil
		} else if dm_build_1632.dm_build_1624 {
			return 0, dm_build_1632.dm_build_1617
		}

		if dm_build_1632.dm_build_1622 != dm_build_1632.dm_build_1623 || dm_build_1632.dm_build_1617 != nil {
			dm_build_1632.dm_build_1619 = 0
			dm_build_1632.dm_build_1620, dm_build_1634, dm_build_1635 = dm_build_1632.dm_build_1616.Transform(dm_build_1632.dm_build_1618, dm_build_1632.dm_build_1621[dm_build_1632.dm_build_1622:dm_build_1632.dm_build_1623], dm_build_1632.dm_build_1617 == io.EOF)
			dm_build_1632.dm_build_1622 += dm_build_1634

			switch {
			case dm_build_1635 == nil:
				if dm_build_1632.dm_build_1622 != dm_build_1632.dm_build_1623 {
					dm_build_1632.dm_build_1617 = nil
				}

				dm_build_1632.dm_build_1624 = dm_build_1632.dm_build_1617 != nil
				continue
			case dm_build_1635 == transform.ErrShortDst && (dm_build_1632.dm_build_1620 != 0 || dm_build_1634 != 0):

				continue
			case dm_build_1635 == transform.ErrShortSrc && dm_build_1632.dm_build_1623-dm_build_1632.dm_build_1622 != len(dm_build_1632.dm_build_1621) && dm_build_1632.dm_build_1617 == nil:

			default:
				dm_build_1632.dm_build_1624 = true

				if dm_build_1632.dm_build_1617 == nil || dm_build_1632.dm_build_1617 == io.EOF {
					dm_build_1632.dm_build_1617 = dm_build_1635
				}
				continue
			}
		}

		if dm_build_1632.dm_build_1622 != 0 {
			dm_build_1632.dm_build_1622, dm_build_1632.dm_build_1623 = 0, copy(dm_build_1632.dm_build_1621, dm_build_1632.dm_build_1621[dm_build_1632.dm_build_1622:dm_build_1632.dm_build_1623])
		}
		dm_build_1634, dm_build_1632.dm_build_1617 = dm_build_1632.dm_build_1615.Read(dm_build_1632.dm_build_1621[dm_build_1632.dm_build_1623:])
		dm_build_1632.dm_build_1623 += dm_build_1634
	}
}
