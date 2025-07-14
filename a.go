/*
 * Copyright (c) 2000-2018, 达梦数据库有限公司.
 * All rights reserved.
 */
package dm

import (
	"bytes"
	"context"
	"crypto/tls"
	"fmt"
	"github.com/gaoyuan98/dm/security"
	"net"
	"strconv"
	"time"
	"unicode/utf8"
)

const (
	Dm_build_1345 = 8192
	Dm_build_1346 = 2 * time.Second
)

type dm_build_1347 struct {
	dm_build_1348 net.Conn
	dm_build_1349 *tls.Conn
	dm_build_1350 *Dm_build_1011
	dm_build_1351 *DmConnection
	dm_build_1352 security.Cipher
	dm_build_1353 bool
	dm_build_1354 bool
	dm_build_1355 *security.DhKey

	dm_build_1356 bool
	dm_build_1357 string
	dm_build_1358 bool
}

func dm_build_1359(dm_build_1360 context.Context, dm_build_1361 *DmConnection) (*dm_build_1347, error) {
	var dm_build_1362 net.Conn
	var dm_build_1363 error

	dialsLock.RLock()
	dm_build_1364, dm_build_1365 := dials[dm_build_1361.dmConnector.dialName]
	dialsLock.RUnlock()
	if dm_build_1365 {
		dm_build_1362, dm_build_1363 = dm_build_1364(dm_build_1360, dm_build_1361.dmConnector.host+":"+strconv.Itoa(int(dm_build_1361.dmConnector.port)))
	} else {
		dm_build_1362, dm_build_1363 = dm_build_1367(dm_build_1361.dmConnector.host+":"+strconv.Itoa(int(dm_build_1361.dmConnector.port)), time.Duration(dm_build_1361.dmConnector.socketTimeout)*time.Second)
	}
	if dm_build_1363 != nil {
		return nil, dm_build_1363
	}

	dm_build_1366 := dm_build_1347{}
	dm_build_1366.dm_build_1348 = dm_build_1362
	dm_build_1366.dm_build_1350 = Dm_build_1014(Dm_build_14)
	dm_build_1366.dm_build_1351 = dm_build_1361
	dm_build_1366.dm_build_1353 = false
	dm_build_1366.dm_build_1354 = false
	dm_build_1366.dm_build_1356 = false
	dm_build_1366.dm_build_1357 = ""
	dm_build_1366.dm_build_1358 = false
	dm_build_1361.Access = &dm_build_1366

	return &dm_build_1366, nil
}

func dm_build_1367(dm_build_1368 string, dm_build_1369 time.Duration) (net.Conn, error) {
	dm_build_1370, dm_build_1371 := net.DialTimeout("tcp", dm_build_1368, dm_build_1369)
	if dm_build_1371 != nil {
		return &net.TCPConn{}, ECGO_COMMUNITION_ERROR.addDetail("\tdial address: " + dm_build_1368).throw()
	}

	if tcpConn, ok := dm_build_1370.(*net.TCPConn); ok {
		tcpConn.SetKeepAlive(true)
		tcpConn.SetKeepAlivePeriod(Dm_build_1346)
		tcpConn.SetNoDelay(true)

	}
	return dm_build_1370, nil
}

func (dm_build_1373 *dm_build_1347) dm_build_1372(dm_build_1374 dm_build_135) bool {
	var dm_build_1375 = dm_build_1373.dm_build_1351.dmConnector.compress
	if dm_build_1374.dm_build_150() == Dm_build_42 || dm_build_1375 == Dm_build_91 {
		return false
	}

	if dm_build_1375 == Dm_build_89 {
		return true
	} else if dm_build_1375 == Dm_build_90 {
		return !dm_build_1373.dm_build_1351.Local && dm_build_1374.dm_build_148() > Dm_build_88
	}

	return false
}

func (dm_build_1377 *dm_build_1347) dm_build_1376(dm_build_1378 dm_build_135) bool {
	var dm_build_1379 = dm_build_1377.dm_build_1351.dmConnector.compress
	if dm_build_1378.dm_build_150() == Dm_build_42 || dm_build_1379 == Dm_build_91 {
		return false
	}

	if dm_build_1379 == Dm_build_89 {
		return true
	} else if dm_build_1379 == Dm_build_90 {
		return dm_build_1377.dm_build_1350.Dm_build_1278(Dm_build_50) == 1
	}

	return false
}

func (dm_build_1381 *dm_build_1347) dm_build_1380(dm_build_1382 dm_build_135) (err error) {
	defer func() {
		if p := recover(); p != nil {
			if _, ok := p.(string); ok {
				err = ECGO_COMMUNITION_ERROR.addDetail("\t" + p.(string)).throw()
			} else {
				err = fmt.Errorf("internal error: %v", p)
			}
		}
	}()

	dm_build_1384 := dm_build_1382.dm_build_148()

	if dm_build_1384 > 0 {

		if dm_build_1381.dm_build_1372(dm_build_1382) {
			var retBytes, err = Compress(dm_build_1381.dm_build_1350, Dm_build_43, int(dm_build_1384), int(dm_build_1381.dm_build_1351.dmConnector.compressID))
			if err != nil {
				return err
			}

			dm_build_1381.dm_build_1350.Dm_build_1025(Dm_build_43)

			dm_build_1381.dm_build_1350.Dm_build_1066(dm_build_1384)

			dm_build_1381.dm_build_1350.Dm_build_1094(retBytes)

			dm_build_1382.dm_build_149(int32(len(retBytes)) + ULINT_SIZE)

			dm_build_1381.dm_build_1350.Dm_build_1198(Dm_build_50, 1)
		}

		if dm_build_1381.dm_build_1354 {
			dm_build_1384 = dm_build_1382.dm_build_148()
			var retBytes = dm_build_1381.dm_build_1352.Encrypt(dm_build_1381.dm_build_1350.Dm_build_1305(Dm_build_43, int(dm_build_1384)), true)

			dm_build_1381.dm_build_1350.Dm_build_1025(Dm_build_43)

			dm_build_1381.dm_build_1350.Dm_build_1094(retBytes)

			dm_build_1382.dm_build_149(int32(len(retBytes)))
		}
	}

	if dm_build_1381.dm_build_1350.Dm_build_1023() > Dm_build_15 {
		return ECGO_MSG_TOO_LONG.throw()
	}

	dm_build_1382.dm_build_144()
	if dm_build_1381.dm_build_1623(dm_build_1382) {
		if dm_build_1381.dm_build_1349 != nil {
			dm_build_1381.dm_build_1350.Dm_build_1028(0)
			if _, err := dm_build_1381.dm_build_1350.Dm_build_1047(dm_build_1381.dm_build_1349); err != nil {
				return err
			}
		}
	} else {
		dm_build_1381.dm_build_1350.Dm_build_1028(0)
		if _, err := dm_build_1381.dm_build_1350.Dm_build_1047(dm_build_1381.dm_build_1348); err != nil {
			return err
		}
	}
	return nil
}

func (dm_build_1386 *dm_build_1347) dm_build_1385(dm_build_1387 dm_build_135) (err error) {
	defer func() {
		if p := recover(); p != nil {
			if _, ok := p.(string); ok {
				err = ECGO_COMMUNITION_ERROR.addDetail("\t" + p.(string)).throw()
			} else {
				err = fmt.Errorf("internal error: %v", p)
			}
		}
	}()

	dm_build_1389 := int32(0)
	if dm_build_1386.dm_build_1623(dm_build_1387) {
		if dm_build_1386.dm_build_1349 != nil {
			dm_build_1386.dm_build_1350.Dm_build_1025(0)
			if _, err := dm_build_1386.dm_build_1350.Dm_build_1041(dm_build_1386.dm_build_1349, Dm_build_43); err != nil {
				return err
			}

			dm_build_1389 = dm_build_1387.dm_build_148()
			if dm_build_1389 > 0 {
				if _, err := dm_build_1386.dm_build_1350.Dm_build_1041(dm_build_1386.dm_build_1349, int(dm_build_1389)); err != nil {
					return err
				}
			}
		}
	} else {

		dm_build_1386.dm_build_1350.Dm_build_1025(0)
		if _, err := dm_build_1386.dm_build_1350.Dm_build_1041(dm_build_1386.dm_build_1348, Dm_build_43); err != nil {
			return err
		}
		dm_build_1389 = dm_build_1387.dm_build_148()

		if dm_build_1389 > 0 {
			if _, err := dm_build_1386.dm_build_1350.Dm_build_1041(dm_build_1386.dm_build_1348, int(dm_build_1389)); err != nil {
				return err
			}
		}
	}

	dm_build_1387.dm_build_145()

	dm_build_1389 = dm_build_1387.dm_build_148()
	if dm_build_1389 <= 0 {
		return nil
	}

	if dm_build_1386.dm_build_1354 {
		ebytes := dm_build_1386.dm_build_1350.Dm_build_1305(Dm_build_43, int(dm_build_1389))
		bytes, err := dm_build_1386.dm_build_1352.Decrypt(ebytes, true)
		if err != nil {
			return err
		}
		dm_build_1386.dm_build_1350.Dm_build_1025(Dm_build_43)
		dm_build_1386.dm_build_1350.Dm_build_1094(bytes)
		dm_build_1387.dm_build_149(int32(len(bytes)))
	}

	if dm_build_1386.dm_build_1376(dm_build_1387) {

		dm_build_1389 = dm_build_1387.dm_build_148()
		cbytes := dm_build_1386.dm_build_1350.Dm_build_1305(Dm_build_43+ULINT_SIZE, int(dm_build_1389-ULINT_SIZE))
		bytes, err := UnCompress(cbytes, int(dm_build_1386.dm_build_1351.dmConnector.compressID))
		if err != nil {
			return err
		}
		dm_build_1386.dm_build_1350.Dm_build_1025(Dm_build_43)
		dm_build_1386.dm_build_1350.Dm_build_1094(bytes)
		dm_build_1387.dm_build_149(int32(len(bytes)))
	}
	return nil
}

func (dm_build_1391 *dm_build_1347) dm_build_1390(dm_build_1392 dm_build_135) (dm_build_1393 interface{}, dm_build_1394 error) {
	if dm_build_1391.dm_build_1358 {
		return nil, ECGO_CONNECTION_CLOSED.throw()
	}
	dm_build_1395 := dm_build_1391.dm_build_1351
	dm_build_1395.mu.Lock()
	defer dm_build_1395.mu.Unlock()
	dm_build_1394 = dm_build_1392.dm_build_139(dm_build_1392)
	if dm_build_1394 != nil {
		return nil, dm_build_1394
	}

	dm_build_1394 = dm_build_1391.dm_build_1380(dm_build_1392)
	if dm_build_1394 != nil {
		return nil, dm_build_1394
	}

	dm_build_1394 = dm_build_1391.dm_build_1385(dm_build_1392)
	if dm_build_1394 != nil {
		return nil, dm_build_1394
	}

	return dm_build_1392.dm_build_143(dm_build_1392)
}

func (dm_build_1397 *dm_build_1347) dm_build_1396() (*dm_build_594, error) {

	Dm_build_1398 := dm_build_600(dm_build_1397)
	_, dm_build_1399 := dm_build_1397.dm_build_1390(Dm_build_1398)
	if dm_build_1399 != nil {
		return nil, dm_build_1399
	}

	return Dm_build_1398, nil
}

func (dm_build_1401 *dm_build_1347) dm_build_1400() error {

	dm_build_1402 := dm_build_459(dm_build_1401)
	_, dm_build_1403 := dm_build_1401.dm_build_1390(dm_build_1402)
	if dm_build_1403 != nil {
		return dm_build_1403
	}

	return nil
}

func (dm_build_1405 *dm_build_1347) dm_build_1404() error {

	var dm_build_1406 *dm_build_594
	var err error
	if dm_build_1406, err = dm_build_1405.dm_build_1396(); err != nil {
		return err
	}

	if dm_build_1405.dm_build_1351.sslEncrypt == 2 {
		if err = dm_build_1405.dm_build_1619(false); err != nil {
			return ECGO_INIT_SSL_FAILED.addDetail("\n" + err.Error()).throw()
		}
	} else if dm_build_1405.dm_build_1351.sslEncrypt == 1 {
		if err = dm_build_1405.dm_build_1619(true); err != nil {
			return ECGO_INIT_SSL_FAILED.addDetail("\n" + err.Error()).throw()
		}
	}

	if dm_build_1405.dm_build_1354 || dm_build_1405.dm_build_1353 {
		k, err := dm_build_1405.dm_build_1609()
		if err != nil {
			return err
		}
		sessionKey := security.ComputeSessionKey(k, dm_build_1406.Dm_build_598)
		encryptType := dm_build_1406.dm_build_596
		hashType := int(dm_build_1406.Dm_build_597)
		if encryptType == -1 {
			encryptType = security.DES_CFB
		}
		if hashType == -1 {
			hashType = security.MD5
		}
		err = dm_build_1405.dm_build_1612(encryptType, sessionKey, dm_build_1405.dm_build_1351.dmConnector.cipherPath, hashType)
		if err != nil {
			return err
		}
	}

	if err := dm_build_1405.dm_build_1400(); err != nil {
		return err
	}
	return nil
}

func (dm_build_1409 *dm_build_1347) Dm_build_1408(dm_build_1410 *DmStatement) error {
	dm_build_1411 := dm_build_623(dm_build_1409, dm_build_1410)
	_, dm_build_1412 := dm_build_1409.dm_build_1390(dm_build_1411)
	if dm_build_1412 != nil {
		return dm_build_1412
	}

	return nil
}

func (dm_build_1414 *dm_build_1347) Dm_build_1413(dm_build_1415 int32) error {
	dm_build_1416 := dm_build_633(dm_build_1414, dm_build_1415)
	_, dm_build_1417 := dm_build_1414.dm_build_1390(dm_build_1416)
	if dm_build_1417 != nil {
		return dm_build_1417
	}

	return nil
}

func (dm_build_1419 *dm_build_1347) Dm_build_1418(dm_build_1420 *DmStatement, dm_build_1421 bool, dm_build_1422 int16) (*execRetInfo, error) {
	dm_build_1423 := dm_build_500(dm_build_1419, dm_build_1420, dm_build_1421, dm_build_1422)
	dm_build_1424, dm_build_1425 := dm_build_1419.dm_build_1390(dm_build_1423)
	if dm_build_1425 != nil {
		return nil, dm_build_1425
	}
	return dm_build_1424.(*execRetInfo), nil
}

func (dm_build_1427 *dm_build_1347) Dm_build_1426(dm_build_1428 *DmStatement, dm_build_1429 int16) (*execRetInfo, error) {
	return dm_build_1427.Dm_build_1418(dm_build_1428, false, Dm_build_95)
}

func (dm_build_1431 *dm_build_1347) Dm_build_1430(dm_build_1432 *DmStatement, dm_build_1433 []OptParameter) (*execRetInfo, error) {
	dm_build_1434, dm_build_1435 := dm_build_1431.dm_build_1390(dm_build_238(dm_build_1431, dm_build_1432, dm_build_1433))
	if dm_build_1435 != nil {
		return nil, dm_build_1435
	}

	return dm_build_1434.(*execRetInfo), nil
}

func (dm_build_1437 *dm_build_1347) Dm_build_1436(dm_build_1438 *DmStatement, dm_build_1439 int16) (*execRetInfo, error) {
	return dm_build_1437.Dm_build_1418(dm_build_1438, true, dm_build_1439)
}

func (dm_build_1441 *dm_build_1347) Dm_build_1440(dm_build_1442 *DmStatement, dm_build_1443 [][]interface{}) (*execRetInfo, error) {
	dm_build_1444 := dm_build_270(dm_build_1441, dm_build_1442, dm_build_1443)
	dm_build_1445, dm_build_1446 := dm_build_1441.dm_build_1390(dm_build_1444)
	if dm_build_1446 != nil {
		return nil, dm_build_1446
	}
	return dm_build_1445.(*execRetInfo), nil
}

func (dm_build_1448 *dm_build_1347) Dm_build_1447(dm_build_1449 *DmStatement, dm_build_1450 [][]interface{}, dm_build_1451 bool) (*execRetInfo, error) {
	var dm_build_1452, dm_build_1453 = 0, 0
	var dm_build_1454 = len(dm_build_1450)
	var dm_build_1455 [][]interface{}
	var dm_build_1456 = NewExceInfo()
	dm_build_1456.updateCounts = make([]int64, dm_build_1454)
	var dm_build_1457 = false
	for dm_build_1452 < dm_build_1454 {
		for dm_build_1453 = dm_build_1452; dm_build_1453 < dm_build_1454; dm_build_1453++ {
			paramData := dm_build_1450[dm_build_1453]
			bindData := make([]interface{}, dm_build_1449.paramCount)
			dm_build_1457 = false
			for icol := 0; icol < int(dm_build_1449.paramCount); icol++ {
				if dm_build_1449.bindParams[icol].ioType == IO_TYPE_OUT {
					continue
				}
				if dm_build_1448.dm_build_1592(bindData, paramData, icol) {
					dm_build_1457 = true
					break
				}
			}

			if dm_build_1457 {
				break
			}
			dm_build_1455 = append(dm_build_1455, bindData)
		}

		if dm_build_1453 != dm_build_1452 {
			tmpExecInfo, err := dm_build_1448.Dm_build_1440(dm_build_1449, dm_build_1455)
			if err != nil {
				return nil, err
			}
			dm_build_1455 = dm_build_1455[0:0]
			dm_build_1456.union(tmpExecInfo, dm_build_1452, dm_build_1453-dm_build_1452)
		}

		if dm_build_1453 < dm_build_1454 {
			tmpExecInfo, err := dm_build_1448.Dm_build_1466(dm_build_1449, dm_build_1450[dm_build_1453], dm_build_1451)
			if err != nil {
				return nil, err
			}

			dm_build_1451 = true
			dm_build_1456.union(tmpExecInfo, dm_build_1453, 1)
		}

		dm_build_1452 = dm_build_1453 + 1
	}
	for _, i := range dm_build_1456.updateCounts {
		if i > 0 {
			dm_build_1456.updateCount += i
		}
	}
	return dm_build_1456, nil
}

func (dm_build_1459 *dm_build_1347) dm_build_1458(dm_build_1460 *DmStatement, dm_build_1461 []parameter) error {
	if !dm_build_1460.prepared {
		retInfo, err := dm_build_1459.Dm_build_1418(dm_build_1460, false, Dm_build_95)
		if err != nil {
			return nil
		}
		dm_build_1460.serverParams = retInfo.serverParams
		dm_build_1460.paramCount = int32(len(dm_build_1460.serverParams))
		dm_build_1460.prepared = true
	}

	dm_build_1462 := dm_build_489(dm_build_1459, dm_build_1460, dm_build_1460.bindParams)
	dm_build_1463, err := dm_build_1459.dm_build_1390(dm_build_1462)
	if err != nil {
		return nil
	}
	retInfo := dm_build_1463.(*execRetInfo)
	if retInfo.serverParams != nil && len(retInfo.serverParams) > 0 {
		dm_build_1460.serverParams = retInfo.serverParams
		dm_build_1460.paramCount = int32(len(dm_build_1460.serverParams))
	}
	dm_build_1460.preExec = true
	return nil
}

func (dm_build_1467 *dm_build_1347) Dm_build_1466(dm_build_1468 *DmStatement, dm_build_1469 []interface{}, dm_build_1470 bool) (*execRetInfo, error) {

	var dm_build_1471 = make([]interface{}, dm_build_1468.paramCount)
	for icol := 0; icol < int(dm_build_1468.paramCount); icol++ {
		if dm_build_1468.bindParams[icol].ioType == IO_TYPE_OUT {
			continue
		}
		if dm_build_1467.dm_build_1592(dm_build_1471, dm_build_1469, icol) {

			if !dm_build_1470 {
				dm_build_1467.dm_build_1458(dm_build_1468, dm_build_1468.bindParams)

				dm_build_1470 = true
			}

			dm_build_1467.dm_build_1598(dm_build_1468, dm_build_1468.bindParams[icol], icol, dm_build_1469[icol].(iOffRowBinder))
			dm_build_1471[icol] = ParamDataEnum_OFF_ROW
		}
	}

	var dm_build_1472 = make([][]interface{}, 1, 1)
	dm_build_1472[0] = dm_build_1471

	dm_build_1473 := dm_build_270(dm_build_1467, dm_build_1468, dm_build_1472)
	dm_build_1474, dm_build_1475 := dm_build_1467.dm_build_1390(dm_build_1473)
	if dm_build_1475 != nil {
		return nil, dm_build_1475
	}
	return dm_build_1474.(*execRetInfo), nil
}

func (dm_build_1477 *dm_build_1347) Dm_build_1476(dm_build_1478 *DmStatement, dm_build_1479 int16) (*execRetInfo, error) {
	dm_build_1480 := dm_build_476(dm_build_1477, dm_build_1478, dm_build_1479)

	dm_build_1481, dm_build_1482 := dm_build_1477.dm_build_1390(dm_build_1480)
	if dm_build_1482 != nil {
		return nil, dm_build_1482
	}
	return dm_build_1481.(*execRetInfo), nil
}

func (dm_build_1484 *dm_build_1347) Dm_build_1483(dm_build_1485 *innerRows, dm_build_1486 int64) (*execRetInfo, error) {
	dm_build_1487 := dm_build_377(dm_build_1484, dm_build_1485, dm_build_1486, INT64_MAX)
	dm_build_1488, dm_build_1489 := dm_build_1484.dm_build_1390(dm_build_1487)
	if dm_build_1489 != nil {
		return nil, dm_build_1489
	}
	return dm_build_1488.(*execRetInfo), nil
}

func (dm_build_1491 *dm_build_1347) Commit() error {
	dm_build_1492 := dm_build_223(dm_build_1491)
	_, dm_build_1493 := dm_build_1491.dm_build_1390(dm_build_1492)
	if dm_build_1493 != nil {
		return dm_build_1493
	}

	return nil
}

func (dm_build_1495 *dm_build_1347) Rollback() error {
	dm_build_1496 := dm_build_538(dm_build_1495)
	_, dm_build_1497 := dm_build_1495.dm_build_1390(dm_build_1496)
	if dm_build_1497 != nil {
		return dm_build_1497
	}

	return nil
}

func (dm_build_1499 *dm_build_1347) Dm_build_1498(dm_build_1500 *DmConnection) error {
	dm_build_1501 := dm_build_543(dm_build_1499, dm_build_1500.IsoLevel)
	_, dm_build_1502 := dm_build_1499.dm_build_1390(dm_build_1501)
	if dm_build_1502 != nil {
		return dm_build_1502
	}

	return nil
}

func (dm_build_1504 *dm_build_1347) Dm_build_1503(dm_build_1505 *DmStatement, dm_build_1506 string) error {
	dm_build_1507 := dm_build_228(dm_build_1504, dm_build_1505, dm_build_1506)
	_, dm_build_1508 := dm_build_1504.dm_build_1390(dm_build_1507)
	if dm_build_1508 != nil {
		return dm_build_1508
	}

	return nil
}

func (dm_build_1510 *dm_build_1347) Dm_build_1509(dm_build_1511 []uint32) ([]int64, error) {
	dm_build_1512 := dm_build_641(dm_build_1510, dm_build_1511)
	dm_build_1513, dm_build_1514 := dm_build_1510.dm_build_1390(dm_build_1512)
	if dm_build_1514 != nil {
		return nil, dm_build_1514
	}
	return dm_build_1513.([]int64), nil
}

func (dm_build_1516 *dm_build_1347) Close() error {
	if dm_build_1516.dm_build_1358 {
		return nil
	}

	dm_build_1517 := dm_build_1516.dm_build_1348.Close()
	if dm_build_1517 != nil {
		return dm_build_1517
	}

	dm_build_1516.dm_build_1351 = nil
	dm_build_1516.dm_build_1358 = true
	return nil
}

func (dm_build_1519 *dm_build_1347) dm_build_1518(dm_build_1520 *lob) (int64, error) {
	dm_build_1521 := dm_build_410(dm_build_1519, dm_build_1520)
	dm_build_1522, dm_build_1523 := dm_build_1519.dm_build_1390(dm_build_1521)
	if dm_build_1523 != nil {
		return 0, dm_build_1523
	}
	return dm_build_1522.(int64), nil
}

func (dm_build_1525 *dm_build_1347) dm_build_1524(dm_build_1526 *lob, dm_build_1527 int32, dm_build_1528 int32) (*lobRetInfo, error) {
	dm_build_1529 := dm_build_395(dm_build_1525, dm_build_1526, int(dm_build_1527), int(dm_build_1528))
	dm_build_1530, dm_build_1531 := dm_build_1525.dm_build_1390(dm_build_1529)
	if dm_build_1531 != nil {
		return nil, dm_build_1531
	}
	return dm_build_1530.(*lobRetInfo), nil
}

func (dm_build_1533 *dm_build_1347) dm_build_1532(dm_build_1534 *DmBlob, dm_build_1535 int32, dm_build_1536 int32) ([]byte, error) {
	var dm_build_1537 = make([]byte, dm_build_1536)
	var dm_build_1538 int32 = 0
	var dm_build_1539 int32 = 0
	var dm_build_1540 *lobRetInfo
	var dm_build_1541 []byte
	var dm_build_1542 error
	for dm_build_1538 < dm_build_1536 {
		dm_build_1539 = dm_build_1536 - dm_build_1538
		if dm_build_1539 > Dm_build_128 {
			dm_build_1539 = Dm_build_128
		}
		dm_build_1540, dm_build_1542 = dm_build_1533.dm_build_1524(&dm_build_1534.lob, dm_build_1535+dm_build_1538, dm_build_1539)
		if dm_build_1542 != nil {
			return nil, dm_build_1542
		}
		dm_build_1541 = dm_build_1540.data
		if dm_build_1541 == nil || len(dm_build_1541) == 0 {
			break
		}
		Dm_build_652.Dm_build_708(dm_build_1537, int(dm_build_1538), dm_build_1541, 0, len(dm_build_1541))
		dm_build_1538 += int32(len(dm_build_1541))
		if dm_build_1534.readOver {
			break
		}
	}
	return dm_build_1537, nil
}

func (dm_build_1544 *dm_build_1347) dm_build_1543(dm_build_1545 *DmClob, dm_build_1546 int32, dm_build_1547 int32) (string, error) {
	var dm_build_1548 bytes.Buffer
	var dm_build_1549 int32 = 0
	var dm_build_1550 int32 = 0
	var dm_build_1551 *lobRetInfo
	var dm_build_1552 []byte
	var dm_build_1553 string
	var dm_build_1554 error
	for dm_build_1549 < dm_build_1547 {
		dm_build_1550 = dm_build_1547 - dm_build_1549
		if dm_build_1550 > Dm_build_128/2 {
			dm_build_1550 = Dm_build_128 / 2
		}
		dm_build_1551, dm_build_1554 = dm_build_1544.dm_build_1524(&dm_build_1545.lob, dm_build_1546+dm_build_1549, dm_build_1550)
		if dm_build_1554 != nil {
			return "", dm_build_1554
		}
		dm_build_1552 = dm_build_1551.data
		if dm_build_1552 == nil || len(dm_build_1552) == 0 {
			break
		}
		dm_build_1553 = Dm_build_652.Dm_build_809(dm_build_1552, 0, len(dm_build_1552), dm_build_1545.serverEncoding, dm_build_1544.dm_build_1351)

		dm_build_1548.WriteString(dm_build_1553)
		var strLen = dm_build_1551.charLen
		if strLen == -1 {
			strLen = int64(utf8.RuneCountInString(dm_build_1553))
		}
		dm_build_1549 += int32(strLen)
		if dm_build_1545.readOver {
			break
		}
	}
	return dm_build_1548.String(), nil
}

func (dm_build_1556 *dm_build_1347) dm_build_1555(dm_build_1557 *DmClob, dm_build_1558 int, dm_build_1559 string, dm_build_1560 string) (int, error) {
	var dm_build_1561 = Dm_build_652.Dm_build_868(dm_build_1559, dm_build_1560, dm_build_1556.dm_build_1351)
	var dm_build_1562 = 0
	var dm_build_1563 = len(dm_build_1561)
	var dm_build_1564 = 0
	var dm_build_1565 = 0
	var dm_build_1566 = 0
	var dm_build_1567 = dm_build_1563/Dm_build_127 + 1
	var dm_build_1568 byte = 0
	var dm_build_1569 byte = 0x01
	var dm_build_1570 byte = 0x02
	for i := 0; i < dm_build_1567; i++ {
		dm_build_1568 = 0
		if i == 0 {
			dm_build_1568 |= dm_build_1569
		}
		if i == dm_build_1567-1 {
			dm_build_1568 |= dm_build_1570
		}
		dm_build_1566 = dm_build_1563 - dm_build_1565
		if dm_build_1566 > Dm_build_127 {
			dm_build_1566 = Dm_build_127
		}

		setLobData := dm_build_557(dm_build_1556, &dm_build_1557.lob, dm_build_1568, dm_build_1558, dm_build_1561, dm_build_1562, dm_build_1566)
		ret, err := dm_build_1556.dm_build_1390(setLobData)
		if err != nil {
			return 0, err
		}
		tmp := ret.(int32)
		if err != nil {
			return -1, err
		}
		if tmp <= 0 {
			return dm_build_1564, nil
		} else {
			dm_build_1558 += int(tmp)
			dm_build_1564 += int(tmp)
			dm_build_1565 += dm_build_1566
			dm_build_1562 += dm_build_1566
		}
	}
	return dm_build_1564, nil
}

func (dm_build_1572 *dm_build_1347) dm_build_1571(dm_build_1573 *DmBlob, dm_build_1574 int, dm_build_1575 []byte) (int, error) {
	var dm_build_1576 = 0
	var dm_build_1577 = len(dm_build_1575)
	var dm_build_1578 = 0
	var dm_build_1579 = 0
	var dm_build_1580 = 0
	var dm_build_1581 = dm_build_1577/Dm_build_127 + 1
	var dm_build_1582 byte = 0
	var dm_build_1583 byte = 0x01
	var dm_build_1584 byte = 0x02
	for i := 0; i < dm_build_1581; i++ {
		dm_build_1582 = 0
		if i == 0 {
			dm_build_1582 |= dm_build_1583
		}
		if i == dm_build_1581-1 {
			dm_build_1582 |= dm_build_1584
		}
		dm_build_1580 = dm_build_1577 - dm_build_1579
		if dm_build_1580 > Dm_build_127 {
			dm_build_1580 = Dm_build_127
		}

		setLobData := dm_build_557(dm_build_1572, &dm_build_1573.lob, dm_build_1582, dm_build_1574, dm_build_1575, dm_build_1576, dm_build_1580)
		ret, err := dm_build_1572.dm_build_1390(setLobData)
		if err != nil {
			return 0, err
		}
		tmp := ret.(int32)
		if tmp <= 0 {
			return dm_build_1578, nil
		} else {
			dm_build_1574 += int(tmp)
			dm_build_1578 += int(tmp)
			dm_build_1579 += dm_build_1580
			dm_build_1576 += dm_build_1580
		}
	}
	return dm_build_1578, nil
}

func (dm_build_1586 *dm_build_1347) dm_build_1585(dm_build_1587 *lob, dm_build_1588 int) (int64, error) {
	dm_build_1589 := dm_build_421(dm_build_1586, dm_build_1587, dm_build_1588)
	dm_build_1590, dm_build_1591 := dm_build_1586.dm_build_1390(dm_build_1589)
	if dm_build_1591 != nil {
		return dm_build_1587.length, dm_build_1591
	}
	return dm_build_1590.(int64), nil
}

func (dm_build_1593 *dm_build_1347) dm_build_1592(dm_build_1594 []interface{}, dm_build_1595 []interface{}, dm_build_1596 int) bool {
	var dm_build_1597 = false
	dm_build_1594[dm_build_1596] = dm_build_1595[dm_build_1596]

	if binder, ok := dm_build_1595[dm_build_1596].(iOffRowBinder); ok {
		dm_build_1597 = true
		dm_build_1594[dm_build_1596] = make([]byte, 0)
		var lob lob
		if l, ok := binder.getObj().(DmBlob); ok {
			lob = l.lob
		} else if l, ok := binder.getObj().(DmClob); ok {
			lob = l.lob
		}
		if &lob != nil && lob.canOptimized(dm_build_1593.dm_build_1351) {
			dm_build_1594[dm_build_1596] = &lobCtl{lob.buildCtlData()}
			dm_build_1597 = false
		}
	} else {
		dm_build_1594[dm_build_1596] = dm_build_1595[dm_build_1596]
	}
	return dm_build_1597
}

func (dm_build_1599 *dm_build_1347) dm_build_1598(dm_build_1600 *DmStatement, dm_build_1601 parameter, dm_build_1602 int, dm_build_1603 iOffRowBinder) error {
	var dm_build_1604 = Dm_build_937()
	dm_build_1603.read(dm_build_1604)
	var dm_build_1605 = 0
	for !dm_build_1603.isReadOver() || dm_build_1604.Dm_build_938() > 0 {
		if !dm_build_1603.isReadOver() && dm_build_1604.Dm_build_938() < Dm_build_127 {
			dm_build_1603.read(dm_build_1604)
		}
		if dm_build_1604.Dm_build_938() > Dm_build_127 {
			dm_build_1605 = Dm_build_127
		} else {
			dm_build_1605 = dm_build_1604.Dm_build_938()
		}

		putData := dm_build_528(dm_build_1599, dm_build_1600, int16(dm_build_1602), dm_build_1604, int32(dm_build_1605))
		_, err := dm_build_1599.dm_build_1390(putData)
		if err != nil {
			return err
		}
	}
	return nil
}

func (dm_build_1607 *dm_build_1347) dm_build_1606() ([]byte, error) {
	var dm_build_1608 error
	if dm_build_1607.dm_build_1355 == nil {
		if dm_build_1607.dm_build_1355, dm_build_1608 = security.NewClientKeyPair(); dm_build_1608 != nil {
			return nil, dm_build_1608
		}
	}
	return security.Bn2Bytes(dm_build_1607.dm_build_1355.GetY(), security.DH_KEY_LENGTH), nil
}

func (dm_build_1610 *dm_build_1347) dm_build_1609() (*security.DhKey, error) {
	var dm_build_1611 error
	if dm_build_1610.dm_build_1355 == nil {
		if dm_build_1610.dm_build_1355, dm_build_1611 = security.NewClientKeyPair(); dm_build_1611 != nil {
			return nil, dm_build_1611
		}
	}
	return dm_build_1610.dm_build_1355, nil
}

func (dm_build_1613 *dm_build_1347) dm_build_1612(dm_build_1614 int, dm_build_1615 []byte, dm_build_1616 string, dm_build_1617 int) (dm_build_1618 error) {
	if dm_build_1614 > 0 && dm_build_1614 < security.MIN_EXTERNAL_CIPHER_ID && dm_build_1615 != nil {
		dm_build_1613.dm_build_1352, dm_build_1618 = security.NewSymmCipher(dm_build_1614, dm_build_1615)
	} else if dm_build_1614 >= security.MIN_EXTERNAL_CIPHER_ID {
		if dm_build_1613.dm_build_1352, dm_build_1618 = security.NewThirdPartCipher(dm_build_1614, dm_build_1615, dm_build_1616, dm_build_1617); dm_build_1618 != nil {
			dm_build_1618 = THIRD_PART_CIPHER_INIT_FAILED.addDetailln(dm_build_1618.Error()).throw()
		}
	}
	return
}

func (dm_build_1620 *dm_build_1347) dm_build_1619(dm_build_1621 bool) (dm_build_1622 error) {
	if dm_build_1620.dm_build_1349, dm_build_1622 = security.NewTLSFromTCP(dm_build_1620.dm_build_1348, dm_build_1620.dm_build_1351.dmConnector.sslCertPath, dm_build_1620.dm_build_1351.dmConnector.sslKeyPath, dm_build_1620.dm_build_1351.dmConnector.user); dm_build_1622 != nil {
		return
	}
	if !dm_build_1621 {
		dm_build_1620.dm_build_1349 = nil
	}
	return
}

func (dm_build_1624 *dm_build_1347) dm_build_1623(dm_build_1625 dm_build_135) bool {
	return dm_build_1625.dm_build_150() != Dm_build_42 && dm_build_1624.dm_build_1351.sslEncrypt == 1
}
