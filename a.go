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
	Dm_build_412 = 8192
	Dm_build_413 = 2 * time.Second
)

type dm_build_414 struct {
	dm_build_415 net.Conn
	dm_build_416 *tls.Conn
	dm_build_417 *Dm_build_78
	dm_build_418 *DmConnection
	dm_build_419 security.Cipher
	dm_build_420 bool
	dm_build_421 bool
	dm_build_422 *security.DhKey

	dm_build_423 bool
	dm_build_424 string
	dm_build_425 bool
}

func dm_build_426(dm_build_427 context.Context, dm_build_428 *DmConnection) (*dm_build_414, error) {
	var dm_build_429 net.Conn
	var dm_build_430 error

	dialsLock.RLock()
	dm_build_431, dm_build_432 := dials[dm_build_428.dmConnector.dialName]
	dialsLock.RUnlock()
	if dm_build_432 {
		dm_build_429, dm_build_430 = dm_build_431(dm_build_427, dm_build_428.dmConnector.host+":"+strconv.Itoa(int(dm_build_428.dmConnector.port)))
	} else {
		dm_build_429, dm_build_430 = dm_build_434(dm_build_428.dmConnector.host+":"+strconv.Itoa(int(dm_build_428.dmConnector.port)), time.Duration(dm_build_428.dmConnector.socketTimeout)*time.Second)
	}
	if dm_build_430 != nil {
		return nil, dm_build_430
	}

	dm_build_433 := dm_build_414{}
	dm_build_433.dm_build_415 = dm_build_429
	dm_build_433.dm_build_417 = Dm_build_81(Dm_build_709)
	dm_build_433.dm_build_418 = dm_build_428
	dm_build_433.dm_build_420 = false
	dm_build_433.dm_build_421 = false
	dm_build_433.dm_build_423 = false
	dm_build_433.dm_build_424 = ""
	dm_build_433.dm_build_425 = false
	dm_build_428.Access = &dm_build_433

	return &dm_build_433, nil
}

func dm_build_434(dm_build_435 string, dm_build_436 time.Duration) (net.Conn, error) {
	dm_build_437, dm_build_438 := net.DialTimeout("tcp", dm_build_435, dm_build_436)
	if dm_build_438 != nil {
		return &net.TCPConn{}, ECGO_COMMUNITION_ERROR.addDetail("\tdial address: " + dm_build_435).throw()
	}

	if tcpConn, ok := dm_build_437.(*net.TCPConn); ok {
		tcpConn.SetKeepAlive(true)
		tcpConn.SetKeepAlivePeriod(Dm_build_413)
		tcpConn.SetNoDelay(true)

	}
	return dm_build_437, nil
}

func (dm_build_440 *dm_build_414) dm_build_439(dm_build_441 dm_build_831) bool {
	var dm_build_442 = dm_build_440.dm_build_418.dmConnector.compress
	if dm_build_441.dm_build_846() == Dm_build_737 || dm_build_442 == Dm_build_786 {
		return false
	}

	if dm_build_442 == Dm_build_784 {
		return true
	} else if dm_build_442 == Dm_build_785 {
		return !dm_build_440.dm_build_418.Local && dm_build_441.dm_build_844() > Dm_build_783
	}

	return false
}

func (dm_build_444 *dm_build_414) dm_build_443(dm_build_445 dm_build_831) bool {
	var dm_build_446 = dm_build_444.dm_build_418.dmConnector.compress
	if dm_build_445.dm_build_846() == Dm_build_737 || dm_build_446 == Dm_build_786 {
		return false
	}

	if dm_build_446 == Dm_build_784 {
		return true
	} else if dm_build_446 == Dm_build_785 {
		return dm_build_444.dm_build_417.Dm_build_345(Dm_build_745) == 1
	}

	return false
}

func (dm_build_448 *dm_build_414) dm_build_447(dm_build_449 dm_build_831) (err error) {
	defer func() {
		if p := recover(); p != nil {
			if _, ok := p.(string); ok {
				err = ECGO_COMMUNITION_ERROR.addDetail("\t" + p.(string)).throw()
			} else {
				err = fmt.Errorf("internal error: %v", p)
			}
		}
	}()

	dm_build_451 := dm_build_449.dm_build_844()

	if dm_build_451 > 0 {

		if dm_build_448.dm_build_439(dm_build_449) {
			var retBytes, err = Compress(dm_build_448.dm_build_417, Dm_build_738, int(dm_build_451), int(dm_build_448.dm_build_418.dmConnector.compressID))
			if err != nil {
				return err
			}

			dm_build_448.dm_build_417.Dm_build_92(Dm_build_738)

			dm_build_448.dm_build_417.Dm_build_133(dm_build_451)

			dm_build_448.dm_build_417.Dm_build_161(retBytes)

			dm_build_449.dm_build_845(int32(len(retBytes)) + ULINT_SIZE)

			dm_build_448.dm_build_417.Dm_build_265(Dm_build_745, 1)
		}

		if dm_build_448.dm_build_421 {
			dm_build_451 = dm_build_449.dm_build_844()
			var retBytes = dm_build_448.dm_build_419.Encrypt(dm_build_448.dm_build_417.Dm_build_372(Dm_build_738, int(dm_build_451)), true)

			dm_build_448.dm_build_417.Dm_build_92(Dm_build_738)

			dm_build_448.dm_build_417.Dm_build_161(retBytes)

			dm_build_449.dm_build_845(int32(len(retBytes)))
		}
	}

	if dm_build_448.dm_build_417.Dm_build_90() > Dm_build_710 {
		return ECGO_MSG_TOO_LONG.throw()
	}

	dm_build_449.dm_build_840()
	if dm_build_448.dm_build_691(dm_build_449) {
		if dm_build_448.dm_build_416 != nil {
			dm_build_448.dm_build_417.Dm_build_95(0)
			if _, err := dm_build_448.dm_build_417.Dm_build_114(dm_build_448.dm_build_416); err != nil {
				return err
			}
		}
	} else {
		dm_build_448.dm_build_417.Dm_build_95(0)
		if _, err := dm_build_448.dm_build_417.Dm_build_114(dm_build_448.dm_build_415); err != nil {
			return err
		}
	}
	return nil
}

func (dm_build_453 *dm_build_414) dm_build_452(dm_build_454 dm_build_831) (err error) {
	defer func() {
		if p := recover(); p != nil {
			if _, ok := p.(string); ok {
				err = ECGO_COMMUNITION_ERROR.addDetail("\t" + p.(string)).throw()
			} else {
				err = fmt.Errorf("internal error: %v", p)
			}
		}
	}()

	dm_build_456 := int32(0)
	if dm_build_453.dm_build_691(dm_build_454) {
		if dm_build_453.dm_build_416 != nil {
			dm_build_453.dm_build_417.Dm_build_92(0)
			if _, err := dm_build_453.dm_build_417.Dm_build_108(dm_build_453.dm_build_416, Dm_build_738); err != nil {
				return err
			}

			dm_build_456 = dm_build_454.dm_build_844()
			if dm_build_456 > 0 {
				if _, err := dm_build_453.dm_build_417.Dm_build_108(dm_build_453.dm_build_416, int(dm_build_456)); err != nil {
					return err
				}
			}
		}
	} else {

		dm_build_453.dm_build_417.Dm_build_92(0)
		if _, err := dm_build_453.dm_build_417.Dm_build_108(dm_build_453.dm_build_415, Dm_build_738); err != nil {
			return err
		}
		dm_build_456 = dm_build_454.dm_build_844()

		if dm_build_456 > 0 {
			if _, err := dm_build_453.dm_build_417.Dm_build_108(dm_build_453.dm_build_415, int(dm_build_456)); err != nil {
				return err
			}
		}
	}

	dm_build_454.dm_build_841()

	dm_build_456 = dm_build_454.dm_build_844()
	if dm_build_456 <= 0 {
		return nil
	}

	if dm_build_453.dm_build_421 {
		ebytes := dm_build_453.dm_build_417.Dm_build_372(Dm_build_738, int(dm_build_456))
		bytes, err := dm_build_453.dm_build_419.Decrypt(ebytes, true)
		if err != nil {
			return err
		}
		dm_build_453.dm_build_417.Dm_build_92(Dm_build_738)
		dm_build_453.dm_build_417.Dm_build_161(bytes)
		dm_build_454.dm_build_845(int32(len(bytes)))
	}

	if dm_build_453.dm_build_443(dm_build_454) {

		dm_build_456 = dm_build_454.dm_build_844()
		cbytes := dm_build_453.dm_build_417.Dm_build_372(Dm_build_738+ULINT_SIZE, int(dm_build_456-ULINT_SIZE))
		bytes, err := UnCompress(cbytes, int(dm_build_453.dm_build_418.dmConnector.compressID))
		if err != nil {
			return err
		}
		dm_build_453.dm_build_417.Dm_build_92(Dm_build_738)
		dm_build_453.dm_build_417.Dm_build_161(bytes)
		dm_build_454.dm_build_845(int32(len(bytes)))
	}
	return nil
}

func (dm_build_458 *dm_build_414) dm_build_457(dm_build_459 dm_build_831) (dm_build_460 interface{}, dm_build_461 error) {
	if dm_build_458.dm_build_425 {
		return nil, ECGO_CONNECTION_CLOSED.throw()
	}
	dm_build_462 := dm_build_458.dm_build_418
	dm_build_462.mu.Lock()
	defer dm_build_462.mu.Unlock()
	dm_build_461 = dm_build_459.dm_build_835(dm_build_459)
	if dm_build_461 != nil {
		return nil, dm_build_461
	}

	dm_build_461 = dm_build_458.dm_build_447(dm_build_459)
	if dm_build_461 != nil {
		return nil, dm_build_461
	}

	dm_build_461 = dm_build_458.dm_build_452(dm_build_459)
	if dm_build_461 != nil {
		return nil, dm_build_461
	}

	return dm_build_459.dm_build_839(dm_build_459)
}

func (dm_build_464 *dm_build_414) dm_build_463() (*dm_build_1293, error) {

	Dm_build_465 := dm_build_1299(dm_build_464)
	_, dm_build_466 := dm_build_464.dm_build_457(Dm_build_465)
	if dm_build_466 != nil {
		return nil, dm_build_466
	}

	return Dm_build_465, nil
}

func (dm_build_468 *dm_build_414) dm_build_467() error {

	dm_build_469 := dm_build_1157(dm_build_468)
	_, dm_build_470 := dm_build_468.dm_build_457(dm_build_469)
	if dm_build_470 != nil {
		return dm_build_470
	}

	return nil
}

func (dm_build_472 *dm_build_414) dm_build_471() error {

	var dm_build_473 *dm_build_1293
	var err error
	if dm_build_473, err = dm_build_472.dm_build_463(); err != nil {
		return err
	}

	if dm_build_472.dm_build_418.sslEncrypt == 2 {
		if err = dm_build_472.dm_build_686(false, true); err != nil {
			return ECGO_INIT_SSL_FAILED.addDetail("\n" + err.Error()).throw()
		}
	} else if dm_build_472.dm_build_418.sslEncrypt == 1 || dm_build_472.dm_build_418.sslEncrypt == 4 {
		if err = dm_build_472.dm_build_686(true, false); err != nil {
			return ECGO_INIT_SSL_FAILED.addDetail("\n" + err.Error()).throw()
		}
	}

	if dm_build_472.dm_build_421 || dm_build_472.dm_build_420 {
		k, err := dm_build_472.dm_build_676()
		if err != nil {
			return err
		}
		sessionKey := security.ComputeSessionKey(k, dm_build_473.Dm_build_1297)
		encryptType := dm_build_473.dm_build_1295
		hashType := int(dm_build_473.Dm_build_1296)
		if encryptType == -1 {
			encryptType = security.DES_CFB
		}
		if hashType == -1 {
			hashType = security.MD5
		}
		err = dm_build_472.dm_build_679(encryptType, sessionKey, dm_build_472.dm_build_418.dmConnector.cipherPath, hashType)
		if err != nil {
			return err
		}
	}

	if err := dm_build_472.dm_build_467(); err != nil {
		return err
	}
	return nil
}

func (dm_build_476 *dm_build_414) Dm_build_475(dm_build_477 *DmStatement) error {
	dm_build_478 := dm_build_1326(dm_build_476, dm_build_477)
	_, dm_build_479 := dm_build_476.dm_build_457(dm_build_478)
	if dm_build_479 != nil {
		return dm_build_479
	}

	return nil
}

func (dm_build_481 *dm_build_414) Dm_build_480(dm_build_482 int32) error {
	dm_build_483 := dm_build_1336(dm_build_481, dm_build_482)
	_, dm_build_484 := dm_build_481.dm_build_457(dm_build_483)
	if dm_build_484 != nil {
		return dm_build_484
	}

	return nil
}

func (dm_build_486 *dm_build_414) Dm_build_485(dm_build_487 *DmStatement, dm_build_488 bool, dm_build_489 int16) (*execRetInfo, error) {
	dm_build_490 := dm_build_1199(dm_build_486, dm_build_487, dm_build_488, dm_build_489)
	dm_build_491, dm_build_492 := dm_build_486.dm_build_457(dm_build_490)
	if dm_build_492 != nil {
		return nil, dm_build_492
	}
	return dm_build_491.(*execRetInfo), nil
}

func (dm_build_494 *dm_build_414) Dm_build_493(dm_build_495 *DmStatement, dm_build_496 int16) (*execRetInfo, error) {
	return dm_build_494.Dm_build_485(dm_build_495, false, Dm_build_790)
}

func (dm_build_498 *dm_build_414) Dm_build_497(dm_build_499 *DmStatement, dm_build_500 []OptParameter) (*execRetInfo, error) {
	dm_build_501, dm_build_502 := dm_build_498.dm_build_457(dm_build_934(dm_build_498, dm_build_499, dm_build_500))
	if dm_build_502 != nil {
		return nil, dm_build_502
	}

	return dm_build_501.(*execRetInfo), nil
}

func (dm_build_504 *dm_build_414) Dm_build_503(dm_build_505 *DmStatement, dm_build_506 int16) (*execRetInfo, error) {
	return dm_build_504.Dm_build_485(dm_build_505, true, dm_build_506)
}

func (dm_build_508 *dm_build_414) Dm_build_507(dm_build_509 *DmStatement, dm_build_510 [][]interface{}) (*execRetInfo, error) {
	dm_build_511 := dm_build_966(dm_build_508, dm_build_509, dm_build_510)
	dm_build_512, dm_build_513 := dm_build_508.dm_build_457(dm_build_511)
	if dm_build_513 != nil {
		return nil, dm_build_513
	}
	return dm_build_512.(*execRetInfo), nil
}

func (dm_build_515 *dm_build_414) Dm_build_514(dm_build_516 *DmStatement, dm_build_517 [][]interface{}, dm_build_518 bool) (*execRetInfo, error) {
	var dm_build_519, dm_build_520 = 0, 0
	var dm_build_521 = len(dm_build_517)
	var dm_build_522 [][]interface{}
	var dm_build_523 = NewExceInfo()
	dm_build_523.updateCounts = make([]int64, dm_build_521)
	var dm_build_524 = false
	for dm_build_519 < dm_build_521 {
		for dm_build_520 = dm_build_519; dm_build_520 < dm_build_521; dm_build_520++ {
			paramData := dm_build_517[dm_build_520]
			bindData := make([]interface{}, dm_build_516.paramCount)
			dm_build_524 = false
			for icol := 0; icol < int(dm_build_516.paramCount); icol++ {
				if dm_build_516.bindParams[icol].ioType == IO_TYPE_OUT {
					continue
				}
				if dm_build_515.dm_build_659(bindData, paramData, icol) {
					dm_build_524 = true
					break
				}
			}

			if dm_build_524 {
				break
			}
			dm_build_522 = append(dm_build_522, bindData)
		}

		if dm_build_520 != dm_build_519 {
			tmpExecInfo, err := dm_build_515.Dm_build_507(dm_build_516, dm_build_522)
			if err != nil {
				return nil, err
			}
			dm_build_522 = dm_build_522[0:0]
			dm_build_523.union(tmpExecInfo, dm_build_519, dm_build_520-dm_build_519)
		}

		if dm_build_520 < dm_build_521 {
			tmpExecInfo, err := dm_build_515.Dm_build_533(dm_build_516, dm_build_517[dm_build_520], dm_build_518)
			if err != nil {
				return nil, err
			}

			dm_build_518 = true
			dm_build_523.union(tmpExecInfo, dm_build_520, 1)
		}

		dm_build_519 = dm_build_520 + 1
	}
	for _, i := range dm_build_523.updateCounts {
		if i > 0 {
			dm_build_523.updateCount += i
		}
	}
	return dm_build_523, nil
}

func (dm_build_526 *dm_build_414) dm_build_525(dm_build_527 *DmStatement, dm_build_528 []parameter) error {
	if !dm_build_527.prepared {
		retInfo, err := dm_build_526.Dm_build_485(dm_build_527, false, Dm_build_790)
		if err != nil {
			return nil
		}
		dm_build_527.serverParams = retInfo.serverParams
		dm_build_527.paramCount = int32(len(dm_build_527.serverParams))
		dm_build_527.prepared = true
	}

	dm_build_529 := dm_build_1188(dm_build_526, dm_build_527, dm_build_527.bindParams)
	dm_build_530, err := dm_build_526.dm_build_457(dm_build_529)
	if err != nil {
		return nil
	}
	retInfo := dm_build_530.(*execRetInfo)
	if retInfo.serverParams != nil && len(retInfo.serverParams) > 0 {
		dm_build_527.serverParams = retInfo.serverParams
		dm_build_527.paramCount = int32(len(dm_build_527.serverParams))
	}
	dm_build_527.preExec = true
	return nil
}

func (dm_build_534 *dm_build_414) Dm_build_533(dm_build_535 *DmStatement, dm_build_536 []interface{}, dm_build_537 bool) (*execRetInfo, error) {

	var dm_build_538 = make([]interface{}, dm_build_535.paramCount)
	for icol := 0; icol < int(dm_build_535.paramCount); icol++ {
		if dm_build_535.bindParams[icol].ioType == IO_TYPE_OUT {
			continue
		}
		if dm_build_534.dm_build_659(dm_build_538, dm_build_536, icol) {

			if !dm_build_537 {
				dm_build_534.dm_build_525(dm_build_535, dm_build_535.bindParams)

				dm_build_537 = true
			}

			dm_build_534.dm_build_665(dm_build_535, dm_build_535.bindParams[icol], icol, dm_build_536[icol].(iOffRowBinder))
			dm_build_538[icol] = ParamDataEnum_OFF_ROW
		}
	}

	var dm_build_539 = make([][]interface{}, 1, 1)
	dm_build_539[0] = dm_build_538

	dm_build_540 := dm_build_966(dm_build_534, dm_build_535, dm_build_539)
	dm_build_541, dm_build_542 := dm_build_534.dm_build_457(dm_build_540)
	if dm_build_542 != nil {
		return nil, dm_build_542
	}
	return dm_build_541.(*execRetInfo), nil
}

func (dm_build_544 *dm_build_414) Dm_build_543(dm_build_545 *DmStatement, dm_build_546 int16) (*execRetInfo, error) {
	dm_build_547 := dm_build_1175(dm_build_544, dm_build_545, dm_build_546)

	dm_build_548, dm_build_549 := dm_build_544.dm_build_457(dm_build_547)
	if dm_build_549 != nil {
		return nil, dm_build_549
	}
	return dm_build_548.(*execRetInfo), nil
}

func (dm_build_551 *dm_build_414) Dm_build_550(dm_build_552 *innerRows, dm_build_553 int64) (*execRetInfo, error) {
	dm_build_554 := dm_build_1074(dm_build_551, dm_build_552, dm_build_553, INT64_MAX)
	dm_build_555, dm_build_556 := dm_build_551.dm_build_457(dm_build_554)
	if dm_build_556 != nil {
		return nil, dm_build_556
	}
	return dm_build_555.(*execRetInfo), nil
}

func (dm_build_558 *dm_build_414) Commit() error {
	dm_build_559 := dm_build_919(dm_build_558)
	_, dm_build_560 := dm_build_558.dm_build_457(dm_build_559)
	if dm_build_560 != nil {
		return dm_build_560
	}

	return nil
}

func (dm_build_562 *dm_build_414) Rollback() error {
	dm_build_563 := dm_build_1237(dm_build_562)
	_, dm_build_564 := dm_build_562.dm_build_457(dm_build_563)
	if dm_build_564 != nil {
		return dm_build_564
	}

	return nil
}

func (dm_build_566 *dm_build_414) Dm_build_565(dm_build_567 *DmConnection) error {
	dm_build_568 := dm_build_1242(dm_build_566, dm_build_567.IsoLevel)
	_, dm_build_569 := dm_build_566.dm_build_457(dm_build_568)
	if dm_build_569 != nil {
		return dm_build_569
	}

	return nil
}

func (dm_build_571 *dm_build_414) Dm_build_570(dm_build_572 *DmStatement, dm_build_573 string) error {
	dm_build_574 := dm_build_924(dm_build_571, dm_build_572, dm_build_573)
	_, dm_build_575 := dm_build_571.dm_build_457(dm_build_574)
	if dm_build_575 != nil {
		return dm_build_575
	}

	return nil
}

func (dm_build_577 *dm_build_414) Dm_build_576(dm_build_578 []uint32) ([]int64, error) {
	dm_build_579 := dm_build_1344(dm_build_577, dm_build_578)
	dm_build_580, dm_build_581 := dm_build_577.dm_build_457(dm_build_579)
	if dm_build_581 != nil {
		return nil, dm_build_581
	}
	return dm_build_580.([]int64), nil
}

func (dm_build_583 *dm_build_414) Close() error {
	if dm_build_583.dm_build_425 {
		return nil
	}

	dm_build_584 := dm_build_583.dm_build_415.Close()
	if dm_build_584 != nil {
		return dm_build_584
	}

	dm_build_583.dm_build_425 = true
	return nil
}

func (dm_build_586 *dm_build_414) dm_build_585(dm_build_587 *lob) (int64, error) {
	dm_build_588 := dm_build_1107(dm_build_586, dm_build_587)
	dm_build_589, dm_build_590 := dm_build_586.dm_build_457(dm_build_588)
	if dm_build_590 != nil {
		return 0, dm_build_590
	}
	return dm_build_589.(int64), nil
}

func (dm_build_592 *dm_build_414) dm_build_591(dm_build_593 *lob, dm_build_594 int32, dm_build_595 int32) (*lobRetInfo, error) {
	dm_build_596 := dm_build_1092(dm_build_592, dm_build_593, int(dm_build_594), int(dm_build_595))
	dm_build_597, dm_build_598 := dm_build_592.dm_build_457(dm_build_596)
	if dm_build_598 != nil {
		return nil, dm_build_598
	}
	return dm_build_597.(*lobRetInfo), nil
}

func (dm_build_600 *dm_build_414) dm_build_599(dm_build_601 *DmBlob, dm_build_602 int32, dm_build_603 int32) ([]byte, error) {
	var dm_build_604 = make([]byte, dm_build_603)
	var dm_build_605 int32 = 0
	var dm_build_606 int32 = 0
	var dm_build_607 *lobRetInfo
	var dm_build_608 []byte
	var dm_build_609 error
	for dm_build_605 < dm_build_603 {
		dm_build_606 = dm_build_603 - dm_build_605
		if dm_build_606 > Dm_build_824 {
			dm_build_606 = Dm_build_824
		}
		dm_build_607, dm_build_609 = dm_build_600.dm_build_591(&dm_build_601.lob, dm_build_602+dm_build_605, dm_build_606)
		if dm_build_609 != nil {
			return nil, dm_build_609
		}
		dm_build_608 = dm_build_607.data
		if dm_build_608 == nil || len(dm_build_608) == 0 {
			break
		}
		Dm_build_1355.Dm_build_1411(dm_build_604, int(dm_build_605), dm_build_608, 0, len(dm_build_608))
		dm_build_605 += int32(len(dm_build_608))
		if dm_build_601.readOver {
			break
		}
	}
	return dm_build_604, nil
}

func (dm_build_611 *dm_build_414) dm_build_610(dm_build_612 *DmClob, dm_build_613 int32, dm_build_614 int32) (string, error) {
	var dm_build_615 bytes.Buffer
	var dm_build_616 int32 = 0
	var dm_build_617 int32 = 0
	var dm_build_618 *lobRetInfo
	var dm_build_619 []byte
	var dm_build_620 string
	var dm_build_621 error
	for dm_build_616 < dm_build_614 {
		dm_build_617 = dm_build_614 - dm_build_616
		if dm_build_617 > Dm_build_824/2 {
			dm_build_617 = Dm_build_824 / 2
		}
		dm_build_618, dm_build_621 = dm_build_611.dm_build_591(&dm_build_612.lob, dm_build_613+dm_build_616, dm_build_617)
		if dm_build_621 != nil {
			return "", dm_build_621
		}
		dm_build_619 = dm_build_618.data
		if dm_build_619 == nil || len(dm_build_619) == 0 {
			break
		}
		dm_build_620 = Dm_build_1355.Dm_build_1512(dm_build_619, 0, len(dm_build_619), dm_build_612.serverEncoding, dm_build_611.dm_build_418)

		dm_build_615.WriteString(dm_build_620)
		var strLen = dm_build_618.charLen
		if strLen == -1 {
			strLen = int64(utf8.RuneCountInString(dm_build_620))
		}
		dm_build_616 += int32(strLen)
		if dm_build_612.readOver {
			break
		}
	}
	return dm_build_615.String(), nil
}

func (dm_build_623 *dm_build_414) dm_build_622(dm_build_624 *DmClob, dm_build_625 int, dm_build_626 string, dm_build_627 string) (int, error) {
	var dm_build_628 = Dm_build_1355.Dm_build_1571(dm_build_626, dm_build_627, dm_build_623.dm_build_418)
	var dm_build_629 = 0
	var dm_build_630 = len(dm_build_628)
	var dm_build_631 = 0
	var dm_build_632 = 0
	var dm_build_633 = 0
	var dm_build_634 = dm_build_630/Dm_build_823 + 1
	var dm_build_635 byte = 0
	var dm_build_636 byte = 0x01
	var dm_build_637 byte = 0x02
	for i := 0; i < dm_build_634; i++ {
		dm_build_635 = 0
		if i == 0 {
			dm_build_635 |= dm_build_636
		}
		if i == dm_build_634-1 {
			dm_build_635 |= dm_build_637
		}
		dm_build_633 = dm_build_630 - dm_build_632
		if dm_build_633 > Dm_build_823 {
			dm_build_633 = Dm_build_823
		}

		setLobData := dm_build_1256(dm_build_623, &dm_build_624.lob, dm_build_635, dm_build_625, dm_build_628, dm_build_629, dm_build_633)
		ret, err := dm_build_623.dm_build_457(setLobData)
		if err != nil {
			return 0, err
		}
		tmp := ret.(int32)
		if err != nil {
			return -1, err
		}
		if tmp <= 0 {
			return dm_build_631, nil
		} else {
			dm_build_625 += int(tmp)
			dm_build_631 += int(tmp)
			dm_build_632 += dm_build_633
			dm_build_629 += dm_build_633
		}
	}
	return dm_build_631, nil
}

func (dm_build_639 *dm_build_414) dm_build_638(dm_build_640 *DmBlob, dm_build_641 int, dm_build_642 []byte) (int, error) {
	var dm_build_643 = 0
	var dm_build_644 = len(dm_build_642)
	var dm_build_645 = 0
	var dm_build_646 = 0
	var dm_build_647 = 0
	var dm_build_648 = dm_build_644/Dm_build_823 + 1
	var dm_build_649 byte = 0
	var dm_build_650 byte = 0x01
	var dm_build_651 byte = 0x02
	for i := 0; i < dm_build_648; i++ {
		dm_build_649 = 0
		if i == 0 {
			dm_build_649 |= dm_build_650
		}
		if i == dm_build_648-1 {
			dm_build_649 |= dm_build_651
		}
		dm_build_647 = dm_build_644 - dm_build_646
		if dm_build_647 > Dm_build_823 {
			dm_build_647 = Dm_build_823
		}

		setLobData := dm_build_1256(dm_build_639, &dm_build_640.lob, dm_build_649, dm_build_641, dm_build_642, dm_build_643, dm_build_647)
		ret, err := dm_build_639.dm_build_457(setLobData)
		if err != nil {
			return 0, err
		}
		tmp := ret.(int32)
		if tmp <= 0 {
			return dm_build_645, nil
		} else {
			dm_build_641 += int(tmp)
			dm_build_645 += int(tmp)
			dm_build_646 += dm_build_647
			dm_build_643 += dm_build_647
		}
	}
	return dm_build_645, nil
}

func (dm_build_653 *dm_build_414) dm_build_652(dm_build_654 *lob, dm_build_655 int) (int64, error) {
	dm_build_656 := dm_build_1118(dm_build_653, dm_build_654, dm_build_655)
	dm_build_657, dm_build_658 := dm_build_653.dm_build_457(dm_build_656)
	if dm_build_658 != nil {
		return dm_build_654.length, dm_build_658
	}
	return dm_build_657.(int64), nil
}

func (dm_build_660 *dm_build_414) dm_build_659(dm_build_661 []interface{}, dm_build_662 []interface{}, dm_build_663 int) bool {
	var dm_build_664 = false
	dm_build_661[dm_build_663] = dm_build_662[dm_build_663]

	if binder, ok := dm_build_662[dm_build_663].(iOffRowBinder); ok {
		dm_build_664 = true
		dm_build_661[dm_build_663] = make([]byte, 0)
		var lob lob
		if l, ok := binder.getObj().(DmBlob); ok {
			lob = l.lob
		} else if l, ok := binder.getObj().(DmClob); ok {
			lob = l.lob
		}
		if &lob != nil && lob.canOptimized(dm_build_660.dm_build_418) {
			dm_build_661[dm_build_663] = &lobCtl{lob.buildCtlData()}
			dm_build_664 = false
		}
	} else {
		dm_build_661[dm_build_663] = dm_build_662[dm_build_663]
	}
	return dm_build_664
}

func (dm_build_666 *dm_build_414) dm_build_665(dm_build_667 *DmStatement, dm_build_668 parameter, dm_build_669 int, dm_build_670 iOffRowBinder) error {
	var dm_build_671 = Dm_build_4()
	dm_build_670.read(dm_build_671)
	var dm_build_672 = 0
	for !dm_build_670.isReadOver() || dm_build_671.Dm_build_5() > 0 {
		if !dm_build_670.isReadOver() && dm_build_671.Dm_build_5() < Dm_build_823 {
			dm_build_670.read(dm_build_671)
		}
		if dm_build_671.Dm_build_5() > Dm_build_823 {
			dm_build_672 = Dm_build_823
		} else {
			dm_build_672 = dm_build_671.Dm_build_5()
		}

		putData := dm_build_1227(dm_build_666, dm_build_667, int16(dm_build_669), dm_build_671, int32(dm_build_672))
		_, err := dm_build_666.dm_build_457(putData)
		if err != nil {
			return err
		}
	}
	return nil
}

func (dm_build_674 *dm_build_414) dm_build_673() ([]byte, error) {
	var dm_build_675 error
	if dm_build_674.dm_build_422 == nil {
		if dm_build_674.dm_build_422, dm_build_675 = security.NewClientKeyPair(); dm_build_675 != nil {
			return nil, dm_build_675
		}
	}
	return security.Bn2Bytes(dm_build_674.dm_build_422.GetY(), security.DH_KEY_LENGTH), nil
}

func (dm_build_677 *dm_build_414) dm_build_676() (*security.DhKey, error) {
	var dm_build_678 error
	if dm_build_677.dm_build_422 == nil {
		if dm_build_677.dm_build_422, dm_build_678 = security.NewClientKeyPair(); dm_build_678 != nil {
			return nil, dm_build_678
		}
	}
	return dm_build_677.dm_build_422, nil
}

func (dm_build_680 *dm_build_414) dm_build_679(dm_build_681 int, dm_build_682 []byte, dm_build_683 string, dm_build_684 int) (dm_build_685 error) {
	if dm_build_681 > 0 && dm_build_681 < security.MIN_EXTERNAL_CIPHER_ID && dm_build_682 != nil {
		dm_build_680.dm_build_419, dm_build_685 = security.NewSymmCipher(dm_build_681, dm_build_682)
	} else if dm_build_681 >= security.MIN_EXTERNAL_CIPHER_ID {
		if dm_build_680.dm_build_419, dm_build_685 = security.NewThirdPartCipher(dm_build_681, dm_build_682, dm_build_683, dm_build_684); dm_build_685 != nil {
			dm_build_685 = THIRD_PART_CIPHER_INIT_FAILED.addDetailln(dm_build_685.Error()).throw()
		}
	}
	return
}

func (dm_build_687 *dm_build_414) dm_build_686(dm_build_688 bool, dm_build_689 bool) (dm_build_690 error) {
	if dm_build_687.dm_build_416, dm_build_690 = security.NewTLSFromTCP(dm_build_687.dm_build_415, dm_build_687.dm_build_418.dmConnector.sslCertPath, dm_build_687.dm_build_418.dmConnector.sslKeyPath, dm_build_687.dm_build_418.dmConnector.user, dm_build_689); dm_build_690 != nil {
		return
	}
	if !dm_build_688 {
		dm_build_687.dm_build_416 = nil
	}
	return
}

func (dm_build_692 *dm_build_414) dm_build_691(dm_build_693 dm_build_831) bool {
	return dm_build_693.dm_build_846() != Dm_build_737 && (dm_build_692.dm_build_418.sslEncrypt == 1 || dm_build_692.dm_build_418.sslEncrypt == 4)
}
